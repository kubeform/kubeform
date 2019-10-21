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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SnsTopicPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicPolicySpec   `json:"spec,omitempty"`
	Status            SnsTopicPolicyStatus `json:"status,omitempty"`
}

type SnsTopicPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Arn    string `json:"arn" tf:"arn"`
	Policy string `json:"policy" tf:"policy"`
}

type SnsTopicPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnsTopicPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsTopicPolicyList is a list of SnsTopicPolicys
type SnsTopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsTopicPolicy CRD objects
	Items []SnsTopicPolicy `json:"items,omitempty"`
}