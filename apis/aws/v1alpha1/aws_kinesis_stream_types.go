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

type AwsKinesisStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKinesisStreamSpec   `json:"spec,omitempty"`
	Status            AwsKinesisStreamStatus `json:"status,omitempty"`
}

type AwsKinesisStreamSpec struct {
	ShardCount              int               `json:"shard_count"`
	ShardLevelMetrics       []string          `json:"shard_level_metrics"`
	EncryptionType          string            `json:"encryption_type"`
	KmsKeyId                string            `json:"kms_key_id"`
	Name                    string            `json:"name"`
	RetentionPeriod         int               `json:"retention_period"`
	EnforceConsumerDeletion bool              `json:"enforce_consumer_deletion"`
	Arn                     string            `json:"arn"`
	Tags                    map[string]string `json:"tags"`
}

type AwsKinesisStreamStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKinesisStreamList is a list of AwsKinesisStreams
type AwsKinesisStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKinesisStream CRD objects
	Items []AwsKinesisStream `json:"items,omitempty"`
}
