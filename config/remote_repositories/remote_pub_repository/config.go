package remote_pub_repository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the remote Pub repository.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_pub_repository", func(r *config.Resource) {
		// r.ShortGroup = ""
		// r.Kind = "RemotePubRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
