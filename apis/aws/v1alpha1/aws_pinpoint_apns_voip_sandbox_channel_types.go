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

type AwsPinpointApnsVoipSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsVoipSandboxChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsVoipSandboxChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsVoipSandboxChannelSpec struct {
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	Enabled                     bool   `json:"enabled"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
}



type AwsPinpointApnsVoipSandboxChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointApnsVoipSandboxChannelList is a list of AwsPinpointApnsVoipSandboxChannels
type AwsPinpointApnsVoipSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsVoipSandboxChannel CRD objects
	Items []AwsPinpointApnsVoipSandboxChannel `json:"items,omitempty"`
}