package remotecargorepository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = ""

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_cargo_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RemoteCargoRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
