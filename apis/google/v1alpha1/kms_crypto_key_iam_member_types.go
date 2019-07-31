package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KmsCryptoKeyIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsCryptoKeyIamMemberSpec   `json:"spec,omitempty"`
	Status            KmsCryptoKeyIamMemberStatus `json:"status,omitempty"`
}

type KmsCryptoKeyIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	CryptoKeyID string `json:"cryptoKeyID" tf:"crypto_key_id"`
	Member      string `json:"member" tf:"member"`
	Role        string `json:"role" tf:"role"`
}

type KmsCryptoKeyIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsCryptoKeyIamMemberList is a list of KmsCryptoKeyIamMembers
type KmsCryptoKeyIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsCryptoKeyIamMember CRD objects
	Items []KmsCryptoKeyIamMember `json:"items,omitempty"`
}
