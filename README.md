<!-- markdownlint-disable  no-hard-tabs -->
# Provider Artifactory

This document provides an overview and guidance for using the `provider-artifactory`, a Crossplane provider for managing Artifactory resources.

`provider-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

The repo was created from [crossplane/upjet-provider-template@7311f9f](https://github.com/crossplane/upjet-provider-template/tree/7311f9f9baa87f4431702ba209dffbc6067ce74b) template.

- [Provider Artifactory](#provider-artifactory)
  - [Getting Started](#getting-started)
  - [Naming convention decision](#naming-convention-decision)
    - [Options](#options)
  - [Supported resources](#supported-resources)
    - [Artifact](#artifact)
    - [Configuration](#configuration)
    - [Federated Repositories](#federated-repositories)
    - [Lifecycle](#lifecycle)
    - [Local Repositories](#local-repositories)
    - [Remote Repositories](#remote-repositories)
    - [Replication](#replication)
    - [Security](#security)
    - [User](#user)
    - [Virtual Repositories](#virtual-repositories)
    - [Webhook](#webhook)
    - [Known issues](#known-issues)
      - [`artifactory_item_properties`](#artifactory_item_properties)
        - [Properties must exist](#properties-must-exist)
        - [Properties are stored as string](#properties-are-stored-as-string)
      - [`keypair`](#keypair)
        - [Fails with error `cannot find pair_name in tfstate`](#fails-with-error-cannot-find-pair_name-in-tfstate)
  - [Build provider from scratch](#build-provider-from-scratch)
  - [Developing](#developing)
  - [Report a Bug](#report-a-bug)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-artifactory):

```bash
up ctp provider install hmlkao/provider-artifactory:v0.1.0
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-artifactory
spec:
  package: hmlkao/provider-artifactory:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/hmlkao/provider-artifactory).

## Naming convention decision

There are more Terraform providers developed by JFrog, e.g.:

- [`jfrog/artifactory`](https://registry.terraform.io/providers/jfrog/artifactory)
- [`jfrog/platform`](https://registry.terraform.io/providers/jfrog/platform)
- [`jfrog/project`](https://registry.terraform.io/providers/jfrog/project)
- etc.

So, I decided to use `jfrog.crossplane.io` base group for this Crossplane provider.

### Options

1. (*current naming convention*) `jfrog.crossplane.io` as a base group and set `ShortGroup` as `artifactory` to all resources which will produce `artifactory.jfrog.crossplane.io` and K8s kind is just `LocalOCIRepository`

    - :heavy_minus_sign: All resources must have specified `ShortGroup` as `artifactory`

2. `artifactory.jfrog.crossplane.io` as base group and resource is just `LocalOCIRepository` (no need to set up `ShortGroup`)

    - :heavy_plus_sign: Get rid of `ShortGroup` config and use `artifactory.jfrog.crossplane.io` for initial config instead
    - Allows to use `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc.

3. `jfrog.crossplane.io` as a base group and resource prefix to be `artifactory`, like `ArtifactoryLocalOCIRepository`

    - :heavy_minus_sign: All resources contains `Artifactory` which is not needed

Not sure, which one is the best.

## Supported resources

List of all resources of [Terraform provider version 12.9.1](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs).

### Artifact

- *Resource Import Not Implemented* - Terraform resource doesn't allow import and it's not possible to set external name, because there is no id set in terraform state, e.g.:

    ```bash
    $ terraform import artifactory_artifact.my-local-artifact artifact
    artifactory_artifact.my-local-artifact: Importing from ID "artifact"...
    ╷
    │ Error: Resource Import Not Implemented
    │
    │ This resource does not support import. Please contact the provider developer for additional information.
    ╵
    ```

| Resource                      | Supported                                                         | Kind             |
|-------------------------------|-------------------------------------------------------------------|------------------|
| `artifactory_artifact`        | :x: (Resource Import Not Implemented)                             |                  |
| `artifactory_item_properties` | :heavy_check_mark: ([known issues](#artifactory_item_properties)) | `ItemProperties` |

### Configuration

| Resource                             | Supported          | Kind                           |
|--------------------------------------|--------------------|--------------------------------|
| `artifactory_archive_policy`         | :x:                |                                |
| `artifactory_backup`                 | :x:                |                                |
| `artifactory_general_security`       | :x:                |                                |
| `artifactory_ldap_group_setting`     | :x:                |                                |
| `artifactory_ldap_group_setting_v2`  | :x:                |                                |
| `artifactory_ldap_setting`           | :x:                |                                |
| `artifactory_ldap_setting_v2`        | :x:                |                                |
| `artifactory_mail_server`            | :x:                |                                |
| `artifactory_oauth_settings`         | :x:                |                                |
| `artifactory_package_cleanup_policy` | :x:                |                                |
| `artifactory_property_set`           | :x:                |                                |
| `artifactory_proxy`                  | :x:                |                                |
| `artifactory_repository_layout`      | :x:                |                                |
| `artifactory_vault_configuration`    | :x:                |                                |

### Federated Repositories

| Resource                                              | Supported          | Kind                           |
|-------------------------------------------------------|--------------------|--------------------------------|
| `artifactory_federated_alpine_repository`             | :x:                |                                |
| `artifactory_federated_ansible_repository`            | :x:                |                                |
| `artifactory_federated_bower_repository`              | :x:                |                                |
| `artifactory_federated_cargo_repository`              | :x:                |                                |
| `artifactory_federated_chef_repository`               | :x:                |                                |
| `artifactory_federated_cocoapods_repository`          | :x:                |                                |
| `artifactory_federated_composer_repository`           | :x:                |                                |
| `artifactory_federated_conan_repository`              | :x:                |                                |
| `artifactory_federated_conda_repository`              | :x:                |                                |
| `artifactory_federated_cran_repository`               | :x:                |                                |
| `artifactory_federated_debian_repository`             | :x:                |                                |
| `artifactory_federated_docker_repository`             | :x:                |                                |
| `artifactory_federated_docker_v1_repository`          | :x:                |                                |
| `artifactory_federated_docker_v2_repository`          | :x:                |                                |
| `artifactory_federated_gems_repository`               | :x:                |                                |
| `artifactory_federated_generic_repository`            | :x:                |                                |
| `artifactory_federated_gitlfs_repository`             | :x:                |                                |
| `artifactory_federated_go_repository`                 | :x:                |                                |
| `artifactory_federated_gradle_repository`             | :x:                |                                |
| `artifactory_federated_helm_repository`               | :x:                |                                |
| `artifactory_federated_helmoci_repository`            | :x:                |                                |
| `artifactory_federated_huggingfaceml_repository`      | :x:                |                                |
| `artifactory_federated_ivy_repository`                | :x:                |                                |
| `artifactory_federated_maven_repository`              | :x:                |                                |
| `artifactory_federated_npm_repository`                | :x:                |                                |
| `artifactory_federated_nuget_repository`              | :x:                |                                |
| `artifactory_federated_oci_repository`                | :x:                |                                |
| `artifactory_federated_opkg_repository`               | :x:                |                                |
| `artifactory_federated_puppet_repository`             | :x:                |                                |
| `artifactory_federated_pypi_repository`               | :x:                |                                |
| `artifactory_federated_rpm_repository`                | :x:                |                                |
| `artifactory_federated_sbt_repository`                | :x:                |                                |
| `artifactory_federated_swift_repository`              | :x:                |                                |
| `artifactory_federated_terraform_module_repository`   | :x:                |                                |
| `artifactory_federated_terraform_provider_repository` | :x:                |                                |
| `artifactory_federated_vagrant_repository`            | :x:                |                                |

### Lifecycle

| Resource                                  | Supported          | Kind                           |
|-------------------------------------------|--------------------|--------------------------------|
| `artifactory_release_bundle_v2`           | :x:                |                                |
| `artifactory_release_bundle_v2_promotion` | :x:                |                                |

### Local Repositories

| Resource                                          | Supported          | Kind                               |
|---------------------------------------------------|--------------------|------------------------------------|
| `artifactory_local_alpine_repository`             | :x:                |                                    |
| `artifactory_local_ansible_repository`            | :heavy_check_mark: | `LocalAnsibleRepository`           |
| `artifactory_local_bower_repository`              | :x:                |                                    |
| `artifactory_local_cargo_repository`              | :x:                |                                    |
| `artifactory_local_chef_repository`               | :x:                |                                    |
| `artifactory_local_cocoapods_repository`          | :x:                |                                    |
| `artifactory_local_composer_repository`           | :x:                |                                    |
| `artifactory_local_conan_repository`              | :x:                |                                    |
| `artifactory_local_conda_repository`              | :x:                |                                    |
| `artifactory_local_cran_repository`               | :x:                |                                    |
| `artifactory_local_debian_repository`             | :x:                |                                    |
| `artifactory_local_docker_v1_repository`          | :x:                |                                    |
| `artifactory_local_docker_v2_repository`          | :heavy_check_mark: | `LocalDockerV2Repository`          |
| `artifactory_local_gems_repository`               | :x:                |                                    |
| `artifactory_local_generic_repository`            | :heavy_check_mark: | `LocalGenericRepository`           |
| `artifactory_local_gitlfs_repository`             | :x:                |                                    |
| `artifactory_local_go_repository`                 | :x:                |                                    |
| `artifactory_local_gradle_repository`             | :x:                |                                    |
| `artifactory_local_helm_repository`               | :x:                |                                    |
| `artifactory_local_helmoci_repository`            | :x:                |                                    |
| `artifactory_local_huggingfaceml_repository`      | :x:                |                                    |
| `artifactory_local_ivy_repository`                | :x:                |                                    |
| `artifactory_local_machinelearning_repository`    | :x:                |                                    |
| `artifactory_local_maven_repository`              | :x:                |                                    |
| `artifactory_local_npm_repository`                | :x:                |                                    |
| `artifactory_local_nuget_repository`              | :x:                |                                    |
| `artifactory_local_oci_repository`                | :heavy_check_mark: | `LocalOCIRepository`               |
| `artifactory_local_opkg_repository`               | :x:                |                                    |
| `artifactory_local_pub_repository`                | :x:                |                                    |
| `artifactory_local_puppet_repository`             | :x:                |                                    |
| `artifactory_local_pypi_repository`               | :x:                |                                    |
| `artifactory_local_rpm_repository`                | :x:                |                                    |
| `artifactory_local_sbt_repository`                | :x:                |                                    |
| `artifactory_local_swift_repository`              | :x:                |                                    |
| `artifactory_local_terraform_module_repository`   | :x:                |                                    |
| `artifactory_local_terraform_provider_repository` | :heavy_check_mark: | `LocalTerraformProviderRepository` |
| `artifactory_local_terraformbackend_repository`   | :x:                |                                    |
| `artifactory_local_vagrant_repository`            | :x:                |                                    |

### Remote Repositories

| Resource                                           | Supported          | Kind                           |
|----------------------------------------------------|--------------------|--------------------------------|
| `artifactory_remote_alpine_repository`             | :x:                |                                |
| `artifactory_remote_ansible_repository`            | :x:                |                                |
| `artifactory_remote_bower_repository`              | :x:                |                                |
| `artifactory_remote_cargo_repository`              | :x:                |                                |
| `artifactory_remote_chef_repository`               | :x:                |                                |
| `artifactory_remote_cocoapods_repository`          | :x:                |                                |
| `artifactory_remote_composer_repository`           | :x:                |                                |
| `artifactory_remote_conan_repository`              | :x:                |                                |
| `artifactory_remote_conda_repository`              | :x:                |                                |
| `artifactory_remote_cran_repository`               | :x:                |                                |
| `artifactory_remote_debian_repository`             | :x:                |                                |
| `artifactory_remote_docker_repository`             | :x:                |                                |
| `artifactory_remote_gems_repository`               | :x:                |                                |
| `artifactory_remote_generic_repository`            | :x:                |                                |
| `artifactory_remote_gitlfs_repository`             | :x:                |                                |
| `artifactory_remote_go_repository`                 | :x:                |                                |
| `artifactory_remote_gradle_repository`             | :x:                |                                |
| `artifactory_remote_helm_repository`               | :x:                |                                |
| `artifactory_remote_helmoci_repository`            | :x:                |                                |
| `artifactory_remote_huggingfaceml_repository`      | :x:                |                                |
| `artifactory_remote_ivy_repository`                | :x:                |                                |
| `artifactory_remote_maven_repository`              | :x:                |                                |
| `artifactory_remote_npm_repository`                | :x:                |                                |
| `artifactory_remote_nuget_repository`              | :x:                |                                |
| `artifactory_remote_oci_repository`                | :x:                |                                |
| `artifactory_remote_opkg_repository`               | :x:                |                                |
| `artifactory_remote_p2_repository`                 | :x:                |                                |
| `artifactory_remote_pub_repository`                | :x:                |                                |
| `artifactory_remote_puppet_repository`             | :x:                |                                |
| `artifactory_remote_pypi_repository`               | :x:                |                                |
| `artifactory_remote_rpm_repository`                | :x:                |                                |
| `artifactory_remote_sbt_repository`                | :x:                |                                |
| `artifactory_remote_swift_repository`              | :x:                |                                |
| `artifactory_remote_terraform_repository`          | :x:                |                                |
| `artifactory_remote_terraform_provider_repository` | :x:                |                                |
| `artifactory_remote_vcs_repository`                | :x:                |                                |

### Replication

| Resource                                          | Supported          | Kind                           |
|---------------------------------------------------|--------------------|--------------------------------|
| `artifactory_local_repository_multi_replication`  | :x:                |                                |
| `artifactory_local_repository_single_replication` | :x:                |                                |
| `artifactory_remote_repository_replication`       | :x:                |                                |

### Security

| Resource                                 | Supported                                     | Kind                           |
|------------------------------------------|-----------------------------------------------|--------------------------------|
| `artifactory_certificate`                | :x:                                           |                                |
| `artifactory_distribution_public_key`    | :x:                                           |                                |
| `artifactory_global_environment`         | :x:                                           |                                |
| `artifactory_keypair`                    | :heavy_check_mark: ([known issues](#keypair)) | `Keypair`                      |
| `artifactory_password_expiration_policy` | :x:                                           |                                |
| `artifactory_scoped_token`               | :x:                                           |                                |
| `artifactory_user_lock_policy`           | :x:                                           |                                |

### User

| Resource                     | Supported          | Kind                           |
|------------------------------|--------------------|--------------------------------|
| `artifactory_anonymous_user` | :x:                |                                |
| `artifactory_managed_user`   | :x:                |                                |
| `artifactory_unmanaged_user` | :x:                |                                |
| `artifactory_user`           | :x:                |                                |

### Virtual Repositories

| Resource                                   | Supported          | Kind                           |
|--------------------------------------------|--------------------|--------------------------------|
| `artifactory_virtual_alpine_repository`    | :x:                |                                |
| `artifactory_virtual_ansible_repository`   | :x:                |                                |
| `artifactory_virtual_bower_repository`     | :x:                |                                |
| `artifactory_virtual_chef_repository`      | :x:                |                                |
| `artifactory_virtual_cocoapods_repository` | :x:                |                                |
| `artifactory_virtual_composer_repository`  | :x:                |                                |
| `artifactory_virtual_conan_repository`     | :x:                |                                |
| `artifactory_virtual_conda_repository`     | :x:                |                                |
| `artifactory_virtual_cran_repository`      | :x:                |                                |
| `artifactory_virtual_debian_repository`    | :x:                |                                |
| `artifactory_virtual_docker_repository`    | :x:                |                                |
| `artifactory_virtual_gems_repository`      | :x:                |                                |
| `artifactory_virtual_generic_repository`   | :x:                |                                |
| `artifactory_virtual_gitlfs_repository`    | :x:                |                                |
| `artifactory_virtual_go_repository`        | :x:                |                                |
| `artifactory_virtual_gradle_repository`    | :x:                |                                |
| `artifactory_virtual_helm_repository`      | :x:                |                                |
| `artifactory_virtual_helmoci_repository`   | :x:                |                                |
| `artifactory_virtual_ivy_repository`       | :x:                |                                |
| `artifactory_virtual_maven_repository`     | :x:                |                                |
| `artifactory_virtual_npm_repository`       | :x:                |                                |
| `artifactory_virtual_nuget_repository`     | :x:                |                                |
| `artifactory_virtual_oci_repository`       | :x:                |                                |
| `artifactory_virtual_p2_repository`        | :x:                |                                |
| `artifactory_virtual_pub_repository`       | :x:                |                                |
| `artifactory_virtual_puppet_repository`    | :x:                |                                |
| `artifactory_virtual_pypi_repository`      | :x:                |                                |
| `artifactory_virtual_rpm_repository`       | :x:                |                                |
| `artifactory_virtual_sbt_repository`       | :x:                |                                |
| `artifactory_virtual_swift_repository`     | :x:                |                                |
| `artifactory_virtual_terraform_repository` | :x:                |                                |

### Webhook

| Resource                                                 | Supported          | Kind                           |
|----------------------------------------------------------|--------------------|--------------------------------|
| `artifactory_artifact_custom_webhook`                    | :x:                |                                |
| `artifactory_artifact_lifecycle_custom_webhook`          | :x:                |                                |
| `artifactory_artifact_lifecycle_webhook`                 | :x:                |                                |
| `artifactory_artifact_property_custom_webhook`           | :x:                |                                |
| `artifactory_artifact_property_webhook`                  | :x:                |                                |
| `artifactory_artifact_webhook`                           | :x:                |                                |
| `artifactory_artifactory_release_bundle_custom_webhook`  | :x:                |                                |
| `artifactory_artifactory_release_bundle_webhook`         | :x:                |                                |
| `artifactory_build_custom_webhook`                       | :x:                |                                |
| `artifactory_build_webhook`                              | :x:                |                                |
| `artifactory_destination_custom_webhook`                 | :x:                |                                |
| `artifactory_destination_webhook`                        | :x:                |                                |
| `artifactory_distribution_custom_webhook`                | :x:                |                                |
| `artifactory_distribution_webhook`                       | :x:                |                                |
| `artifactory_local_docker_v2_repository`                 | :x:                |                                |
| `artifactory_docker_webhook`                             | :x:                |                                |
| `artifactory_release_bundle_v2_custom_webhook`           | :x:                |                                |
| `artifactory_release_bundle_v2_promotion_custom_webhook` | :x:                |                                |
| `artifactory_release_bundle_v2_promotion_webhook`        | :x:                |                                |
| `artifactory_release_bundle_v2_webhook`                  | :x:                |                                |
| `artifactory_user_custom_webhook`                        | :x:                |                                |
| `artifactory_user_webhook`                               | :x:                |                                |

### Known issues

#### `artifactory_item_properties`

##### Properties must exist

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

##### Properties are stored as string

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

You can find the whole reconciliation process investigation of provider log [here](./docs/artifact/item_properties/reconsiliation-process-investigation.log).

> [!NOTE] Solution
> Don't use more than one value for the key even if it allows to use a set of strings.

#### `keypair`

##### Fails with error `cannot find pair_name in tfstate`

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

```
Error: json: cannot unmarshal array into Go value of type security.KeyPairAPIModel
```

This is also mentioned in Terraform Artifactory provider description for `artifactory_keypair` resource

> Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.

So, it seems it's possible to CREATE or UPDATE resource, but **it's not possible to refresh Terraform state**.

More details from the investigation are [here](./docs/security/keypair/cannot-find-pair-name-in-tfstate.md).

## Build provider from scratch

Check [`BUILD_FROM_SCRATCH.md`](./BUILD_FROM_SCRATCH.md) for notes how was this provider built using [crossplane/upjet tool](https://github.com/crossplane/upjet) step-by-step.

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/hmlkao/provider-artifactory/issues).
