<!-- markdownlint-disable no-hard-tabs -->
# Provider JFrog Artifactory

This document provides an overview and guidance for using the `provider-jfrog-artifactory`, a Crossplane provider for managing Artifactory resources.

`provider-jfrog-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

The repo was created from [crossplane/upjet-provider-template@7311f9f](https://github.com/crossplane/upjet-provider-template/tree/7311f9f9baa87f4431702ba209dffbc6067ce74b) template.

Provider is generated from Terraform provider [jfrog/artifactory v12.9.1](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs).

- [Provider JFrog Artifactory](#provider-jfrog-artifactory)
  - [Getting Started](#getting-started)
  - [Naming convention decision](#naming-convention-decision)
    - [Options](#options)
  - [Supported resources](#supported-resources)
    - [Artifact](#artifact)
    - [Configuration](#configuration)
    - [Federated Repositories](#federated-repositories)
    - [Lifecycle](#lifecycle)
    - [Local Repositories](#local-repositories)
    - [Remote Repositories](#remote-repositories)
    - [Replication](#replication)
    - [Security](#security)
    - [User](#user)
    - [Virtual Repositories](#virtual-repositories)
    - [Webhook](#webhook)
  - [Build provider from scratch](#build-provider-from-scratch)
  - [Developing](#developing)
  - [Report a Bug](#report-a-bug)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-jfrog-artifactory):

```bash
up ctp provider install hmlkao/provider-jfrog-artifactory:v0.2.0
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-artifactory
spec:
  package: hmlkao/provider-jfrog-artifactory:v0.2.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/hmlkao/provider-jfrog-artifactory).

## Naming convention decision

There are more Terraform providers developed by JFrog, e.g.:

- [`jfrog/artifactory`](https://registry.terraform.io/providers/jfrog/artifactory)
- [`jfrog/platform`](https://registry.terraform.io/providers/jfrog/platform)
- [`jfrog/project`](https://registry.terraform.io/providers/jfrog/project)
- etc.

So, I decided to use `artifactory.jfrog.crossplane.io` base group for this Crossplane provider.

### Options

1. (*current naming convention*) `artifactory.jfrog.crossplane.io` as base group and resource is `LocalOCIRepository`

    - :heavy_plus_sign: Get rid of `ShortGroup` config and use `artifactory.jfrog.crossplane.io` for initial config instead
    - Auto generated `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider (by default)
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc.
      - :heavy_minus_sign: All resources must have specified `ShortGroup` as `""` (empty string) to override default behavior

2. `jfrog.crossplane.io` as a base group and set `ShortGroup` as `artifactory` to all resources which will produce `artifactory.jfrog.crossplane.io` and K8s kind is just `LocalOCIRepository`

    - :heavy_minus_sign: All resources must have specified `ShortGroup` as `artifactory`

3. `artifactory.crossplane.io` as a base group and resource prefix to be `artifactory`, like `LocalOCIRepository`

    - :heavy_minus_sign: It's not clear for other JFrog providers, e.g. for `jfrog/platform` would be `platform.crossplane.io`

Not sure, which one is the best.

## Supported resources

List of all resources of [Terraform provider v12.9.1](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs).

### Artifact

| Resource                      | Supported                                                                                  | Kind             |
|-------------------------------|--------------------------------------------------------------------------------------------|------------------|
| `artifactory_artifact`        | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                  |
| `artifactory_item_properties` | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#artifactory_item_properties))         | `ItemProperties` |

### Configuration

| Resource                             | Supported                                              | Kind                           |
|--------------------------------------|--------------------------------------------------------|--------------------------------|
| `artifactory_archive_policy`         | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) |                                |
| `artifactory_backup`                 | :x:                                                    |                                |
| `artifactory_general_security`       | :x:                                                    |                                |
| `artifactory_ldap_group_setting`     | :x:                                                    |                                |
| `artifactory_ldap_group_setting_v2`  | :x:                                                    |                                |
| `artifactory_ldap_setting`           | :x:                                                    |                                |
| `artifactory_ldap_setting_v2`        | :x:                                                    |                                |
| `artifactory_mail_server`            | :x:                                                    |                                |
| `artifactory_oauth_settings`         | :x:                                                    |                                |
| `artifactory_package_cleanup_policy` | :x:                                                    |                                |
| `artifactory_property_set`           | :x:                                                    |                                |
| `artifactory_proxy`                  | :x:                                                    |                                |
| `artifactory_repository_layout`      | :x:                                                    |                                |
| `artifactory_vault_configuration`    | :x:                                                    |                                |

### Federated Repositories

| Resource                                              | Supported          | Kind                           |
|-------------------------------------------------------|--------------------|--------------------------------|
| `artifactory_federated_alpine_repository`             | :x:                |                                |
| `artifactory_federated_ansible_repository`            | :x:                |                                |
| `artifactory_federated_bower_repository`              | :x:                |                                |
| `artifactory_federated_cargo_repository`              | :x:                |                                |
| `artifactory_federated_chef_repository`               | :x:                |                                |
| `artifactory_federated_cocoapods_repository`          | :x:                |                                |
| `artifactory_federated_composer_repository`           | :x:                |                                |
| `artifactory_federated_conan_repository`              | :x:                |                                |
| `artifactory_federated_conda_repository`              | :x:                |                                |
| `artifactory_federated_cran_repository`               | :x:                |                                |
| `artifactory_federated_debian_repository`             | :x:                |                                |
| `artifactory_federated_docker_repository`             | :x:                |                                |
| `artifactory_federated_docker_v1_repository`          | :x:                |                                |
| `artifactory_federated_docker_v2_repository`          | :x:                |                                |
| `artifactory_federated_gems_repository`               | :x:                |                                |
| `artifactory_federated_generic_repository`            | :x:                |                                |
| `artifactory_federated_gitlfs_repository`             | :x:                |                                |
| `artifactory_federated_go_repository`                 | :x:                |                                |
| `artifactory_federated_gradle_repository`             | :x:                |                                |
| `artifactory_federated_helm_repository`               | :x:                |                                |
| `artifactory_federated_helmoci_repository`            | :x:                |                                |
| `artifactory_federated_huggingfaceml_repository`      | :x:                |                                |
| `artifactory_federated_ivy_repository`                | :x:                |                                |
| `artifactory_federated_maven_repository`              | :x:                |                                |
| `artifactory_federated_npm_repository`                | :x:                |                                |
| `artifactory_federated_nuget_repository`              | :x:                |                                |
| `artifactory_federated_oci_repository`                | :x:                |                                |
| `artifactory_federated_opkg_repository`               | :x:                |                                |
| `artifactory_federated_puppet_repository`             | :x:                |                                |
| `artifactory_federated_pypi_repository`               | :x:                |                                |
| `artifactory_federated_rpm_repository`                | :x:                |                                |
| `artifactory_federated_sbt_repository`                | :x:                |                                |
| `artifactory_federated_swift_repository`              | :x:                |                                |
| `artifactory_federated_terraform_module_repository`   | :x:                |                                |
| `artifactory_federated_terraform_provider_repository` | :x:                |                                |
| `artifactory_federated_vagrant_repository`            | :x:                |                                |

### Lifecycle

| Resource                                  | Supported          | Kind                           |
|-------------------------------------------|--------------------|--------------------------------|
| `artifactory_release_bundle_v2`           | :x:                |                                |
| `artifactory_release_bundle_v2_promotion` | :x:                |                                |

### Local Repositories

| Resource                                          | Supported          | Kind                               |
|---------------------------------------------------|--------------------|------------------------------------|
| `artifactory_local_alpine_repository`             | :x:                |                                    |
| `artifactory_local_ansible_repository`            | :heavy_check_mark: | `LocalAnsibleRepository`           |
| `artifactory_local_bower_repository`              | :x:                |                                    |
| `artifactory_local_cargo_repository`              | :x:                |                                    |
| `artifactory_local_chef_repository`               | :x:                |                                    |
| `artifactory_local_cocoapods_repository`          | :x:                |                                    |
| `artifactory_local_composer_repository`           | :x:                |                                    |
| `artifactory_local_conan_repository`              | :x:                |                                    |
| `artifactory_local_conda_repository`              | :x:                |                                    |
| `artifactory_local_cran_repository`               | :x:                |                                    |
| `artifactory_local_debian_repository`             | :x:                |                                    |
| `artifactory_local_docker_v1_repository`          | :x:                |                                    |
| `artifactory_local_docker_v2_repository`          | :heavy_check_mark: | `LocalDockerV2Repository`          |
| `artifactory_local_gems_repository`               | :x:                |                                    |
| `artifactory_local_generic_repository`            | :heavy_check_mark: | `LocalGenericRepository`           |
| `artifactory_local_gitlfs_repository`             | :x:                |                                    |
| `artifactory_local_go_repository`                 | :x:                |                                    |
| `artifactory_local_gradle_repository`             | :x:                |                                    |
| `artifactory_local_helm_repository`               | :x:                |                                    |
| `artifactory_local_helmoci_repository`            | :x:                |                                    |
| `artifactory_local_huggingfaceml_repository`      | :x:                |                                    |
| `artifactory_local_ivy_repository`                | :x:                |                                    |
| `artifactory_local_machinelearning_repository`    | :x:                |                                    |
| `artifactory_local_maven_repository`              | :x:                |                                    |
| `artifactory_local_npm_repository`                | :x:                |                                    |
| `artifactory_local_nuget_repository`              | :x:                |                                    |
| `artifactory_local_oci_repository`                | :heavy_check_mark: | `LocalOCIRepository`               |
| `artifactory_local_opkg_repository`               | :x:                |                                    |
| `artifactory_local_pub_repository`                | :x:                |                                    |
| `artifactory_local_puppet_repository`             | :x:                |                                    |
| `artifactory_local_pypi_repository`               | :x:                |                                    |
| `artifactory_local_rpm_repository`                | :x:                |                                    |
| `artifactory_local_sbt_repository`                | :x:                |                                    |
| `artifactory_local_swift_repository`              | :x:                |                                    |
| `artifactory_local_terraform_module_repository`   | :x:                |                                    |
| `artifactory_local_terraform_provider_repository` | :heavy_check_mark: | `LocalTerraformProviderRepository` |
| `artifactory_local_terraformbackend_repository`   | :x:                |                                    |
| `artifactory_local_vagrant_repository`            | :x:                |                                    |

### Remote Repositories

| Resource                                           | Supported          | Kind                           |
|----------------------------------------------------|--------------------|--------------------------------|
| `artifactory_remote_alpine_repository`             | :x:                |                                |
| `artifactory_remote_ansible_repository`            | :x:                |                                |
| `artifactory_remote_bower_repository`              | :x:                |                                |
| `artifactory_remote_cargo_repository`              | :x:                |                                |
| `artifactory_remote_chef_repository`               | :x:                |                                |
| `artifactory_remote_cocoapods_repository`          | :x:                |                                |
| `artifactory_remote_composer_repository`           | :x:                |                                |
| `artifactory_remote_conan_repository`              | :x:                |                                |
| `artifactory_remote_conda_repository`              | :x:                |                                |
| `artifactory_remote_cran_repository`               | :x:                |                                |
| `artifactory_remote_debian_repository`             | :x:                |                                |
| `artifactory_remote_docker_repository`             | :x:                |                                |
| `artifactory_remote_gems_repository`               | :x:                |                                |
| `artifactory_remote_generic_repository`            | :x:                |                                |
| `artifactory_remote_gitlfs_repository`             | :x:                |                                |
| `artifactory_remote_go_repository`                 | :x:                |                                |
| `artifactory_remote_gradle_repository`             | :x:                |                                |
| `artifactory_remote_helm_repository`               | :x:                |                                |
| `artifactory_remote_helmoci_repository`            | :x:                |                                |
| `artifactory_remote_huggingfaceml_repository`      | :x:                |                                |
| `artifactory_remote_ivy_repository`                | :x:                |                                |
| `artifactory_remote_maven_repository`              | :x:                |                                |
| `artifactory_remote_npm_repository`                | :x:                |                                |
| `artifactory_remote_nuget_repository`              | :x:                |                                |
| `artifactory_remote_oci_repository`                | :x:                |                                |
| `artifactory_remote_opkg_repository`               | :x:                |                                |
| `artifactory_remote_p2_repository`                 | :x:                |                                |
| `artifactory_remote_pub_repository`                | :x:                |                                |
| `artifactory_remote_puppet_repository`             | :x:                |                                |
| `artifactory_remote_pypi_repository`               | :x:                |                                |
| `artifactory_remote_rpm_repository`                | :x:                |                                |
| `artifactory_remote_sbt_repository`                | :x:                |                                |
| `artifactory_remote_swift_repository`              | :x:                |                                |
| `artifactory_remote_terraform_repository`          | :x:                |                                |
| `artifactory_remote_terraform_provider_repository` | :x:                |                                |
| `artifactory_remote_vcs_repository`                | :x:                |                                |

### Replication

| Resource                                          | Supported          | Kind                           |
|---------------------------------------------------|--------------------|--------------------------------|
| `artifactory_local_repository_multi_replication`  | :x:                |                                |
| `artifactory_local_repository_single_replication` | :x:                |                                |
| `artifactory_remote_repository_replication`       | :x:                |                                |

### Security

| Resource                                 | Supported                                                      | Kind                           |
|------------------------------------------|----------------------------------------------------------------|--------------------------------|
| `artifactory_certificate`                | :x:                                                            |                                |
| `artifactory_distribution_public_key`    | :x:                                                            |                                |
| `artifactory_global_environment`         | :x:                                                            |                                |
| `artifactory_keypair`                    | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#keypair)) | `Keypair`                      |
| `artifactory_password_expiration_policy` | :x:                                                            |                                |
| `artifactory_scoped_token`               | :x:                                                            |                                |
| `artifactory_user_lock_policy`           | :x:                                                            |                                |

### User

| Resource                     | Supported                                              | Kind            |
|------------------------------|--------------------------------------------------------|-----------------|
| `artifactory_anonymous_user` | :heavy_check_mark:                                     | `AnonymousUser` |
| `artifactory_managed_user`   | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) |                 |
| `artifactory_unmanaged_user` | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) |                 |
| `artifactory_user`           | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) |                 |

### Virtual Repositories

| Resource                                   | Supported          | Kind                           |
|--------------------------------------------|--------------------|--------------------------------|
| `artifactory_virtual_alpine_repository`    | :x:                |                                |
| `artifactory_virtual_ansible_repository`   | :x:                |                                |
| `artifactory_virtual_bower_repository`     | :x:                |                                |
| `artifactory_virtual_chef_repository`      | :x:                |                                |
| `artifactory_virtual_cocoapods_repository` | :x:                |                                |
| `artifactory_virtual_composer_repository`  | :x:                |                                |
| `artifactory_virtual_conan_repository`     | :x:                |                                |
| `artifactory_virtual_conda_repository`     | :x:                |                                |
| `artifactory_virtual_cran_repository`      | :x:                |                                |
| `artifactory_virtual_debian_repository`    | :x:                |                                |
| `artifactory_virtual_docker_repository`    | :x:                |                                |
| `artifactory_virtual_gems_repository`      | :x:                |                                |
| `artifactory_virtual_generic_repository`   | :x:                |                                |
| `artifactory_virtual_gitlfs_repository`    | :x:                |                                |
| `artifactory_virtual_go_repository`        | :x:                |                                |
| `artifactory_virtual_gradle_repository`    | :x:                |                                |
| `artifactory_virtual_helm_repository`      | :x:                |                                |
| `artifactory_virtual_helmoci_repository`   | :x:                |                                |
| `artifactory_virtual_ivy_repository`       | :x:                |                                |
| `artifactory_virtual_maven_repository`     | :x:                |                                |
| `artifactory_virtual_npm_repository`       | :x:                |                                |
| `artifactory_virtual_nuget_repository`     | :x:                |                                |
| `artifactory_virtual_oci_repository`       | :x:                |                                |
| `artifactory_virtual_p2_repository`        | :x:                |                                |
| `artifactory_virtual_pub_repository`       | :x:                |                                |
| `artifactory_virtual_puppet_repository`    | :x:                |                                |
| `artifactory_virtual_pypi_repository`      | :x:                |                                |
| `artifactory_virtual_rpm_repository`       | :x:                |                                |
| `artifactory_virtual_sbt_repository`       | :x:                |                                |
| `artifactory_virtual_swift_repository`     | :x:                |                                |
| `artifactory_virtual_terraform_repository` | :x:                |                                |

### Webhook

| Resource                                                 | Supported          | Kind                           |
|----------------------------------------------------------|--------------------|--------------------------------|
| `artifactory_artifact_custom_webhook`                    | :x:                |                                |
| `artifactory_artifact_lifecycle_custom_webhook`          | :x:                |                                |
| `artifactory_artifact_lifecycle_webhook`                 | :x:                |                                |
| `artifactory_artifact_property_custom_webhook`           | :x:                |                                |
| `artifactory_artifact_property_webhook`                  | :x:                |                                |
| `artifactory_artifact_webhook`                           | :x:                |                                |
| `artifactory_artifactory_release_bundle_custom_webhook`  | :x:                |                                |
| `artifactory_artifactory_release_bundle_webhook`         | :x:                |                                |
| `artifactory_build_custom_webhook`                       | :x:                |                                |
| `artifactory_build_webhook`                              | :x:                |                                |
| `artifactory_destination_custom_webhook`                 | :x:                |                                |
| `artifactory_destination_webhook`                        | :x:                |                                |
| `artifactory_distribution_custom_webhook`                | :x:                |                                |
| `artifactory_distribution_webhook`                       | :x:                |                                |
| `artifactory_local_docker_v2_repository`                 | :x:                |                                |
| `artifactory_docker_webhook`                             | :x:                |                                |
| `artifactory_release_bundle_v2_custom_webhook`           | :x:                |                                |
| `artifactory_release_bundle_v2_promotion_custom_webhook` | :x:                |                                |
| `artifactory_release_bundle_v2_promotion_webhook`        | :x:                |                                |
| `artifactory_release_bundle_v2_webhook`                  | :x:                |                                |
| `artifactory_user_custom_webhook`                        | :x:                |                                |
| `artifactory_user_webhook`                               | :x:                |                                |

## Build provider from scratch

Check [`BUILD_FROM_SCRATCH.md`](./BUILD_FROM_SCRATCH.md) for notes how was this provider built using [crossplane/upjet tool](https://github.com/crossplane/upjet) step-by-step.

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/hmlkao/provider-jfrog-artifactory/issues).
