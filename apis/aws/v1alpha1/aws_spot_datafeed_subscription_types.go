package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotDatafeedSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

type AwsSpotDatafeedSubscriptionSpec struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

type AwsSpotDatafeedSubscriptionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotDatafeedSubscriptionList is a list of AwsSpotDatafeedSubscriptions
type AwsSpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotDatafeedSubscription CRD objects
	Items []AwsSpotDatafeedSubscription `json:"items,omitempty"`
}
