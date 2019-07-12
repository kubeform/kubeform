package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNeptuneEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsNeptuneEventSubscriptionSpec struct {
	SnsTopicArn     string            `json:"sns_topic_arn"`
	SourceIds       []string          `json:"source_ids"`
	Tags            map[string]string `json:"tags"`
	EventCategories []string          `json:"event_categories"`
	SourceType      string            `json:"source_type"`
	Enabled         bool              `json:"enabled"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
}



type AwsNeptuneEventSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNeptuneEventSubscriptionList is a list of AwsNeptuneEventSubscriptions
type AwsNeptuneEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneEventSubscription CRD objects
	Items []AwsNeptuneEventSubscription `json:"items,omitempty"`
}