package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementProductGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductGroupSpec   `json:"spec,omitempty"`
	Status            ApiManagementProductGroupStatus `json:"status,omitempty"`
}

type ApiManagementProductGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	GroupName         string `json:"groupName" tf:"group_name"`
	ProductID         string `json:"productID" tf:"product_id"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type ApiManagementProductGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementProductGroupList is a list of ApiManagementProductGroups
type ApiManagementProductGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementProductGroup CRD objects
	Items []ApiManagementProductGroup `json:"items,omitempty"`
}
