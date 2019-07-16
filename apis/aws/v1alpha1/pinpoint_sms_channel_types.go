package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PinpointSmsChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointSmsChannelSpec   `json:"spec,omitempty"`
	Status            PinpointSmsChannelStatus `json:"status,omitempty"`
}

type PinpointSmsChannelSpec struct {
	ApplicationId string `json:"application_id"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	SenderId string `json:"sender_id,omitempty"`
	// +optional
	ShortCode string `json:"short_code,omitempty"`
}

type PinpointSmsChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointSmsChannelList is a list of PinpointSmsChannels
type PinpointSmsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointSmsChannel CRD objects
	Items []PinpointSmsChannel `json:"items,omitempty"`
}
