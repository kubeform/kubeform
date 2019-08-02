package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PubsubTopicIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubTopicIamMemberSpec   `json:"spec,omitempty"`
	Status            PubsubTopicIamMemberStatus `json:"status,omitempty"`
}

type PubsubTopicIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Member string `json:"member" tf:"member"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Role    string `json:"role" tf:"role"`
	Topic   string `json:"topic" tf:"topic"`
}

type PubsubTopicIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *PubsubTopicIamMemberSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubTopicIamMemberList is a list of PubsubTopicIamMembers
type PubsubTopicIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubTopicIamMember CRD objects
	Items []PubsubTopicIamMember `json:"items,omitempty"`
}
