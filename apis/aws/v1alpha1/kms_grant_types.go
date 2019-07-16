package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	EncryptionContextEquals map[string]string `json:"encryption_context_equals,omitempty"`
	// +optional
	EncryptionContextSubset map[string]string `json:"encryption_context_subset,omitempty"`
}

type KmsGrantSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Constraints *[]KmsGrantSpec `json:"constraints,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GrantCreationTokens []string `json:"grant_creation_tokens,omitempty"`
	GranteePrincipal    string   `json:"grantee_principal"`
	KeyId               string   `json:"key_id"`
	// +optional
	Name string `json:"name,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Operations []string `json:"operations"`
	// +optional
	RetireOnDelete bool `json:"retire_on_delete,omitempty"`
	// +optional
	RetiringPrincipal string `json:"retiring_principal,omitempty"`
}

type KmsGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
