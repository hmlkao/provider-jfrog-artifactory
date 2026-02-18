package certificate

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_certificate", func(r *config.Resource) {
		r.ShortGroup = "security" // Otherwise, '' is used
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			// Certificate can be imported using the alias, e.g.
			//   terraform import artifactory_certificate.my-cert my-cert-alias
			if id, ok := tfstate["alias"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'alias' in tfstate")
		}
	})
}
