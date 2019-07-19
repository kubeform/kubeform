package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KeyVaultKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultKeySpec   `json:"spec,omitempty"`
	Status            KeyVaultKeyStatus `json:"status,omitempty"`
}

type KeyVaultKeySpec struct {
	// +optional
	Curve   string   `json:"curve,omitempty" tf:"curve,omitempty"`
	KeyOpts []string `json:"keyOpts" tf:"key_opts"`
	// +optional
	KeySize int    `json:"keySize,omitempty" tf:"key_size,omitempty"`
	KeyType string `json:"keyType" tf:"key_type"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Deprecated
	VaultURI    string                    `json:"vaultURI,omitempty" tf:"vault_uri,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type KeyVaultKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultKeyList is a list of KeyVaultKeys
type KeyVaultKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultKey CRD objects
	Items []KeyVaultKey `json:"items,omitempty"`
}
