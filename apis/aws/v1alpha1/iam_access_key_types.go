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

type IamAccessKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccessKeySpec   `json:"spec,omitempty"`
	Status            IamAccessKeyStatus `json:"status,omitempty"`
}

type IamAccessKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EncryptedSecret string `json:"encryptedSecret,omitempty" tf:"encrypted_secret,omitempty"`
	// +optional
	KeyFingerprint string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`
	// +optional
	PgpKey string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
	// +optional
	// Deprecated
	Secret string `json:"secret,omitempty" tf:"secret,omitempty"`
	// +optional
	SesSMTPPassword string `json:"sesSMTPPassword,omitempty" tf:"ses_smtp_password,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	User   string `json:"user" tf:"user"`
}

type IamAccessKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamAccessKeySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamAccessKeyList is a list of IamAccessKeys
type IamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamAccessKey CRD objects
	Items []IamAccessKey `json:"items,omitempty"`
}
