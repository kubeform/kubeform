package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsDbEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsDbEventSubscriptionSpec struct {
	Name            string            `json:"name"`
	SnsTopic        string            `json:"sns_topic"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Enabled         bool              `json:"enabled"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Arn             string            `json:"arn"`
	NamePrefix      string            `json:"name_prefix"`
	EventCategories []string          `json:"event_categories"`
	Tags            map[string]string `json:"tags"`
}



type AwsDbEventSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbEventSubscriptionList is a list of AwsDbEventSubscriptions
type AwsDbEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbEventSubscription CRD objects
	Items []AwsDbEventSubscription `json:"items,omitempty"`
}