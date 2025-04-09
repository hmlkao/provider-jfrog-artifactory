package keypair

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = "artifactory"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_keypair", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		// Specify Kubernetes kind
		r.Kind = "Keypair"
		// Set custom func to get external name because there is no 'id' stored in Terraform state
		// Keypair can be imported using the pair name, e.g.
		//   terraform import artifactory_keypair.my-keypair my-keypair-name
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["pair_name"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find pair_name in tfstate")
		}
	})
}
