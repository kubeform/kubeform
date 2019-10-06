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

type ComputeGlobalAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeGlobalAddressSpec   `json:"spec,omitempty"`
	Status            ComputeGlobalAddressStatus `json:"status,omitempty"`
}

type ComputeGlobalAddressSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	// +optional
	AddressType string `json:"addressType,omitempty" tf:"address_type,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IpVersion string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`
	// +optional
	// Deprecated
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	// Deprecated
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	// Deprecated
	PrefixLength int `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	Purpose string `json:"purpose,omitempty" tf:"purpose,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeGlobalAddressStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeGlobalAddressSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeGlobalAddressList is a list of ComputeGlobalAddresss
type ComputeGlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeGlobalAddress CRD objects
	Items []ComputeGlobalAddress `json:"items,omitempty"`
}
