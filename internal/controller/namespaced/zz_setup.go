// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/artifact/itemproperties"
	backup "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/configuration/backup"
	alpinerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/alpinerepository"
	ansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/ansiblerepository"
	bowerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/bowerrepository"
	cargorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/cargorepository"
	chefrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/chefrepository"
	cocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/cocoapodsrepository"
	composerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/composerrepository"
	conanrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/conanrepository"
	condarepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/condarepository"
	cranrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/cranrepository"
	debianrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/debianrepository"
	dockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/dockerv1repository"
	dockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/dockerv2repository"
	gemsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/gemsrepository"
	genericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/genericrepository"
	gitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/gitlfsrepository"
	gorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/gorepository"
	gradlerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/gradlerepository"
	helmocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/helmocirepository"
	helmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/helmrepository"
	huggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/huggingfacemlrepository"
	ivyrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/ivyrepository"
	machinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/machinelearningrepository"
	mavenrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/mavenrepository"
	npmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/npmrepository"
	nugetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/nugetrepository"
	ocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/ocirepository"
	opkgrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/opkgrepository"
	pubrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/pubrepository"
	puppetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/puppetrepository"
	pypirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/pypirepository"
	rpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/rpmrepository"
	sbtrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/sbtrepository"
	swiftrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/swiftrepository"
	terraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/terraformbackendrepository"
	terraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/terraformmodulerepository"
	terraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/terraformproviderrepository"
	vagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/local/vagrantrepository"
	providerconfig "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/providerconfig"
	alpinerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/alpinerepository"
	ansiblerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/ansiblerepository"
	bowerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/bowerrepository"
	cargorepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/cargorepository"
	chefrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/chefrepository"
	cocoapodsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/cocoapodsrepository"
	composerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/composerrepository"
	conanrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/conanrepository"
	condarepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/condarepository"
	cranrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/cranrepository"
	debianrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/debianrepository"
	dockerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/dockerrepository"
	gemsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/gemsrepository"
	genericrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/genericrepository"
	gitlfsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/gitlfsrepository"
	gorepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/gorepository"
	gradlerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/gradlerepository"
	helmocirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/helmocirepository"
	helmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/helmrepository"
	huggingfacemlrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/huggingfacemlrepository"
	ivyrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/ivyrepository"
	mavenrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/mavenrepository"
	npmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/npmrepository"
	nugetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/nugetrepository"
	ocirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/ocirepository"
	opkgrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/opkgrepository"
	p2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/p2repository"
	pubrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/pubrepository"
	puppetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/puppetrepository"
	pypirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/pypirepository"
	rpmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/rpmrepository"
	sbtrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/sbtrepository"
	swiftrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/swiftrepository"
	terraformrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/terraformrepository"
	vcsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/remote/vcsrepository"
	keypair "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/security/keypair"
	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/user/anonymoususer"
	alpinerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/alpinerepository"
	ansiblerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/ansiblerepository"
	bowerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/bowerrepository"
	cocoapodsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/cocoapodsrepository"
	composerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/composerrepository"
	conanrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/conanrepository"
	condarepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/condarepository"
	cranrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/cranrepository"
	debianrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/debianrepository"
	dockerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/dockerrepository"
	gemsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/gemsrepository"
	genericrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/genericrepository"
	gitlfsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/gitlfsrepository"
	gorepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/gorepository"
	gradlerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/gradlerepository"
	helmocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/helmocirepository"
	helmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/helmrepository"
	ivyrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/ivyrepository"
	mavenrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/mavenrepository"
	npmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/npmrepository"
	nugetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/nugetrepository"
	ocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/ocirepository"
	p2repositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/p2repository"
	pubrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/pubrepository"
	puppetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/puppetrepository"
	pypirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/pypirepository"
	rpmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/rpmrepository"
	sbtrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/sbtrepository"
	swiftrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/swiftrepository"
	terraformrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/namespaced/virtual/terraformrepository"
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
		machinelearningrepository.Setup,
		mavenrepository.Setup,
		npmrepository.Setup,
		nugetrepository.Setup,
		ocirepository.Setup,
		opkgrepository.Setup,
		pubrepository.Setup,
		puppetrepository.Setup,
		pypirepository.Setup,
		rpmrepository.Setup,
		sbtrepository.Setup,
		swiftrepository.Setup,
		terraformbackendrepository.Setup,
		terraformmodulerepository.Setup,
		terraformproviderrepository.Setup,
		vagrantrepository.Setup,
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
		dockerrepository.Setup,
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
		machinelearningrepository.SetupGated,
		mavenrepository.SetupGated,
		npmrepository.SetupGated,
		nugetrepository.SetupGated,
		ocirepository.SetupGated,
		opkgrepository.SetupGated,
		pubrepository.SetupGated,
		puppetrepository.SetupGated,
		pypirepository.SetupGated,
		rpmrepository.SetupGated,
		sbtrepository.SetupGated,
		swiftrepository.SetupGated,
		terraformbackendrepository.SetupGated,
		terraformmodulerepository.SetupGated,
		terraformproviderrepository.SetupGated,
		vagrantrepository.SetupGated,
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
		dockerrepository.SetupGated,
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
