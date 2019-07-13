package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermSecurityCenterSubscriptionPricing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSecurityCenterSubscriptionPricingSpec   `json:"spec,omitempty"`
	Status            AzurermSecurityCenterSubscriptionPricingStatus `json:"status,omitempty"`
}

type AzurermSecurityCenterSubscriptionPricingSpec struct {
	Tier string `json:"tier"`
}



type AzurermSecurityCenterSubscriptionPricingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSecurityCenterSubscriptionPricingList is a list of AzurermSecurityCenterSubscriptionPricings
type AzurermSecurityCenterSubscriptionPricingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSecurityCenterSubscriptionPricing CRD objects
	Items []AzurermSecurityCenterSubscriptionPricing `json:"items,omitempty"`
}