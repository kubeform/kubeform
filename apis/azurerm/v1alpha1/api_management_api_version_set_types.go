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

type ApiManagementApiVersionSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiVersionSetSpec   `json:"spec,omitempty"`
	Status            ApiManagementApiVersionSetStatus `json:"status,omitempty"`
}

type ApiManagementApiVersionSetSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	Description       string `json:"description,omitempty"`
	DisplayName       string `json:"display_name"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	VersionHeaderName string `json:"version_header_name,omitempty"`
	// +optional
	VersionQueryName string `json:"version_query_name,omitempty"`
	VersioningScheme string `json:"versioning_scheme"`
}

type ApiManagementApiVersionSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementApiVersionSetList is a list of ApiManagementApiVersionSets
type ApiManagementApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementApiVersionSet CRD objects
	Items []ApiManagementApiVersionSet `json:"items,omitempty"`
}
