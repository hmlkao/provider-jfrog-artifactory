package localocirepository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = "artifactory"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_oci_repository", func(r *config.Resource) {
		// TODO: Get rid of ShortGroup config and use artifactory.jfrog.crossplane.io for initial config instead
		//   This is naming convention dilema beween:
		//     - jfrog.crossplane.io as base group and set ShortGroup to artifactory which will produce artifactory.jfrog.crossplane.io anyway and resource is just LocalOCIRepository (current solution)
		//     - jfrog.crossplane.io as base group and resource prefix to be artifactory, like ArtifactoryLocalOCIRepository
		//     - artifactory.jfrog.crossplane.io as base group and resource is just LocalOCIRepository (no need to set up ShortGroup)
		r.ShortGroup = shortGroup
		// Specify Kubernetes kind
		r.Kind = "LocalOCIRepository"
		// Fix an issue with ExternalName configuration
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find id in tfstate")
		}
	})
}
