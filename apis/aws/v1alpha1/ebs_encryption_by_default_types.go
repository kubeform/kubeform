package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type EbsEncryptionByDefault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsEncryptionByDefaultSpec   `json:"spec,omitempty"`
	Status            EbsEncryptionByDefaultStatus `json:"status,omitempty"`
}

type EbsEncryptionByDefaultSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type EbsEncryptionByDefaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EbsEncryptionByDefaultList is a list of EbsEncryptionByDefaults
type EbsEncryptionByDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EbsEncryptionByDefault CRD objects
	Items []EbsEncryptionByDefault `json:"items,omitempty"`
}
