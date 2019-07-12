package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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

type AwsDynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpecLocalSecondaryIndex struct {
	Name             string   `json:"name"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
}

type AwsDynamodbTableSpecGlobalSecondaryIndex struct {
	ReadCapacity     int      `json:"read_capacity"`
	HashKey          string   `json:"hash_key"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
	WriteCapacity    int      `json:"write_capacity"`
}

type AwsDynamodbTableSpecTtl struct {
	AttributeName string `json:"attribute_name"`
	Enabled       bool   `json:"enabled"`
}

type AwsDynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpec struct {
	Name                 string                 `json:"name"`
	Attribute            []AwsDynamodbTableSpec `json:"attribute"`
	StreamEnabled        bool                   `json:"stream_enabled"`
	StreamArn            string                 `json:"stream_arn"`
	ServerSideEncryption []AwsDynamodbTableSpec `json:"server_side_encryption"`
	RangeKey             string                 `json:"range_key"`
	LocalSecondaryIndex  []AwsDynamodbTableSpec `json:"local_secondary_index"`
	GlobalSecondaryIndex []AwsDynamodbTableSpec `json:"global_secondary_index"`
	Arn                  string                 `json:"arn"`
	Ttl                  []AwsDynamodbTableSpec `json:"ttl"`
	StreamLabel          string                 `json:"stream_label"`
	PointInTimeRecovery  []AwsDynamodbTableSpec `json:"point_in_time_recovery"`
	HashKey              string                 `json:"hash_key"`
	BillingMode          string                 `json:"billing_mode"`
	WriteCapacity        int                    `json:"write_capacity"`
	ReadCapacity         int                    `json:"read_capacity"`
	StreamViewType       string                 `json:"stream_view_type"`
	Tags                 map[string]string      `json:"tags"`
}

type AwsDynamodbTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTables
type AwsDynamodbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDynamodbTable CRD objects
	Items []AwsDynamodbTable `json:"items,omitempty"`
}
