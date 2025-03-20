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
    export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://releases.hashicorp.com/$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)/$(TERRAFORM_PROVIDER_VERSION)
    export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-artifactory_v12.9.1
    export TERRAFORM_DOCS_PATH ?= docs/resources
    ```
