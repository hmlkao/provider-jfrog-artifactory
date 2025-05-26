#!/usr/bin/env bash

# Bump version of Terraform provider

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_version=${1}
if [[ -z "$current_version" ]]; then
  echo "::error Current version is empty"
  exit 1
fi
new_version=${2}
if [[ -z "$new_version" ]]; then
  echo "::error New version is empty"
  exit 1
fi

main() {
  sed -i "s|export TERRAFORM_PROVIDER_VERSION ?= ${current_version}|export TERRAFORM_PROVIDER_VERSION ?= ${new_version}|" "$dir/../Makefile"
  sed -i "s|terraform-provider-artifactory_v${current_version}|terraform-provider-artifactory_v${new_version}|" "$dir/../Makefile"
  sed -i "s|jfrog/artifactory v${current_version}|jfrog/artifactory v${new_version}|g" "$dir/../README.md"
  sed -i "s|Terraform provider v${current_version}|Terraform provider v${new_version}|g" "$dir/../README.md"
  sed -i "s|jfrog/artifactory/${current_version}|jfrog/artifactory/${new_version}|" "$dir/../README.md"
}

main "$@"
