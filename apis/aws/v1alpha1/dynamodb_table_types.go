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

type DynamodbTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbTableSpec   `json:"spec,omitempty"`
	Status            DynamodbTableStatus `json:"status,omitempty"`
}

type DynamodbTableSpecAttribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DynamodbTableSpecGlobalSecondaryIndex struct {
	HashKey string `json:"hash_key"`
	Name    string `json:"name"`
	// +optional
	NonKeyAttributes []string `json:"non_key_attributes,omitempty"`
	ProjectionType   string   `json:"projection_type"`
	// +optional
	RangeKey string `json:"range_key,omitempty"`
	// +optional
	ReadCapacity int `json:"read_capacity,omitempty"`
	// +optional
	WriteCapacity int `json:"write_capacity,omitempty"`
}

type DynamodbTableSpecLocalSecondaryIndex struct {
	Name string `json:"name"`
	// +optional
	NonKeyAttributes []string `json:"non_key_attributes,omitempty"`
	ProjectionType   string   `json:"projection_type"`
	RangeKey         string   `json:"range_key"`
}

type DynamodbTableSpecTtl struct {
	AttributeName string `json:"attribute_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type DynamodbTableSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	Attribute []DynamodbTableSpec `json:"attribute"`
	// +optional
	BillingMode string `json:"billing_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GlobalSecondaryIndex *[]DynamodbTableSpec `json:"global_secondary_index,omitempty"`
	HashKey              string               `json:"hash_key"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LocalSecondaryIndex *[]DynamodbTableSpec `json:"local_secondary_index,omitempty"`
	Name                string               `json:"name"`
	// +optional
	RangeKey string `json:"range_key,omitempty"`
	// +optional
	ReadCapacity int `json:"read_capacity,omitempty"`
	// +optional
	StreamEnabled bool `json:"stream_enabled,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ttl *[]DynamodbTableSpec `json:"ttl,omitempty"`
	// +optional
	WriteCapacity int `json:"write_capacity,omitempty"`
}

type DynamodbTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DynamodbTableList is a list of DynamodbTables
type DynamodbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DynamodbTable CRD objects
	Items []DynamodbTable `json:"items,omitempty"`
}
