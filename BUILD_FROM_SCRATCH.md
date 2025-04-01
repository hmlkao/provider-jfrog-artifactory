<!-- markdownlint-disable no-duplicate-heading no-hard-tabs -->
# Build provider from scratch

`provider-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

> [!WARNING]
> This repo targets to do as less changes as possible in source code to successfully build the provider with explanation.

- [Build provider from scratch](#build-provider-from-scratch)
  - [Steps to build using template](#steps-to-build-using-template)
  - [Testing the generated resources](#testing-the-generated-resources)
    - [Debugging](#debugging)
  - [Publish provider](#publish-provider)
  - [Troubleshooting](#troubleshooting)
    - [`make generate` fails with error](#make-generate-fails-with-error)
      - [Solution](#solution)
    - [`cannot find id in tfstate` in provider output](#cannot-find-id-in-tfstate-in-provider-output)
      - [Solution](#solution-1)
    - [`make build` fails with error](#make-build-fails-with-error)
      - [Solution](#solution-2)

## Steps to build using template

Followed steps in [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)

1. Clone repo to your localhost and change directory into it

    ```bash
    git clone git@github.com:hmlkao/provider-artifactory.git
    cd provider-artifactory
    ```

2. Fetch submodule (no change in source code)

    ```bash
    make submodules
    ```

3. Setup provider name and group

    ```bash
    $ ./hack/prepare.sh
    Lower case provider name (ex. github): artifactory
    Normal case provider name (ex. GitHub): Artifactory
    Organization (ex. upbound, my-org-name): hmlkao
    CRD rootGroup (ex. upbound.io, crossplane.io): jfrog.crossplane.io
    ```

4. Configure Terraform provider in `Makefile`

    ```bash
    export TERRAFORM_PROVIDER_SOURCE ?= jfrog/artifactory
    export TERRAFORM_PROVIDER_REPO ?= https://github.com/jfrog/terraform-provider-artifactory
    export TERRAFORM_PROVIDER_VERSION ?= 12.9.1
    export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-artifactory
    export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/jfrog/$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
    export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-artifactory_v12.9.1
    export TERRAFORM_DOCS_PATH ?= docs/resources
    ```

5. Add `ProviderConfig` logic to `internal/clients/artifactory.go` according to [Terraform provider argument reference](https://registry.terraform.io/providers/jfrog/artifactory/12.9.0/docs#argument-reference)

    It means to add provider configuration arguments, for Artifactory Terraform provider it means to set optional `url`, `access_token`, `oidc_provider_name` and `tfc_credential_tag_name`, we can omit `api_key` as it's deprecated (in v12.9.1).

6. Add external name to `config/external_name.go` file according to Terraform import reference, examples:

   - E.g., for `artifactory_local_oci_repository` I want to use K8s resource `metadata.name` as a `key` parameter, otherwise we will have resource `name` and OCI repository `key` different which can be confusing, so we set `config.ParameterAsIdentifier("key")` which cause that the `key` parameter is set to the same as `metadata.name` of K8s resource

        ```text
        terraform import artifactory_local_oci_repository.my-oci-local-repo name-of-oci-local-repo
                                                          ~~~~~~~~~~~~~~~~~ ~~~~~~~~~~~~~~~~~~~~~~
                                                                \                   \
                                                                 `-------------------`------------   These fields may differ, so we need to set `"artifactory_local_oci_repository": config.ParameterAsIdentifier("key"),` in `ExternalNameConfigs` variable
        ```

7. Create folder and add custom configuration to `config/` folder for Artifactory resource(s) and remove `null` resource

    > [!NOTE]
    > More details are in [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        mkdir config/local_oci_repository
        ```

8. Update go modules, otherwise it failed with missing modules

    ```sh
    go mod tidy
    ```

9. Install required Go tools

    ```bash
    go get golang.org/x/tools/cmd/goimports@latest
    go install golang.org/x/tools/cmd/goimports@latest
    ```

10. Build provider

    (check [Troubleshooting](#make-generate-fails-with-error) in case it fails)

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
        mkdir examples/local_oci_repository
        ```

4. Create K8s resource YAML file

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        cat <<EOF > examples/local_oci_repository/local_oci_repository.yaml
        apiVersion: artifactory.jfrog.upbound.io/v1alpha1
        kind: LocalOCIRepository
        metadata:
          name: crossplane-oci
        spec:
          forProvider:
            key: crossplane-oci-repository
        providerConfigRef:
          name: default
        EOF
        ```

5. Create K8s Secret from template

    ```bash
    $ ARTIFACTORY_URL=https://artifactory.site.com/artifactory
    $ read -r ARTIFACTORY_TOKEN
    <put-your-token-here>
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

    It will block the terminal.

8. Follow commands must be run in different terminal due to block of the first one
9. Deploy `ProviderConfig` to the cluster

    ```bash
    kubectl apply -f examples/providerconfig/
    ```

10. Deploy Artifactory provider resource

    - E.g., for `artifactory_local_oci_repository`

        ```bash
        kubectl apply -f examples/local_oci_repository/
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

Steps to publish provider to Upbound Marketplace according to [these instructions](https://docs.upbound.io/upbound-marketplace/packages/).

1. [Upbound account](https://docs.upbound.io/operate/accounts/identity-management/users/#create-an-account) is required
2. Create API token in Upbound account
3. Create new repository according to provider name, e.g. `provider-artifactory`
4. Set these GitHub secrets in the repository

    - `UPBOUND_MARKETPLACE_PUSH_ROBOT_USR` - Upbound account username, like `homolkao`
    - `UPBOUND_MARKETPLACE_PUSH_ROBOT_PWD` - Upbound account password

5. Update registry domain on all places in `Makefile`, e.g. replace `xpkg.upbound.io/upbound` with `xpkg.upbound.io/homolkao`

## Troubleshooting

### `make generate` fails with error

```text
scraper: error: Failed to scrape Terraform provider metadata: cannot scrape Terraform registry: failed to scrape resource metadata from path: ../.work/jfrog/artifactory/docs/resources/anonymous_user.md: failed to find the prelude of the document using the xpath expressions: //text()[contains(., "description") and contains(., "page_title")]
exit status 1
```

Upjet tool requires `page_title` and `description` fields in each file in Terraform Markdown documentation file and generated docs unfortunately contain them.

#### Solution

This is a workaround which generates valid Markdown files which are unfortunately missing siginificant part of origin files.

I've created an issue with question how it looks with a fix:
    <!-- markdownlint-disable-next-line no-bare-urls -->
    https://github.com/jfrog/terraform-provider-artifactory/issues/1229

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

Can be solved by own `GetExternalNameFn` function, just it to resource configuration in `config/local_oci_repository/config.go`.

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
