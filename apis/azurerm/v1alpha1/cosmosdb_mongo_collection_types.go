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

type CosmosdbMongoCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbMongoCollectionSpec   `json:"spec,omitempty"`
	Status            CosmosdbMongoCollectionStatus `json:"status,omitempty"`
}

type CosmosdbMongoCollectionSpecIndexes struct {
	Key string `json:"key"`
	// +optional
	Unique bool `json:"unique,omitempty"`
}

type CosmosdbMongoCollectionSpec struct {
	AccountName  string `json:"account_name"`
	DatabaseName string `json:"database_name"`
	// +optional
	DefaultTtlSeconds int `json:"default_ttl_seconds,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Indexes           *[]CosmosdbMongoCollectionSpec `json:"indexes,omitempty"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	// +optional
	ShardKey string `json:"shard_key,omitempty"`
}

type CosmosdbMongoCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
