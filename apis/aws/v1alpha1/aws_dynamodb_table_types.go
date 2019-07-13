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

type AwsDynamodbTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDynamodbTableSpec   `json:"spec,omitempty"`
	Status            AwsDynamodbTableStatus `json:"status,omitempty"`
}

type AwsDynamodbTableSpecAttribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AwsDynamodbTableSpecTtl struct {
	AttributeName string `json:"attribute_name"`
	Enabled       bool   `json:"enabled"`
}

type AwsDynamodbTableSpecGlobalSecondaryIndex struct {
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
	WriteCapacity    int      `json:"write_capacity"`
	ReadCapacity     int      `json:"read_capacity"`
	HashKey          string   `json:"hash_key"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
}

type AwsDynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpecLocalSecondaryIndex struct {
	Name             string   `json:"name"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
}

type AwsDynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpec struct {
	Attribute            []AwsDynamodbTableSpec `json:"attribute"`
	Ttl                  []AwsDynamodbTableSpec `json:"ttl"`
	GlobalSecondaryIndex []AwsDynamodbTableSpec `json:"global_secondary_index"`
	StreamViewType       string                 `json:"stream_view_type"`
	PointInTimeRecovery  []AwsDynamodbTableSpec `json:"point_in_time_recovery"`
	StreamArn            string                 `json:"stream_arn"`
	StreamLabel          string                 `json:"stream_label"`
	Name                 string                 `json:"name"`
	HashKey              string                 `json:"hash_key"`
	RangeKey             string                 `json:"range_key"`
	BillingMode          string                 `json:"billing_mode"`
	WriteCapacity        int                    `json:"write_capacity"`
	StreamEnabled        bool                   `json:"stream_enabled"`
	Tags                 map[string]string      `json:"tags"`
	Arn                  string                 `json:"arn"`
	ReadCapacity         int                    `json:"read_capacity"`
	LocalSecondaryIndex  []AwsDynamodbTableSpec `json:"local_secondary_index"`
	ServerSideEncryption []AwsDynamodbTableSpec `json:"server_side_encryption"`
}



type AwsDynamodbTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDynamodbTableList is a list of AwsDynamodbTables
type AwsDynamodbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDynamodbTable CRD objects
	Items []AwsDynamodbTable `json:"items,omitempty"`
}