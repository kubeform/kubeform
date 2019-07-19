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

type ApiManagementGroupUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementGroupUserSpec   `json:"spec,omitempty"`
	Status            ApiManagementGroupUserStatus `json:"status,omitempty"`
}

type ApiManagementGroupUserSpec struct {
	ApiManagementName string                    `json:"apiManagementName" tf:"api_management_name"`
	GroupName         string                    `json:"groupName" tf:"group_name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	UserID            string                    `json:"userID" tf:"user_id"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiManagementGroupUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementGroupUserList is a list of ApiManagementGroupUsers
type ApiManagementGroupUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementGroupUser CRD objects
	Items []ApiManagementGroupUser `json:"items,omitempty"`
}
