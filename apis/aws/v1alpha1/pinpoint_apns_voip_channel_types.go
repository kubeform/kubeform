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

type PinpointApnsVoipChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsVoipChannelSpec   `json:"spec,omitempty"`
	Status            PinpointApnsVoipChannelStatus `json:"status,omitempty"`
}

type PinpointApnsVoipChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	ApplicationID string `json:"applicationID" tf:"application_id"`
	// +optional
	// +optional
	// +optional
	DefaultAuthenticationMethod string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +optional
	// +optional
	// +optional
}

type PinpointApnsVoipChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointApnsVoipChannelList is a list of PinpointApnsVoipChannels
type PinpointApnsVoipChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointApnsVoipChannel CRD objects
	Items []PinpointApnsVoipChannel `json:"items,omitempty"`
}
