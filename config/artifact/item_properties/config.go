package itemproperties

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_item_properties", func(r *config.Resource) {
		r.ShortGroup = "artifact" // Otherwise, 'item' is used
		r.Kind = "ItemProperties" // Otherwise, 'Properties' is used
		// Set custom func to get external name because there is no 'id' stored in Terraform state
		// Import is supported using the following syntax:
		//   terraform import artifactory_item_properties.my-repo-properties repo_key
		//   terraform import artifactory_item_properties.my-folder-properties repo_key:folder/subfolder
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			itemPath, _ := tfstate["item_path"].(string)
			repoKey, _ := tfstate["repo_key"].(string)
			if repoKey != "" {
				if itemPath != "" {
					// Return id for an item if item_path is set
					return repoKey + ":" + itemPath, nil
				}
				// Return id for a repo
				return repoKey, nil
			}
			return "", errors.New("cannot find 'repo_key' in tfstate")
		}
	})
}
