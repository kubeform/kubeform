package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	Name string `json:"name" tf:"name"`
	Type string `json:"type" tf:"type"`
}

type DynamodbTableSpecGlobalSecondaryIndex struct {
	HashKey string `json:"hashKey" tf:"hash_key"`
	Name    string `json:"name" tf:"name"`
	// +optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`
	ProjectionType   string   `json:"projectionType" tf:"projection_type"`
	// +optional
	RangeKey string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`
	// +optional
	ReadCapacity int `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`
	// +optional
	WriteCapacity int `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}

type DynamodbTableSpecLocalSecondaryIndex struct {
	Name string `json:"name" tf:"name"`
	// +optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`
	ProjectionType   string   `json:"projectionType" tf:"projection_type"`
	RangeKey         string   `json:"rangeKey" tf:"range_key"`
}

type DynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type DynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type DynamodbTableSpecTtl struct {
	AttributeName string `json:"attributeName" tf:"attribute_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DynamodbTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn       string                       `json:"arn,omitempty" tf:"arn,omitempty"`
	Attribute []DynamodbTableSpecAttribute `json:"attribute" tf:"attribute"`
	// +optional
	BillingMode string `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`
	// +optional
	GlobalSecondaryIndex []DynamodbTableSpecGlobalSecondaryIndex `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index,omitempty"`
	HashKey              string                                  `json:"hashKey" tf:"hash_key"`
	// +optional
	LocalSecondaryIndex []DynamodbTableSpecLocalSecondaryIndex `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index,omitempty"`
	Name                string                                 `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PointInTimeRecovery []DynamodbTableSpecPointInTimeRecovery `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`
	// +optional
	RangeKey string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`
	// +optional
	ReadCapacity int `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerSideEncryption []DynamodbTableSpecServerSideEncryption `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`
	// +optional
	StreamArn string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
	// +optional
	StreamEnabled bool `json:"streamEnabled,omitempty" tf:"stream_enabled,omitempty"`
	// +optional
	StreamLabel string `json:"streamLabel,omitempty" tf:"stream_label,omitempty"`
	// +optional
	StreamViewType string `json:"streamViewType,omitempty" tf:"stream_view_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ttl []DynamodbTableSpecTtl `json:"ttl,omitempty" tf:"ttl,omitempty"`
	// +optional
	WriteCapacity int `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}

type DynamodbTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DynamodbTableSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
