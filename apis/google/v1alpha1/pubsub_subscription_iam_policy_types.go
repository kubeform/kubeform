package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PubsubSubscriptionIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionIamPolicySpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionIamPolicyStatus `json:"status,omitempty"`
}

type PubsubSubscriptionIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	PolicyData string `json:"policyData" tf:"policy_data"`
	// +optional
	Project      string `json:"project,omitempty" tf:"project,omitempty"`
	Subscription string `json:"subscription" tf:"subscription"`
}

type PubsubSubscriptionIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionIamPolicyList is a list of PubsubSubscriptionIamPolicys
type PubsubSubscriptionIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscriptionIamPolicy CRD objects
	Items []PubsubSubscriptionIamPolicy `json:"items,omitempty"`
}
