package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ApiManagementAPIVersionSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAPIVersionSetSpec   `json:"spec,omitempty"`
	Status            ApiManagementAPIVersionSetStatus `json:"status,omitempty"`
}

type ApiManagementAPIVersionSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	Description       string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName       string `json:"displayName" tf:"display_name"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	VersionHeaderName string `json:"versionHeaderName,omitempty" tf:"version_header_name,omitempty"`
	// +optional
	VersionQueryName string `json:"versionQueryName,omitempty" tf:"version_query_name,omitempty"`
	VersioningScheme string `json:"versioningScheme" tf:"versioning_scheme"`
}

type ApiManagementAPIVersionSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementAPIVersionSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAPIVersionSetList is a list of ApiManagementAPIVersionSets
type ApiManagementAPIVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAPIVersionSet CRD objects
	Items []ApiManagementAPIVersionSet `json:"items,omitempty"`
}
