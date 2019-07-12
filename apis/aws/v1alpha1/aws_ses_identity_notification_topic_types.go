package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesIdentityNotificationTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesIdentityNotificationTopicSpec   `json:"spec,omitempty"`
	Status            AwsSesIdentityNotificationTopicStatus `json:"status,omitempty"`
}

type AwsSesIdentityNotificationTopicSpec struct {
	Identity         string `json:"identity"`
	TopicArn         string `json:"topic_arn"`
	NotificationType string `json:"notification_type"`
}

type AwsSesIdentityNotificationTopicStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesIdentityNotificationTopicList is a list of AwsSesIdentityNotificationTopics
type AwsSesIdentityNotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesIdentityNotificationTopic CRD objects
	Items []AwsSesIdentityNotificationTopic `json:"items,omitempty"`
}
