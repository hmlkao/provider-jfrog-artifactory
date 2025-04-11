<!-- markdownlint-disable no-duplicate-heading no-hard-tabs -->
# Build provider from scratch

`provider-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

> [!WARNING]
> This repo targets to do as fewer changes as possible in source code to successfully build the provider with explanation.

- [Build provider from scratch](#build-provider-from-scratch)
  - [Steps to build using template](#steps-to-build-using-template)
  - [Testing the generated resources](#testing-the-generated-resources)
    - [Debugging](#debugging)
  - [Publish provider](#publish-provider)
  - [Publish to Upbound Marketplace](#publish-to-upbound-marketplace)
  - [Troubleshooting](#troubleshooting)
    - [`make generate` fails with error](#make-generate-fails-with-error)
      - [Solution](#solution)
    - [`cannot find id in tfstate` in provider output](#cannot-find-id-in-tfstate-in-provider-output)
      - [Solution](#solution-1)
    - [`make build` fails with error](#make-build-fails-with-error)
      - [Solution](#solution-2)
    - [Tag GH workflow fails](#tag-gh-workflow-fails)
      - [Solution](#solution-3)
    - [Failed to instantiate provider "xxx" to obtain schema](#failed-to-instantiate-provider-xxx-to-obtain-schema)
      - [Solution](#solution-4)

## Steps to build using template

Followed steps in [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)

1. Create a new repo from template [crossplane/upjet-provider-template](https://github.com/crossplane/upjet-provider-template)
2. Clone repo to your localhost and change directory into it

    ```bash
    git clone git@github.com:hmlkao/provider-artifactory.git
    cd provider-artifactory
    ```

3. Fetch submodule (no change in source code)

    > [!WARNING]
    > Submodule `Avarei/build` is used for the `build/` path, which contains essential scripts and configurations for the build process. This submodule is necessary because [this PR](https://github.com/crossplane/build/pull/14), which integrates these changes into the main repository, has not been merged yet.

    ```bash
    make submodules
    ```

4. Setup provider name and group

    ```bash
    $ ./hack/prepare.sh
    Lower case provider name (ex. github): artifactory
    Normal case provider name (ex. GitHub): Artifactory
    Organization (ex. upbound, my-org-name): hmlkao
    CRD rootGroup (ex. upbound.io, crossplane.io): artifactory.jfrog.crossplane.io
    ```

5. Configure Terraform provider in `Makefile`

    ```bash
    export TERRAFORM_PROVIDER_SOURCE ?= jfrog/artifactory
    export TERRAFORM_PROVIDER_REPO ?= https://github.com/jfrog/terraform-provider-artifactory
    export TERRAFORM_PROVIDER_VERSION ?= 12.9.1
    export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-artifactory
    export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/jfrog/$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
    export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-artifactory_v12.9.1
    export TERRAFORM_DOCS_PATH ?= docs/resources
    ```

6. Add `ProviderConfig` logic to `internal/clients/artifactory.go` according to [Terraform provider argument reference](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs#argument-reference)

    It means to add provider configuration arguments. For the Artifactory Terraform provider, the arguments `url`, `access_token`, `oidc_provider_name`, and `tfc_credential_tag_name` are optional. The `api_key` argument can be omitted as it is deprecated (in v12.9.1).

7. Add external name to `config/external_name.go` file according to the Terraform import reference.

   - For example, with `artifactory_local_oci_repository`, use the Kubernetes resource `metadata.name` as the `key` parameter. This ensures that the resource `name` and the OCI repository `key` are identical, avoiding confusion. To achieve this, set `config.ParameterAsIdentifier("key")`, which maps the `key` parameter to the `metadata.name` of the Kubernetes resource.

        ```text
        terraform import artifactory_local_oci_repository.my-oci-local-repo name-of-oci-local-repo
                                                          ~~~~~~~~~~~~~~~~~ ~~~~~~~~~~~~~~~~~~~~~~
                                                                \                   \
                                                                 `-------------------`------------   These fields may differ, if we set `"artifactory_local_oci_repository": config.ParameterAsIdentifier("key"),` in `ExternalNameConfigs` variable the key is always used as resource name
        ```

8. Create a folder and add custom configuration to the `config/` folder for Artifactory resource(s).

    > [!NOTE]
    > More details are in [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        mkdir config/local_oci_repository
        ```

9. Remove the `null` resource as it is included in the template for demonstration purposes and does not apply to the Artifactory provider. Keeping it may lead to unnecessary clutter or errors during the build process.

    ```bash
    rm -rf config/null
    ```

    And remove this reference from `config/provider.go`

10. Add line to `config/provider.go`, e.g. for `artifactory_local_oci_repository`

    > [!NOTE]
    > It's not necessary to have it here if you don't need to override the default resource.
    > If you don't need to configure it, it will generate resource according to default behavior.

    ```go
    for _, configure := range []func(provider *ujconfig.Provider){
      ...
      localocirepository.Configure,
    } {
      configure(pc)
    }
    ```

11. Update Go modules to ensure all required dependencies are downloaded and properly resolved, avoiding errors caused by missing modules.

    ```sh
    go mod tidy
    ```

12. Install required Go tools

    `goimports` is a tool that formats Go code and automatically manages import statements, ensuring consistency and reducing manual effort during the build process.

    ```bash
    go get golang.org/x/tools/cmd/goimports@latest
    go install golang.org/x/tools/cmd/goimports@latest
    ```

13. Build provider

    The `make generate` command is used to generate the necessary Crossplane Custom Resource Definitions (CRDs) and other artifacts required for the provider. This step ensures that the provider is correctly configured and ready for deployment. (check [Troubleshooting](#make-generate-fails-with-error) in case it fails)

    ```sh
    make generate
    ```

## Testing the generated resources

According to [crossplane/upjet documentation](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md#testing-the-generated-resources).

1. Remove examples for `null` provider (from template)

    ```bash
    rm -rf examples/null
    rm -rf examples/storeconfig
    ```

2. Create provider configuration

    ```bash
    cat <<EOF > examples/providerconfig/secret.yaml.tmpl
    apiVersion: v1
    kind: Secret
    metadata:
      name: artifactory-creds
      namespace: crossplane-system
    type: Opaque
    stringData:
      credentials: |
        {
          "url": "y0ur-url",
          "access_token": "y0ur-t0k3n"
        }
    EOF
    ```

3. Create directory for resource examples

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        mkdir -p examples/local_repositories/local_oci_repository
        ```

4. Create K8s resource YAML file

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        cat <<EOF > examples/local_repositories/local_oci_repository/local_oci_repository.yaml
        apiVersion: artifactory.jfrog.upbound.io/v1alpha1
        kind: LocalOCIRepository
        metadata:
          name: oci-crossplane-repository
        spec:
          forProvider: {}
        providerConfigRef:
          name: default
        EOF
        ```

5. Create K8s Secret from template

    ```bash
    $ ARTIFACTORY_URL=https://artifactory.site.com/artifactory
    $ read -r ARTIFACTORY_TOKEN
    <put-your-token-here + enter>
    $ cat examples/providerconfig/secret.yaml.tmpl | sed -e "s/y0ur-t0k3n/${ARTIFACTORY_TOKEN}/g" -e "s^y0ur-url^${ARTIFACTORY_URL}^g" > examples/providerconfig/secret.yaml
    ```

6. Apply CRDs

    ```bash
    kubectl apply -f package/crds/
    ```

7. Run provider on local machine

    > [!NOTE]
    > Please make sure Terraform is installed before running the "make run" command, you can check [this guide](https://developer.hashicorp.com/terraform/downloads).

    ```bash
    make run
    ```

    It will block the terminal, you can check provider logs there.

8. Follow commands must be run in different terminal due to block of the first one
9. Deploy `ProviderConfig` to the cluster

    ```bash
    kubectl apply -f examples/providerconfig/
    ```

10. Deploy Artifactory provider resource

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        kubectl apply -f examples/local_repositories/local_oci_repository/
        ```

11. Check the resource in K8s cluster

    (check [Troubleshooting](#cannot-find-id-in-tfstate-in-provider-output) in case of failing resource)

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        kubectl get localocirepositories.artifactory.jfrog.crossplane.io
        kubectl get managed
        ```

### Debugging

When you do any change in `config/` folder, regenarate CRDs and apply them to cluster to be sure that you use actual resource specifications.

```bash
make generate
kubectl apply -f package/crds/
```

When you run `make run`, Terraform state from Crossplane provider is stored in `/var/folders/` folder, check logs from provider to get exact path of Terraform workspace, so you can check what exactly is applied and what is stored in Terraform state.

## Publish provider

> [!NOTE]
> To be able to write artifact to GitHub, allow Read/Write permissions to GitHub Action Workflow
> You can enable it generally in GH repository Settings > Actions > General > Workflow permissions > Read and write permissions
>
> --OR--
>
> Set this permission directly in `.github/workflows/ci.yaml` workflow
>
> ```yaml
> permissions:
>   packages: write
> ```

Steps to publish provider to Upbound Marketplace according to [these instructions](https://docs.upbound.io/upbound-marketplace/packages/).

1. [Upbound account](https://docs.upbound.io/operate/accounts/identity-management/users/#create-an-account) is required
2. Create an Organization within Upbound account, e.g. `hmlkao`
3. Craete a Team in Organization settings, e.g. `Artifactory`
4. Create a Robot account in Organization settings
   1. Name - e.g., `provider-artifactory-ci`
   2. Create a token
      1. Name - e.g., `github-ci`
5. Assign the Robot account to the Team
6. Create a new Repository according to provider name, e.g. `provider-artifactory` (MUST match with provider name!)
7. Assign a Write permissions for the Team in the Repository settings
8. Set these GitHub secrets in the repository

    - `UPBOUND_MARKETPLACE_PUSH_ROBOT_USR` - Robot token Access ID, e.g. `a064335d-4b4e-78da-89fb-0d0bd9e832ee`
    - `UPBOUND_MARKETPLACE_PUSH_ROBOT_PSW` - Robot token itself

9. Update registry domain to `xpkg.upbound.io/<org-name>` on all places (`REGISTRY_ORGS`, `XPKG_REG_ORGS` and `XPKG_REG_ORGS_NO_PROMOTE`) in `Makefile`, e.g. replace `xpkg.upbound.io/upbound` with `xpkg.upbound.io/hmlkao`
10. Push changes to GitHub, CI workflow will be triggered and package pushed to Upbound repository

## Publish to Upbound Marketplace

You can do the initial publish from the commandline. No need to write to `#upbound` channel on Crossplane Slack as mentioned in [Upbound Marketplace documentation](https://docs.upbound.io/build/repositories/publish-packages/#publishing-public-packages).

1. Create an access token for your account
   1. My Account > API tokens
2. Install `up` command on your local machine

    ```bash
    curl -sL https://cli.upbound.io | CHANNEL=main sh
    ```

3. Login to your Upbound account with token

    ```bash
    $ read -r UP_TOKEN
    <paste-your-token-here+enter>
    $ export UP_TOKEN
    $ up login -u <access-token-id> --organization=<org-id>

    # Example
    $ up login -u f41ef3fc-f7d6-4ee4-ca51-9e68cce10e01 --organization=hmlkao
    ```

4. Enable auto-publishing on the repository

    ```bash
    up repo update --organization=<org-id> --private=false --publish <repo-name>

    # Example
    up repo update --organization=hmlkao --private=false --publish provider-artifactory
    ```

All current and new versions with semver tag will be published.

## Troubleshooting

### `make generate` fails with error

```text
scraper: error: Failed to scrape Terraform provider metadata: cannot scrape Terraform registry: failed to scrape resource metadata from path: ../.work/jfrog/artifactory/docs/resources/anonymous_user.md: failed to find the prelude of the document using the xpath expressions: //text()[contains(., "description") and contains(., "page_title")]
exit status 1
```

Upjet tool requires `page_title` and `description` fields in each file in Terraform Markdown documentation file and generated docs unfortunately contain them.

#### Solution

This is a workaround which generates valid Markdown files which are unfortunately missing siginificant part of origin files.

> [!NOTE]
> I created [an issue](https://github.com/jfrog/terraform-provider-artifactory/issues/1229) asking what the fix looks like

1. Clone repo of the Terraform provider and change directory into it

    ```sh
    git clone git@github.com:jfrog/terraform-provider-artifactory.git
    cd terraform-provider-artifactory
    ```

2. Checkout to the same version as you specify in [`Makefile`](./Makefile) variable `TERRAFORM_PROVIDER_VERSION`

    ```sh
    # Example
    git checkout v12.9.1
    ```

3. Check version of `github.com/hashicorp/terraform-plugin-docs` in `go.mod` file

    ```sh
    grep 'hashicorp/terraform-plugin-docs' go.mod
    ```

4. (*optional*) Install [`hashicorp/terraform-plugin-docs`](https://github.com/hashicorp/terraform-plugin-docs?tab=readme-ov-file#installation) CLI tool

    ```sh
    go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.21.0
    go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.21.0
    # Verify that it is possible to find binary
    which tfplugindocs
    ```

5. Generate docs

    ```sh
    make doc
    ```

6. Copy generated doc to the path of Terraform provider in Crossplane provider

    ```sh
    rm -rf ../provider-artifactory/.work/jfrog/artifactory/docs/
    cp -R docs/ ../provider-artifactory/.work/jfrog/artifactory/docs/
    ```

### `cannot find id in tfstate` in provider output

```text
2025-03-26T17:15:32+01:00	DEBUG	events	cannot set critical annotations: cannot get external name: cannot find id in tfstate	{"type": "Warning", "object": {"kind":"LocalOCIRepository","name":"crossplane-oci","uid":"dbbdd029-85e7-4f87-826e-72071d3c6d99","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"952939"}, "reason": "CannotObserveExternalResource"}
```

- E.g., for `artifactory_local_oci_repository` the problem is caused by Artifactory Terraform provider not using `id` attribute for identifying resources in Terraform state while Upjet counts with that.

    ```json
    {
      "version": 4,
      "terraform_version": "1.11.2",
      "serial": 2,
      "lineage": "473524e6-61f9-40c9-b1f4-9724f73c36e7",
      "outputs": {},
      "resources": [
        {
          "mode": "managed",
          "type": "artifactory_local_oci_repository",
          "name": "oci-crossplane-repository",
          "provider": "provider[\"registry.terraform.io/jfrog/artifactory\"]",
          "instances": [
            {
              "schema_version": 1,
              "attributes": {
                // <<<<<<< NO 'id' attribute here
              },
              "sensitive_attributes": []
            }
          ]
        }
      ],
      "check_results": null
    }
    ```

#### Solution

Can be solved by own `GetExternalNameFn` function, just add it to the resource configuration in, e.g. `config/local_oci_repository/config.go`.

TODO: However, I think that there is better solution than this, but requires more investigation.

```go
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_oci_repository", func(r *config.Resource) {
		...
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find id in tfstate")
		}
	})
}
```

### `make build` fails with error

When you use Podman as container runtime on MacOS.

```log
up: error: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
16:52:09 [FAIL]
make[3]: *** [xpkg.build.provider-artifactory] Error 1
make[2]: *** [do.build.artifacts.linux_arm64] Error 2
make[1]: *** [build.all] Error 2
make: *** [build] Error 2
```

#### Solution

Follow [these instructions](https://podman-desktop.io/docs/migrating-from-docker/using-the-docker_host-environment-variable#procedure) in Podman docs.

1. Get current Podman machine socket

    ```bash
    podman machine inspect --format '{{.ConnectionInfo.PodmanSocket.Path}}' | pbcopy
    ## Example
    # podman machine inspect --format '{{.ConnectionInfo.PodmanSocket.Path}}'
    # /var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/podman/podman-machine-default-api.sock
    ```

2. Set `DOCKER_HOST`

    ```bash
    export DOCKER_HOST=unix://<paste-socket-path-here>
    ## Example
    # export DOCKER_HOST=unix:///var/folders/sj/hzvqvb413qgd6mx110wmbr6m0000gn/T/podman/podman-machine-default-api.sock
    ```

### Tag GH workflow fails

```text
Resource not accessible by integration
```

#### Solution

Add this block to the `.github/workflows/tag.yaml`:

```yaml
permissions:
  contents: write
```

### Failed to instantiate provider "xxx" to obtain schema

When you run `make run` and you'll get following error message:

```log
2025-04-08T13:05:11+02:00	DEBUG	events	cannot run refresh: refresh failed: failed to read schema for artifactory_item_properties.my-folder-properties in registry.terraform.io/jfrog/artifactory: failed to instantiate provider "registry.terraform.io/jfrog/artifactory" to obtain schema: Unrecognized remote plugin message:

This usually means that the plugin is either invalid or simply
needs to be recompiled to support the latest protocol.: 	{"type": "Warning", "object": {"kind":"ItemProperties","name":"my-folder-properties","uid":"c4312fe6-b815-4282-b123-49ea5577a6f5","apiVersion":"artifactory.jfrog.crossplane.io/v1alpha1","resourceVersion":"130449"}, "reason": "CannotObserveExternalResource"}
```

#### Solution

1. Try to run Terraform directly with debug as it is in [`Makefile`](./Makefile)

    ```bash
    $ TF_LOG=debug .cache/tools/darwin_arm64/terraform-1.5.7 -chdir=.work/terraform/ providers schema -json=true
    ...
    2025-04-08T14:14:52.068+0200 [DEBUG] provider: plugin started: path=.terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1 pid=25512
    2025-04-08T14:14:52.068+0200 [DEBUG] provider: waiting for RPC address: path=.terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    ╷
    │ Error: Failed to load plugin schemas
    │
    │ Error while loading schemas for plugin components: Failed to obtain provider schema: Could not load the schema for provider registry.terraform.io/jfrog/artifactory: failed to instantiate provider "registry.terraform.io/jfrog/artifactory" to obtain schema: timeout while waiting for plugin to start..
    ╵

    2025-04-08T14:15:52.071+0200 [WARN]  provider: plugin failed to exit gracefully
    ```

2. Check running terraform processes

    ```bash
    $ ps ax | grep terraform
    34519 s002  S+     0:00.00 grep terraform
    60678 s002  S      0:04.03 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    7399 s005  S      0:00.53 /Users/homolkao/Work/provider-artifactory/.cache/tools/darwin_arm64/terraform-1.5.7 -chdir=/Users/homolkao/Work/provider-artifactory/.work/terraform providers schema -json=true
    7400 s005  UE     0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    8404 s005  S      0:00.51 /Users/homolkao/Work/provider-artifactory/.cache/tools/darwin_arm64/terraform-1.5.7 -chdir=/Users/homolkao/Work/provider-artifactory/.work/terraform providers schema -json=true
    8407 s005  UE     0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    24394 s005  S      0:00.25 /Users/homolkao/Work/provider-artifactory/.cache/tools/darwin_arm64/terraform-1.5.7 -chdir=/Users/homolkao/Work/provider-artifactory/.work/terraform providers schema -json=true
    24395 s005  UE     0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    25509 s005  S+     0:00.21 .cache/tools/darwin_arm64/terraform-1.5.7 -chdir=.work/terraform/ providers schema -json=true
    25512 s005  UE+    0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    92484 s005  S      0:00.76 /Users/homolkao/Work/provider-artifactory/.cache/tools/darwin_arm64/terraform-1.5.7 -chdir=/Users/homolkao/Work/provider-artifactory/.work/terraform providers schema -json=true
    92485 s005  UE     0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    82693 s010  S+     0:00.84 terraform plan
    82694 s010  UE+    0:00.00 .terraform/providers/registry.terraform.io/jfrog/artifactory/12.9.1/darwin_arm64/terraform-provider-artifactory_v12.9.1
    ```

    From `man ps`, `UE` stat means

      - `U       Marks a process in uninterruptible wait.`
      - `E       The process is trying to exit.`

    These processes cannot be killed (even using `kill -9`). The only way I found to get rid of it is to restart laptop...

    However, you can kill at least terraform processes `kill -9 terraform`, it may unblock laptop to run `make run` again.
