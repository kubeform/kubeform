package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SqsQueuePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqsQueuePolicySpec   `json:"spec,omitempty"`
	Status            SqsQueuePolicyStatus `json:"status,omitempty"`
}

type SqsQueuePolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Policy   string `json:"policy" tf:"policy"`
	QueueURL string `json:"queueURL" tf:"queue_url"`
}

type SqsQueuePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SqsQueuePolicySpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqsQueuePolicyList is a list of SqsQueuePolicys
type SqsQueuePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqsQueuePolicy CRD objects
	Items []SqsQueuePolicy `json:"items,omitempty"`
}
