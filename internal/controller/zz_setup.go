// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	anonymoususer "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/anonymoususer"
	itemproperties "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/itemproperties"
	keypair "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/keypair"
	localansiblerepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localansiblerepository"
	localdockerv2repository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localdockerv2repository"
	localgenericrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localgenericrepository"
	localocirepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localocirepository"
	localterraformproviderrepository "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/artifactory/localterraformproviderrepository"
	providerconfig "github.com/hmlkao/provider-jfrog-artifactory/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		anonymoususer.Setup,
		itemproperties.Setup,
		keypair.Setup,
		localansiblerepository.Setup,
		localdockerv2repository.Setup,
		localgenericrepository.Setup,
		localocirepository.Setup,
		localterraformproviderrepository.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
