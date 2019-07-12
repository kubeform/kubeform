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

type AwsSnsTopicSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsTopicSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSnsTopicSubscriptionStatus `json:"status,omitempty"`
}

type AwsSnsTopicSubscriptionSpec struct {
	Endpoint                     string `json:"endpoint"`
	EndpointAutoConfirms         bool   `json:"endpoint_auto_confirms"`
	ConfirmationTimeoutInMinutes int    `json:"confirmation_timeout_in_minutes"`
	Arn                          string `json:"arn"`
	Protocol                     string `json:"protocol"`
	TopicArn                     string `json:"topic_arn"`
	DeliveryPolicy               string `json:"delivery_policy"`
	RawMessageDelivery           bool   `json:"raw_message_delivery"`
	FilterPolicy                 string `json:"filter_policy"`
}

type AwsSnsTopicSubscriptionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSnsTopicSubscriptionList is a list of AwsSnsTopicSubscriptions
type AwsSnsTopicSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsTopicSubscription CRD objects
	Items []AwsSnsTopicSubscription `json:"items,omitempty"`
}
