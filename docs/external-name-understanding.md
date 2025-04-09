<!-- markdownlint-disable no-hard-tabs -->
# External name understanding

Crossplane was primarily developed to manage infrastructure on Cloud platforms (AWS, GCP, Azure).

Terraform state of resources of these platforms usualy set `id` parameter, e.g.

`LocalGenericRepository` resource as is defined by user

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

Real `LocalGenericRepository` resource as is stored in Kubernetes

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

## Processing of resource

1. User pushes K8s resource to the cluster
2. Initial reconciliation is triggered

    ```log
    DEBUG	provider-artifactory	Calling the inner handler for Create event.
    ```

3. Terraform code is generated by provider
4. Terraform is initialized

    ```log
    DEBUG	provider-artifactory	Running terraform
    DEBUG	provider-artifactory	init ended
    ```

5. Terraform state is refreshed

    ```log
    DEBUG	provider-artifactory	Running terraform
    DEBUG	provider-artifactory	refresh ended
    ```

6. Provider fails on `cannot set critical annotations: cannot get external name: cannot find id in tfstate`

    ```log
    DEBUG	provider-artifactory	Cannot observe external resource
    ```

    1. Function `GetExternalNameFn()` is called from `SetCriticalAnnotations()`
