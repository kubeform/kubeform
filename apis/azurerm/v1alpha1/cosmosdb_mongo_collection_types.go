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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type CosmosdbMongoCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbMongoCollectionSpec   `json:"spec,omitempty"`
	Status            CosmosdbMongoCollectionStatus `json:"status,omitempty"`
}

type CosmosdbMongoCollectionSpecIndexes struct {
	Key string `json:"key" tf:"key"`
	// +optional
	Unique bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type CosmosdbMongoCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName  string `json:"accountName" tf:"account_name"`
	DatabaseName string `json:"databaseName" tf:"database_name"`
	// +optional
	DefaultTtlSeconds int `json:"defaultTtlSeconds,omitempty" tf:"default_ttl_seconds,omitempty"`
	// +optional
	Indexes           []CosmosdbMongoCollectionSpecIndexes `json:"indexes,omitempty" tf:"indexes,omitempty"`
	Name              string                               `json:"name" tf:"name"`
	ResourceGroupName string                               `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ShardKey string `json:"shardKey,omitempty" tf:"shard_key,omitempty"`
}

type CosmosdbMongoCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CosmosdbMongoCollectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbMongoCollectionList is a list of CosmosdbMongoCollections
type CosmosdbMongoCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbMongoCollection CRD objects
	Items []CosmosdbMongoCollection `json:"items,omitempty"`
}
