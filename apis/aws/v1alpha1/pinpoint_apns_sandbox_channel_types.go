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

type PinpointApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsSandboxChannelSpec   `json:"spec,omitempty"`
	Status            PinpointApnsSandboxChannelStatus `json:"status,omitempty"`
}

type PinpointApnsSandboxChannelSpec struct {
	ApplicationID string `json:"applicationID" tf:"application_id"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	BundleID core.LocalObjectReference `json:"bundleID,omitempty" tf:"bundle_id,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Certificate core.LocalObjectReference `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	DefaultAuthenticationMethod string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	PrivateKey core.LocalObjectReference `json:"privateKey,omitempty" tf:"private_key,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	TeamID core.LocalObjectReference `json:"teamID,omitempty" tf:"team_id,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	TokenKey core.LocalObjectReference `json:"tokenKey,omitempty" tf:"token_key,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	TokenKeyID  core.LocalObjectReference `json:"tokenKeyID,omitempty" tf:"token_key_id,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PinpointApnsSandboxChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointApnsSandboxChannelList is a list of PinpointApnsSandboxChannels
type PinpointApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointApnsSandboxChannel CRD objects
	Items []PinpointApnsSandboxChannel `json:"items,omitempty"`
}
