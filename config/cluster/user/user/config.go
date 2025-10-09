package user

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_user", func(r *config.Resource) {
		r.ShortGroup = "user" // Otherwise, '' is used
	})
}
