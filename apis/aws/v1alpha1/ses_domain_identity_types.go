package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesDomainIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesDomainIdentitySpec   `json:"spec,omitempty"`
	Status            SesDomainIdentityStatus `json:"status,omitempty"`
}

type SesDomainIdentitySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Domain string `json:"domain" tf:"domain"`
}

type SesDomainIdentityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SesDomainIdentitySpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesDomainIdentityList is a list of SesDomainIdentitys
type SesDomainIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesDomainIdentity CRD objects
	Items []SesDomainIdentity `json:"items,omitempty"`
}
