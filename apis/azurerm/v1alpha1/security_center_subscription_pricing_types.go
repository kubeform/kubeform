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

type SecurityCenterSubscriptionPricing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterSubscriptionPricingSpec   `json:"spec,omitempty"`
	Status            SecurityCenterSubscriptionPricingStatus `json:"status,omitempty"`
}

type SecurityCenterSubscriptionPricingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Tier string `json:"tier" tf:"tier"`
}

type SecurityCenterSubscriptionPricingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SecurityCenterSubscriptionPricingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityCenterSubscriptionPricingList is a list of SecurityCenterSubscriptionPricings
type SecurityCenterSubscriptionPricingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityCenterSubscriptionPricing CRD objects
	Items []SecurityCenterSubscriptionPricing `json:"items,omitempty"`
}
