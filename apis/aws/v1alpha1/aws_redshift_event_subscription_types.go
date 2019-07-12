package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsRedshiftEventSubscriptionSpec struct {
	Enabled         bool              `json:"enabled"`
	Tags            map[string]string `json:"tags"`
	SnsTopicArn     string            `json:"sns_topic_arn"`
	Status          string            `json:"status"`
	EventCategories []string          `json:"event_categories"`
	Severity        string            `json:"severity"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Name            string            `json:"name"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
}



type AwsRedshiftEventSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftEventSubscriptionList is a list of AwsRedshiftEventSubscriptions
type AwsRedshiftEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftEventSubscription CRD objects
	Items []AwsRedshiftEventSubscription `json:"items,omitempty"`
}