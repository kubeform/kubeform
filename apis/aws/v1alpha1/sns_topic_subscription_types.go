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

type SnsTopicSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicSubscriptionSpec   `json:"spec,omitempty"`
	Status            SnsTopicSubscriptionStatus `json:"status,omitempty"`
}

type SnsTopicSubscriptionSpec struct {
	// +optional
	ConfirmationTimeoutInMinutes int `json:"confirmation_timeout_in_minutes,omitempty"`
	// +optional
	DeliveryPolicy string `json:"delivery_policy,omitempty"`
	Endpoint       string `json:"endpoint"`
	// +optional
	EndpointAutoConfirms bool `json:"endpoint_auto_confirms,omitempty"`
	// +optional
	FilterPolicy string `json:"filter_policy,omitempty"`
	Protocol     string `json:"protocol"`
	// +optional
	RawMessageDelivery bool   `json:"raw_message_delivery,omitempty"`
	TopicArn           string `json:"topic_arn"`
}

type SnsTopicSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsTopicSubscriptionList is a list of SnsTopicSubscriptions
type SnsTopicSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsTopicSubscription CRD objects
	Items []SnsTopicSubscription `json:"items,omitempty"`
}
