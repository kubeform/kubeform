package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type KmsKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeySpec   `json:"spec,omitempty"`
	Status            KmsKeyStatus `json:"status,omitempty"`
}

type KmsKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnableKeyRotation bool `json:"enableKeyRotation,omitempty" tf:"enable_key_rotation,omitempty"`
	// +optional
	IsEnabled bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
	// +optional
	KeyID string `json:"keyID,omitempty" tf:"key_id,omitempty"`
	// +optional
	KeyUsage string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KmsKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KmsKeySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsKeyList is a list of KmsKeys
type KmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsKey CRD objects
	Items []KmsKey `json:"items,omitempty"`
}
