package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	PgpKey                string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
	PrivateKey            string `json:"-" sensitive:"true" tf:"private_key"`
	PrivateKeyEncrypted   string `json:"privateKeyEncrypted" tf:"private_key_encrypted"`
	PrivateKeyFingerprint string `json:"privateKeyFingerprint" tf:"private_key_fingerprint"`
	// +optional
	PrivateKeyType string `json:"privateKeyType,omitempty" tf:"private_key_type,omitempty"`
	PublicKey      string `json:"publicKey" tf:"public_key"`
	// +optional
	PublicKeyType    string `json:"publicKeyType,omitempty" tf:"public_key_type,omitempty"`
	ServiceAccountID string `json:"serviceAccountID" tf:"service_account_id"`
	ValidAfter       string `json:"validAfter" tf:"valid_after"`
	ValidBefore      string `json:"validBefore" tf:"valid_before"`
}

type ServiceAccountKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServiceAccountKeySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
