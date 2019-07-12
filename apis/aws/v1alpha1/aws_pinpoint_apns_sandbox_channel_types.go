package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsSandboxChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsSandboxChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsSandboxChannelSpec struct {
	Certificate                 string `json:"certificate"`
	TokenKey                    string `json:"token_key"`
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKeyId                  string `json:"token_key_id"`
}



type AwsPinpointApnsSandboxChannelStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointApnsSandboxChannelList is a list of AwsPinpointApnsSandboxChannels
type AwsPinpointApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsSandboxChannel CRD objects
	Items []AwsPinpointApnsSandboxChannel `json:"items,omitempty"`
}