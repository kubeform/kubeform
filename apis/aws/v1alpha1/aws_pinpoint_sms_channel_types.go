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

type AwsPinpointSmsChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointSmsChannelSpec   `json:"spec,omitempty"`
	Status            AwsPinpointSmsChannelStatus `json:"status,omitempty"`
}

type AwsPinpointSmsChannelSpec struct {
	ApplicationId                  string `json:"application_id"`
	Enabled                        bool   `json:"enabled"`
	SenderId                       string `json:"sender_id"`
	ShortCode                      string `json:"short_code"`
	PromotionalMessagesPerSecond   int    `json:"promotional_messages_per_second"`
	TransactionalMessagesPerSecond int    `json:"transactional_messages_per_second"`
}



type AwsPinpointSmsChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPinpointSmsChannelList is a list of AwsPinpointSmsChannels
type AwsPinpointSmsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointSmsChannel CRD objects
	Items []AwsPinpointSmsChannel `json:"items,omitempty"`
}