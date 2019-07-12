package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSqsQueueSpec   `json:"spec,omitempty"`
	Status            AwsSqsQueueStatus `json:"status,omitempty"`
}

type AwsSqsQueueSpec struct {
	RedrivePolicy                string            `json:"redrive_policy"`
	Arn                          string            `json:"arn"`
	Tags                         map[string]string `json:"tags"`
	NamePrefix                   string            `json:"name_prefix"`
	MaxMessageSize               int               `json:"max_message_size"`
	KmsDataKeyReusePeriodSeconds int               `json:"kms_data_key_reuse_period_seconds"`
	Name                         string            `json:"name"`
	MessageRetentionSeconds      int               `json:"message_retention_seconds"`
	ReceiveWaitTimeSeconds       int               `json:"receive_wait_time_seconds"`
	FifoQueue                    bool              `json:"fifo_queue"`
	KmsMasterKeyId               string            `json:"kms_master_key_id"`
	DelaySeconds                 int               `json:"delay_seconds"`
	Policy                       string            `json:"policy"`
	ContentBasedDeduplication    bool              `json:"content_based_deduplication"`
	VisibilityTimeoutSeconds     int               `json:"visibility_timeout_seconds"`
}



type AwsSqsQueueStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSqsQueueList is a list of AwsSqsQueues
type AwsSqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSqsQueue CRD objects
	Items []AwsSqsQueue `json:"items,omitempty"`
}