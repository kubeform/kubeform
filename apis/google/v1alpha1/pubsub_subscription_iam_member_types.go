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

type PubsubSubscriptionIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionIamMemberSpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionIamMemberStatus `json:"status,omitempty"`
}

type PubsubSubscriptionIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Member string `json:"member" tf:"member"`
	// +optional
	Project      string `json:"project,omitempty" tf:"project,omitempty"`
	Role         string `json:"role" tf:"role"`
	Subscription string `json:"subscription" tf:"subscription"`
}

type PubsubSubscriptionIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionIamMemberList is a list of PubsubSubscriptionIamMembers
type PubsubSubscriptionIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscriptionIamMember CRD objects
	Items []PubsubSubscriptionIamMember `json:"items,omitempty"`
}
