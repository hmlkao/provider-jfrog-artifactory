/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Artifact
	"artifactory_item_properties": config.IdentifierFromProvider,
	// Configuration
	// "artifactory_archive_policy": config.IdentifierFromProvider,
	"artifactory_backup": config.ParameterAsIdentifier("key"),
	// Local Repositories
	"artifactory_local_alpine_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_local_ansible_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_local_bower_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_cargo_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_chef_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_cocoapods_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_local_composer_repository":           config.ParameterAsIdentifier("key"),
	"artifactory_local_conan_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_conda_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_cran_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_debian_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_local_docker_v1_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_local_docker_v2_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_local_gems_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_generic_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_local_gitlfs_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_local_go_repository":                 config.ParameterAsIdentifier("key"),
	"artifactory_local_gradle_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_local_helm_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_helmoci_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_local_huggingfaceml_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_local_ivy_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_machinelearning_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_local_maven_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_npm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_nuget_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_oci_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_opkg_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_pub_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_puppet_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_local_pypi_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_local_rpm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_sbt_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_swift_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_local_terraform_module_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_local_terraform_provider_repository": config.ParameterAsIdentifier("key"),
	"artifactory_local_terraformbackend_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_local_vagrant_repository":            config.ParameterAsIdentifier("key"),
	// Security
	"artifactory_keypair":      config.ParameterAsIdentifier("pair_name"),
	"artifactory_scoped_token": config.IdentifierFromProvider,
	// Remote Repositories
	"artifactory_remote_alpine_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_ansible_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_remote_bower_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_cargo_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_chef_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_cocoapods_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_remote_composer_repository":           config.ParameterAsIdentifier("key"),
	"artifactory_remote_conan_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_conda_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_cran_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_debian_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_docker_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_gems_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_generic_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_remote_gitlfs_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_go_repository":                 config.ParameterAsIdentifier("key"),
	"artifactory_remote_gradle_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_helm_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_helmoci_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_remote_huggingfaceml_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_remote_ivy_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_maven_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_npm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_nuget_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_oci_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_opkg_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_p2_repository":                 config.ParameterAsIdentifier("key"),
	"artifactory_remote_pub_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_puppet_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_remote_pypi_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_remote_rpm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_sbt_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_remote_swift_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_remote_terraform_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_remote_terraform_provider_repository": config.ParameterAsIdentifier("key"),
	"artifactory_remote_vcs_repository":                config.ParameterAsIdentifier("key"),
	// User
	"artifactory_anonymous_user": config.NameAsIdentifier,
	// "artifactory_managed_user": config.NameAsIdentifier,
	// "artifactory_unmanaged_user": config.NameAsIdentifier,
	"artifactory_user": config.NameAsIdentifier,
	// Virtual Repositories
	"artifactory_virtual_alpine_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_ansible_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_virtual_bower_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_cocoapods_repository": config.ParameterAsIdentifier("key"),
	"artifactory_virtual_composer_repository":  config.ParameterAsIdentifier("key"),
	"artifactory_virtual_conan_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_conda_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_cran_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_virtual_debian_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_docker_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_gems_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_virtual_generic_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_virtual_gitlfs_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_go_repository":        config.ParameterAsIdentifier("key"),
	"artifactory_virtual_gradle_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_helm_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_virtual_helmoci_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_virtual_ivy_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_maven_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_npm_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_nuget_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_oci_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_pub_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_p2_repository":        config.ParameterAsIdentifier("key"),
	"artifactory_virtual_puppet_repository":    config.ParameterAsIdentifier("key"),
	"artifactory_virtual_pypi_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_virtual_rpm_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_sbt_repository":       config.ParameterAsIdentifier("key"),
	"artifactory_virtual_swift_repository":     config.ParameterAsIdentifier("key"),
	"artifactory_virtual_terraform_repository": config.ParameterAsIdentifier("key"),
	// Federated Repositories
	"artifactory_federated_alpine_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_ansible_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_federated_bower_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_cargo_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_chef_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_cocoapods_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_federated_composer_repository":           config.ParameterAsIdentifier("key"),
	"artifactory_federated_conan_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_conda_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_cran_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_debian_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_docker_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_docker_v1_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_federated_docker_v2_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_federated_gems_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_generic_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_federated_gitlfs_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_go_repository":                 config.ParameterAsIdentifier("key"),
	"artifactory_federated_gradle_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_helm_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_helmoci_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_federated_huggingfaceml_repository":      config.ParameterAsIdentifier("key"),
	"artifactory_federated_ivy_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_federated_maven_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_npm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_federated_nuget_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_oci_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_federated_opkg_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_puppet_repository":             config.ParameterAsIdentifier("key"),
	"artifactory_federated_pypi_repository":               config.ParameterAsIdentifier("key"),
	"artifactory_federated_rpm_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_federated_sbt_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_federated_swift_repository":              config.ParameterAsIdentifier("key"),
	"artifactory_federated_terraform_module_repository":   config.ParameterAsIdentifier("key"),
	"artifactory_federated_terraform_provider_repository": config.ParameterAsIdentifier("key"),
	"artifactory_federated_vagrant_repository":            config.ParameterAsIdentifier("key"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
