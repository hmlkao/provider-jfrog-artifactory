// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ItemPropertiesInitParameters struct {

	// (Boolean) Add this property to the selected folder and to all of artifacts and folders under this folder. Default to false
	// Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
	IsRecursive *bool `json:"isRecursive,omitempty" tf:"is_recursive,omitempty"`

	// (String) The relative path of the item (file/folder/repository). Leave unset for repository.
	// The relative path of the item (file/folder/repository). Leave unset for repository.
	ItemPath *string `json:"itemPath,omitempty" tf:"item_path,omitempty"`

	// (Map of Set of String) Map of key and list of values.
	// Map of key and list of values.
	//
	// ~>Keys are limited up to 255 characters and values are limited up to 2,400 characters. Using properties with values over this limit might cause backend issues.
	//
	// ~>The following special characters are forbidden in the key field: `)(}{][*+^$/~“!@#%&<>;=,±§` and the space character.
	Properties map[string][]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// (String) Respository key.
	// Respository key.
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`
}

type ItemPropertiesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Add this property to the selected folder and to all of artifacts and folders under this folder. Default to false
	// Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
	IsRecursive *bool `json:"isRecursive,omitempty" tf:"is_recursive,omitempty"`

	// (String) The relative path of the item (file/folder/repository). Leave unset for repository.
	// The relative path of the item (file/folder/repository). Leave unset for repository.
	ItemPath *string `json:"itemPath,omitempty" tf:"item_path,omitempty"`

	// (Map of Set of String) Map of key and list of values.
	// Map of key and list of values.
	//
	// ~>Keys are limited up to 255 characters and values are limited up to 2,400 characters. Using properties with values over this limit might cause backend issues.
	//
	// ~>The following special characters are forbidden in the key field: `)(}{][*+^$/~“!@#%&<>;=,±§` and the space character.
	Properties map[string][]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// (String) Respository key.
	// Respository key.
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`
}

type ItemPropertiesParameters struct {

	// (Boolean) Add this property to the selected folder and to all of artifacts and folders under this folder. Default to false
	// Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
	// +kubebuilder:validation:Optional
	IsRecursive *bool `json:"isRecursive,omitempty" tf:"is_recursive,omitempty"`

	// (String) The relative path of the item (file/folder/repository). Leave unset for repository.
	// The relative path of the item (file/folder/repository). Leave unset for repository.
	// +kubebuilder:validation:Optional
	ItemPath *string `json:"itemPath,omitempty" tf:"item_path,omitempty"`

	// (Map of Set of String) Map of key and list of values.
	// Map of key and list of values.
	//
	// ~>Keys are limited up to 255 characters and values are limited up to 2,400 characters. Using properties with values over this limit might cause backend issues.
	//
	// ~>The following special characters are forbidden in the key field: `)(}{][*+^$/~“!@#%&<>;=,±§` and the space character.
	// +kubebuilder:validation:Optional
	Properties map[string][]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// (String) Respository key.
	// Respository key.
	// +kubebuilder:validation:Optional
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`
}

// ItemPropertiesSpec defines the desired state of ItemProperties
type ItemPropertiesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ItemPropertiesParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ItemPropertiesInitParameters `json:"initProvider,omitempty"`
}

// ItemPropertiesStatus defines the observed state of ItemProperties.
type ItemPropertiesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ItemPropertiesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ItemProperties is the Schema for the ItemPropertiess API. Provides a resource for managaing item (file, folder, or repository) properties. When a folder is used property attachment is recursive by default. See JFrog documentation https://jfrog.com/help/r/jfrog-artifactory-documentation/working-with-jfrog-properties for more details.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,artifactory}
type ItemProperties struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.properties) || (has(self.initProvider) && has(self.initProvider.properties))",message="spec.forProvider.properties is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repoKey) || (has(self.initProvider) && has(self.initProvider.repoKey))",message="spec.forProvider.repoKey is a required parameter"
	Spec   ItemPropertiesSpec   `json:"spec"`
	Status ItemPropertiesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ItemPropertiesList contains a list of ItemPropertiess
type ItemPropertiesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ItemProperties `json:"items"`
}

// Repository type metadata.
var (
	ItemProperties_Kind             = "ItemProperties"
	ItemProperties_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ItemProperties_Kind}.String()
	ItemProperties_KindAPIVersion   = ItemProperties_Kind + "." + CRDGroupVersion.String()
	ItemProperties_GroupVersionKind = CRDGroupVersion.WithKind(ItemProperties_Kind)
)

func init() {
	SchemeBuilder.Register(&ItemProperties{}, &ItemPropertiesList{})
}
