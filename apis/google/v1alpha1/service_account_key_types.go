package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountKeySpec   `json:"spec,omitempty"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty"`
}

type ServiceAccountKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`
	// +optional
	PgpKey string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
	// +optional
	PrivateKeyType string `json:"privateKeyType,omitempty" tf:"private_key_type,omitempty"`
	// +optional
	PublicKeyType    string `json:"publicKeyType,omitempty" tf:"public_key_type,omitempty"`
	ServiceAccountID string `json:"serviceAccountID" tf:"service_account_id"`
}

type ServiceAccountKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountKeyList is a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccountKey CRD objects
	Items []ServiceAccountKey `json:"items,omitempty"`
}
