/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	localdockerv2repository "github.com/hmlkao/provider-artifactory/config/local_docker_v2_repository"
	localgenericrepository "github.com/hmlkao/provider-artifactory/config/local_generic_repository"
	localocirepository "github.com/hmlkao/provider-artifactory/config/local_oci_repository"
	localterraformproviderrepository "github.com/hmlkao/provider-artifactory/config/local_terraform_provider_repository"
)

const (
	resourcePrefix = "artifactory"
	modulePath     = "github.com/hmlkao/provider-artifactory"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("jfrog.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		localdockerv2repository.Configure,
		localgenericrepository.Configure,
		localocirepository.Configure,
		localterraformproviderrepository.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
