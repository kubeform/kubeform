package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointApnsVoipChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsVoipChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsVoipChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsVoipChannelSpec struct {
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	PrivateKey                  string `json:"private_key"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
}



type AwsPinpointApnsVoipChannelStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointApnsVoipChannelList is a list of AwsPinpointApnsVoipChannels
type AwsPinpointApnsVoipChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsVoipChannel CRD objects
	Items []AwsPinpointApnsVoipChannel `json:"items,omitempty"`
}