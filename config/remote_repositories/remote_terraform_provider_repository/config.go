package remote_terraform_provider_repository

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the remote Terraform Provider repository.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_terraform_provider_repository", func(r *config.Resource) {
		// r.ShortGroup = ""
		// r.Kind = "RemoteTerraformProviderRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
