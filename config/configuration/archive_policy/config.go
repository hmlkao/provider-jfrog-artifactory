package archivepolicy

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_archive_policy", func(r *config.Resource) {
		// Specify Kubernetes kind, otherwise, 'AnsibleRepository' is used
		r.Kind = "ArchivePolicy"
		// Import is supported using the following syntax
		//   terraform import artifactory_archive_policy.my-archive-policy my-policy
		//   terraform import artifactory_archive_policy.my-archive-policy my-policy:myproj
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			key, _ := tfstate["key"].(string)
			projectKey, _ := tfstate["project_key"].(string)
			if key != "" {
				if projectKey != "" {
					// Return policy id for project if project_key is set
					return key + ":" + projectKey, nil
				}
				// Return policy id
				return key, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}

		// TODO: How to set reference from another Terraform provider?
		//   In this case it's jfrog/project (https://registry.terraform.io/providers/jfrog/project/latest/docs)
		// r.References["project_key"] = config.Reference{
		// 	TerraformName: "project_repository",
		// }
	})
}
