package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BillingAccountIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountIamPolicySpec   `json:"spec,omitempty"`
	Status            BillingAccountIamPolicyStatus `json:"status,omitempty"`
}

type BillingAccountIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BillingAccountID string `json:"billingAccountID" tf:"billing_account_id"`
	PolicyData       string `json:"policyData" tf:"policy_data"`
}

type BillingAccountIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *BillingAccountIamPolicySpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BillingAccountIamPolicyList is a list of BillingAccountIamPolicys
type BillingAccountIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BillingAccountIamPolicy CRD objects
	Items []BillingAccountIamPolicy `json:"items,omitempty"`
}
