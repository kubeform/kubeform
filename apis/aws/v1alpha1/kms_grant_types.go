package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KmsGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsGrantSpec   `json:"spec,omitempty"`
	Status            KmsGrantStatus `json:"status,omitempty"`
}

type KmsGrantSpecConstraints struct {
	// +optional
	EncryptionContextEquals map[string]string `json:"encryptionContextEquals,omitempty" tf:"encryption_context_equals,omitempty"`
	// +optional
	EncryptionContextSubset map[string]string `json:"encryptionContextSubset,omitempty" tf:"encryption_context_subset,omitempty"`
}

type KmsGrantSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Constraints []KmsGrantSpecConstraints `json:"constraints,omitempty" tf:"constraints,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GrantCreationTokens []string `json:"grantCreationTokens,omitempty" tf:"grant_creation_tokens,omitempty"`
	GranteePrincipal    string   `json:"granteePrincipal" tf:"grantee_principal"`
	KeyID               string   `json:"keyID" tf:"key_id"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Operations []string `json:"operations" tf:"operations"`
	// +optional
	RetireOnDelete bool `json:"retireOnDelete,omitempty" tf:"retire_on_delete,omitempty"`
	// +optional
	RetiringPrincipal string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

type KmsGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *KmsGrantSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsGrantList is a list of KmsGrants
type KmsGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsGrant CRD objects
	Items []KmsGrant `json:"items,omitempty"`
}
