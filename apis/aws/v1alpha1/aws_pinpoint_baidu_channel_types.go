package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointBaiduChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointBaiduChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointBaiduChannelStatus `json:"status,omitempty"`
}

type AwsPinpointBaiduChannelSpec struct {
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	ApiKey        string `json:"api_key"`
	SecretKey     string `json:"secret_key"`
}

type AwsPinpointBaiduChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointBaiduChannelList is a list of AwsPinpointBaiduChannels
type AwsPinpointBaiduChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointBaiduChannel CRD objects
	Items []AwsPinpointBaiduChannel `json:"items,omitempty"`
}
