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

type AwsPinpointApnsVoipChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointApnsVoipChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointApnsVoipChannelStatus `json:"status,omitempty"`
}

type AwsPinpointApnsVoipChannelSpec struct {
	Enabled                     bool   `json:"enabled"`
	PrivateKey                  string `json:"private_key"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
	TeamId                      string `json:"team_id"`
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
}



type AwsPinpointApnsVoipChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointApnsVoipChannelList is a list of AwsPinpointApnsVoipChannels
type AwsPinpointApnsVoipChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointApnsVoipChannel CRD objects
	Items []AwsPinpointApnsVoipChannel `json:"items,omitempty"`
}