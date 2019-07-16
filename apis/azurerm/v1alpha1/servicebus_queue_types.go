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

type ServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueSpec   `json:"spec,omitempty"`
	Status            ServicebusQueueStatus `json:"status,omitempty"`
}

type ServicebusQueueSpec struct {
	// +optional
	DeadLetteringOnMessageExpiration bool `json:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	// Deprecated
	EnableBatchedOperations bool `json:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress bool `json:"enable_express,omitempty"`
	// +optional
	EnablePartitioning bool `json:"enable_partitioning,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty"`
	// +optional
	MaxDeliveryCount int    `json:"max_delivery_count,omitempty"`
	Name             string `json:"name"`
	NamespaceName    string `json:"namespace_name"`
	// +optional
	RequiresDuplicateDetection bool `json:"requires_duplicate_detection,omitempty"`
	// +optional
	RequiresSession   bool   `json:"requires_session,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	// Deprecated
	SupportOrdering bool `json:"support_ordering,omitempty"`
}

type ServicebusQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusQueueList is a list of ServicebusQueues
type ServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusQueue CRD objects
	Items []ServicebusQueue `json:"items,omitempty"`
}
