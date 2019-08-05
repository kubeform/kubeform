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

type KmsExternalKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsExternalKeySpec   `json:"spec,omitempty"`
	Status            KmsExternalKeyStatus `json:"status,omitempty"`
}

type KmsExternalKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled         bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	ExpirationModel string `json:"expirationModel" tf:"expiration_model"`
	// +optional
	KeyMaterialBase64 string `json:"-" sensitive:"true" tf:"key_material_base64,omitempty"`
	KeyState          string `json:"keyState" tf:"key_state"`
	KeyUsage          string `json:"keyUsage" tf:"key_usage"`
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
	State *apis.State `json:"state,omitempty"`
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
