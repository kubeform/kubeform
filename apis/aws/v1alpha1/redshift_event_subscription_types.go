package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            RedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

type RedshiftEventSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomerAwsID string `json:"customerAwsID,omitempty" tf:"customer_aws_id,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EventCategories []string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`
	Name            string   `json:"name" tf:"name"`
	// +optional
	Severity    string `json:"severity,omitempty" tf:"severity,omitempty"`
	SnsTopicArn string `json:"snsTopicArn" tf:"sns_topic_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceIDS []string `json:"sourceIDS,omitempty" tf:"source_ids,omitempty"`
	// +optional
	SourceType string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedshiftEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedshiftEventSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftEventSubscriptionList is a list of RedshiftEventSubscriptions
type RedshiftEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftEventSubscription CRD objects
	Items []RedshiftEventSubscription `json:"items,omitempty"`
}
