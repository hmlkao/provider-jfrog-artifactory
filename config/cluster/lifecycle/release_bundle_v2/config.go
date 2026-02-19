package releasebundlev2

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_release_bundle_v2", func(r *config.Resource) {
		r.ShortGroup = "lifecycle" // Otherwise, "release" is used
		r.Kind = "ReleaseBundleV2" // Otherwise, "BundleV2" is used
	})
}
