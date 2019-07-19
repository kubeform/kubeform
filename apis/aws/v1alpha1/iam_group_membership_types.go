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

type IamGroupMembership struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamGroupMembershipSpec   `json:"spec,omitempty"`
	Status            IamGroupMembershipStatus `json:"status,omitempty"`
}

type IamGroupMembershipSpec struct {
	Group string `json:"group" tf:"group"`
	Name  string `json:"name" tf:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Users       []string                  `json:"users" tf:"users"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IamGroupMembershipStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamGroupMembershipList is a list of IamGroupMemberships
type IamGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamGroupMembership CRD objects
	Items []IamGroupMembership `json:"items,omitempty"`
}
