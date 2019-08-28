package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamUserLoginProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserLoginProfileSpec   `json:"spec,omitempty"`
	Status            IamUserLoginProfileStatus `json:"status,omitempty"`
}

type IamUserLoginProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EncryptedPassword string `json:"encryptedPassword,omitempty" tf:"encrypted_password,omitempty"`
	// +optional
	KeyFingerprint string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`
	// +optional
	PasswordLength int `json:"passwordLength,omitempty" tf:"password_length,omitempty"`
	// +optional
	PasswordResetRequired bool   `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`
	PgpKey                string `json:"pgpKey" tf:"pgp_key"`
	User                  string `json:"user" tf:"user"`
}



type IamUserLoginProfileStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *IamUserLoginProfileSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamUserLoginProfileList is a list of IamUserLoginProfiles
type IamUserLoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamUserLoginProfile CRD objects
	Items []IamUserLoginProfile `json:"items,omitempty"`
}