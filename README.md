<!-- markdownlint-disable no-hard-tabs -->
# Provider JFrog Artifactory

This document provides an overview and guidance for using the `provider-jfrog-artifactory`, a Crossplane provider for managing Artifactory resources.

`provider-jfrog-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

The repo was created from [crossplane/upjet-provider-template@7311f9f](https://github.com/crossplane/upjet-provider-template/tree/7311f9f9baa87f4431702ba209dffbc6067ce74b) template.

Provider is generated from Terraform provider [jfrog/artifactory v12.10.1](https://registry.terraform.io/providers/jfrog/artifactory/12.10.1/docs).

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
  - [JFrog Artifactory icon](#jfrog-artifactory-icon)
  - [Report a Bug](#report-a-bug)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-jfrog-artifactory):

```bash
up ctp provider install hmlkao/provider-jfrog-artifactory:v0.9.1
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-jfrog-artifactory
spec:
  package: hmlkao/provider-jfrog-artifactory:v0.9.1
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference at [doc.crds.dev](https://doc.crds.dev/github.com/hmlkao/provider-jfrog-artifactory).

## Naming convention decision

There are more Terraform providers developed by JFrog, e.g.:

- [`jfrog/artifactory`](https://registry.terraform.io/providers/jfrog/artifactory)
- [`jfrog/platform`](https://registry.terraform.io/providers/jfrog/platform)
- [`jfrog/project`](https://registry.terraform.io/providers/jfrog/project)
- etc.

So, I decided to use `artifactory.jfrog.crossplane.io` base group for this Crossplane provider.

### Options

1. (*current naming convention*) `artifactory.jfrog.crossplane.io` as base group and resource is `OCIRepository` for each short group (`local`, `remote`, `virtual`, `federated`)

    - :heavy_plus_sign: `ShortGroup` doesn't have to be set for all resources, just for specific resources
    - Auto generated `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider (by default)
      - :heavy_plus_sign: Crossplane v2 introduced [managed resource definitions (MRDs)](https://docs.crossplane.io/latest/managed-resources/managed-resource-definitions/) for selective activation of provider resources, reducing [cluster overhead](https://docs.crossplane.io/v2.0/guides/disabling-unused-managed-resources/#the-problem-resource-overhead). It allows to activate only specific CRDs, so it's helpful to have separated groups like `local.artifactory.jfrog.crossplane.io` or `virtual.artifactory.jfrog.crossplane.io` to be able to activate only some of them in case users don't need them all.
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc. (all resources must have specified `ShortGroup` as `""` (empty string) to override default behavior and have `AlpineRepository` with `artifactory.jfrog.crossplane.io` group)

2. `artifactory.jfrog.crossplane.io` as base group and resource is `OCIRepository`

    - :heavy_plus_sign: `ShortGroup` doesn't have to be set for all resources
    - :heavy_minus_sign: Still need to specify `ShortGroup` for the most of resources
    - Auto generated `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider (by default)
      - :heavy_plus_sign: Crossplane v2 introduced [managed resource definitions (MRDs)](https://docs.crossplane.io/latest/managed-resources/managed-resource-definitions/) for selective activation of provider resources, reducing [cluster overhead](https://docs.crossplane.io/v2.0/guides/disabling-unused-managed-resources/#the-problem-resource-overhead). It allows to activate only specific CRDs, so it's helpful to have separated groups like `local.artifactory.jfrog.crossplane.io` or `virtual.artifactory.jfrog.crossplane.io` to be able to activate only some of them in case users don't need them all.
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc. (all resources must have specified `ShortGroup` as `""` (empty string) to override default behavior and have `AlpineRepository` with `artifactory.jfrog.crossplane.io` group)

3. `jfrog.crossplane.io` as a base group and set `ShortGroup` as `artifactory` to all resources which will produce `artifactory.jfrog.crossplane.io` and K8s kind is just `OCIRepository`

    - :heavy_minus_sign: All resources must have specified `ShortGroup` as `artifactory`

4. `artifactory.crossplane.io` as a base group and resource prefix to be `artifactory`, like `OCIRepository`

    - :heavy_minus_sign: It's not clear for other JFrog providers, e.g. for `jfrog/platform` would be `platform.crossplane.io`

Not sure, which one is the best.

## Supported resources

List of all resources of [Terraform provider v12.10.1](https://registry.terraform.io/providers/jfrog/artifactory/12.10.1/docs).

### Artifact

Short group is `artifact`, so the `apiGroup` is:

- `artifact.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `artifact.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                      | Supported                                                                                  | Kind             |
|-------------------------------|--------------------------------------------------------------------------------------------|------------------|
| `artifactory_artifact`        | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                  |
| `artifactory_item_properties` | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#artifactory_item_properties))         | `ItemProperties` |

### Configuration

Short group is `configuration`, so the `apiGroup` is:

- `configuration.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `configuration.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                             | Supported                                              | Kind                           |
|--------------------------------------|--------------------------------------------------------|--------------------------------|
| `artifactory_archive_policy`         | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) |                                |
| `artifactory_backup`                 | :heavy_check_mark:                                     | `Backup`                       |
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

Short group is `federated`, so the `apiGroup` is:

- `federated.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `federated.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                              | Supported          | Kind                          |
| ----------------------------------------------------- | ------------------ | ----------------------------- |
| `artifactory_federated_alpine_repository`             | :heavy_check_mark: | `AlpineRepository`            |
| `artifactory_federated_ansible_repository`            | :heavy_check_mark: | `AnsibleRepository`           |
| `artifactory_federated_bower_repository`              | :heavy_check_mark: | `BowerRepository`             |
| `artifactory_federated_cargo_repository`              | :heavy_check_mark: | `CargoRepository`             |
| `artifactory_federated_chef_repository`               | :heavy_check_mark: | `ChefRepository`              |
| `artifactory_federated_cocoapods_repository`          | :heavy_check_mark: | `CocoaPodsRepository`         |
| `artifactory_federated_composer_repository`           | :heavy_check_mark: | `ComposerRepository`          |
| `artifactory_federated_conan_repository`              | :heavy_check_mark: | `ConanRepository`             |
| `artifactory_federated_conda_repository`              | :heavy_check_mark: | `CondaRepository`             |
| `artifactory_federated_cran_repository`               | :heavy_check_mark: | `CRANRepository`              |
| `artifactory_federated_debian_repository`             | :heavy_check_mark: | `DebianRepository`            |
| `artifactory_federated_docker_repository`             | :heavy_check_mark: | `DockerRepository`            |
| `artifactory_federated_docker_v1_repository`          | :heavy_check_mark: | `DockerV1Repository`          |
| `artifactory_federated_docker_v2_repository`          | :heavy_check_mark: | `DockerV2Repository`          |
| `artifactory_federated_gems_repository`               | :heavy_check_mark: | `GemsRepository`              |
| `artifactory_federated_generic_repository`            | :heavy_check_mark: | `GenericRepository`           |
| `artifactory_federated_gitlfs_repository`             | :heavy_check_mark: | `GitLFSRepository`            |
| `artifactory_federated_go_repository`                 | :heavy_check_mark: | `GoRepository`                |
| `artifactory_federated_gradle_repository`             | :heavy_check_mark: | `GradleRepository`            |
| `artifactory_federated_helm_repository`               | :heavy_check_mark: | `HelmRepository`              |
| `artifactory_federated_helmoci_repository`            | :heavy_check_mark: | `HelmOCIRepository`           |
| `artifactory_federated_huggingfaceml_repository`      | :heavy_check_mark: | `HuggingFaceMLRepository`     |
| `artifactory_federated_ivy_repository`                | :heavy_check_mark: | `IvyRepository`               |
| `artifactory_federated_maven_repository`              | :heavy_check_mark: | `MavenRepository`             |
| `artifactory_federated_npm_repository`                | :heavy_check_mark: | `NPMRepository`               |
| `artifactory_federated_nuget_repository`              | :heavy_check_mark: | `NuGetRepository`             |
| `artifactory_federated_oci_repository`                | :heavy_check_mark: | `OCIRepository`               |
| `artifactory_federated_opkg_repository`               | :heavy_check_mark: | `OpkgRepository`              |
| `artifactory_federated_puppet_repository`             | :heavy_check_mark: | `PuppetRepository`            |
| `artifactory_federated_pypi_repository`               | :heavy_check_mark: | `PypiRepository`              |
| `artifactory_federated_rpm_repository`                | :heavy_check_mark: | `RPMRepository`               |
| `artifactory_federated_sbt_repository`                | :heavy_check_mark: | `SbtRepository`               |
| `artifactory_federated_swift_repository`              | :heavy_check_mark: | `SwiftRepository`             |
| `artifactory_federated_terraform_module_repository`   | :heavy_check_mark: | `TerraformModuleRepository`   |
| `artifactory_federated_terraform_provider_repository` | :heavy_check_mark: | `TerraformProviderRepository` |
| `artifactory_federated_vagrant_repository`            | :heavy_check_mark: | `VagrantRepository`           |

### Lifecycle

Short group is `lifecycle`, so the `apiGroup` is:

- `lifecycle.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `lifecycle.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                  | Supported          | Kind                           |
|-------------------------------------------|--------------------|--------------------------------|
| `artifactory_release_bundle_v2`           | :x:                |                                |
| `artifactory_release_bundle_v2_promotion` | :x:                |                                |

### Local Repositories

Short group is `local`, so the `apiGroup` is

- `local.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `local.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                          | Supported          | Kind                         |
|---------------------------------------------------|--------------------|------------------------------|
| `artifactory_local_alpine_repository`             | :heavy_check_mark: | `AlpineRepository`           |
| `artifactory_local_ansible_repository`            | :heavy_check_mark: | `AnsibleRepository`          |
| `artifactory_local_bower_repository`              | :heavy_check_mark: | `BowerRepository`            |
| `artifactory_local_cargo_repository`              | :heavy_check_mark: | `CargoRepository`            |
| `artifactory_local_chef_repository`               | :heavy_check_mark: | `ChefRepository`             |
| `artifactory_local_cocoapods_repository`          | :heavy_check_mark: | `CocoaPodsRepository`        |
| `artifactory_local_composer_repository`           | :heavy_check_mark: | `ComposerRepository`         |
| `artifactory_local_conan_repository`              | :heavy_check_mark: | `ConanRepository`            |
| `artifactory_local_conda_repository`              | :heavy_check_mark: | `CondaRepository`            |
| `artifactory_local_cran_repository`               | :heavy_check_mark: | `CRANRepository`             |
| `artifactory_local_debian_repository`             | :heavy_check_mark: | `DebianRepository`           |
| `artifactory_local_docker_v1_repository`          | :heavy_check_mark: | `DockerV1Repository`         |
| `artifactory_local_docker_v2_repository`          | :heavy_check_mark: | `DockerV2Repository`         |
| `artifactory_local_gems_repository`               | :heavy_check_mark: | `GemsRepository`             |
| `artifactory_local_generic_repository`            | :heavy_check_mark: | `GenericRepository`          |
| `artifactory_local_gitlfs_repository`             | :heavy_check_mark: | `GitLFSRepository`           |
| `artifactory_local_go_repository`                 | :heavy_check_mark: | `GoRepository`               |
| `artifactory_local_gradle_repository`             | :heavy_check_mark: | `GradleRepository`           |
| `artifactory_local_helm_repository`               | :heavy_check_mark: | `HelmRepository`             |
| `artifactory_local_helmoci_repository`            | :heavy_check_mark: | `HelmOCIRepository`          |
| `artifactory_local_huggingfaceml_repository`      | :heavy_check_mark: | `HuggingFaceMLRepository`    |
| `artifactory_local_ivy_repository`                | :heavy_check_mark: | `IvyRepository`              |
| `artifactory_local_machinelearning_repository`    | :heavy_check_mark: | `MachinelearningRepository`  |
| `artifactory_local_maven_repository`              | :heavy_check_mark: | `MavenRepository`            |
| `artifactory_local_npm_repository`                | :heavy_check_mark: | `NPMRepository`              |
| `artifactory_local_nuget_repository`              | :heavy_check_mark: | `NuGetRepository`            |
| `artifactory_local_oci_repository`                | :heavy_check_mark: | `OCIRepository`              |
| `artifactory_local_opkg_repository`               | :heavy_check_mark: | `OPKGRepository`             |
| `artifactory_local_pub_repository`                | :heavy_check_mark: | `PubRepository`              |
| `artifactory_local_puppet_repository`             | :heavy_check_mark: | `PuppetRepository`           |
| `artifactory_local_pypi_repository`               | :heavy_check_mark: | `PyPIRepository`             |
| `artifactory_local_rpm_repository`                | :heavy_check_mark: | `RPMRepository`              |
| `artifactory_local_sbt_repository`                | :heavy_check_mark: | `SBTRepository`              |
| `artifactory_local_swift_repository`              | :heavy_check_mark: | `SwiftRepository`            |
| `artifactory_local_terraform_module_repository`   | :heavy_check_mark: | `TerraformModuleRepository`  |
| `artifactory_local_terraformbackend_repository`   | :heavy_check_mark: | `TerraformbackendRepository` |
| `artifactory_local_vagrant_repository`            | :heavy_check_mark: | `VagrantRepository`          |

### Remote Repositories

Short group is `remote`, so the `apiGroup` is:

- `remote.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `remote.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                           | Supported          | Kind                          |
|----------------------------------------------------|--------------------|-------------------------------|
| `artifactory_remote_alpine_repository`             | :heavy_check_mark: | `AlpineRepository`            |
| `artifactory_remote_ansible_repository`            | :heavy_check_mark: | `AnsibleRepository`           |
| `artifactory_remote_bower_repository`              | :heavy_check_mark: | `BowerRepository`             |
| `artifactory_remote_cargo_repository`              | :heavy_check_mark: | `CargoRepository`             |
| `artifactory_remote_chef_repository`               | :heavy_check_mark: | `ChefRepository`              |
| `artifactory_remote_cocoapods_repository`          | :heavy_check_mark: | `CocoaPodsRepository`         |
| `artifactory_remote_composer_repository`           | :heavy_check_mark: | `ComposerRepository`          |
| `artifactory_remote_conan_repository`              | :heavy_check_mark: | `ConanRepository`             |
| `artifactory_remote_conda_repository`              | :heavy_check_mark: | `CondaRepository`             |
| `artifactory_remote_cran_repository`               | :heavy_check_mark: | `CRANRepository`              |
| `artifactory_remote_debian_repository`             | :heavy_check_mark: | `DebianRepository`            |
| `artifactory_remote_docker_repository`             | :heavy_check_mark: | `DockerRepository`            |
| `artifactory_remote_gems_repository`               | :heavy_check_mark: | `GemsRepository`              |
| `artifactory_remote_generic_repository`            | :heavy_check_mark: | `GenericRepository`           |
| `artifactory_remote_gitlfs_repository`             | :heavy_check_mark: | `GitLFSRepository`            |
| `artifactory_remote_go_repository`                 | :heavy_check_mark: | `GoRepository`                |
| `artifactory_remote_gradle_repository`             | :heavy_check_mark: | `GradleRepository`            |
| `artifactory_remote_helm_repository`               | :heavy_check_mark: | `HelmRepository`              |
| `artifactory_remote_helmoci_repository`            | :heavy_check_mark: | `HelmOCIRepository`           |
| `artifactory_remote_huggingfaceml_repository`      | :heavy_check_mark: | `HuggingFaceMLRepository`     |
| `artifactory_remote_ivy_repository`                | :heavy_check_mark: | `IvyRepository`               |
| `artifactory_remote_maven_repository`              | :heavy_check_mark: | `MavenRepository`             |
| `artifactory_remote_npm_repository`                | :heavy_check_mark: | `NPMRepository`               |
| `artifactory_remote_nuget_repository`              | :heavy_check_mark: | `NuGetRepository`             |
| `artifactory_remote_oci_repository`                | :heavy_check_mark: | `OCIRepository`               |
| `artifactory_remote_opkg_repository`               | :heavy_check_mark: | `OPKGRepository`              |
| `artifactory_remote_p2_repository`                 | :heavy_check_mark: | `P2Repository`                |
| `artifactory_remote_pub_repository`                | :heavy_check_mark: | `PubRepository`               |
| `artifactory_remote_puppet_repository`             | :heavy_check_mark: | `PuppetRepository`            |
| `artifactory_remote_pypi_repository`               | :heavy_check_mark: | `PyPIRepository`              |
| `artifactory_remote_rpm_repository`                | :heavy_check_mark: | `RPMRepository`               |
| `artifactory_remote_sbt_repository`                | :heavy_check_mark: | `SBTRepository`               |
| `artifactory_remote_swift_repository`              | :heavy_check_mark: | `SwiftRepository`             |
| `artifactory_remote_terraform_repository`          | :heavy_check_mark: | `TerraformRepository`         |
| `artifactory_remote_terraform_provider_repository` | :heavy_check_mark: | `TerraformProviderRepository` |
| `artifactory_remote_vcs_repository`                | :heavy_check_mark: | `VCSRepository`               |

### Replication

Short group is `replication`, so the `apiGroup` is :

- `replication.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `replication.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                          | Supported          | Kind                           |
|---------------------------------------------------|--------------------|--------------------------------|
| `artifactory_local_repository_multi_replication`  | :x:                |                                |
| `artifactory_local_repository_single_replication` | :x:                |                                |
| `artifactory_remote_repository_replication`       | :x:                |                                |

### Security

Short group is `security`, so the `apiGroup` is:

- `security.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `security.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                 | Supported                                                      | Kind                           |
|------------------------------------------|----------------------------------------------------------------|--------------------------------|
| `artifactory_certificate`                | :x:                                                            |                                |
| `artifactory_distribution_public_key`    | :x:                                                            |                                |
| `artifactory_global_environment`         | :x:                                                            |                                |
| `artifactory_keypair`                    | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#keypair)) | `Keypair`                      |
| `artifactory_password_expiration_policy` | :x:                                                            |                                |
| `artifactory_scoped_token`               | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#tokens))  | `ScopedToken`                  |
| `artifactory_user_lock_policy`           | :x:                                                            |                                |

### User

Short group is `user`, so the `apiGroup` is:

- `user.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `user.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                     | Supported                                                             | Kind            |
|------------------------------|-----------------------------------------------------------------------|-----------------|
| `artifactory_anonymous_user` | :heavy_check_mark:                                                    | `AnonymousUser` |
| `artifactory_managed_user`   | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                |                 |
| `artifactory_unmanaged_user` | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                |                 |
| `artifactory_user`           | :heavy_check_mark: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema)) | `User`          |

### Virtual Repositories

Short group is `virtual`, so the `apiGroup` is:

- `virtual.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `virtual.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

| Resource                                   | Supported          | Kind                  |
|--------------------------------------------|--------------------|-----------------------|
| `artifactory_virtual_alpine_repository`    | :heavy_check_mark: | `AlpineRepository`    |
| `artifactory_virtual_ansible_repository`   | :heavy_check_mark: | `AnsibleRepository`   |
| `artifactory_virtual_bower_repository`     | :heavy_check_mark: | `BowerRepository`     |
| `artifactory_virtual_chef_repository`      | :heavy_check_mark: | `ChefRepository`      |
| `artifactory_virtual_cocoapods_repository` | :heavy_check_mark: | `CocoaPodsRepository` |
| `artifactory_virtual_composer_repository`  | :heavy_check_mark: | `ComposerRepository`  |
| `artifactory_virtual_conan_repository`     | :heavy_check_mark: | `ConanRepository`     |
| `artifactory_virtual_conda_repository`     | :heavy_check_mark: | `CondaRepository`     |
| `artifactory_virtual_cran_repository`      | :heavy_check_mark: | `CRANRepository`      |
| `artifactory_virtual_debian_repository`    | :heavy_check_mark: | `DebianRepository`    |
| `artifactory_virtual_docker_repository`    | :heavy_check_mark: | `DockerRepository`    |
| `artifactory_virtual_gems_repository`      | :heavy_check_mark: | `GemsRepository`      |
| `artifactory_virtual_generic_repository`   | :heavy_check_mark: | `GenericRepository`   |
| `artifactory_virtual_gitlfs_repository`    | :heavy_check_mark: | `GitLFSRepository`    |
| `artifactory_virtual_go_repository`        | :heavy_check_mark: | `GoRepository`        |
| `artifactory_virtual_gradle_repository`    | :heavy_check_mark: | `GradleRepository`    |
| `artifactory_virtual_helm_repository`      | :heavy_check_mark: | `HelmRepository`      |
| `artifactory_virtual_helmoci_repository`   | :heavy_check_mark: | `HelmOCIRepository`   |
| `artifactory_virtual_ivy_repository`       | :heavy_check_mark: | `IvyRepository`       |
| `artifactory_virtual_maven_repository`     | :heavy_check_mark: | `MavenRepository`     |
| `artifactory_virtual_npm_repository`       | :heavy_check_mark: | `NPMRepository`       |
| `artifactory_virtual_nuget_repository`     | :heavy_check_mark: | `NuGetRepository`     |
| `artifactory_virtual_oci_repository`       | :heavy_check_mark: | `OCIRepository`       |
| `artifactory_virtual_p2_repository`        | :heavy_check_mark: | `P2Repository`        |
| `artifactory_virtual_pub_repository`       | :heavy_check_mark: | `PubRepository`       |
| `artifactory_virtual_puppet_repository`    | :heavy_check_mark: | `PuppetRepository`    |
| `artifactory_virtual_pypi_repository`      | :heavy_check_mark: | `PyPIRepository`      |
| `artifactory_virtual_rpm_repository`       | :heavy_check_mark: | `RPMRepository`       |
| `artifactory_virtual_sbt_repository`       | :heavy_check_mark: | `SBTRepository`       |
| `artifactory_virtual_swift_repository`     | :heavy_check_mark: | `SwiftRepository`     |
| `artifactory_virtual_terraform_repository` | :heavy_check_mark: | `TerraformRepository` |

### Webhook

Short group is `webhook`, so the `apiGroup` is:

- `webhook.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**
- `webhook.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**

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

Check [`BUILD_FROM_SCRATCH.md`](./BUILD_FROM_SCRATCH.md) for notes on how this provider was built step-by-step using the [crossplane/upjet tool](https://github.com/crossplane/upjet).

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

## JFrog Artifactory icon

Package icon was pulled from [JFrog Brand Guidelines](https://jfrog.com/brand-guidelines/).

Icon is stored in [`extensions/icons/`](./extensions/icons/) folder according to instructions [Add documentation, icons, and other assets to your package](https://docs.upbound.io/manuals/marketplace/packages#add-documentation-icons-and-other-assets-to-your-package).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/hmlkao/provider-jfrog-artifactory/issues).
