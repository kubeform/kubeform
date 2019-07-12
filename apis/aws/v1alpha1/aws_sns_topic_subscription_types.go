package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsTopicSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSnsTopicSubscriptionStatus `json:"status,omitempty"`
}

type AwsSnsTopicSubscriptionSpec struct {
	Endpoint                     string `json:"endpoint"`
	TopicArn                     string `json:"topic_arn"`
	FilterPolicy                 string `json:"filter_policy"`
	Arn                          string `json:"arn"`
	Protocol                     string `json:"protocol"`
	EndpointAutoConfirms         bool   `json:"endpoint_auto_confirms"`
	ConfirmationTimeoutInMinutes int    `json:"confirmation_timeout_in_minutes"`
	DeliveryPolicy               string `json:"delivery_policy"`
	RawMessageDelivery           bool   `json:"raw_message_delivery"`
}



type AwsSnsTopicSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicSubscriptionList is a list of AwsSnsTopicSubscriptions
type AwsSnsTopicSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsTopicSubscription CRD objects
	Items []AwsSnsTopicSubscription `json:"items,omitempty"`
}