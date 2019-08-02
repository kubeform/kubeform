package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesReceiptFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesReceiptFilterSpec   `json:"spec,omitempty"`
	Status            SesReceiptFilterStatus `json:"status,omitempty"`
}

type SesReceiptFilterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Cidr   string `json:"cidr" tf:"cidr"`
	Name   string `json:"name" tf:"name"`
	Policy string `json:"policy" tf:"policy"`
}

type SesReceiptFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SesReceiptFilterSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesReceiptFilterList is a list of SesReceiptFilters
type SesReceiptFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesReceiptFilter CRD objects
	Items []SesReceiptFilter `json:"items,omitempty"`
}
