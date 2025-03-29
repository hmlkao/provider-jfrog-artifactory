# Provider Artifactory

`provider-artifactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-artifactory):

```bash
up ctp provider install hmlkao/provider-artifactory:v0.1.0
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-artifactory
spec:
  package: hmlkao/provider-artifactory:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/hmlkao/provider-artifactory).

## Resources

List of all resources of [Terraform provider version 12.9.1](https://registry.terraform.io/providers/jfrog/artifactory/12.9.1/docs).

### Artifact

| Resource                      | Implemented | Kind                           |
|-------------------------------|-------------|--------------------------------|
| `artifactory_artifact`        | [ ]         |                                |
| `artifactory_item_properties` | [ ]         |                                |

### Configuration

| Resource                             | Implemented | Kind                           |
|--------------------------------------|-------------|--------------------------------|
| `artifactory_archive_policy`         | [ ]         |                                |
| `artifactory_backup`                 | [ ]         |                                |
| `artifactory_general_security`       | [ ]         |                                |
| `artifactory_ldap_group_setting`     | [ ]         |                                |
| `artifactory_ldap_group_setting_v2`  | [ ]         |                                |
| `artifactory_ldap_setting`           | [ ]         |                                |
| `artifactory_ldap_setting_v2`        | [ ]         |                                |
| `artifactory_mail_server`            | [ ]         |                                |
| `artifactory_oauth_settings`         | [ ]         |                                |
| `artifactory_package_cleanup_policy` | [ ]         |                                |
| `artifactory_property_set`           | [ ]         |                                |
| `artifactory_proxy`                  | [ ]         |                                |
| `artifactory_repository_layout`      | [ ]         |                                |
| `artifactory_vault_configuration`    | [ ]         |                                |

### Federated Repositories

| Resource                                              | Implemented | Kind                           |
|-------------------------------------------------------|-------------|--------------------------------|
| `artifactory_federated_alpine_repository`             | [ ]         |                                |
| `artifactory_federated_ansible_repository`            | [ ]         |                                |
| `artifactory_federated_bower_repository`              | [ ]         |                                |
| `artifactory_federated_cargo_repository`              | [ ]         |                                |
| `artifactory_federated_chef_repository`               | [ ]         |                                |
| `artifactory_federated_cocoapods_repository`          | [ ]         |                                |
| `artifactory_federated_composer_repository`           | [ ]         |                                |
| `artifactory_federated_conan_repository`              | [ ]         |                                |
| `artifactory_federated_conda_repository`              | [ ]         |                                |
| `artifactory_federated_cran_repository`               | [ ]         |                                |
| `artifactory_federated_debian_repository`             | [ ]         |                                |
| `artifactory_federated_docker_repository`             | [ ]         |                                |
| `artifactory_federated_docker_v1_repository`          | [ ]         |                                |
| `artifactory_federated_docker_v2_repository`          | [ ]         |                                |
| `artifactory_federated_gems_repository`               | [ ]         |                                |
| `artifactory_federated_generic_repository`            | [ ]         |                                |
| `artifactory_federated_gitlfs_repository`             | [ ]         |                                |
| `artifactory_federated_go_repository`                 | [ ]         |                                |
| `artifactory_federated_gradle_repository`             | [ ]         |                                |
| `artifactory_federated_helm_repository`               | [ ]         |                                |
| `artifactory_federated_helmoci_repository`            | [ ]         |                                |
| `artifactory_federated_huggingfaceml_repository`      | [ ]         |                                |
| `artifactory_federated_ivy_repository`                | [ ]         |                                |
| `artifactory_federated_maven_repository`              | [ ]         |                                |
| `artifactory_federated_npm_repository`                | [ ]         |                                |
| `artifactory_federated_nuget_repository`              | [ ]         |                                |
| `artifactory_federated_oci_repository`                | [ ]         |                                |
| `artifactory_federated_opkg_repository`               | [ ]         |                                |
| `artifactory_federated_puppet_repository`             | [ ]         |                                |
| `artifactory_federated_pypi_repository`               | [ ]         |                                |
| `artifactory_federated_rpm_repository`                | [ ]         |                                |
| `artifactory_federated_sbt_repository`                | [ ]         |                                |
| `artifactory_federated_swift_repository`              | [ ]         |                                |
| `artifactory_federated_terraform_module_repository`   | [ ]         |                                |
| `artifactory_federated_terraform_provider_repository` | [ ]         |                                |
| `artifactory_federated_vagrant_repository`            | [ ]         |                                |

### Lifecycle

| Resource                                  | Implemented | Kind                           |
|-------------------------------------------|-------------|--------------------------------|
| `artifactory_release_bundle_v2`           | [ ]         |                                |
| `artifactory_release_bundle_v2_promotion` | [ ]         |                                |

### Local Repositories

| Resource                                          | Implemented | Kind                           |
|---------------------------------------------------|-------------|--------------------------------|
| `artifactory_local_alpine_repository`             | [ ]         |                                |
| `artifactory_local_ansible_repository`            | [ ]         |                                |
| `artifactory_local_bower_repository`              | [ ]         |                                |
| `artifactory_local_cargo_repository`              | [ ]         |                                |
| `artifactory_local_chef_repository`               | [ ]         |                                |
| `artifactory_local_cocoapods_repository`          | [ ]         |                                |
| `artifactory_local_composer_repository`           | [ ]         |                                |
| `artifactory_local_conan_repository`              | [ ]         |                                |
| `artifactory_local_conda_repository`              | [ ]         |                                |
| `artifactory_local_cran_repository`               | [ ]         |                                |
| `artifactory_local_debian_repository`             | [ ]         |                                |
| `artifactory_local_docker_v1_repository`          | [ ]         |                                |
| `artifactory_local_docker_v2_repository`          | [ ]         |                                |
| `artifactory_local_gems_repository`               | [ ]         |                                |
| `artifactory_local_generic_repository`            | [ ]         |                                |
| `artifactory_local_gitlfs_repository`             | [ ]         |                                |
| `artifactory_local_go_repository`                 | [ ]         |                                |
| `artifactory_local_gradle_repository`             | [ ]         |                                |
| `artifactory_local_helm_repository`               | [ ]         |                                |
| `artifactory_local_helmoci_repository`            | [ ]         |                                |
| `artifactory_local_huggingfaceml_repository`      | [ ]         |                                |
| `artifactory_local_ivy_repository`                | [ ]         |                                |
| `artifactory_local_machinelearning_repository`    | [ ]         |                                |
| `artifactory_local_maven_repository`              | [ ]         |                                |
| `artifactory_local_npm_repository`                | [ ]         |                                |
| `artifactory_local_nuget_repository`              | [ ]         |                                |
| `artifactory_local_oci_repository`                | [x]         | `LocalOCIRepository`           |
| `artifactory_local_opkg_repository`               | [ ]         |                                |
| `artifactory_local_pub_repository`                | [ ]         |                                |
| `artifactory_local_puppet_repository`             | [ ]         |                                |
| `artifactory_local_pypi_repository`               | [ ]         |                                |
| `artifactory_local_rpm_repository`                | [ ]         |                                |
| `artifactory_local_sbt_repository`                | [ ]         |                                |
| `artifactory_local_swift_repository`              | [ ]         |                                |
| `artifactory_local_terraform_module_repository`   | [ ]         |                                |
| `artifactory_local_terraform_provider_repository` | [ ]         |                                |
| `artifactory_local_terraformbackend_repository`   | [ ]         |                                |
| `artifactory_local_vagrant_repository`            | [ ]         |                                |

### Remote Repositories

| Resource                                           | Implemented | Kind                           |
|----------------------------------------------------|-------------|--------------------------------|
| `artifactory_remote_alpine_repository`             | [ ]         |                                |
| `artifactory_remote_ansible_repository`            | [ ]         |                                |
| `artifactory_remote_bower_repository`              | [ ]         |                                |
| `artifactory_remote_cargo_repository`              | [ ]         |                                |
| `artifactory_remote_chef_repository`               | [ ]         |                                |
| `artifactory_remote_cocoapods_repository`          | [ ]         |                                |
| `artifactory_remote_composer_repository`           | [ ]         |                                |
| `artifactory_remote_conan_repository`              | [ ]         |                                |
| `artifactory_remote_conda_repository`              | [ ]         |                                |
| `artifactory_remote_cran_repository`               | [ ]         |                                |
| `artifactory_remote_debian_repository`             | [ ]         |                                |
| `artifactory_remote_docker_repository`             | [ ]         |                                |
| `artifactory_remote_gems_repository`               | [ ]         |                                |
| `artifactory_remote_generic_repository`            | [ ]         |                                |
| `artifactory_remote_gitlfs_repository`             | [ ]         |                                |
| `artifactory_remote_go_repository`                 | [ ]         |                                |
| `artifactory_remote_gradle_repository`             | [ ]         |                                |
| `artifactory_remote_helm_repository`               | [ ]         |                                |
| `artifactory_remote_helmoci_repository`            | [ ]         |                                |
| `artifactory_remote_huggingfaceml_repository`      | [ ]         |                                |
| `artifactory_remote_ivy_repository`                | [ ]         |                                |
| `artifactory_remote_maven_repository`              | [ ]         |                                |
| `artifactory_remote_npm_repository`                | [ ]         |                                |
| `artifactory_remote_nuget_repository`              | [ ]         |                                |
| `artifactory_remote_oci_repository`                | [ ]         |                                |
| `artifactory_remote_opkg_repository`               | [ ]         |                                |
| `artifactory_remote_p2_repository`                 | [ ]         |                                |
| `artifactory_remote_pub_repository`                | [ ]         |                                |
| `artifactory_remote_puppet_repository`             | [ ]         |                                |
| `artifactory_remote_pypi_repository`               | [ ]         |                                |
| `artifactory_remote_rpm_repository`                | [ ]         |                                |
| `artifactory_remote_sbt_repository`                | [ ]         |                                |
| `artifactory_remote_swift_repository`              | [ ]         |                                |
| `artifactory_remote_terraform_repository`          | [ ]         |                                |
| `artifactory_remote_terraform_provider_repository` | [ ]         |                                |
| `artifactory_remote_vcs_repository`                | [ ]         |                                |

### Replication

| Resource                                          | Implemented | Kind                           |
|---------------------------------------------------|-------------|--------------------------------|
| `artifactory_local_repository_multi_replication`  | [ ]         |                                |
| `artifactory_local_repository_single_replication` | [ ]         |                                |
| `artifactory_remote_repository_replication`       | [ ]         |                                |

### Security

| Resource                                 | Implemented | Kind                           |
|------------------------------------------|-------------|--------------------------------|
| `artifactory_certificate`                | [ ]         |                                |
| `artifactory_distribution_public_key`    | [ ]         |                                |
| `artifactory_global_environment`         | [ ]         |                                |
| `artifactory_keypair`                    | [ ]         |                                |
| `artifactory_password_expiration_policy` | [ ]         |                                |
| `artifactory_scoped_token`               | [ ]         |                                |
| `artifactory_user_lock_policy`           | [ ]         |                                |

### User

| Resource                     | Implemented | Kind                           |
|------------------------------|-------------|--------------------------------|
| `artifactory_anonymous_user` | [ ]         |                                |
| `artifactory_managed_user`   | [ ]         |                                |
| `artifactory_unmanaged_user` | [ ]         |                                |
| `artifactory_user`           | [ ]         |                                |

### Virtual Repositories

| Resource                                   | Implemented | Kind                           |
|--------------------------------------------|-------------|--------------------------------|
| `artifactory_virtual_alpine_repository`    | [ ]         |                                |
| `artifactory_virtual_ansible_repository`   | [ ]         |                                |
| `artifactory_virtual_bower_repository`     | [ ]         |                                |
| `artifactory_virtual_chef_repository`      | [ ]         |                                |
| `artifactory_virtual_cocoapods_repository` | [ ]         |                                |
| `artifactory_virtual_composer_repository`  | [ ]         |                                |
| `artifactory_virtual_conan_repository`     | [ ]         |                                |
| `artifactory_virtual_conda_repository`     | [ ]         |                                |
| `artifactory_virtual_cran_repository`      | [ ]         |                                |
| `artifactory_virtual_debian_repository`    | [ ]         |                                |
| `artifactory_virtual_docker_repository`    | [ ]         |                                |
| `artifactory_virtual_gems_repository`      | [ ]         |                                |
| `artifactory_virtual_generic_repository`   | [ ]         |                                |
| `artifactory_virtual_gitlfs_repository`    | [ ]         |                                |
| `artifactory_virtual_go_repository`        | [ ]         |                                |
| `artifactory_virtual_gradle_repository`    | [ ]         |                                |
| `artifactory_virtual_helm_repository`      | [ ]         |                                |
| `artifactory_virtual_helmoci_repository`   | [ ]         |                                |
| `artifactory_virtual_ivy_repository`       | [ ]         |                                |
| `artifactory_virtual_maven_repository`     | [ ]         |                                |
| `artifactory_virtual_npm_repository`       | [ ]         |                                |
| `artifactory_virtual_nuget_repository`     | [ ]         |                                |
| `artifactory_virtual_oci_repository`       | [ ]         |                                |
| `artifactory_virtual_p2_repository`        | [ ]         |                                |
| `artifactory_virtual_pub_repository`       | [ ]         |                                |
| `artifactory_virtual_puppet_repository`    | [ ]         |                                |
| `artifactory_virtual_pypi_repository`      | [ ]         |                                |
| `artifactory_virtual_rpm_repository`       | [ ]         |                                |
| `artifactory_virtual_sbt_repository`       | [ ]         |                                |
| `artifactory_virtual_swift_repository`     | [ ]         |                                |
| `artifactory_virtual_terraform_repository` | [ ]         |                                |

### Webhook

| Resource                                                 | Implemented | Kind                           |
|----------------------------------------------------------|-------------|--------------------------------|
| `artifactory_artifact_custom_webhook`                    | [ ]         |                                |
| `artifactory_artifact_lifecycle_custom_webhook`          | [ ]         |                                |
| `artifactory_artifact_lifecycle_webhook`                 | [ ]         |                                |
| `artifactory_artifact_property_custom_webhook`           | [ ]         |                                |
| `artifactory_artifact_property_webhook`                  | [ ]         |                                |
| `artifactory_artifact_webhook`                           | [ ]         |                                |
| `artifactory_artifactory_release_bundle_custom_webhook`  | [ ]         |                                |
| `artifactory_artifactory_release_bundle_webhook`         | [ ]         |                                |
| `artifactory_build_custom_webhook`                       | [ ]         |                                |
| `artifactory_build_webhook`                              | [ ]         |                                |
| `artifactory_destination_custom_webhook`                 | [ ]         |                                |
| `artifactory_destination_webhook`                        | [ ]         |                                |
| `artifactory_distribution_custom_webhook`                | [ ]         |                                |
| `artifactory_distribution_webhook`                       | [ ]         |                                |
| `artifactory_local_docker_v2_repository`                 | [ ]         |                                |
| `artifactory_docker_webhook`                             | [ ]         |                                |
| `artifactory_release_bundle_v2_custom_webhook`           | [ ]         |                                |
| `artifactory_release_bundle_v2_promotion_custom_webhook` | [ ]         |                                |
| `artifactory_release_bundle_v2_promotion_webhook`        | [ ]         |                                |
| `artifactory_release_bundle_v2_webhook`                  | [ ]         |                                |
| `artifactory_user_custom_webhook`                        | [ ]         |                                |
| `artifactory_user_webhook`                               | [ ]         |                                |

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

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/hmlkao/provider-artifactory/issues).
