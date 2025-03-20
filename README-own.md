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
