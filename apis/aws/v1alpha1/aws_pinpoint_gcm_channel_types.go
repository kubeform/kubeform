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

type AwsPinpointGcmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointGcmChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointGcmChannelStatus `json:"status,omitempty"`
}

type AwsPinpointGcmChannelSpec struct {
	ApplicationId string `json:"application_id"`
	ApiKey        string `json:"api_key"`
	Enabled       bool   `json:"enabled"`
}

type AwsPinpointGcmChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointGcmChannelList is a list of AwsPinpointGcmChannels
type AwsPinpointGcmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointGcmChannel CRD objects
	Items []AwsPinpointGcmChannel `json:"items,omitempty"`
}
