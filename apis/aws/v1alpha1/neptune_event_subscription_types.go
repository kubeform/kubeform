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

type NeptuneEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            NeptuneEventSubscriptionStatus `json:"status,omitempty"`
}

type NeptuneEventSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CustomerAwsID string `json:"customerAwsID,omitempty" tf:"customer_aws_id,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	EventCategories []string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix  string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	SnsTopicArn string `json:"snsTopicArn" tf:"sns_topic_arn"`
	// +optional
	SourceIDS []string `json:"sourceIDS,omitempty" tf:"source_ids,omitempty"`
	// +optional
	SourceType string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NeptuneEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NeptuneEventSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneEventSubscriptionList is a list of NeptuneEventSubscriptions
type NeptuneEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneEventSubscription CRD objects
	Items []NeptuneEventSubscription `json:"items,omitempty"`
}
