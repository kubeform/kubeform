package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KmsCiphertext struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsCiphertextSpec   `json:"spec,omitempty"`
	Status            KmsCiphertextStatus `json:"status,omitempty"`
}

type KmsCiphertextSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Context   map[string]string `json:"context,omitempty" tf:"context,omitempty"`
	KeyID     string            `json:"keyID" tf:"key_id"`
	Plaintext string            `json:"-" sensitive:"true" tf:"plaintext"`
}

type KmsCiphertextStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsCiphertextList is a list of KmsCiphertexts
type KmsCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsCiphertext CRD objects
	Items []KmsCiphertext `json:"items,omitempty"`
}
