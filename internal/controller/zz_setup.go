// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/anonymoususer"
	backup "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/backup"
	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/itemproperties"
	keypair "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/keypair"
	localalpinerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localalpinerepository"
	localansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localansiblerepository"
	localbowerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localbowerrepository"
	localcargorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localcargorepository"
	localchefrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localchefrepository"
	localcocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localcocoapodsrepository"
	localcomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localcomposerrepository"
	localconanrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localconanrepository"
	localcondarepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localcondarepository"
	localcranrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localcranrepository"
	localdebianrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localdebianrepository"
	localdockerv1repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localdockerv1repository"
	localdockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localdockerv2repository"
	localgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localgenericrepository"
	localgradlerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localgradlerepository"
	localhuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localhuggingfacemlrepository"
	localivyrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localivyrepository"
	localmachinelearningrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localmachinelearningrepository"
	localmavenrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localmavenrepository"
	localnpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localnpmrepository"
	localnugetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localnugetrepository"
	localocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localocirepository"
	localopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localopkgrepository"
	localpubrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localpubrepository"
	localpuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localpuppetrepository"
	localpypirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localpypirepository"
	localrpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localrpmrepository"
	localsbtrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localsbtrepository"
	localswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localswiftrepository"
	localterraformbackendrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localterraformbackendrepository"
	localterraformmodulerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localterraformmodulerepository"
	localterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localterraformproviderrepository"
	localvagrantrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localvagrantrepository"
	remotealpinerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotealpinerepository"
	remoteansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteansiblerepository"
	remotebowerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotebowerrepository"
	remotecargorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotecargorepository"
	remotechefrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotechefrepository"
	remotecocoapodsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotecocoapodsrepository"
	remotecomposerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotecomposerrepository"
	remotecondarepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotecondarepository"
	remotecranrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotecranrepository"
	remotedebianrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotedebianrepository"
	remotedockerrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotedockerrepository"
	remotegemsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotegemsrepository"
	remotegenericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotegenericrepository"
	remotegitlfsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotegitlfsrepository"
	remotegorepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotegorepository"
	remotegradlerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotegradlerepository"
	remotehelmocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotehelmocirepository"
	remotehelmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotehelmrepository"
	remotehuggingfacemlrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotehuggingfacemlrepository"
	remoteivyrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteivyrepository"
	remotemavenrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotemavenrepository"
	remotenpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotenpmrepository"
	remotenugetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotenugetrepository"
	remoteocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteocirepository"
	remoteopkgrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteopkgrepository"
	remotep2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotep2repository"
	remotepubrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotepubrepository"
	remotepuppetrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotepuppetrepository"
	remotepypirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotepypirepository"
	remoterpmrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoterpmrepository"
	remotesbtrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotesbtrepository"
	remoteswiftrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteswiftrepository"
	remoteterraformrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remoteterraformrepository"
	remotevcsrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/remotevcsrepository"
	providerconfig "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		anonymoususer.Setup,
		backup.Setup,
		itemproperties.Setup,
		keypair.Setup,
		localalpinerepository.Setup,
		localansiblerepository.Setup,
		localbowerrepository.Setup,
		localcargorepository.Setup,
		localchefrepository.Setup,
		localcocoapodsrepository.Setup,
		localcomposerrepository.Setup,
		localconanrepository.Setup,
		localcondarepository.Setup,
		localcranrepository.Setup,
		localdebianrepository.Setup,
		localdockerv1repository.Setup,
		localdockerv2repository.Setup,
		localgenericrepository.Setup,
		localgradlerepository.Setup,
		localhuggingfacemlrepository.Setup,
		localivyrepository.Setup,
		localmachinelearningrepository.Setup,
		localmavenrepository.Setup,
		localnpmrepository.Setup,
		localnugetrepository.Setup,
		localocirepository.Setup,
		localopkgrepository.Setup,
		localpubrepository.Setup,
		localpuppetrepository.Setup,
		localpypirepository.Setup,
		localrpmrepository.Setup,
		localsbtrepository.Setup,
		localswiftrepository.Setup,
		localterraformbackendrepository.Setup,
		localterraformmodulerepository.Setup,
		localterraformproviderrepository.Setup,
		localvagrantrepository.Setup,
		remotealpinerepository.Setup,
		remoteansiblerepository.Setup,
		remotebowerrepository.Setup,
		remotecargorepository.Setup,
		remotechefrepository.Setup,
		remotecocoapodsrepository.Setup,
		remotecomposerrepository.Setup,
		remotecondarepository.Setup,
		remotecranrepository.Setup,
		remotedebianrepository.Setup,
		remotedockerrepository.Setup,
		remotegemsrepository.Setup,
		remotegenericrepository.Setup,
		remotegitlfsrepository.Setup,
		remotegorepository.Setup,
		remotegradlerepository.Setup,
		remotehelmocirepository.Setup,
		remotehelmrepository.Setup,
		remotehuggingfacemlrepository.Setup,
		remoteivyrepository.Setup,
		remotemavenrepository.Setup,
		remotenpmrepository.Setup,
		remotenugetrepository.Setup,
		remoteocirepository.Setup,
		remoteopkgrepository.Setup,
		remotep2repository.Setup,
		remotepubrepository.Setup,
		remotepuppetrepository.Setup,
		remotepypirepository.Setup,
		remoterpmrepository.Setup,
		remotesbtrepository.Setup,
		remoteswiftrepository.Setup,
		remoteterraformrepository.Setup,
		remotevcsrepository.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
