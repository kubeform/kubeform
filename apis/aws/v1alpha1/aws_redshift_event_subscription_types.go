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

type AwsRedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsRedshiftEventSubscriptionSpec struct {
	Status          string            `json:"status"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Tags            map[string]string `json:"tags"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
	SnsTopicArn     string            `json:"sns_topic_arn"`
	EventCategories []string          `json:"event_categories"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Enabled         bool              `json:"enabled"`
	Severity        string            `json:"severity"`
}



type AwsRedshiftEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRedshiftEventSubscriptionList is a list of AwsRedshiftEventSubscriptions
type AwsRedshiftEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftEventSubscription CRD objects
	Items []AwsRedshiftEventSubscription `json:"items,omitempty"`
}