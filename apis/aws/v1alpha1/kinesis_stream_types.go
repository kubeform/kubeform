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

type KinesisStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisStreamSpec   `json:"spec,omitempty"`
	Status            KinesisStreamStatus `json:"status,omitempty"`
}

type KinesisStreamSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	EncryptionType string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	RetentionPeriod int `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
	ShardCount      int `json:"shardCount" tf:"shard_count"`
	// +optional
	ShardLevelMetrics []string `json:"shardLevelMetrics,omitempty" tf:"shard_level_metrics,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KinesisStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KinesisStreamSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
