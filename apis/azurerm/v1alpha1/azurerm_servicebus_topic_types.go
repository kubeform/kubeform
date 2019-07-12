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

type AzurermServicebusTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusTopicSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusTopicStatus `json:"status,omitempty"`
}

type AzurermServicebusTopicSpec struct {
	DuplicateDetectionHistoryTimeWindow     string `json:"duplicate_detection_history_time_window"`
	EnableBatchedOperations                 bool   `json:"enable_batched_operations"`
	SupportOrdering                         bool   `json:"support_ordering"`
	Status                                  string `json:"status"`
	EnableExpress                           bool   `json:"enable_express"`
	EnableFilteringMessagesBeforePublishing bool   `json:"enable_filtering_messages_before_publishing"`
	Name                                    string `json:"name"`
	NamespaceName                           string `json:"namespace_name"`
	Location                                string `json:"location"`
	MaxSizeInMegabytes                      int    `json:"max_size_in_megabytes"`
	RequiresDuplicateDetection              bool   `json:"requires_duplicate_detection"`
	AutoDeleteOnIdle                        string `json:"auto_delete_on_idle"`
	DefaultMessageTtl                       string `json:"default_message_ttl"`
	EnablePartitioning                      bool   `json:"enable_partitioning"`
	ResourceGroupName                       string `json:"resource_group_name"`
}

type AzurermServicebusTopicStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusTopicList is a list of AzurermServicebusTopics
type AzurermServicebusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusTopic CRD objects
	Items []AzurermServicebusTopic `json:"items,omitempty"`
}
