# How did I generate it?

Followed steps in [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)

1. Fetch submodule (no change in source code)

    ```bash
    make submodules
    ```

2. Setup provider name and group

    ```bash
    $ ./hack/prepare.sh
    Lower case provider name (ex. github): artifactory
    Normal case provider name (ex. GitHub): Artifactory
    Organization (ex. upbound, my-org-name): hmlkao
    CRD rootGroup (ex. upbound.io, crossplane.io): jfrog.crossplane.io
    ```

3. Configure Terraform provider in `Makefile`

    ```bash
    export TERRAFORM_PROVIDER_SOURCE ?= jfrog/artifactory
    export TERRAFORM_PROVIDER_REPO ?= https://github.com/jfrog/terraform-provider-artifactory
    export TERRAFORM_PROVIDER_VERSION ?= 12.9.1
    export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-artifactory
    export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/jfrog/$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
    export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-artifactory_v12.9.1
    export TERRAFORM_DOCS_PATH ?= docs/resources
    ```

4. Add ProviderConfig logic to `internal/clients/artifactory.go` according to [Terraform provider argument reference](https://registry.terraform.io/providers/jfrog/artifactory/12.9.0/docs#argument-reference)
5. Add external name to `config/external_name.go` file according to Terraform import reference, examples:

   - for `artifactory_local_oci_repository` it's `my-oci-local` which is the same as Terraform resource name

        ```sh
        $ terraform import artifactory_local_oci_repository.my-oci-local my-oci-local
        ```

6. Add custom configuration for Artifactory resource(s) and remove `null` resource

7. `make generate` fails with error

    ```text
    scraper: error: Failed to scrape Terraform provider metadata: cannot scrape Terraform registry: failed to scrape resource metadata from path: ../.work/jfrog/artifactory/docs/resources/anonymous_user.md: failed to find the prelude of the document using the xpath expressions: //text()[contains(., "description") and contains(., "page_title")]
    exit status 1
    ```

    Upjet tool requires `page_title` and `description` fields in each file in Terraform Markdown documentation file and generated docs unfortunately contain them.

    This is a workaround which generates valid Markdown files which are unfortunately missing siginificant part of origin files.

    I've created an issue with question how it looks with a fix:
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

8. Update go modules

    ```sh
    go mod tidy
    ```

9. Run command

    ```sh
    make generate
    ```
