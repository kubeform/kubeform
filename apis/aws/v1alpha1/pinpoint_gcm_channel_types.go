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

type PinpointGcmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointGcmChannelSpec   `json:"spec,omitempty"`
	Status            PinpointGcmChannelStatus `json:"status,omitempty"`
}

type PinpointGcmChannelSpec struct {
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type PinpointGcmChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
