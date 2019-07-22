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

type ApiManagementUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementUserSpec   `json:"spec,omitempty"`
	Status            ApiManagementUserStatus `json:"status,omitempty"`
}

type ApiManagementUserSpec struct {
	Secret            *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef       core.LocalObjectReference  `json:"providerRef" tf:"-"`
	ApiManagementName string                     `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	Confirmation string `json:"confirmation,omitempty" tf:"confirmation,omitempty"`
	Email        string `json:"email" tf:"email"`
	FirstName    string `json:"firstName" tf:"first_name"`
	LastName     string `json:"lastName" tf:"last_name"`
	// +optional
	Note string `json:"note,omitempty" tf:"note,omitempty"`
	// +optional
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	State  string `json:"state,omitempty" tf:"state,omitempty"`
	UserID string `json:"userID" tf:"user_id"`
}

type ApiManagementUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementUserList is a list of ApiManagementUsers
type ApiManagementUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementUser CRD objects
	Items []ApiManagementUser `json:"items,omitempty"`
}
