/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/artifact/item_properties"
	backup "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/configuration/backup"
	federatedalpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_alpine_repository"
	federatedansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_ansible_repository"
	federatedbowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_bower_repository"
	federatedcargorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_cargo_repository"
	federatedchefrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_chef_repository"
	federatedcocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_cocoapods_repository"
	federatedcomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_composer_repository"
	federatedconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_conan_repository"
	federatedcondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_conda_repository"
	federatedcranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_cran_repository"
	federateddebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_debian_repository"
	federateddockerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_docker_repository"
	federateddockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_docker_v1_repository"
	federateddockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_docker_v2_repository"
	federatedgemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_gems_repository"
	federatedgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_generic_repository"
	federatedgitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_gitlfs_repository"
	federatedgorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_go_repository"
	federatedgradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_gradle_repository"
	federatedhelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_helm_repository"
	federatedhelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_helmoci_repository"
	federatedhuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_huggingfaceml_repository"
	federatedivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_ivy_repository"
	federatedmavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_maven_repository"
	federatednpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_npm_repository"
	federatednugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_nuget_repository"
	federatedocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_oci_repository"
	federatedopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_opkg_repository"
	federatedpuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_puppet_repository"
	federatedpypirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_pypi_repository"
	federatedrpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_rpm_repository"
	federatedsbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_sbt_repository"
	federatedswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_swift_repository"
	federatedterraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_terraform_module_repository"
	federatedterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_terraform_provider_repository"
	federatedvagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/federated_repositories/federated_vagrant_repository"
	localalpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_alpine_repository"
	localansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_ansible_repository"
	localbowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_bower_repository"
	localcargorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_cargo_repository"
	localchefrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_chef_repository"
	localcocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_cocoapods_repository"
	localcomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_composer_repository"
	localconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_conan_repository"
	localcondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_conda_repository"
	localcranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_cran_repository"
	localdebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_debian_repository"
	localdockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_docker_v1_repository"
	localdockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_docker_v2_repository"
	localgemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_gems_repository"
	localgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_generic_repository"
	localgitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_gitlfs_repository"
	localgorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_go_repository"
	localgradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_gradle_repository"
	localhelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_helm_repository"
	localhelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_helmoci_repository"
	localhuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_huggingfaceml_repository"
	localivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_ivy_repository"
	localmachinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_machinelearning_repository"
	localmavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_maven_repository"
	localnpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_npm_repository"
	localnugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_nuget_repository"
	localocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_oci_repository"
	localopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_opkg_repository"
	localpubrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_pub_repository"
	localpuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_puppet_repository"
	localpypirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_pypi_repository"
	localrpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_rpm_repository"
	localsbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_sbt_repository"
	localswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_swift_repository"
	localterraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_terraform_module_repository"
	localterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_terraform_provider_repository"
	localterraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_terraformbackend_repository"
	localvagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/local_repositories/local_vagrant_repository"
	remotealpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_alpine_repository"
	remoteansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_ansible_repository"
	remotebowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_bower_repository"
	remotecargorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_cargo_repository"
	remotechefrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_chef_repository"
	remotecocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_cocoapods_repository"
	remotecomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_composer_repository"
	remoteconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_conan_repository"
	remotecondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_conda_repository"
	remotecranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_cran_repository"
	remotedebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_debian_repository"
	remotedockerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_docker_repository"
	remotegemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_gems_repository"
	remotegenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_generic_repository"
	remotegitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_gitlfs_repository"
	remotegorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_go_repository"
	remotegradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_gradle_repository"
	remotehelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_helm_repository"
	remotehelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_helmoci_repository"
	remotehuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_huggingfaceml_repository"
	remoteivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_ivy_repository"
	remotemavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_maven_repository"
	remotenpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_npm_repository"
	remotenugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_nuget_repository"
	remoteocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_oci_repository"
	remoteopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_opkg_repository"
	remotep2repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_p2_repository"
	remotepubrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_pub_repository"
	remotepuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_puppet_repository"
	remotepypirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_pypi_repository"
	remoterpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_rpm_repository"
	remotesbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_sbt_repository"
	remoteswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_swift_repository"
	remoteterrraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_terraform_provider_repository"
	remoteterrraformrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_terraform_repository"
	remotevcsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/remote_repositories/remote_vcs_repository"
	"github.com/hmlkao/provider-jfrog-artifactory/config/cluster/security/keypair"
	scopedtoken "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/security/scoped_token"
	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/user/anonymous_user"
	"github.com/hmlkao/provider-jfrog-artifactory/config/cluster/user/user"
	virtualalpinerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_alpine_repository"
	virtualansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_ansible_repository"
	virtualbowerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_bower_repository"
	virtualcocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_cocoapods_repository"
	virtualcomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_composer_repository"
	virtualconanrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_conan_repository"
	virtualcondarepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_conda_repository"
	virtualcranrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_cran_repository"
	virtualdebianrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_debian_repository"
	virtualdockerrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_docker_repository"
	virtualgemsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_gems_repository"
	virtualgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_generic_repository"
	virtualgitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_gitlfs_repository"
	virtualgorepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_go_repository"
	virtualgradlerepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_gradle_repository"
	virtualhelmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_helm_repository"
	virtualhelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_helmoci_repository"
	virtualivyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_ivy_repository"
	virtualmavenrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_maven_repository"
	virtualnpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_npm_repository"
	virtualnugetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_nuget_repository"
	virtualocirepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_oci_repository"
	virtualp2repository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_p2_repository"
	virtualpubrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_pub_repository"
	virtualpuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_puppet_repository"
	virtualpyyrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_pypi_repository"
	virtualrpmrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_rpm_repository"
	virtualsbtrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_sbt_repository"
	virtualswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_swift_repository"
	virtualterraformrepository "github.com/hmlkao/provider-jfrog-artifactory/config/cluster/virtual_repositories/virtual_terraform_repository"
)

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
		backup.Configure,
		// Federated Repositories
		federatedalpinerepository.Configure,
		federatedansiblerepository.Configure,
		federatedbowerrepository.Configure,
		federatedcargorepository.Configure,
		federatedchefrepository.Configure,
		federatedcocoapodsrepository.Configure,
		federatedcomposerrepository.Configure,
		federatedconanrepository.Configure,
		federatedcondarepository.Configure,
		federatedcranrepository.Configure,
		federateddebianrepository.Configure,
		federateddockerrepository.Configure,
		federateddockerv1repository.Configure,
		federateddockerv2repository.Configure,
		federatedgemsrepository.Configure,
		federatedgenericrepository.Configure,
		federatedgitlfsrepository.Configure,
		federatedgorepository.Configure,
		federatedgradlerepository.Configure,
		federatedhelmrepository.Configure,
		federatedhelmocirepository.Configure,
		federatedhuggingfacemlrepository.Configure,
		federatedivyrepository.Configure,
		federatedmavenrepository.Configure,
		federatednpmrepository.Configure,
		federatednugetrepository.Configure,
		federatedocirepository.Configure,
		federatedopkgrepository.Configure,
		federatedpuppetrepository.Configure,
		federatedpypirepository.Configure,
		federatedrpmrepository.Configure,
		federatedsbtrepository.Configure,
		federatedswiftrepository.Configure,
		federatedterraformmodulerepository.Configure,
		federatedterraformproviderrepository.Configure,
		federatedvagrantrepository.Configure,
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
		scopedtoken.Configure,
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
		user.Configure,
		// Virtual Repositories
		virtualalpinerepository.Configure,
		virtualansiblerepository.Configure,
		virtualbowerrepository.Configure,
		virtualcocoapodsrepository.Configure,
		virtualcomposerrepository.Configure,
		virtualconanrepository.Configure,
		virtualcondarepository.Configure,
		virtualcranrepository.Configure,
		virtualdebianrepository.Configure,
		virtualdockerrepository.Configure,
		virtualgemsrepository.Configure,
		virtualgenericrepository.Configure,
		virtualgitlfsrepository.Configure,
		virtualgorepository.Configure,
		virtualgradlerepository.Configure,
		virtualhelmrepository.Configure,
		virtualhelmocirepository.Configure,
		virtualivyrepository.Configure,
		virtualnpmrepository.Configure,
		virtualnugetrepository.Configure,
		virtualmavenrepository.Configure,
		virtualocirepository.Configure,
		virtualp2repository.Configure,
		virtualpubrepository.Configure,
		virtualpuppetrepository.Configure,
		virtualpyyrepository.Configure,
		virtualrpmrepository.Configure,
		virtualsbtrepository.Configure,
		virtualswiftrepository.Configure,
		virtualterraformrepository.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
