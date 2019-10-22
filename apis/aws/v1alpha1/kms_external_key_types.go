package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type KmsExternalKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsExternalKeySpec   `json:"spec,omitempty"`
	Status            KmsExternalKeyStatus `json:"status,omitempty"`
}

type KmsExternalKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	ExpirationModel string `json:"expirationModel,omitempty" tf:"expiration_model,omitempty"`
	// +optional
	KeyMaterialBase64 string `json:"-" sensitive:"true" tf:"key_material_base64,omitempty"`
	// +optional
	KeyState string `json:"keyState,omitempty" tf:"key_state,omitempty"`
	// +optional
	KeyUsage string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ValidTo string `json:"validTo,omitempty" tf:"valid_to,omitempty"`
}

type KmsExternalKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KmsExternalKeySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsExternalKeyList is a list of KmsExternalKeys
type KmsExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsExternalKey CRD objects
	Items []KmsExternalKey `json:"items,omitempty"`
}
