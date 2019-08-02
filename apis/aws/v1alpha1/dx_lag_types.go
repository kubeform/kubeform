package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DxLag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxLagSpec   `json:"spec,omitempty"`
	Status            DxLagStatus `json:"status,omitempty"`
}

type DxLagSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ConnectionsBandwidth string `json:"connectionsBandwidth" tf:"connections_bandwidth"`
	// +optional
	ForceDestroy bool   `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	Location     string `json:"location" tf:"location"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DxLagStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *DxLagSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxLagList is a list of DxLags
type DxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxLag CRD objects
	Items []DxLag `json:"items,omitempty"`
}
