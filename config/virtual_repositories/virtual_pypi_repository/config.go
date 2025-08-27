package virtualpyyrepository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_virtual_pypi_repository", func(r *config.Resource) {
		// r.ShortGroup = ""
		r.Kind = "PyPIRepository" // Otherwise, "PypiRepository" is used
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
