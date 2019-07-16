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

type RedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            RedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

type RedshiftEventSubscriptionSpec struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EventCategories []string `json:"event_categories,omitempty"`
	Name            string   `json:"name"`
	// +optional
	Severity    string `json:"severity,omitempty"`
	SnsTopicArn string `json:"sns_topic_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceIds []string `json:"source_ids,omitempty"`
	// +optional
	SourceType string `json:"source_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RedshiftEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
