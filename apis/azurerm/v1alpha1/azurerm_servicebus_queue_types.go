package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusQueueSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusQueueStatus `json:"status,omitempty"`
}

type AzurermServicebusQueueSpec struct {
	AutoDeleteOnIdle                    string `json:"auto_delete_on_idle"`
	RequiresDuplicateDetection          bool   `json:"requires_duplicate_detection"`
	SupportOrdering                     bool   `json:"support_ordering"`
	MaxDeliveryCount                    int    `json:"max_delivery_count"`
	Name                                string `json:"name"`
	Location                            string `json:"location"`
	ResourceGroupName                   string `json:"resource_group_name"`
	DefaultMessageTtl                   string `json:"default_message_ttl"`
	LockDuration                        string `json:"lock_duration"`
	RequiresSession                     bool   `json:"requires_session"`
	DeadLetteringOnMessageExpiration    bool   `json:"dead_lettering_on_message_expiration"`
	EnableExpress                       bool   `json:"enable_express"`
	MaxSizeInMegabytes                  int    `json:"max_size_in_megabytes"`
	EnableBatchedOperations             bool   `json:"enable_batched_operations"`
	NamespaceName                       string `json:"namespace_name"`
	DuplicateDetectionHistoryTimeWindow string `json:"duplicate_detection_history_time_window"`
	EnablePartitioning                  bool   `json:"enable_partitioning"`
}

type AzurermServicebusQueueStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermServicebusQueueList is a list of AzurermServicebusQueues
type AzurermServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusQueue CRD objects
	Items []AzurermServicebusQueue `json:"items,omitempty"`
}
