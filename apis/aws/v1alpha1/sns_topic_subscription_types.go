package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type SnsTopicSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicSubscriptionSpec   `json:"spec,omitempty"`
	Status            SnsTopicSubscriptionStatus `json:"status,omitempty"`
}

type SnsTopicSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	ConfirmationTimeoutInMinutes int `json:"confirmationTimeoutInMinutes,omitempty" tf:"confirmation_timeout_in_minutes,omitempty"`
	// +optional
	DeliveryPolicy string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`
	Endpoint       string `json:"endpoint" tf:"endpoint"`
	// +optional
	EndpointAutoConfirms bool `json:"endpointAutoConfirms,omitempty" tf:"endpoint_auto_confirms,omitempty"`
	// +optional
	FilterPolicy string `json:"filterPolicy,omitempty" tf:"filter_policy,omitempty"`
	Protocol     string `json:"protocol" tf:"protocol"`
	// +optional
	RawMessageDelivery bool   `json:"rawMessageDelivery,omitempty" tf:"raw_message_delivery,omitempty"`
	TopicArn           string `json:"topicArn" tf:"topic_arn"`
}

type SnsTopicSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnsTopicSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
