// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/artifact/itemproperties"
	backup "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/configuration/backup"
	alpinerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/alpinerepository"
	ansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/ansiblerepository"
	bowerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/bowerrepository"
	cargorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/cargorepository"
	chefrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/chefrepository"
	cocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/cocoapodsrepository"
	composerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/composerrepository"
	conanrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/conanrepository"
	condarepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/condarepository"
	cranrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/cranrepository"
	debianrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/debianrepository"
	dockerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/dockerrepository"
	dockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/dockerv1repository"
	dockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/dockerv2repository"
	gemsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/gemsrepository"
	genericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/genericrepository"
	gitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/gitlfsrepository"
	gorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/gorepository"
	gradlerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/gradlerepository"
	helmocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/helmocirepository"
	helmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/helmrepository"
	huggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/huggingfacemlrepository"
	ivyrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/ivyrepository"
	mavenrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/mavenrepository"
	npmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/npmrepository"
	nugetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/nugetrepository"
	ocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/ocirepository"
	opkgrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/opkgrepository"
	puppetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/puppetrepository"
	pypirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/pypirepository"
	rpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/rpmrepository"
	sbtrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/sbtrepository"
	swiftrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/swiftrepository"
	terraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/terraformmodulerepository"
	terraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/terraformproviderrepository"
	vagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/federated/vagrantrepository"
	alpinerepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/alpinerepository"
	ansiblerepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/ansiblerepository"
	bowerrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/bowerrepository"
	cargorepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/cargorepository"
	chefrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/chefrepository"
	cocoapodsrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/cocoapodsrepository"
	composerrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/composerrepository"
	conanrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/conanrepository"
	condarepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/condarepository"
	cranrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/cranrepository"
	debianrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/debianrepository"
	dockerv1repositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/dockerv1repository"
	dockerv2repositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/dockerv2repository"
	gemsrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/gemsrepository"
	genericrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/genericrepository"
	gitlfsrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/gitlfsrepository"
	gorepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/gorepository"
	gradlerepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/gradlerepository"
	helmocirepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/helmocirepository"
	helmrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/helmrepository"
	huggingfacemlrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/huggingfacemlrepository"
	ivyrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/ivyrepository"
	machinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/machinelearningrepository"
	mavenrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/mavenrepository"
	npmrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/npmrepository"
	nugetrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/nugetrepository"
	ocirepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/ocirepository"
	opkgrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/opkgrepository"
	pubrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/pubrepository"
	puppetrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/puppetrepository"
	pypirepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/pypirepository"
	rpmrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/rpmrepository"
	sbtrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/sbtrepository"
	swiftrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/swiftrepository"
	terraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/terraformbackendrepository"
	terraformmodulerepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/terraformmodulerepository"
	terraformproviderrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/terraformproviderrepository"
	vagrantrepositorylocal "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/local/vagrantrepository"
	providerconfig "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/providerconfig"
	alpinerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/alpinerepository"
	ansiblerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/ansiblerepository"
	bowerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/bowerrepository"
	cargorepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/cargorepository"
	chefrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/chefrepository"
	cocoapodsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/cocoapodsrepository"
	composerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/composerrepository"
	conanrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/conanrepository"
	condarepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/condarepository"
	cranrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/cranrepository"
	debianrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/debianrepository"
	dockerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/dockerrepository"
	gemsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/gemsrepository"
	genericrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/genericrepository"
	gitlfsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/gitlfsrepository"
	gorepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/gorepository"
	gradlerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/gradlerepository"
	helmocirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/helmocirepository"
	helmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/helmrepository"
	huggingfacemlrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/huggingfacemlrepository"
	ivyrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/ivyrepository"
	mavenrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/mavenrepository"
	npmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/npmrepository"
	nugetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/nugetrepository"
	ocirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/ocirepository"
	opkgrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/opkgrepository"
	p2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/p2repository"
	pubrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/pubrepository"
	puppetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/puppetrepository"
	pypirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/pypirepository"
	rpmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/rpmrepository"
	sbtrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/sbtrepository"
	swiftrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/swiftrepository"
	terraformrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/terraformrepository"
	vcsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/remote/vcsrepository"
	keypair "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/security/keypair"
	scopedtoken "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/security/scopedtoken"
	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/user/anonymoususer"
	alpinerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/alpinerepository"
	ansiblerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/ansiblerepository"
	bowerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/bowerrepository"
	cocoapodsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/cocoapodsrepository"
	composerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/composerrepository"
	conanrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/conanrepository"
	condarepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/condarepository"
	cranrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/cranrepository"
	debianrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/debianrepository"
	dockerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/dockerrepository"
	gemsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/gemsrepository"
	genericrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/genericrepository"
	gitlfsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/gitlfsrepository"
	gorepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/gorepository"
	gradlerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/gradlerepository"
	helmocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/helmocirepository"
	helmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/helmrepository"
	ivyrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/ivyrepository"
	mavenrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/mavenrepository"
	npmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/npmrepository"
	nugetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/nugetrepository"
	ocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/ocirepository"
	p2repositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/p2repository"
	pubrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/pubrepository"
	puppetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/puppetrepository"
	pypirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/pypirepository"
	rpmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/rpmrepository"
	sbtrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/sbtrepository"
	swiftrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/swiftrepository"
	terraformrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/cluster/virtual/terraformrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		itemproperties.Setup,
		backup.Setup,
		alpinerepository.Setup,
		ansiblerepository.Setup,
		bowerrepository.Setup,
		cargorepository.Setup,
		chefrepository.Setup,
		cocoapodsrepository.Setup,
		composerrepository.Setup,
		conanrepository.Setup,
		condarepository.Setup,
		cranrepository.Setup,
		debianrepository.Setup,
		dockerrepository.Setup,
		dockerv1repository.Setup,
		dockerv2repository.Setup,
		gemsrepository.Setup,
		genericrepository.Setup,
		gitlfsrepository.Setup,
		gorepository.Setup,
		gradlerepository.Setup,
		helmocirepository.Setup,
		helmrepository.Setup,
		huggingfacemlrepository.Setup,
		ivyrepository.Setup,
		mavenrepository.Setup,
		npmrepository.Setup,
		nugetrepository.Setup,
		ocirepository.Setup,
		opkgrepository.Setup,
		puppetrepository.Setup,
		pypirepository.Setup,
		rpmrepository.Setup,
		sbtrepository.Setup,
		swiftrepository.Setup,
		terraformmodulerepository.Setup,
		terraformproviderrepository.Setup,
		vagrantrepository.Setup,
		alpinerepositorylocal.Setup,
		ansiblerepositorylocal.Setup,
		bowerrepositorylocal.Setup,
		cargorepositorylocal.Setup,
		chefrepositorylocal.Setup,
		cocoapodsrepositorylocal.Setup,
		composerrepositorylocal.Setup,
		conanrepositorylocal.Setup,
		condarepositorylocal.Setup,
		cranrepositorylocal.Setup,
		debianrepositorylocal.Setup,
		dockerv1repositorylocal.Setup,
		dockerv2repositorylocal.Setup,
		gemsrepositorylocal.Setup,
		genericrepositorylocal.Setup,
		gitlfsrepositorylocal.Setup,
		gorepositorylocal.Setup,
		gradlerepositorylocal.Setup,
		helmocirepositorylocal.Setup,
		helmrepositorylocal.Setup,
		huggingfacemlrepositorylocal.Setup,
		ivyrepositorylocal.Setup,
		machinelearningrepository.Setup,
		mavenrepositorylocal.Setup,
		npmrepositorylocal.Setup,
		nugetrepositorylocal.Setup,
		ocirepositorylocal.Setup,
		opkgrepositorylocal.Setup,
		pubrepository.Setup,
		puppetrepositorylocal.Setup,
		pypirepositorylocal.Setup,
		rpmrepositorylocal.Setup,
		sbtrepositorylocal.Setup,
		swiftrepositorylocal.Setup,
		terraformbackendrepository.Setup,
		terraformmodulerepositorylocal.Setup,
		terraformproviderrepositorylocal.Setup,
		vagrantrepositorylocal.Setup,
		providerconfig.Setup,
		alpinerepositoryremote.Setup,
		ansiblerepositoryremote.Setup,
		bowerrepositoryremote.Setup,
		cargorepositoryremote.Setup,
		chefrepositoryremote.Setup,
		cocoapodsrepositoryremote.Setup,
		composerrepositoryremote.Setup,
		conanrepositoryremote.Setup,
		condarepositoryremote.Setup,
		cranrepositoryremote.Setup,
		debianrepositoryremote.Setup,
		dockerrepositoryremote.Setup,
		gemsrepositoryremote.Setup,
		genericrepositoryremote.Setup,
		gitlfsrepositoryremote.Setup,
		gorepositoryremote.Setup,
		gradlerepositoryremote.Setup,
		helmocirepositoryremote.Setup,
		helmrepositoryremote.Setup,
		huggingfacemlrepositoryremote.Setup,
		ivyrepositoryremote.Setup,
		mavenrepositoryremote.Setup,
		npmrepositoryremote.Setup,
		nugetrepositoryremote.Setup,
		ocirepositoryremote.Setup,
		opkgrepositoryremote.Setup,
		p2repository.Setup,
		pubrepositoryremote.Setup,
		puppetrepositoryremote.Setup,
		pypirepositoryremote.Setup,
		rpmrepositoryremote.Setup,
		sbtrepositoryremote.Setup,
		swiftrepositoryremote.Setup,
		terraformrepository.Setup,
		vcsrepository.Setup,
		keypair.Setup,
		scopedtoken.Setup,
		anonymoususer.Setup,
		alpinerepositoryvirtual.Setup,
		ansiblerepositoryvirtual.Setup,
		bowerrepositoryvirtual.Setup,
		cocoapodsrepositoryvirtual.Setup,
		composerrepositoryvirtual.Setup,
		conanrepositoryvirtual.Setup,
		condarepositoryvirtual.Setup,
		cranrepositoryvirtual.Setup,
		debianrepositoryvirtual.Setup,
		dockerrepositoryvirtual.Setup,
		gemsrepositoryvirtual.Setup,
		genericrepositoryvirtual.Setup,
		gitlfsrepositoryvirtual.Setup,
		gorepositoryvirtual.Setup,
		gradlerepositoryvirtual.Setup,
		helmocirepositoryvirtual.Setup,
		helmrepositoryvirtual.Setup,
		ivyrepositoryvirtual.Setup,
		mavenrepositoryvirtual.Setup,
		npmrepositoryvirtual.Setup,
		nugetrepositoryvirtual.Setup,
		ocirepositoryvirtual.Setup,
		p2repositoryvirtual.Setup,
		pubrepositoryvirtual.Setup,
		puppetrepositoryvirtual.Setup,
		pypirepositoryvirtual.Setup,
		rpmrepositoryvirtual.Setup,
		sbtrepositoryvirtual.Setup,
		swiftrepositoryvirtual.Setup,
		terraformrepositoryvirtual.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		itemproperties.SetupGated,
		backup.SetupGated,
		alpinerepository.SetupGated,
		ansiblerepository.SetupGated,
		bowerrepository.SetupGated,
		cargorepository.SetupGated,
		chefrepository.SetupGated,
		cocoapodsrepository.SetupGated,
		composerrepository.SetupGated,
		conanrepository.SetupGated,
		condarepository.SetupGated,
		cranrepository.SetupGated,
		debianrepository.SetupGated,
		dockerrepository.SetupGated,
		dockerv1repository.SetupGated,
		dockerv2repository.SetupGated,
		gemsrepository.SetupGated,
		genericrepository.SetupGated,
		gitlfsrepository.SetupGated,
		gorepository.SetupGated,
		gradlerepository.SetupGated,
		helmocirepository.SetupGated,
		helmrepository.SetupGated,
		huggingfacemlrepository.SetupGated,
		ivyrepository.SetupGated,
		mavenrepository.SetupGated,
		npmrepository.SetupGated,
		nugetrepository.SetupGated,
		ocirepository.SetupGated,
		opkgrepository.SetupGated,
		puppetrepository.SetupGated,
		pypirepository.SetupGated,
		rpmrepository.SetupGated,
		sbtrepository.SetupGated,
		swiftrepository.SetupGated,
		terraformmodulerepository.SetupGated,
		terraformproviderrepository.SetupGated,
		vagrantrepository.SetupGated,
		alpinerepositorylocal.SetupGated,
		ansiblerepositorylocal.SetupGated,
		bowerrepositorylocal.SetupGated,
		cargorepositorylocal.SetupGated,
		chefrepositorylocal.SetupGated,
		cocoapodsrepositorylocal.SetupGated,
		composerrepositorylocal.SetupGated,
		conanrepositorylocal.SetupGated,
		condarepositorylocal.SetupGated,
		cranrepositorylocal.SetupGated,
		debianrepositorylocal.SetupGated,
		dockerv1repositorylocal.SetupGated,
		dockerv2repositorylocal.SetupGated,
		gemsrepositorylocal.SetupGated,
		genericrepositorylocal.SetupGated,
		gitlfsrepositorylocal.SetupGated,
		gorepositorylocal.SetupGated,
		gradlerepositorylocal.SetupGated,
		helmocirepositorylocal.SetupGated,
		helmrepositorylocal.SetupGated,
		huggingfacemlrepositorylocal.SetupGated,
		ivyrepositorylocal.SetupGated,
		machinelearningrepository.SetupGated,
		mavenrepositorylocal.SetupGated,
		npmrepositorylocal.SetupGated,
		nugetrepositorylocal.SetupGated,
		ocirepositorylocal.SetupGated,
		opkgrepositorylocal.SetupGated,
		pubrepository.SetupGated,
		puppetrepositorylocal.SetupGated,
		pypirepositorylocal.SetupGated,
		rpmrepositorylocal.SetupGated,
		sbtrepositorylocal.SetupGated,
		swiftrepositorylocal.SetupGated,
		terraformbackendrepository.SetupGated,
		terraformmodulerepositorylocal.SetupGated,
		terraformproviderrepositorylocal.SetupGated,
		vagrantrepositorylocal.SetupGated,
		providerconfig.SetupGated,
		alpinerepositoryremote.SetupGated,
		ansiblerepositoryremote.SetupGated,
		bowerrepositoryremote.SetupGated,
		cargorepositoryremote.SetupGated,
		chefrepositoryremote.SetupGated,
		cocoapodsrepositoryremote.SetupGated,
		composerrepositoryremote.SetupGated,
		conanrepositoryremote.SetupGated,
		condarepositoryremote.SetupGated,
		cranrepositoryremote.SetupGated,
		debianrepositoryremote.SetupGated,
		dockerrepositoryremote.SetupGated,
		gemsrepositoryremote.SetupGated,
		genericrepositoryremote.SetupGated,
		gitlfsrepositoryremote.SetupGated,
		gorepositoryremote.SetupGated,
		gradlerepositoryremote.SetupGated,
		helmocirepositoryremote.SetupGated,
		helmrepositoryremote.SetupGated,
		huggingfacemlrepositoryremote.SetupGated,
		ivyrepositoryremote.SetupGated,
		mavenrepositoryremote.SetupGated,
		npmrepositoryremote.SetupGated,
		nugetrepositoryremote.SetupGated,
		ocirepositoryremote.SetupGated,
		opkgrepositoryremote.SetupGated,
		p2repository.SetupGated,
		pubrepositoryremote.SetupGated,
		puppetrepositoryremote.SetupGated,
		pypirepositoryremote.SetupGated,
		rpmrepositoryremote.SetupGated,
		sbtrepositoryremote.SetupGated,
		swiftrepositoryremote.SetupGated,
		terraformrepository.SetupGated,
		vcsrepository.SetupGated,
		keypair.SetupGated,
		scopedtoken.SetupGated,
		anonymoususer.SetupGated,
		alpinerepositoryvirtual.SetupGated,
		ansiblerepositoryvirtual.SetupGated,
		bowerrepositoryvirtual.SetupGated,
		cocoapodsrepositoryvirtual.SetupGated,
		composerrepositoryvirtual.SetupGated,
		conanrepositoryvirtual.SetupGated,
		condarepositoryvirtual.SetupGated,
		cranrepositoryvirtual.SetupGated,
		debianrepositoryvirtual.SetupGated,
		dockerrepositoryvirtual.SetupGated,
		gemsrepositoryvirtual.SetupGated,
		genericrepositoryvirtual.SetupGated,
		gitlfsrepositoryvirtual.SetupGated,
		gorepositoryvirtual.SetupGated,
		gradlerepositoryvirtual.SetupGated,
		helmocirepositoryvirtual.SetupGated,
		helmrepositoryvirtual.SetupGated,
		ivyrepositoryvirtual.SetupGated,
		mavenrepositoryvirtual.SetupGated,
		npmrepositoryvirtual.SetupGated,
		nugetrepositoryvirtual.SetupGated,
		ocirepositoryvirtual.SetupGated,
		p2repositoryvirtual.SetupGated,
		pubrepositoryvirtual.SetupGated,
		puppetrepositoryvirtual.SetupGated,
		pypirepositoryvirtual.SetupGated,
		rpmrepositoryvirtual.SetupGated,
		sbtrepositoryvirtual.SetupGated,
		swiftrepositoryvirtual.SetupGated,
		terraformrepositoryvirtual.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
