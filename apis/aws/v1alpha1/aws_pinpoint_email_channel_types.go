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

type AwsPinpointEmailChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointEmailChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointEmailChannelStatus `json:"status,omitempty"`
}

type AwsPinpointEmailChannelSpec struct {
	ApplicationId     string `json:"application_id"`
	Enabled           bool   `json:"enabled"`
	FromAddress       string `json:"from_address"`
	Identity          string `json:"identity"`
	RoleArn           string `json:"role_arn"`
	MessagesPerSecond int    `json:"messages_per_second"`
}

type AwsPinpointEmailChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointEmailChannelList is a list of AwsPinpointEmailChannels
type AwsPinpointEmailChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointEmailChannel CRD objects
	Items []AwsPinpointEmailChannel `json:"items,omitempty"`
}
