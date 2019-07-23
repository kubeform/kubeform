package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ContentBasedDeduplication bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`
	// +optional
	DelaySeconds int `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`
	// +optional
	FifoQueue bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`
	// +optional
	KmsDataKeyReusePeriodSeconds int `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds,omitempty"`
	// +optional
	KmsMasterKeyID string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id,omitempty"`
	// +optional
	MaxMessageSize int `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`
	// +optional
	MessageRetentionSeconds int `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	ReceiveWaitTimeSeconds int `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`
	// +optional
	RedrivePolicy string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VisibilityTimeoutSeconds int `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

type SqsQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
