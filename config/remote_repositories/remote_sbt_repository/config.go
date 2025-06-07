package remote_sbt_repository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = ""

// Configure configures the remote SBT repository.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_sbt_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RemoteSBTRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
