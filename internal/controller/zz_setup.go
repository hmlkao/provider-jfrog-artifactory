// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifact/itemproperties"
	backup "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/configuration/backup"
	alpinerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/alpinerepository"
	ansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/ansiblerepository"
	bowerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/bowerrepository"
	cargorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/cargorepository"
	chefrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/chefrepository"
	cocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/cocoapodsrepository"
	composerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/composerrepository"
	conanrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/conanrepository"
	condarepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/condarepository"
	cranrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/cranrepository"
	debianrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/debianrepository"
	dockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/dockerv1repository"
	dockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/dockerv2repository"
	genericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/genericrepository"
	gradlerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/gradlerepository"
	huggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/huggingfacemlrepository"
	ivyrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/ivyrepository"
	machinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/machinelearningrepository"
	mavenrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/mavenrepository"
	npmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/npmrepository"
	nugetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/nugetrepository"
	ocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/ocirepository"
	opkgrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/opkgrepository"
	pubrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/pubrepository"
	puppetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/puppetrepository"
	pypirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/pypirepository"
	rpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/rpmrepository"
	sbtrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/sbtrepository"
	swiftrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/swiftrepository"
	terraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/terraformbackendrepository"
	terraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/terraformmodulerepository"
	terraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/terraformproviderrepository"
	vagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/local/vagrantrepository"
	providerconfig "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/providerconfig"
	alpinerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/alpinerepository"
	ansiblerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/ansiblerepository"
	bowerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/bowerrepository"
	cargorepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/cargorepository"
	chefrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/chefrepository"
	cocoapodsrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/cocoapodsrepository"
	composerrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/composerrepository"
	condarepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/condarepository"
	cranrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/cranrepository"
	debianrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/debianrepository"
	dockerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/dockerrepository"
	gemsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/gemsrepository"
	genericrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/genericrepository"
	gitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/gitlfsrepository"
	gorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/gorepository"
	gradlerepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/gradlerepository"
	helmocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/helmocirepository"
	helmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/helmrepository"
	huggingfacemlrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/huggingfacemlrepository"
	ivyrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/ivyrepository"
	mavenrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/mavenrepository"
	npmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/npmrepository"
	nugetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/nugetrepository"
	ocirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/ocirepository"
	opkgrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/opkgrepository"
	p2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/p2repository"
	pubrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/pubrepository"
	puppetrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/puppetrepository"
	pypirepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/pypirepository"
	rpmrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/rpmrepository"
	sbtrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/sbtrepository"
	swiftrepositoryremote "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/swiftrepository"
	terraformrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/terraformrepository"
	vcsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/remote/vcsrepository"
	keypair "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/security/keypair"
	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/user/anonymoususer"
	alpinerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/alpinerepository"
	ansiblerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/ansiblerepository"
	bowerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/bowerrepository"
	cocoapodsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/cocoapodsrepository"
	composerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/composerrepository"
	conanrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/conanrepository"
	condarepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/condarepository"
	cranrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/cranrepository"
	debianrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/debianrepository"
	dockerrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/dockerrepository"
	gemsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/gemsrepository"
	genericrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/genericrepository"
	gitlfsrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/gitlfsrepository"
	gorepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/gorepository"
	gradlerepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/gradlerepository"
	helmocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/helmocirepository"
	helmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/helmrepository"
	ivyrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/ivyrepository"
	mavenrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/mavenrepository"
	npmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/npmrepository"
	nugetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/nugetrepository"
	ocirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/ocirepository"
	p2repositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/p2repository"
	pubrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/pubrepository"
	puppetrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/puppetrepository"
	pypirepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/pypirepository"
	rpmrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/rpmrepository"
	sbtrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/sbtrepository"
	swiftrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/swiftrepository"
	terraformrepositoryvirtual "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/virtual/terraformrepository"
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
		genericrepository.Setup,
		gradlerepository.Setup,
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
		condarepositoryremote.Setup,
		cranrepositoryremote.Setup,
		debianrepositoryremote.Setup,
		dockerrepository.Setup,
		gemsrepository.Setup,
		genericrepositoryremote.Setup,
		gitlfsrepository.Setup,
		gorepository.Setup,
		gradlerepositoryremote.Setup,
		helmocirepository.Setup,
		helmrepository.Setup,
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
