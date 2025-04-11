package localterraformproviderrepository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = "" // Otherwise, 'local' is used

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_terraform_provider_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		// Specify Kubernetes kind, otherwise, 'TerraformProviderRepository' is used
		r.Kind = "LocalTerraformProviderRepository"
		// Set custom func to get external name because there is no 'id' stored in Terraform state
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
