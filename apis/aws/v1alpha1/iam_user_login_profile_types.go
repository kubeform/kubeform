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

type IamUserLoginProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserLoginProfileSpec   `json:"spec,omitempty"`
	Status            IamUserLoginProfileStatus `json:"status,omitempty"`
}

type IamUserLoginProfileSpec struct {
	// +optional
	PasswordLength int `json:"passwordLength,omitempty" tf:"password_length,omitempty"`
	// +optional
	PasswordResetRequired bool                      `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`
	PgpKey                string                    `json:"pgpKey" tf:"pgp_key"`
	User                  string                    `json:"user" tf:"user"`
	ProviderRef           core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IamUserLoginProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
