/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Artifact
	"artifactory_item_properties": config.IdentifierFromProvider,
	// Configuration
	// "artifactory_archive_policy": config.IdentifierFromProvider,
	// Local Repositories
	"artifactory_local_ansible_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_local_docker_v2_repository":          config.ParameterAsIdentifier("key"),
	"artifactory_local_generic_repository":            config.ParameterAsIdentifier("key"),
	"artifactory_local_oci_repository":                config.ParameterAsIdentifier("key"),
	"artifactory_local_terraform_provider_repository": config.ParameterAsIdentifier("key"),
	// Security
	"artifactory_keypair": config.ParameterAsIdentifier("pair_name"),
	// User
	"artifactory_anonymous_user": config.NameAsIdentifier,
	// "artifactory_managed_user": config.NameAsIdentifier,
	// "artifactory_unmanaged_user": config.NameAsIdentifier,
	// "artifactory_user": config.NameAsIdentifier,
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
