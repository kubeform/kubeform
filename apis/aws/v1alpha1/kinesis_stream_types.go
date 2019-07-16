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

type KinesisStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisStreamSpec   `json:"spec,omitempty"`
	Status            KinesisStreamStatus `json:"status,omitempty"`
}

type KinesisStreamSpec struct {
	// +optional
	EncryptionType string `json:"encryption_type,omitempty"`
	// +optional
	EnforceConsumerDeletion bool `json:"enforce_consumer_deletion,omitempty"`
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	Name     string `json:"name"`
	// +optional
	RetentionPeriod int `json:"retention_period,omitempty"`
	ShardCount      int `json:"shard_count"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ShardLevelMetrics []string `json:"shard_level_metrics,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type KinesisStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KinesisStreamList is a list of KinesisStreams
type KinesisStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KinesisStream CRD objects
	Items []KinesisStream `json:"items,omitempty"`
}
