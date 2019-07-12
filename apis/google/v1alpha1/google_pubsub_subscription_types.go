package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GooglePubsubSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubSubscriptionSpec   `json:"spec,omitempty"`
	Status            GooglePubsubSubscriptionStatus `json:"status,omitempty"`
}

type GooglePubsubSubscriptionSpecPushConfig struct {
	Attributes   map[string]string `json:"attributes"`
	PushEndpoint string            `json:"push_endpoint"`
}

type GooglePubsubSubscriptionSpec struct {
	AckDeadlineSeconds int                            `json:"ack_deadline_seconds"`
	Project            string                         `json:"project"`
	Path               string                         `json:"path"`
	PushConfig         []GooglePubsubSubscriptionSpec `json:"push_config"`
	Name               string                         `json:"name"`
	Topic              string                         `json:"topic"`
}



type GooglePubsubSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GooglePubsubSubscriptionList is a list of GooglePubsubSubscriptions
type GooglePubsubSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubSubscription CRD objects
	Items []GooglePubsubSubscription `json:"items,omitempty"`
}