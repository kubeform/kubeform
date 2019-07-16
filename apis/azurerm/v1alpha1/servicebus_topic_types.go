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

type ServicebusTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusTopicSpec   `json:"spec,omitempty"`
	Status            ServicebusTopicStatus `json:"status,omitempty"`
}

type ServicebusTopicSpec struct {
	// +optional
	EnableBatchedOperations bool `json:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress bool `json:"enable_express,omitempty"`
	// +optional
	// Deprecated
	EnableFilteringMessagesBeforePublishing bool `json:"enable_filtering_messages_before_publishing,omitempty"`
	// +optional
	EnablePartitioning bool `json:"enable_partitioning,omitempty"`
	// +optional
	// Deprecated
	Location      string `json:"location,omitempty"`
	Name          string `json:"name"`
	NamespaceName string `json:"namespace_name"`
	// +optional
	RequiresDuplicateDetection bool   `json:"requires_duplicate_detection,omitempty"`
	ResourceGroupName          string `json:"resource_group_name"`
	// +optional
	Status string `json:"status,omitempty"`
	// +optional
	SupportOrdering bool `json:"support_ordering,omitempty"`
}

type ServicebusTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusTopicList is a list of ServicebusTopics
type ServicebusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusTopic CRD objects
	Items []ServicebusTopic `json:"items,omitempty"`
}
