package releasebundlev2promotion

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_release_bundle_v2_promotion", func(r *config.Resource) {
		r.ShortGroup = "lifecycle"          // Otherwise, "release" is used
		r.Kind = "ReleaseBundleV2Promotion" // Otherwise, "BundleV2Promotion" is used
	})
}
