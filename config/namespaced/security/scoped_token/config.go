package scoped_token

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_scoped_token", func(r *config.Resource) {
		r.ShortGroup = "security" // Otherwise, 'scoped' is used
		r.Kind = "ScopedToken"    // Otherwise, 'Token' is used
		r.ExternalName.DisableNameInitializer = false
	})
}
