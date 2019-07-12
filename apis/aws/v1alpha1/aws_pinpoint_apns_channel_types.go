package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointApnsChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsChannelSpec struct {
	Enabled                     bool   `json:"enabled"`
	TokenKey                    string `json:"token_key"`
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKeyId                  string `json:"token_key_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
}

type AwsPinpointApnsChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointApnsChannelList is a list of AwsPinpointApnsChannels
type AwsPinpointApnsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsChannel CRD objects
	Items []AwsPinpointApnsChannel `json:"items,omitempty"`
}
