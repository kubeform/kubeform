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

type AwsDbEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsDbEventSubscriptionStatus `json:"status,omitempty"`
}

type AwsDbEventSubscriptionSpec struct {
	Tags            map[string]string `json:"tags"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	SnsTopic        string            `json:"sns_topic"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Enabled         bool              `json:"enabled"`
	Arn             string            `json:"arn"`
	EventCategories []string          `json:"event_categories"`
	CustomerAwsId   string            `json:"customer_aws_id"`
}



type AwsDbEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbEventSubscriptionList is a list of AwsDbEventSubscriptions
type AwsDbEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbEventSubscription CRD objects
	Items []AwsDbEventSubscription `json:"items,omitempty"`
}