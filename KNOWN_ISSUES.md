
<!-- markdownlint-disable no-hard-tabs -->
# Known issues

- [Known issues](#known-issues)
  - [Resource Import Not Implemented](#resource-import-not-implemented)
  - [Nested Schema](#nested-schema)
  - [`artifactory_item_properties`](#artifactory_item_properties)
    - [Properties must exist](#properties-must-exist)
    - [Properties are stored as string](#properties-are-stored-as-string)
  - [`keypair`](#keypair)
    - [Fails with error `cannot find pair_name in tfstate`](#fails-with-error-cannot-find-pair_name-in-tfstate)
  - [`tokens`](#tokens)

## Resource Import Not Implemented

Terraform resource doesn't allow import and it's not possible to set external name, because there is no id set in terraform state, e.g.:

```bash
# Example
$ terraform import artifactory_artifact.my-local-artifact artifact
artifactory_artifact.my-local-artifact: Importing from ID "artifact"...
╷
│ Error: Resource Import Not Implemented
│
│ This resource does not support import. Please contact the provider developer for additional information.
╵
```

## Nested Schema

> [!NOTE]
> The issue [crossplane/upjet#372](https://github.com/crossplane/upjet/issues/372) was solved and nested schema is not a problem anymore.

The Terraform resource contains a [Nested Schema](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes#nested-attribute-types) and Upjet is not able to generate a provider, it fails with an error.

Example for [`search_criteria` attribute](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs/resources/archive_policy#nested-schema-for-search_criteria) of resource `artifactory_archive_policy`:

```bash
# Example
$ make generate
...
panic: cannot generate crd for resource artifactory_archive_policy: cannot build types for ArchivePolicy: cannot build the Types for resource "artifactory_archive_policy": cannot infer type from schema of field search_criteria: invalid schema type TypeInvalid
```

There is [opened issue](https://github.com/crossplane/upjet/issues/372) on crossplane/upjet, no workaround here for now.

## `artifactory_item_properties`

### Properties must exist

At least one property of repo/item MUST exists, otherwise this resource fails with an error:

```log
2025-04-05T21:00:10+02:00	DEBUG	events	cannot run refresh: refresh failed: Unable to Refresh Resource: An unexpected error occurred while attempting to refresh resource state. Please retry the operation or report this issue to the provider developers.

Error: {
  "errors" : [ {
    "status" : 404,
    "message" : "No properties could be found."
  } ]
}
```

### Properties are stored as string

Terraform requires to set [property values as a set of strings](https://registry.terraform.io/providers/jfrog/artifactory/latest/docs/resources/item_properties#properties-1).

However, Artifactory converts this set of strings to a single string separated by commas. The next reconciliation use Terraform refresh which cause the state is changed to this single string and Terraform apply is triggered again which will **end up with neverending loop**.

```yaml
apiVersion: artifactory.jfrog.crossplane.io/v1alpha1
kind: ItemProperties
metadata:
  name: my-repo-properties
spec:
  forProvider:
    repoKey: generic-crossplane-local
    properties:
      key1: ["value1"]
      key2: ["value2", "value3"]    # <--- This configuration will cause neverending reconciliation loop
  providerConfigRef:
    name: default
```

You can find the whole reconciliation process investigation of provider log in [docs](./docs/artifact/item_properties/reconsiliation-process-investigation.log).

> [!NOTE]
> Don't use more than one value for the key even if it allows to use a set of strings.

## `keypair`

### Fails with error `cannot find pair_name in tfstate`

Provider prints error message

```log
2025-04-09T15:45:04+02:00	DEBUG	events	cannot set critical annotations: cannot get external name: cannot find pair_name in tfstate	{"type": "Warning", "object": {"kind":"Keypair","name":"my-crossplane-keypair","uid":"203fa67f-74a6-41c1-9de8-49b3de5573f7","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"236597"}, "reason": "CannotObserveExternalResource"}
```

After the further investigation I found out that the provider is not able to refresh Terraform state. The reason is that provider does the Terraform refresh first and expects map of [`KeyPairAPIModel` values](https://github.com/jfrog/terraform-provider-artifactory/blob/v12.9.1/pkg/artifactory/resource/security/resource_artifactory_keypair.go#L60-L67).

```go
type KeyPairAPIModel struct {
  PairName   string `json:"pairName"`
  PairType   string `json:"pairType"`
  Alias      string `json:"alias"`
  PrivateKey string `json:"privateKey"`
  Passphrase string `json:"passphrase"`
  PublicKey  string `json:"publicKey"`
}
```

But Artifactory returns the empty list (in case there is no key) or list of keys without `PrivateKey` and `Passphrase` which cause an error

```log
Error: json: cannot unmarshal array into Go value of type security.KeyPairAPIModel
```

This is also [mentioned](https://registry.terraform.io/providers/jfrog/artifactory/latest/docs/resources/keypair#argument-reference) in Terraform Artifactory provider description for `artifactory_keypair` resource:

> Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.

So, it seems it's possible to CREATE or UPDATE resource, but **it's not possible to refresh Terraform state**.

More details from the investigation are [here](./docs/security/keypair/cannot-find-pair-name-in-tfstate.md).

## `tokens`

You must remember to include the `writeConnectionSecretToRef` object in your `Token` definitions in order for the token secrets to be stored in a Kubernetes secret. See [Managed Resources](https://docs.crossplane.io/v2.0/managed-resources/managed-resources/#writeconnectionsecrettoref) for details.

Token would not be saved by Artifactory if `expires_in` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://jfrog.com/help/r/jfrog-platform-administration-documentation/use-the-revocable-and-persistency-thresholds) for details.

You may only create and delete tokens managed using this provider. Existing tokens cannot be imported nor updated. Attempting to change any parameter of the `forProvider` object will fail and cause the controller to retry repeatedly. Tokens that are revoked or refreshed on Artifactory do not cause this provider to attempt to reconcile its state.
