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

type PubsubTopicIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubTopicIamBindingSpec   `json:"spec,omitempty"`
	Status            PubsubTopicIamBindingStatus `json:"status,omitempty"`
}

type PubsubTopicIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag    string   `json:"etag,omitempty" tf:"etag,omitempty"`
	Members []string `json:"members" tf:"members"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Role    string `json:"role" tf:"role"`
	Topic   string `json:"topic" tf:"topic"`
}

type PubsubTopicIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PubsubTopicIamBindingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubTopicIamBindingList is a list of PubsubTopicIamBindings
type PubsubTopicIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubTopicIamBinding CRD objects
	Items []PubsubTopicIamBinding `json:"items,omitempty"`
}
