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

type IamUserGroupMembership struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserGroupMembershipSpec   `json:"spec,omitempty"`
	Status            IamUserGroupMembershipStatus `json:"status,omitempty"`
}

type IamUserGroupMembershipSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +kubebuilder:validation:UniqueItems=true
	Groups []string `json:"groups" tf:"groups"`
	User   string   `json:"user" tf:"user"`
}

type IamUserGroupMembershipStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamUserGroupMembershipList is a list of IamUserGroupMemberships
type IamUserGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamUserGroupMembership CRD objects
	Items []IamUserGroupMembership `json:"items,omitempty"`
}
