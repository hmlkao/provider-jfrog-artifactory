package remote_vcs_repository

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the remote VCS repository.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_vcs_repository", func(r *config.Resource) {
		r.Kind = "VCSRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
