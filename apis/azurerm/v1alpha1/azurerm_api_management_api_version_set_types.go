package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermApiManagementApiVersionSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiVersionSetSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiVersionSetStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiVersionSetSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	DisplayName       string `json:"display_name"`
	VersioningScheme  string `json:"versioning_scheme"`
	Description       string `json:"description"`
	VersionHeaderName string `json:"version_header_name"`
	VersionQueryName  string `json:"version_query_name"`
}



type AzurermApiManagementApiVersionSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementApiVersionSetList is a list of AzurermApiManagementApiVersionSets
type AzurermApiManagementApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiVersionSet CRD objects
	Items []AzurermApiManagementApiVersionSet `json:"items,omitempty"`
}