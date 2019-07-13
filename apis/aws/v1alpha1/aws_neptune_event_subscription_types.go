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

type AwsNeptuneEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsNeptuneEventSubscriptionSpec struct {
	NamePrefix      string            `json:"name_prefix"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
	Enabled         bool              `json:"enabled"`
	Tags            map[string]string `json:"tags"`
	SnsTopicArn     string            `json:"sns_topic_arn"`
	EventCategories []string          `json:"event_categories"`
}



type AwsNeptuneEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNeptuneEventSubscriptionList is a list of AwsNeptuneEventSubscriptions
type AwsNeptuneEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneEventSubscription CRD objects
	Items []AwsNeptuneEventSubscription `json:"items,omitempty"`
}