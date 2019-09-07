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

type PinpointApnsChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsChannelSpec   `json:"spec,omitempty"`
	Status            PinpointApnsChannelStatus `json:"status,omitempty"`
}

type PinpointApnsChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApplicationID string `json:"applicationID" tf:"application_id"`
	// +optional
	BundleID string `json:"-" sensitive:"true" tf:"bundle_id,omitempty"`
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	DefaultAuthenticationMethod string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	TeamID string `json:"-" sensitive:"true" tf:"team_id,omitempty"`
	// +optional
	TokenKey string `json:"-" sensitive:"true" tf:"token_key,omitempty"`
	// +optional
	TokenKeyID string `json:"-" sensitive:"true" tf:"token_key_id,omitempty"`
}

type PinpointApnsChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PinpointApnsChannelSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointApnsChannelList is a list of PinpointApnsChannels
type PinpointApnsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointApnsChannel CRD objects
	Items []PinpointApnsChannel `json:"items,omitempty"`
}