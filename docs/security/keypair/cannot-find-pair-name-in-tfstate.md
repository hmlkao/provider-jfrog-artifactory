<!-- markdownlint-disable no-hard-tabs -->
# Cannot find `pair_name` in tfstate

## Parsed log from provider

```log
2025-04-09T14:01:19+02:00	DEBUG	provider-artifactory	Calling the inner handler for Create event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=Keypair", "name": "my-crossplane-keypair", "queueLength": 0}
2025-04-09T14:01:19+02:00	DEBUG	provider-artifactory	Reconciling	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=keypair", "request": {"name":"my-crossplane-keypair"}}
{
  "controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=keypair",
  "request": {
    "name": "my-crossplane-keypair"
  }
}
2025-04-09T14:01:19+02:00	DEBUG	provider-artifactory	Calling the inner handler for Update event.	{"gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=Keypair", "name": "my-crossplane-keypair", "queueLength": 0}
{
  "gvk": "artifactory.jfrog.crossplane.io/v1alpha1, Kind=Keypair",
  "name": "my-crossplane-keypair",
  "queueLength": 0
}
2025-04-09T14:01:19+02:00	DEBUG	provider-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7", "args": ["init", "-input=false"]}
2025-04-09T14:01:25+02:00	DEBUG	provider-artifactory	init ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7", "out": "\n\u001b[0m\u001b[1mInitializing the backend...\u001b[0m\n\n\u001b[0m\u001b[1mInitializing provider plugins...\u001b[0m\n- Finding jfrog/artifactory versions matching \"12.9.1\"...\n- Installing jfrog/artifactory v12.9.1...\n- Installed jfrog/artifactory v12.9.1 (signed by a HashiCorp partner, key ID \u001b[0m\u001b[1m2FA4D2A520237FA7\u001b[0m\u001b[0m)\n\nPartner and community providers are signed by their developers.\nIf you'd like to know more about provider signing, you can read about it here:\nhttps://www.terraform.io/docs/cli/plugins/signing.html\n\nTerraform has created a lock file \u001b[1m.terraform.lock.hcl\u001b[0m to record the provider\nselections it made above. Include this file in your version control repository\nso that Terraform can guarantee to make the same selections by default when\nyou run \"terraform init\" in the future.\u001b[0m\n\n\u001b[0m\u001b[1m\u001b[32mTerraform has been successfully initialized!\u001b[0m\u001b[32m\u001b[0m\n\u001b[0m\u001b[32m\nYou may now begin working with Terraform. Try running \"terraform plan\" to see\nany changes that are required for your infrastructure. All Terraform commands\nshould now work.\n\nIf you ever set or change modules or backend configuration for Terraform,\nrerun this command to reinitialize your working directory. If you forget, other\ncommands will detect it and remind you to do so if necessary.\u001b[0m\n"}
2025-04-09T14:01:25+02:00	DEBUG	provider-artifactory	Running terraform	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7", "args": ["apply", "-refresh-only", "-auto-approve", "-input=false", "-lock=false", "-json"]}
2025-04-09T14:01:26+02:00	DEBUG	provider-artifactory	refresh ended	{"workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7", "out": "{\"@level\":\"info\",\"@message\":\"Terraform 1.5.7\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:25.324154+02:00\",\"terraform\":\"1.5.7\",\"type\":\"version\",\"ui\":\"1.1\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_keypair.my-crossplane-keypair: Refreshing state...\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.819046+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_keypair.my-crossplane-keypair\",\"module\":\"\",\"resource\":\"artifactory_keypair.my-crossplane-keypair\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_keypair\",\"resource_name\":\"my-crossplane-keypair\",\"resource_key\":null}},\"type\":\"refresh_start\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_keypair.my-crossplane-keypair: Refresh complete\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.949052+02:00\",\"hook\":{\"resource\":{\"addr\":\"artifactory_keypair.my-crossplane-keypair\",\"module\":\"\",\"resource\":\"artifactory_keypair.my-crossplane-keypair\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_keypair\",\"resource_name\":\"my-crossplane-keypair\",\"resource_key\":null}},\"type\":\"refresh_complete\"}\n{\"@level\":\"info\",\"@message\":\"artifactory_keypair.my-crossplane-keypair: Drift detected (update)\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.955892+02:00\",\"change\":{\"resource\":{\"addr\":\"artifactory_keypair.my-crossplane-keypair\",\"module\":\"\",\"resource\":\"artifactory_keypair.my-crossplane-keypair\",\"implied_provider\":\"artifactory\",\"resource_type\":\"artifactory_keypair\",\"resource_name\":\"my-crossplane-keypair\",\"resource_key\":null},\"action\":\"update\"},\"type\":\"resource_drift\"}\n{\"@level\":\"info\",\"@message\":\"Plan: 0 to add, 0 to change, 0 to destroy.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.955988+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"plan\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Apply complete! Resources: 0 added, 0 changed, 0 destroyed.\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.969438+02:00\",\"changes\":{\"add\":0,\"change\":0,\"import\":0,\"remove\":0,\"operation\":\"apply\"},\"type\":\"change_summary\"}\n{\"@level\":\"info\",\"@message\":\"Outputs: 0\",\"@module\":\"terraform.ui\",\"@timestamp\":\"2025-04-09T14:01:26.969482+02:00\",\"outputs\":{},\"type\":\"outputs\"}\n"}
{
  "workspace": "/var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7",
  "out":
    {
      "@level": "info",
      "@message": "Terraform 1.5.7",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:25.324154+02:00",
      "terraform": "1.5.7",
      "type": "version",
      "ui": "1.1"
    }
    {
      "@level": "info",
      "@message": "artifactory_keypair.my-crossplane-keypair: Refreshing state...",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.819046+02:00",
      "hook": {
        "resource": {
          "addr": "artifactory_keypair.my-crossplane-keypair",
          "module": "",
          "resource": "artifactory_keypair.my-crossplane-keypair",
          "implied_provider": "artifactory",
          "resource_type": "artifactory_keypair",
          "resource_name": "my-crossplane-keypair",
          "resource_key": null
        }
      },
      "type": "refresh_start"
    }
    {
      "@level": "info",
      "@message": "artifactory_keypair.my-crossplane-keypair: Refresh complete",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.949052+02:00",
      "hook": {
        "resource": {
          "addr": "artifactory_keypair.my-crossplane-keypair",
          "module": "",
          "resource": "artifactory_keypair.my-crossplane-keypair",
          "implied_provider": "artifactory",
          "resource_type": "artifactory_keypair",
          "resource_name": "my-crossplane-keypair",
          "resource_key": null
        }
      },
      "type": "refresh_complete"
    }
    {
      "@level": "info",
      "@message": "artifactory_keypair.my-crossplane-keypair: Drift detected (update)",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.955892+02:00",
      "change": {
        "resource": {
          "addr": "artifactory_keypair.my-crossplane-keypair",
          "module": "",
          "resource": "artifactory_keypair.my-crossplane-keypair",
          "implied_provider": "artifactory",
          "resource_type": "artifactory_keypair",
          "resource_name": "my-crossplane-keypair",
          "resource_key": null
        },
        "action": "update"
      },
      "type": "resource_drift"
    }
    {
      "@level": "info",
      "@message": "Plan: 0 to add, 0 to change, 0 to destroy.",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.955988+02:00",
      "changes": {
        "add": 0,
        "change": 0,
        "import": 0,
        "remove": 0,
        "operation": "plan"
      },
      "type": "change_summary"
    }
    {
      "@level": "info",
      "@message": "Apply complete! Resources: 0 added, 0 changed, 0 destroyed.",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.969438+02:00",
      "changes": {
        "add": 0,
        "change": 0,
        "import": 0,
        "remove": 0,
        "operation": "apply"
      },
      "type": "change_summary"
    }
    {
      "@level": "info",
      "@message": "Outputs: 0",
      "@module": "terraform.ui",
      "@timestamp": "2025-04-09T14:01:26.969482+02:00",
      "outputs": {},
      "type": "outputs"
    }
}

2025-04-09T14:01:26+02:00	DEBUG	provider-artifactory	Cannot observe external resource	{"controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=keypair", "request": {"name":"my-crossplane-keypair"}, "uid": "203fa67f-74a6-41c1-9de8-49b3de5573f7", "version": "232364", "external-name": "", "error": "cannot set critical annotations: cannot get external name: cannot find pair_name in tfstate", "errorVerbose": "cannot find pair_name in tfstate\ncannot get external name\ngithub.com/crossplane/upjet/pkg/resource.SetCriticalAnnotations\n\tgithub.com/crossplane/upjet@v1.4.1/pkg/resource/lateinit.go:58\ngithub.com/crossplane/upjet/pkg/controller.(*external).Observe\n\tgithub.com/crossplane/upjet@v1.4.1/pkg/controller/external.go:258\ngithub.com/crossplane/crossplane-runtime/pkg/reconciler/managed.(*Reconciler).Reconcile\n\tgithub.com/crossplane/crossplane-runtime@v1.16.0/pkg/reconciler/managed/reconciler.go:914\ngithub.com/crossplane/crossplane-runtime/pkg/ratelimiter.(*Reconciler).Reconcile\n\tgithub.com/crossplane/crossplane-runtime@v1.16.0/pkg/ratelimiter/reconciler.go:54\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Reconcile\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:119\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:316\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:266\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Start.func2.2\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:227\nruntime.goexit\n\truntime/asm_arm64.s:1223\ncannot set critical annotations\ngithub.com/crossplane/upjet/pkg/controller.(*external).Observe\n\tgithub.com/crossplane/upjet@v1.4.1/pkg/controller/external.go:260\ngithub.com/crossplane/crossplane-runtime/pkg/reconciler/managed.(*Reconciler).Reconcile\n\tgithub.com/crossplane/crossplane-runtime@v1.16.0/pkg/reconciler/managed/reconciler.go:914\ngithub.com/crossplane/crossplane-runtime/pkg/ratelimiter.(*Reconciler).Reconcile\n\tgithub.com/crossplane/crossplane-runtime@v1.16.0/pkg/ratelimiter/reconciler.go:54\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Reconcile\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:119\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:316\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:266\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Start.func2.2\n\tsigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:227\nruntime.goexit\n\truntime/asm_arm64.s:1223"}
{
  "controller": "managed/artifactory.jfrog.crossplane.io/v1alpha1, kind=keypair",
  "request": {
    "name": "my-crossplane-keypair"
  },
  "uid": "203fa67f-74a6-41c1-9de8-49b3de5573f7",
  "version": "232364",
  "external-name": "",
  "error": "cannot set critical annotations: cannot get external name: cannot find pair_name in tfstate",
  "errorVerbose":
    cannot find pair_name in tfstate
    cannot get external name
    github.com/crossplane/upjet/pkg/resource.SetCriticalAnnotations
      github.com/crossplane/upjet@v1.4.1/pkg/resource/lateinit.go:58
    github.com/crossplane/upjet/pkg/controller.(*external).Observe
      github.com/crossplane/upjet@v1.4.1/pkg/controller/external.go:258
    github.com/crossplane/crossplane-runtime/pkg/reconciler/managed.(*Reconciler).Reconcile
      github.com/crossplane/crossplane-runtime@v1.16.0/pkg/reconciler/managed/reconciler.go:914
    github.com/crossplane/crossplane-runtime/pkg/ratelimiter.(*Reconciler).Reconcile
      github.com/crossplane/crossplane-runtime@v1.16.0/pkg/ratelimiter/reconciler.go:54
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Reconcile
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:119
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:316
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:266
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Start.func2.2
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:227
    runtime.goexit
      runtime/asm_arm64.s:1223
    cannot set critical annotations
    github.com/crossplane/upjet/pkg/controller.(*external).Observe
      github.com/crossplane/upjet@v1.4.1/pkg/controller/external.go:260
    github.com/crossplane/crossplane-runtime/pkg/reconciler/managed.(*Reconciler).Reconcile
      github.com/crossplane/crossplane-runtime@v1.16.0/pkg/reconciler/managed/reconciler.go:914
    github.com/crossplane/crossplane-runtime/pkg/ratelimiter.(*Reconciler).Reconcile
      github.com/crossplane/crossplane-runtime@v1.16.0/pkg/ratelimiter/reconciler.go:54
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Reconcile
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:119
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:316
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:266
    sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Start.func2.2
      sigs.k8s.io/controller-runtime@v0.17.0/pkg/internal/controller/controller.go:227
    runtime.goexit
      runtime/asm_arm64.s:1223
}
2025-04-09T14:01:26+02:00	DEBUG	events	cannot set critical annotations: cannot get external name: cannot find pair_name in tfstate	{"type": "Warning", "object": {"kind":"Keypair","name":"my-crossplane-keypair","uid":"203fa67f-74a6-41c1-9de8-49b3de5573f7","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"232365"}, "reason": "CannotObserveExternalResource"}
```

## Manual Terraform run

```bash
cd /var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/203fa67f-74a6-41c1-9de8-49b3de5573f7
```

Run Terraform refresh the same way as provider does

```bash
$ terraform apply -refresh-only -auto-approve -input=false -lock=false
artifactory_keypair.my-crossplane-keypair: Refreshing state...
╷
│ Error: Unable to Refresh Resource
│
│   with artifactory_keypair.my-crossplane-keypair,
│   on main.tf.json line 19, in resource.artifactory_keypair.my-crossplane-keypair:
│   19:       }
│
│ An unexpected error occurred while attempting to refresh resource state. Please retry the operation or report this issue to the provider developers.
│
│ Error: json: cannot unmarshal array into Go value of type security.KeyPairAPIModel
```

Run Terraform refresh with debug enabled when there is no key pair in Artifactory

```bash
$ TF_LOG=debug terraform apply -refresh-only -auto-approve -input=false -lock=false
...
2025/04/09 15:10:35.772749 DEBUG RESTY
==============================================================================
~~~ REQUEST ~~~
GET  /artifactory/api/security/keypair/  HTTP/1.1
HOST   : artifacts-testing.merck.com
HEADERS:
        Accept: */*
        Authorization: <REDACTED>
        Content-Type: application/json
        Cookie: AWSALB=zPqpCXCmaJOuMkjk7YRcWlDY8iOEpjHli2Is7s0GFzo3srELNQw8WTJA1j5TaHgmuIc7JJqm2wx2futvwe/V566VL42sS+2GLsgUdU0tGNAjbiu0Xa5RsrnkWmfL; AWSALBCORS=zPqpCXCmaJOuMkjk7YRcWlDY8iOEpjHli2Is7s0GFzo3srELNQw8WTJA1j5TaHgmuIc7JJqm2wx2futvwe/V566VL42sS+2GLsgUdU0tGNAjbiu0Xa5RsrnkWmfL
        User-Agent: jfrog/terraform-provider-artifactory/12.9.1
BODY   :
***** NO CONTENT *****
------------------------------------------------------------------------------
~~~ RESPONSE ~~~
STATUS       : 200 OK
PROTO        : HTTP/2.0
RECEIVED AT  : 2025-04-09T15:10:35.772724+02:00
TIME DURATION: 125.834125ms
HEADERS      :
        Content-Type: application/json
        Date: Wed, 09 Apr 2025 13:10:35 GMT
        Set-Cookie: AWSALB=sT2tpybJ5bSRoUKIdMltRtPGPzW0kif1AdmLbxb065m1KQw2w6mEj/Z3Wb7LEmrZDo8BfQ9/ZB+BS2NeSHaMdxpFNdGa5soiv7UrnddPqQc+scB2iwVUwGNAN029; Expires=Wed, 16 Apr 2025 13:10:35 GMT; Path=/, AWSALBCORS=sT2tpybJ5bSRoUKIdMltRtPGPzW0kif1AdmLbxb065m1KQw2w6mEj/Z3Wb7LEmrZDo8BfQ9/ZB+BS2NeSHaMdxpFNdGa5soiv7UrnddPqQc+scB2iwVUwGNAN029; Expires=Wed, 16 Apr 2025 13:10:35 GMT; Path=/; SameSite=None; Secure
        Strict-Transport-Security: always
        X-Artifactory-Id: fee86edeb04e8db3:4524b271:195a3c747b2:-8000
        X-Artifactory-Node-Id: artifactory-1-test
        X-Jfrog-Version: Artifactory/7.98.9 79809900
BODY         :
[]
=============================================================================="
2025/04/09 15:10:35.772768 ERROR RESTY json: cannot unmarshal array into Go value of type security.KeyPairAPIModel"
```

Run Terraform refresh with debug enabled when there is at least one key pair

```bash
$ TF_LOG=debug terraform apply -refresh-only -auto-approve -input=false -lock=false
2025/04/09 15:17:40.715982 DEBUG RESTY
==============================================================================
~~~ REQUEST ~~~
GET  /artifactory/api/security/keypair/  HTTP/1.1
HOST   : artifacts-testing.merck.com
HEADERS:
        Accept: */*
        Authorization: <REDACTED>
        Content-Type: application/json
        Cookie: AWSALB=C3u0eN6o6gGQxsMettpWjXaLcVnZeQMzHXAOeBrGvEiRevrvHL0GzfE89niW27yETlmwDTF8OnnIaw3MZTWUxDUwAgCCPribecNBxWLjVoywVU/7th+7qwTLxAv8; AWSALBCORS=C3u0eN6o6gGQxsMettpWjXaLcVnZeQMzHXAOeBrGvEiRevrvHL0GzfE89niW27yETlmwDTF8OnnIaw3MZTWUxDUwAgCCPribecNBxWLjVoywVU/7th+7qwTLxAv8
        User-Agent: jfrog/terraform-provider-artifactory/12.9.1
BODY   :
***** NO CONTENT *****
------------------------------------------------------------------------------
~~~ RESPONSE ~~~
STATUS       : 200 OK
PROTO        : HTTP/2.0
RECEIVED AT  : 2025-04-09T15:17:40.71594+02:00
TIME DURATION: 132.710125ms
HEADERS      :
        Content-Type: application/json
        Date: Wed, 09 Apr 2025 13:17:40 GMT
        Set-Cookie: AWSALB=+j/GfmX0oIdWrQR5mAJR5vAfh1M8dluTNnw2YXy8hNceDJZy0Ea7MLvvymxz+1/LTQ51qyXWwzfcNqQeXX+0TNohhhaB8zCD6jy2xWu0yV5NCPL3mgDjo2JSnyFK; Expires=Wed, 16 Apr 2025 13:17:40 GMT; Path=/, AWSALBCORS=+j/GfmX0oIdWrQR5mAJR5vAfh1M8dluTNnw2YXy8hNceDJZy0Ea7MLvvymxz+1/LTQ51qyXWwzfcNqQeXX+0TNohhhaB8zCD6jy2xWu0yV5NCPL3mgDjo2JSnyFK; Expires=Wed, 16 Apr 2025 13:17:40 GMT; Path=/; SameSite=None; Secure
        Strict-Transport-Security: always
        X-Artifactory-Id: fee86edeb04e8db3:4524b271:195a3c747b2:-8000
        X-Artifactory-Node-Id: artifactory-1-test
        X-Jfrog-Version: Artifactory/7.98.9 79809900
BODY         :
[
   {
      "pairName": "asdf",
      "pairType": "RSA",
      "alias": "srvcdbuild",
      "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1gCrNUFUNbYQTtr53li\nT/y5itkQA8Gus/nuFOhy0SyWAZ7+CTJnJmFSWLxqszjkasQ5qHC+8j/+WOD+hgKt\nq4cII/GkeVTQm0zuzACHCzZzcqfo53OxBFpK3YJjjFtF5RCQqEjQnBKz19Hvuweb\nEB9sh+B0rOiM9227VBQ4BVn26twjUfXoP9w5Ulxk+selHZ11e35KMbEULZeahOLh\n1Q+KCEqhGSWOB6jVU3LulOsIOlp0qYdM/muMub8vaSqu5wfUFM4ogjChZCrc9tgn\n8mJZdovAK0J1x4YEkT/J0GwlVuY4sZQYXG50ou7O0TaR7jawz4UpcuHsfwdCKR0W\npQIDAQAB\n-----END PUBLIC KEY-----\n",
      "unavailable": false
   }
]
=============================================================================="
2025/04/09 15:17:40.716013 ERROR RESTY json: cannot unmarshal array into Go value of type security.KeyPairAPIModel"
```

## Weird behavior

Output from `terraform apply -refresh-only` command run by provider is different from when I run it manually from the folder.

If I apply Terraform manually, it passes well.

```bash
$ terraform apply -refresh=false -input=false -lock=false

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # artifactory_keypair.my-crossplane-keypair will be created
  + resource "artifactory_keypair" "my-crossplane-keypair" {
      + alias       = "my-keypair-alias"
      + pair_name   = "my-crossplane-keypair"
      + pair_type   = "RSA"
      + private_key = (sensitive value)
      + public_key  = <<-EOT
            -----BEGIN PUBLIC KEY-----
            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8AMZ8x2rIuUKjZGyk/qF
            RrAppHhix36zQLZWwvVfxeNsgjqwwQdP2VJbSmZfoAnjR/uHMMrkKEMMik6qD5Jd
            wwMtFm8zQblLXZeO8TFbc8a0S+nfKOgpGie755mCg+2yBUZxQUtaaY+n4Cgp/sSd
            8lzuNh+BnZbxLVFEmy7K6sEf8Fuerbnc/WBe8alqQTimNj2fcG7wwIxP0PpzT4o2
            h+ktNSUuWFy1UpJS1KvpPjjsa3KXw/lTK2EN1HBt6oEQDwFJfelI2T2hvBbA+vKS
            VwR5rPubgY1NIeM9TK/N8aP+bbyNbyn8RyhcQepGENS2H4yGO4cIdtzhItHHp+r7
            1wIDAQAB
            -----END PUBLIC KEY-----
        EOT
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

artifactory_keypair.my-crossplane-keypair: Creating...
artifactory_keypair.my-crossplane-keypair: Creation complete after 0s

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

If I try to refresh the state again, it also passes well.

```bash
$ terraform apply -refresh-only -auto-approve -input=false -lock=false
artifactory_keypair.my-crossplane-keypair: Refreshing state...

No changes. Your infrastructure still matches the configuration.

Terraform has checked that the real remote objects still match the result of your most recent changes, and found no differences.

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
```

So, the idea with bad parsing of returned values might be wrong... Otherwise, I would expect that `terraform apply -refresh-only` fails everytime.
