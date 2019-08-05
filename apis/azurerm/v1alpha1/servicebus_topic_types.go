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

type ServicebusTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusTopicSpec   `json:"spec,omitempty"`
	Status            ServicebusTopicStatus `json:"status,omitempty"`
}

type ServicebusTopicSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`
	// +optional
	DuplicateDetectionHistoryTimeWindow string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`
	// +optional
	EnableBatchedOperations bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`
	// +optional
	// Deprecated
	EnableFilteringMessagesBeforePublishing bool `json:"enableFilteringMessagesBeforePublishing,omitempty" tf:"enable_filtering_messages_before_publishing,omitempty"`
	// +optional
	EnablePartitioning bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	MaxSizeInMegabytes int    `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`
	Name               string `json:"name" tf:"name"`
	NamespaceName      string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	RequiresDuplicateDetection bool   `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	SupportOrdering bool `json:"supportOrdering,omitempty" tf:"support_ordering,omitempty"`
}

type ServicebusTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServicebusTopicSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
