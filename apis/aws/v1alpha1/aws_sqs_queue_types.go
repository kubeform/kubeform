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

type AwsSqsQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSqsQueueSpec   `json:"spec,omitempty"`
	Status            AwsSqsQueueStatus `json:"status,omitempty"`
}

type AwsSqsQueueSpec struct {
	ReceiveWaitTimeSeconds       int               `json:"receive_wait_time_seconds"`
	Arn                          string            `json:"arn"`
	KmsMasterKeyId               string            `json:"kms_master_key_id"`
	MessageRetentionSeconds      int               `json:"message_retention_seconds"`
	ContentBasedDeduplication    bool              `json:"content_based_deduplication"`
	VisibilityTimeoutSeconds     int               `json:"visibility_timeout_seconds"`
	RedrivePolicy                string            `json:"redrive_policy"`
	Tags                         map[string]string `json:"tags"`
	NamePrefix                   string            `json:"name_prefix"`
	DelaySeconds                 int               `json:"delay_seconds"`
	Policy                       string            `json:"policy"`
	FifoQueue                    bool              `json:"fifo_queue"`
	KmsDataKeyReusePeriodSeconds int               `json:"kms_data_key_reuse_period_seconds"`
	Name                         string            `json:"name"`
	MaxMessageSize               int               `json:"max_message_size"`
}



type AwsSqsQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSqsQueueList is a list of AwsSqsQueues
type AwsSqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSqsQueue CRD objects
	Items []AwsSqsQueue `json:"items,omitempty"`
}