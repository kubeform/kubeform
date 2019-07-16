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

type PinpointAdmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointAdmChannelSpec   `json:"spec,omitempty"`
	Status            PinpointAdmChannelStatus `json:"status,omitempty"`
}

type PinpointAdmChannelSpec struct {
	ApplicationId string `json:"application_id"`
	ClientId      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type PinpointAdmChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
