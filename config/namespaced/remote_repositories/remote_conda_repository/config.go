package remotecondarepository

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_conda_repository", func(r *config.Resource) {
		// r.ShortGroup = ""
		// r.Kind = "RemoteCondaRepository"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
