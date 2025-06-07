package remote_p2_repository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = ""

// Configure configures the remote P2 repository.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_p2_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RemoteP2Repository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
