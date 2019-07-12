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

type AwsDynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpecLocalSecondaryIndex struct {
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
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
	Name             string   `json:"name"`
	WriteCapacity    int      `json:"write_capacity"`
	ReadCapacity     int      `json:"read_capacity"`
	HashKey          string   `json:"hash_key"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
}

type AwsDynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpec struct {
	ServerSideEncryption []AwsDynamodbTableSpec `json:"server_side_encryption"`
	HashKey              string                 `json:"hash_key"`
	RangeKey             string                 `json:"range_key"`
	WriteCapacity        int                    `json:"write_capacity"`
	LocalSecondaryIndex  []AwsDynamodbTableSpec `json:"local_secondary_index"`
	StreamLabel          string                 `json:"stream_label"`
	Attribute            []AwsDynamodbTableSpec `json:"attribute"`
	Ttl                  []AwsDynamodbTableSpec `json:"ttl"`
	StreamViewType       string                 `json:"stream_view_type"`
	StreamArn            string                 `json:"stream_arn"`
	Tags                 map[string]string      `json:"tags"`
	Arn                  string                 `json:"arn"`
	BillingMode          string                 `json:"billing_mode"`
	ReadCapacity         int                    `json:"read_capacity"`
	GlobalSecondaryIndex []AwsDynamodbTableSpec `json:"global_secondary_index"`
	Name                 string                 `json:"name"`
	StreamEnabled        bool                   `json:"stream_enabled"`
	PointInTimeRecovery  []AwsDynamodbTableSpec `json:"point_in_time_recovery"`
}

type AwsDynamodbTableStatus struct {
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
