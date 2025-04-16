/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/config/artifact/item_properties"
	localansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_ansible_repository"
	localdockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_docker_v2_repository"
	localgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_generic_repository"
	localocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_oci_repository"
	localterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_terraform_provider_repository"
	"github.com/hmlkao/provider-jfrog-artifactory/config/security/keypair"
	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/config/user/anonymous_user"
)

const (
	resourcePrefix = "artifactory"
	modulePath     = "github.com/hmlkao/provider-jfrog-artifactory"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("artifactory.jfrog.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// Artifact
		itemproperties.Configure,
		// Configuration
		// archivepolicy.Configure,
		// Local Repositories
		localansiblerepository.Configure,
		localdockerv2repository.Configure,
		localgenericrepository.Configure,
		localocirepository.Configure,
		localterraformproviderrepository.Configure,
		// Security
		keypair.Configure,
		// User
		anonymoususer.Configure,
		// manageduser.Configure,
		// unmanageduser.Configure,
		// user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
