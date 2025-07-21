#!/usr/bin/env bash
set -euox pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" # Directory of this script
rootdir="$( cd "${dir}/.." && pwd )" # Root directory of the repository
workdir="${rootdir}/.work/terraform-provider-artifactory"

# Clean up
cleanup() {
  echo "Cleaning up..."
  #rm -rf "${workdir}"
}

### Just copy-paste from BUILD_FROM_SCRATCH.md#make-generate-fails-with-error file
generate_docs() {
  # Clone repo of the Terraform provider and change directory into it
  if [[ ! -d "${workdir}" ]]; then
    git clone https://github.com/jfrog/terraform-provider-artifactory.git "${workdir}"
  fi
  {
    cd "${workdir}"

    # Checkout to the same version as you specify in [`Makefile`](./Makefile) variable `TERRAFORM_PROVIDER_VERSION`
    current_version=$(grep 'export TERRAFORM_PROVIDER_VERSION' < "${rootdir}/Makefile" | cut -d' ' -f4)
    git checkout "v${current_version}"

    # Check version of `github.com/hashicorp/terraform-plugin-docs` in `go.mod` file
    grep 'hashicorp/terraform-plugin-docs' go.mod

    # (*optional*) Install [`hashicorp/terraform-plugin-docs`](https://github.com/hashicorp/terraform-plugin-docs?tab=readme-ov-file#installation) CLI tool
    go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.21.0
    go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.21.0

    # Verify that it is possible to find binary
    command -v tfplugindocs

    # Generate docs
    make doc

    # Copy generated doc to the path of Terraform provider in Crossplane provider
    rm -rf "${rootdir}/.work/jfrog/artifactory/docs/"
    cp -R docs "${rootdir}/.work/jfrog/artifactory/"
    ls -la "${rootdir}/.work/jfrog/artifactory/docs/"

    # Fix missing description in the generated docs which was not fixed by make doc command
    # Use -i.bak for compatibility with macOS and GNU sed
    sed -i.bak 's/subcategory: "User"/subcategory: "User"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/anonymous_user.md"
    sed -i.bak 's/subcategory: "Security"/subcategory: "Security"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/global_environment.md"
    sed -i.bak 's/subcategory: "Security"/subcategory: "Security"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/group.md"
    sed -i.bak 's/subcategory: "Configuration"/subcategory: "Configuration"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/ldap_group_setting_v2.md"
    sed -i.bak 's/subcategory: "Configuration"/subcategory: "Configuration"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/ldap_setting_v2.md"
    sed -i.bak 's/subcategory: "Configuration"/subcategory: "Configuration"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/mail_server.md"
    sed -i.bak 's/subcategory: "User"/subcategory: "User"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/managed_user.md"
    sed -i.bak 's/subcategory: "Security"/subcategory: "Security"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/scoped_token.md"
    sed -i.bak 's/subcategory: "User"/subcategory: "User"\ndescription: ""/' "${rootdir}/.work/jfrog/artifactory/docs/resources/user.md"

    # Clean up backup files created by sed
    find "${rootdir}/.work/jfrog/artifactory/docs/resources/" -type f -name "*.md.bak" -exec rm -f {} \;
  }
}

main() {
  trap cleanup EXIT
  echo "Generating Terraform provider docs..."
  generate_docs
}

main "$@"
