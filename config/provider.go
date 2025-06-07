/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/config/artifact/item_properties"
	backup "github.com/hmlkao/provider-jfrog-artifactory/config/configuration/backup"
	localalpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_alpine_repository"
	localansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_ansible_repository"
	localbowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_bower_repository"
	localcargorepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_cargo_repository"
	localchefrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_chef_repository"
	localcocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_cocoapods_repository"
	localcomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_composer_repository"
	localconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_conan_repository"
	localcondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_conda_repository"
	localcranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_cran_repository"
	localdebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_debian_repository"
	localdockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_docker_v1_repository"
	localdockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_docker_v2_repository"
	localgemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_gems_repository"
	localgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_generic_repository"
	localgitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_gitlfs_repository"
	localgorepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_go_repository"
	localgradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_gradle_repository"
	localhelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_helm_repository"
	localhelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_helmoci_repository"
	localhuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_huggingfaceml_repository"
	localivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_ivy_repository"
	localmachinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_machinelearning_repository"
	localmavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_maven_repository"
	localnpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_npm_repository"
	localnugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_nuget_repository"
	localocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_oci_repository"
	localopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_opkg_repository"
	localpubrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_pub_repository"
	localpuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_puppet_repository"
	localpypirepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_pypi_repository"
	localrpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_rpm_repository"
	localsbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_sbt_repository"
	localswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_swift_repository"
	localterraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_terraform_module_repository"
	localterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_terraform_provider_repository"
	localterraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_terraformbackend_repository"
	localvagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/config/local_repositories/local_vagrant_repository"
	remotealpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_alpine_repository"
	remoteansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_ansible_repository"
	remotebowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_bower_repository"
	remotecargorepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_cargo_repository"
	remotechefrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_chef_repository"
	remotecocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_cocoapods_repository"
	remotecomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_composer_repository"
	remoteconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_conan_repository"
	remotecondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_conda_repository"
	remotecranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_cran_repository"
	remotedebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_debian_repository"
	remotedockerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_docker_repository"
	remotegemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_gems_repository"
	remotegenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_generic_repository"
	remotegitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_gitlfs_repository"
	remotegorepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_go_repository"
	remotegradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_gradle_repository"
	remotehelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_helm_repository"
	remotehelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_helmoci_repository"
	remotehuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_huggingfaceml_repository"
	remoteivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_ivy_repository"
	remotemavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_maven_repository"
	remotenpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_npm_repository"
	remotenugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_nuget_repository"
	remoteocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_oci_repository"
	remoteopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_opkg_repository"
	remotep2repository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_p2_repository"
	remotepubrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_pub_repository"
	remotepuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_puppet_repository"
	remotepypirepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_pypi_repository"
	remoterpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_rpm_repository"
	remotesbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_sbt_repository"
	remoteswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_swift_repository"
	remoteterrraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_terraform_provider_repository"
	remoteterrraformrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_terraform_repository"
	remotevcsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/remote_repositories/remote_vcs_repository"
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
		backup.Configure,
		// Local Repositories
		localansiblerepository.Configure,
		localalpinerepository.Configure,
		localbowerrepository.Configure,
		localcargorepository.Configure,
		localchefrepository.Configure,
		localcocoapodsrepository.Configure,
		localcomposerrepository.Configure,
		localconanrepository.Configure,
		localcondarepository.Configure,
		localcranrepository.Configure,
		localdebianrepository.Configure,
		localdockerv1repository.Configure,
		localdockerv2repository.Configure,
		localgemsrepository.Configure,
		localgenericrepository.Configure,
		localgitlfsrepository.Configure,
		localgorepository.Configure,
		localgradlerepository.Configure,
		localhelmrepository.Configure,
		localhelmocirepository.Configure,
		localhuggingfacemlrepository.Configure,
		localivyrepository.Configure,
		localmachinelearningrepository.Configure,
		localmavenrepository.Configure,
		localnpmrepository.Configure,
		localnugetrepository.Configure,
		localocirepository.Configure,
		localopkgrepository.Configure,
		localpubrepository.Configure,
		localpuppetrepository.Configure,
		localpypirepository.Configure,
		localrpmrepository.Configure,
		localsbtrepository.Configure,
		localswiftrepository.Configure,
		localterraformmodulerepository.Configure,
		localterraformproviderrepository.Configure,
		localterraformbackendrepository.Configure,
		localvagrantrepository.Configure,
		// Security
		keypair.Configure,
		// Remote Repositories
		remotealpinerepository.Configure,
		remoteansiblerepository.Configure,
		remotebowerrepository.Configure,
		remotecargorepository.Configure,
		remotechefrepository.Configure,
		remotecocoapodsrepository.Configure,
		remotecomposerrepository.Configure,
		remoteconanrepository.Configure,
		remotecondarepository.Configure,
		remotecranrepository.Configure,
		remotedebianrepository.Configure,
		remotedockerrepository.Configure,
		remotegemsrepository.Configure,
		remotegenericrepository.Configure,
		remotegitlfsrepository.Configure,
		remotegorepository.Configure,
		remotegradlerepository.Configure,
		remotehelmrepository.Configure,
		remotehelmocirepository.Configure,
		remotehuggingfacemlrepository.Configure,
		remoteivyrepository.Configure,
		remotemavenrepository.Configure,
		remotenpmrepository.Configure,
		remotenugetrepository.Configure,
		remoteocirepository.Configure,
		remoteopkgrepository.Configure,
		remotep2repository.Configure,
		remotepubrepository.Configure,
		remotepuppetrepository.Configure,
		remotepypirepository.Configure,
		remoterpmrepository.Configure,
		remotesbtrepository.Configure,
		remoteswiftrepository.Configure,
		remoteterrraformrepository.Configure,
		remoteterrraformproviderrepository.Configure,
		remotevcsrepository.Configure,
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
