package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type PubsubSubscriptionIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionIamBindingSpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionIamBindingStatus `json:"status,omitempty"`
}

type PubsubSubscriptionIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag    string   `json:"etag,omitempty" tf:"etag,omitempty"`
	Members []string `json:"members" tf:"members"`
	// +optional
	Project      string `json:"project,omitempty" tf:"project,omitempty"`
	Role         string `json:"role" tf:"role"`
	Subscription string `json:"subscription" tf:"subscription"`
}

type PubsubSubscriptionIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PubsubSubscriptionIamBindingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionIamBindingList is a list of PubsubSubscriptionIamBindings
type PubsubSubscriptionIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscriptionIamBinding CRD objects
	Items []PubsubSubscriptionIamBinding `json:"items,omitempty"`
}
