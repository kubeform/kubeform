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

type DbEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            DbEventSubscriptionStatus `json:"status,omitempty"`
}

type DbEventSubscriptionSpec struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EventCategories []string `json:"event_categories,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	SnsTopic   string `json:"sns_topic"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceIds []string `json:"source_ids,omitempty"`
	// +optional
	SourceType string `json:"source_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbEventSubscriptionList is a list of DbEventSubscriptions
type DbEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbEventSubscription CRD objects
	Items []DbEventSubscription `json:"items,omitempty"`
}
