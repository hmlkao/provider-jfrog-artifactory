package anonymoususer

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_anonymous_user", func(r *config.Resource) {
		r.ShortGroup = "user" // Otherwise, 'anonymous' is used
		r.Kind = "AnonymousUser"
	})
}
