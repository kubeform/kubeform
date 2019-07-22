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

type PinpointGcmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointGcmChannelSpec   `json:"spec,omitempty"`
	Status            PinpointGcmChannelStatus `json:"status,omitempty"`
}

type PinpointGcmChannelSpec struct {
	Secret        *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef   core.LocalObjectReference  `json:"providerRef" tf:"-"`
	ApplicationID string                     `json:"applicationID" tf:"application_id"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type PinpointGcmChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointGcmChannelList is a list of PinpointGcmChannels
type PinpointGcmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointGcmChannel CRD objects
	Items []PinpointGcmChannel `json:"items,omitempty"`
}
