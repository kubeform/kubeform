package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointAdmChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointAdmChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointAdmChannelStatus `json:"status,omitempty"`
}

type AwsPinpointAdmChannelSpec struct {
	ClientSecret  string `json:"client_secret"`
	Enabled       bool   `json:"enabled"`
	ApplicationId string `json:"application_id"`
	ClientId      string `json:"client_id"`
}

type AwsPinpointAdmChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointAdmChannelList is a list of AwsPinpointAdmChannels
type AwsPinpointAdmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointAdmChannel CRD objects
	Items []AwsPinpointAdmChannel `json:"items,omitempty"`
}
