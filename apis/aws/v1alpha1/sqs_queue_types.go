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

type SqsQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqsQueueSpec   `json:"spec,omitempty"`
	Status            SqsQueueStatus `json:"status,omitempty"`
}

type SqsQueueSpec struct {
	// +optional
	ContentBasedDeduplication bool `json:"content_based_deduplication,omitempty"`
	// +optional
	DelaySeconds int `json:"delay_seconds,omitempty"`
	// +optional
	FifoQueue bool `json:"fifo_queue,omitempty"`
	// +optional
	KmsMasterKeyId string `json:"kms_master_key_id,omitempty"`
	// +optional
	MaxMessageSize int `json:"max_message_size,omitempty"`
	// +optional
	MessageRetentionSeconds int `json:"message_retention_seconds,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	ReceiveWaitTimeSeconds int `json:"receive_wait_time_seconds,omitempty"`
	// +optional
	RedrivePolicy string `json:"redrive_policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	VisibilityTimeoutSeconds int `json:"visibility_timeout_seconds,omitempty"`
}

type SqsQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqsQueueList is a list of SqsQueues
type SqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqsQueue CRD objects
	Items []SqsQueue `json:"items,omitempty"`
}
