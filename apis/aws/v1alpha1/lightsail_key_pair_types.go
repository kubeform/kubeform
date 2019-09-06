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

type LightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailKeyPairSpec   `json:"spec,omitempty"`
	Status            LightsailKeyPairStatus `json:"status,omitempty"`
}

type LightsailKeyPairSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	EncryptedFingerprint string `json:"encryptedFingerprint,omitempty" tf:"encrypted_fingerprint,omitempty"`
	// +optional
	EncryptedPrivateKey string `json:"encryptedPrivateKey,omitempty" tf:"encrypted_private_key,omitempty"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	PgpKey string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
	// +optional
	PrivateKey string `json:"privateKey,omitempty" tf:"private_key,omitempty"`
	// +optional
	PublicKey string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type LightsailKeyPairStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LightsailKeyPairSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailKeyPairList is a list of LightsailKeyPairs
type LightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailKeyPair CRD objects
	Items []LightsailKeyPair `json:"items,omitempty"`
}
