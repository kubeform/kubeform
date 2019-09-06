package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementProperty struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementPropertySpec   `json:"spec,omitempty"`
	Status            ApiManagementPropertyStatus `json:"status,omitempty"`
}

type ApiManagementPropertySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	DisplayName       string `json:"displayName" tf:"display_name"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Secret bool `json:"secret,omitempty" tf:"secret,omitempty"`
	// +optional
	Tags  []string `json:"tags,omitempty" tf:"tags,omitempty"`
	Value string   `json:"value" tf:"value"`
}

type ApiManagementPropertyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementPropertySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementPropertyList is a list of ApiManagementPropertys
type ApiManagementPropertyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementProperty CRD objects
	Items []ApiManagementProperty `json:"items,omitempty"`
}
