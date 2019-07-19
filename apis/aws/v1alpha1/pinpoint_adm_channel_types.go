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

type PinpointAdmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointAdmChannelSpec   `json:"spec,omitempty"`
	Status            PinpointAdmChannelStatus `json:"status,omitempty"`
}

type PinpointAdmChannelSpec struct {
	ApplicationID string `json:"applicationID" tf:"application_id"`
	// Sensitive Data. Provide secret name which contains one value only
	ClientID *core.LocalObjectReference `json:"clientID" tf:"client_id"`
	// Sensitive Data. Provide secret name which contains one value only
	ClientSecret *core.LocalObjectReference `json:"clientSecret" tf:"client_secret"`
	// +optional
	Enabled     bool                      `json:"enabled,omitempty" tf:"enabled,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PinpointAdmChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointAdmChannelList is a list of PinpointAdmChannels
type PinpointAdmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointAdmChannel CRD objects
	Items []PinpointAdmChannel `json:"items,omitempty"`
}
