package localocirepository

import "github.com/crossplane/upjet/pkg/config"

const shortGroup string = "artifactory"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_oci_repository", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "artifactory"
		r.ShortGroup = shortGroup
		r.Kind = "LocalOCIRepository"
	})
}
