package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointApnsVoipSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsVoipSandboxChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsVoipSandboxChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsVoipSandboxChannelSpec struct {
	PrivateKey                  string `json:"private_key"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
	TeamId                      string `json:"team_id"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
}

type AwsPinpointApnsVoipSandboxChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointApnsVoipSandboxChannelList is a list of AwsPinpointApnsVoipSandboxChannels
type AwsPinpointApnsVoipSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsVoipSandboxChannel CRD objects
	Items []AwsPinpointApnsVoipSandboxChannel `json:"items,omitempty"`
}
