<!-- markdownlint-disable no-hard-tabs no-inline-html -->
# External name understanding

Crossplane was primarily developed to manage infrastructure on Cloud platforms (AWS, GCP, Azure).

Terraform state of resources of these platforms usualy set `id` parameter, e.g.

Source:

- [crossplane/upjet docs](https://github.com/crossplane/upjet/blob/v1.5.0/docs/configuring-a-resource.md)
- [Crossplane Docs](https://docs.crossplane.io/master/concepts/managed-resources/#naming-external-resources)
- [Terraform glossary](https://developer.hashicorp.com/terraform/docs/glossary#id)

There are mentioned 3 most common cases:

- [Case 1: Name as External Name and Terraform ID](https://github.com/crossplane/upjet/blob/v1.5.0/docs/configuring-a-resource.md#case-1-name-as-external-name-and-terraform-id)

  Two conditions MUST be met:

  - Terraform resource uses the `name` argument to identify the resources
  - Terraform resource can be imported with `name`, i.e. `id`=`name`

  ```go
  // config.go
  r.ExternalName = config.NameAsIdentifier

  // externalname.go
  // NameAsIdentifier uses "name" field in the arguments as the identifier of
	// the resource.
	NameAsIdentifier = ExternalName{
		SetIdentifierArgumentFn: func(base map[string]any, name string) {
			base["name"] = name
		},
		GetExternalNameFn: IDAsExternalName,    // <--- expects 'id' parameter in Terraform state
		GetIDFn:           ExternalNameAsID,    // <--- external name to be used as ID in TF State file
		OmittedFields: []string{
			"name",
			"name_prefix",
		},
  }
  ```

  So, it cannot be used for Artifactory provider, because, e.g. `terraform import` for `artifactory_local_oci_repository` use `key` value to identify the repository, not `name`.

- [Case 2: Identifier from Provider](https://github.com/crossplane/upjet/blob/v1.5.0/docs/configuring-a-resource.md#case-2-identifier-from-provider)

    ```go
    // config.go
    r.ExternalName = config.IdentifierFromProvider

    // externalname.go
    // IdentifierFromProvider is used in resources whose identifier is assigned by
    // the remote client, such as AWS VPC where it gets an identifier like
    // vpc-2213das instead of letting user choose a name.
    IdentifierFromProvider = ExternalName{
      SetIdentifierArgumentFn: NopSetIdentifierArgument,
      GetExternalNameFn:       IDAsExternalName,    // <--- expects 'id' parameter in Terraform state
      GetIDFn:                 ExternalNameAsID,    // <--- external name to be used as ID in TF State file
      DisableNameInitializer:  true,                // <--- it needs to be disabled for resources whose external identifier is randomly assigned by the provider, e.g. AWS VPC ID
    }
    ```

    So, it cannot be used for Artifactory provider, because the provider doesn't generate custom ID for Artifactory resources.

- [Case 3: Terraform ID as a Formatted String](https://github.com/crossplane/upjet/blob/v1.5.0/docs/configuring-a-resource.md#case-3-terraform-id-as-a-formatted-string)

    There is great example how to use custom functions `GetExternalNameFn` and `GetIDFn` and how is `id` put together from various values for Azure resources in the docs, check it there.

- Case 4: Use parameter as ID (not mentioned in docs)

    Use the given field name in the arguments as the identifier of the resource.

    ```go
    // config.go
    r.ExternalName = config.ParameterAsIdentifier("key"),

    // ParameterAsIdentifier uses the given field name in the arguments as the
    // identifier of the resource.
    func ParameterAsIdentifier(param string) ExternalName {
      e := NameAsIdentifier
      e.SetIdentifierArgumentFn = func(base map[string]any, name string) {
        base[param] = name
      }
      e.OmittedFields = []string{
        param,
        param + "_prefix",
      }
      e.IdentifierFields = []string{param}
      return e
    }
    ```

    This case exactly fits to our needs when is `key` used as identifier, e.g. for `artifactory_local_oci_repository` resource.

    However, we still need to set custom `GetExternalNameFn`, otherwise it still expects `id` field to get external name.

    ```go
    r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
    ```

These functions are used in more cases:

```go
// resource.go
// NopSetIdentifierArgument does nothing. It's useful for cases where the external
// name is calculated by provider and doesn't have any effect on spec fields.
var NopSetIdentifierArgument SetIdentifierArgumentsFn = func(_ map[string]any, _ string) {}

// IDAsExternalName returns the TF State ID as external name.
var IDAsExternalName GetExternalNameFn = func(tfstate map[string]any) (string, error) {
	if id, ok := tfstate["id"].(string); ok && id != "" {
		return id, nil
	}
	return "", errors.New("cannot find id in tfstate")
}

// ExternalNameAsID returns the name to be used as ID in TF State file.
var ExternalNameAsID GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
	return externalName, nil
}
```

<details>
  <summary>Description of all `ExternalName` parameters</summary>

  [GitHub source](https://github.com/crossplane/upjet/blob/v1.5.0/pkg/config/resource.go#L119-L166) for `crossplane/upjet` v1.5.0

  ```go
  // ExternalName contains all information that is necessary for naming operations,
  // such as removal of those fields from spec schema and calling Configure function
  // to fill attributes with information given in external name.
  type ExternalName struct {
    // SetIdentifierArgumentFn sets the name of the resource in Terraform argument
    // map. In many cases, there is a field called "name" in the HCL schema, however,
    // there are cases like RDS DB Cluster where the name field in HCL is called
    // "cluster_identifier". This function is the place that you can take external
    // name and assign it to that specific key for that resource type.
    SetIdentifierArgumentFn SetIdentifierArgumentsFn

    // GetExternalNameFn returns the external name extracted from TF State. In most cases,
    // "id" field contains all the information you need. You'll need to extract
    // the format that is decided for external name annotation to use.
    // For example the following is an Azure resource ID:
    // /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1
    // The function should return "mygroup1" so that it can be used to set external
    // name if it was not set already.
    GetExternalNameFn GetExternalNameFn

    // GetIDFn returns the string that will be used as "id" key in TF state. In
    // many cases, external name format is the same as "id" but when it is not
    // we may need information from other places to construct it. For example,
    // the following is an Azure resource ID:
    // /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1
    // The function here should use information from supplied arguments to
    // construct this ID, i.e. "mygroup1" from external name, subscription ID
    // from terraformProviderConfig, and others from parameters map if needed.
    GetIDFn GetIDFn

    // OmittedFields are the ones you'd like to be removed from the schema since
    // they are specified via external name. For example, if you set
    // "cluster_identifier" in SetIdentifierArgumentFn, then you need to omit
    // that field.
    // You can omit only the top level fields.
    // No field is omitted by default.
    OmittedFields []string

    // DisableNameInitializer allows you to specify whether the name initializer
    // that sets external name to metadata.name if none specified should be disabled.
    // It needs to be disabled for resources whose external identifier is randomly
    // assigned by the provider, like AWS VPC where it gets vpc-21kn123 identifier
    // and not let you name it.
    DisableNameInitializer bool

    // IdentifierFields are the fields that are used to construct external
    // resource identifier. We need to know these fields no matter what the
    // management policy is including the Observe Only, different from other
    // (required) fields.
    IdentifierFields []string
  }
  ```

</details>

## Puzzle pieces

`GenericRepository` resource as is defined by user:

```yaml
apiVersion: artifactory.jfrog.crossplane.io/v1alpha1
kind: LocalGenericRepository
metadata:
  name: generic-crossplane-local           # <---------------- Name of the Kubernetes resource
spec:
  forProvider: {}
    # We don't need to set up the 'key' parameter, even if it is mandatory,
    #   because it is set up as a name of the resource
    # Otherwise, there would be two different parameters,
    #   one for resource name and the second for repository name which is confusing
    #key: generic-crossplane-local         # <---------------- Name of the repository in Artifactory (in this implementation is name of the resource used for this purpose)
  providerConfigRef:
    name: default
```

Generated Terraform code used to apply changes to Artifactory

```json
# main.tf.json
{
  ...
  "resource": {
    "artifactory_local_generic_repository": {
      "generic-crossplane-local": {             # <--------/------- Both resource name and key are the same
        "includes_pattern": "**/*",             #         /
        "key": "generic-crossplane-local",      # <------/
        "lifecycle": {
          "prevent_destroy": true
        },
        "repo_layout_ref": "simple-default"
      }
    }
  }
}
```

```json
# terraform.tfstate
{
  "version": 4,
  "terraform_version": "1.5.7",
  "serial": 2,
  "lineage": "eecbf778-6c54-4044-a8f8-ec7644428062",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "artifactory_local_generic_repository",
      "name": "generic-crossplane-local",                 # <----------------------/-------- Both resource name and key are the same
      "provider": "provider[\"registry.terraform.io/jfrog/artifactory\"]", #      /
      "instances": [                                      #                      /
        {                                                 #                     /
          "schema_version": 1,                            #                    /
          "attributes": {                                 #                   /
            "archive_browsing_enabled": false,            #                  /
            "blacked_out": false,                         #                 /
            "cdn_redirect": false,                        #                /
            "description": "",                            #               /
            "download_direct": false,                     #              /
            "excludes_pattern": "",                       #             /
            "includes_pattern": "**/*",                   #            /
            "key": "generic-crossplane-local",            # <---------/
            "notes": "",
            "priority_resolution": false,
            "project_environments": [],
            "project_key": "",
            "property_sets": null,
            "repo_layout_ref": "simple-default",
            "xray_index": false
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}
```

Real `GenericRepository` resource as is stored in Kubernetes

```yaml
apiVersion: artifactory.jfrog.crossplane.io/v1alpha1
kind: LocalGenericRepository
metadata:
  annotations:
    crossplane.io/external-create-pending: "2025-04-07T17:04:24+02:00"
    crossplane.io/external-create-succeeded: "2025-04-07T17:04:24+02:00"
    crossplane.io/external-name: generic-crossplane-local                   # <------------ Annotation value taken from Terraform state by GetExternalNameFn() func in Go lang
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","kind":"LocalGenericRepository","metadata":{"annotations":{},"name":"generic-crossplane-local"},"spec":{"forProvider":{},"providerConfigRef":{"name":"default"}}}
  creationTimestamp: "2025-04-07T15:04:11Z"
  finalizers:
  - finalizer.managedresource.crossplane.io
  generation: 3
  name: generic-crossplane-local
  resourceVersion: "167900"
  uid: eecbf778-6c54-4044-a8f8-ec7644428062
spec:
  deletionPolicy: Delete
  forProvider:
    includesPattern: '**/*'
    repoLayoutRef: simple-default
  initProvider: {}
  managementPolicies:
  - '*'
  providerConfigRef:
    name: default
status:
  atProvider:
    archiveBrowsingEnabled: false
    blackedOut: false
    cdnRedirect: false
    description: ""
    downloadDirect: false
    excludesPattern: ""
    includesPattern: '**/*'
    notes: ""
    priorityResolution: false
    projectKey: ""
    repoLayoutRef: simple-default
    xrayIndex: false
  conditions:
  - lastTransitionTime: "2025-04-07T15:04:28Z"
    reason: Available
    status: "True"
    type: Ready
  - lastTransitionTime: "2025-04-08T15:39:00Z"
    reason: ReconcileSuccess
    status: "True"
    type: Synced
  - lastTransitionTime: "2025-04-07T15:04:26Z"
    reason: Finished
    status: "True"
    type: AsyncOperation
  - lastTransitionTime: "2025-04-07T15:04:26Z"
    reason: Success
    status: "True"
    type: LastAsyncOperation
```

## Processing of added resource

<details>
  <summary>Full log</summary>

  ```log
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Create event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Update event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "providerconfig/providerconfig.artifactory.jfrog.crossplane.io", "request": {"name":"default"}}
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["init", "-input=false"]}
  2025-04-10T14:59:06+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "providerconfig/providerconfig.artifactory.jfrog.crossplane.io", "request": {"name":"default"}}
  2025-04-10T14:59:12+02:00	DEBUG	provider-jfrog-artifactory	init ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "\n\u001b[0m\u001b[1mInitializing the backend...\u001b[0m\n\n\u001b[0m\u001b[1mInitializing provider plugins...\u001b[0m\n- Finding jfrog/artifactory versions matching \"12.9.1\"...\n- Installing jfrog/artifactory v12.9.1...\n- Installed jfrog/artifactory v12.9.1 (signed by a HashiCorp partner, key ID \u001b[0m\u001b[1m2FA4D2A520237FA7\u001b[0m\u001b[0m)\n\nPartner and community providers are signed by their developers.\nIf you'd like to know more about provider signing, you can read about it here:\nhttps://www.terraform.io/docs/cli/plugins/signing.html\n\nTerraform has created a lock file \u001b[1m.terraform.lock.hcl\u001b[0m to record the provider\nselections it made above. Include this file in your version control repository\nso that Terraform can guarantee to make the same selections by default when\nyou run \"terraform init\" in the future.\u001b[0m\n\n\u001b[0m\u001b[1m\u001b[32mTerraform has been successfully initialized!\u001b[0m\u001b[32m\u001b[0m\n\u001b[0m\u001b[32m\nYou may now begin working with Terraform. Try running \"terraform plan\" to see\nany changes that are required for your infrastructure. All Terraform commands\nshould now work.\n\nIf you ever set or change modules or backend configuration for Terraform,\nrerun this command to reinitialize your working directory. If you forget, other\ncommands will detect it and remind you to do so if necessary.\u001b[0m\n"}
  2025-04-10T14:59:12+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:13+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:12.126588+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.766332+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.907339+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Drift detected (delete)\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.917912+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"delete\"},\"type\":\"resource_drift\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.917976+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.936271+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:13.936308+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:59:13+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["apply", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:14+02:00	DEBUG	provider-jfrog-artifactory	Successfully requested creation of external resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381573", "external-name": "", "external-name": "ansible-crossplane-repository"}
  2025-04-10T14:59:14+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Update event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:59:14+02:00	DEBUG	events	Successfully requested creation of external resource	{"type": "Normal", "object": {"kind":"LocalAnsibleRepository","name":"ansible-crossplane-repository","uid":"c1f2abb9-b5c8-4187-a2e7-290e7a604ad3","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"381596"}, "reason": "CreatedExternalResource"}
  2025-04-10T14:59:14+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:14+02:00	DEBUG	provider-jfrog-artifactory	External resource is up to date	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381597", "external-name": "ansible-crossplane-repository", "requeue-after": "2025-04-10T15:09:14+02:00"}
  2025-04-10T14:59:15+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:15+02:00	DEBUG	provider-jfrog-artifactory	External resource is up to date	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381599", "external-name": "ansible-crossplane-repository", "requeue-after": "2025-04-10T15:09:15+02:00"}
  2025-04-10T14:59:17+02:00	DEBUG	provider-jfrog-artifactory	apply async ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:14.038095+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Plan to create\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:15.682787+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"create\"},\"type\":\"planned_change\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 1 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:15.682887+02:00\",\"changes\":{\"add\":1,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Creating...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:16.931831+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"create\"},\"type\":\"apply_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Creation complete after 0s\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:17.260811+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"create\",\"elapsed_seconds\":0},\"type\":\"apply_complete\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 1 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:17.281454+02:00\",\"changes\":{\"add\":1,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:17.281488+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:59:17+02:00	DEBUG	provider-jfrog-artifactory	Reconcile request has been requeued.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "rateLimiterName": "", "when": "0s"}
  2025-04-10T14:59:17+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:17+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:17.348382+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.728416+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.872554+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Drift detected (update)\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.878410+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"update\"},\"type\":\"resource_drift\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.878476+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.887650+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.887674+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	Resource is marked as available.	{"uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "name": "ansible-crossplane-repository", "gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository"}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	Reconcile request has been requeued.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "rateLimiterName": "status", "when": "5ms"}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	External resource is up to date	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381610", "external-name": "ansible-crossplane-repository", "requeue-after": "2025-04-10T15:09:18+02:00"}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:18+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:18.952086+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.448238+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.600727+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Drift detected (update)\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.611536+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"update\"},\"type\":\"resource_drift\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.611626+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.627979+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.628010+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	Resource is late-initialized.	{"uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "name": "ansible-crossplane-repository", "gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository"}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	External resource is up to date	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381612", "external-name": "ansible-crossplane-repository", "requeue-after": "2025-04-10T15:09:20+02:00"}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Update event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:59:20+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:22+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:20.770133+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.213627+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.360262+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Drift detected (update)\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.368330+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"update\"},\"type\":\"resource_drift\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.368392+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.381209+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.381234+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:59:22+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "args": ["plan", "-refresh=false", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:59:23+02:00	DEBUG	provider-jfrog-artifactory	plan ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:22.421494+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:59:23.734097+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n"}
  2025-04-10T14:59:23+02:00	DEBUG	provider-jfrog-artifactory	Called plan on the resource.	{"uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "name": "ansible-crossplane-repository", "gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "upToDate": true}
  2025-04-10T14:59:23+02:00	DEBUG	provider-jfrog-artifactory	External resource is up to date	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "c1f2abb9-b5c8-4187-a2e7-290e7a604ad3", "version": "381617", "external-name": "ansible-crossplane-repository", "requeue-after": "2025-04-10T15:09:23+02:00"}
  ```

</details>

1. User pushes K8s resource to the cluster
2. Initial reconciliation is triggered

    ```log
    DEBUG	provider-jfrog-artifactory	Calling the inner handler for Create event.
    ```

3. Terraform code is generated by provider
4. Terraform is initialized

    ```log
    DEBUG	provider-jfrog-artifactory	Running terraform
    DEBUG	provider-jfrog-artifactory	init ended
    ```

5. Terraform state is refreshed

    ```log
    DEBUG	provider-jfrog-artifactory	Running terraform
    DEBUG	provider-jfrog-artifactory	refresh ended
    ```

6. Provider sets `crossplane.io/external-name` annotation with `SetCriticalAnnotations` function
   1. Provider gets identificator from Terraform state with `GetExternalNameFn` function

      - Provider might fail on `cannot set critical annotations: cannot get external name: cannot find id in tfstate` when is `GetExternalNameFn` not set properly

        ```log
        DEBUG	provider-jfrog-artifactory	Cannot observe external resource
        ```

## Reconciling of the resource

<details>
  <summary>Full log</summary>

  ```log

  ```

</details>

## Processing of removed resourc

<details>
  <summary>Full log</summary>

  ```log
  2025-04-10T14:12:36+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Update event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:12:36+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:36+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:12:38+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:36.573369+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.090656+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.238299+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.241864+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.253551+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.253567+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:12:38+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "args": ["plan", "-refresh=false", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:12:39+02:00	DEBUG	provider-jfrog-artifactory	plan ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:38.290586+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:39.657380+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n"}
  2025-04-10T14:12:39+02:00	DEBUG	provider-jfrog-artifactory	Called plan on the resource.	{"uid": "5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "name": "ansible-crossplane-repository", "gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "upToDate": true}
  2025-04-10T14:12:39+02:00	DEBUG	provider-jfrog-artifactory	Successfully requested deletion of external resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "version": "375170", "external-name": "ansible-crossplane-repository", "deletion-timestamp": "2025-04-09 12:57:58 +0200 CEST"}
  2025-04-10T14:12:39+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "args": ["destroy", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:12:39+02:00	DEBUG	events	Successfully requested deletion of external resource	{"type": "Normal", "object": {"kind":"LocalAnsibleRepository","name":"ansible-crossplane-repository","uid":"5fb1abac-369f-48a7-973e-4c97d1bc6bc2","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"375170"}, "reason": "DeletedExternalResource"}
  2025-04-10T14:12:40+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:40+02:00	DEBUG	provider-jfrog-artifactory	Successfully requested deletion of external resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "version": "375176", "external-name": "ansible-crossplane-repository", "deletion-timestamp": "2025-04-09 12:57:58 +0200 CEST"}
  2025-04-10T14:12:40+02:00	DEBUG	events	Successfully requested deletion of external resource	{"type": "Normal", "object": {"kind":"LocalAnsibleRepository","name":"ansible-crossplane-repository","uid":"5fb1abac-369f-48a7-973e-4c97d1bc6bc2","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"375176"}, "reason": "DeletedExternalResource"}
  2025-04-10T14:12:42+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:42+02:00	DEBUG	provider-jfrog-artifactory	Successfully requested deletion of external resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "version": "375183", "external-name": "ansible-crossplane-repository", "deletion-timestamp": "2025-04-09 12:57:58 +0200 CEST"}
  2025-04-10T14:12:42+02:00	DEBUG	events	Successfully requested deletion of external resource	{"type": "Normal", "object": {"kind":"LocalAnsibleRepository","name":"ansible-crossplane-repository","uid":"5fb1abac-369f-48a7-973e-4c97d1bc6bc2","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"375183"}, "reason": "DeletedExternalResource"}
  2025-04-10T14:12:43+02:00	DEBUG	provider-jfrog-artifactory	destroy async ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:39.685746+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:41.069864+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:41.215795+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Plan to delete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:42.409540+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"delete\"},\"type\":\"planned_change\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 1 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:42.409604+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":1,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Destroying...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:43.637169+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"delete\"},\"type\":\"apply_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_local_ansible_repository.ansible-crossplane-repository: Destruction complete after 0s\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:43.974663+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"module\":\"\",\"resource\":\"artifactory_local_ansible_repository.ansible-crossplane-repository\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_local_ansible_repository\",\"resource_name\":\"ansible-crossplane-repository\",\"resource_key\":null},\"action\":\"delete\",\"elapsed_seconds\":0},\"type\":\"apply_complete\"}\n{\"@level\":\"info\",\"@message\":\"Destroy complete! Resources: 1 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:43.992205+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":1,\"operation\":\"destroy\"},\"type\":\"change_summary\"}\n"}
  2025-04-10T14:12:44+02:00	DEBUG	provider-jfrog-artifactory	Reconcile request has been requeued.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "rateLimiterName": "", "when": "0s"}
  2025-04-10T14:12:44+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:44+02:00	DEBUG	provider-jfrog-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
  2025-04-10T14:12:45+02:00	DEBUG	provider-jfrog-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:44.055762+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:45.485064+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:45.491747+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-10T14:12:45.491784+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
  2025-04-10T14:12:45+02:00	DEBUG	provider-jfrog-artifactory	Successfully deleted managed resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "uid": "5fb1abac-369f-48a7-973e-4c97d1bc6bc2", "version": "375194", "external-name": "ansible-crossplane-repository", "deletion-timestamp": "2025-04-09 12:57:58 +0200 CEST"}
  2025-04-10T14:12:45+02:00	DEBUG	provider-jfrog-artifactory	Calling the inner handler for Delete event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=LocalAnsibleRepository", "name": "ansible-crossplane-repository", "queueLength": 0}
  2025-04-10T14:12:45+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:45+02:00	DEBUG	provider-jfrog-artifactory	Cannot get managed resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "error": "LocalAnsibleRepository.artifactory.jfrog.crossplane.io \"ansible-crossplane-repository\" not found"}
  2025-04-10T14:12:46+02:00	DEBUG	provider-jfrog-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}}
  2025-04-10T14:12:46+02:00	DEBUG	provider-jfrog-artifactory	Cannot get managed resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=localansiblerepository", "request": {"name":"ansible-crossplane-repository"}, "error": "LocalAnsibleRepository.artifactory.jfrog.crossplane.io \"ansible-crossplane-repository\" not found"}
  ```

</details>
