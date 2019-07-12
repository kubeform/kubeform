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

type AwsPinpointApnsChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsChannelSpec struct {
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	Certificate                 string `json:"certificate"`
	Enabled                     bool   `json:"enabled"`
}

type AwsPinpointApnsChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointApnsChannelList is a list of AwsPinpointApnsChannels
type AwsPinpointApnsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsChannel CRD objects
	Items []AwsPinpointApnsChannel `json:"items,omitempty"`
}
