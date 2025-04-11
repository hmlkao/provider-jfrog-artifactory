package manageduser

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = "" // Otherwise, 'user' is used

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_managed_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
