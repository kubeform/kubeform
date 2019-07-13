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

type AzurermCosmosdbMongoCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbMongoCollectionSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbMongoCollectionStatus `json:"status,omitempty"`
}

type AzurermCosmosdbMongoCollectionSpecIndexes struct {
	Unique bool   `json:"unique"`
	Key    string `json:"key"`
}

type AzurermCosmosdbMongoCollectionSpec struct {
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	AccountName       string                               `json:"account_name"`
	DatabaseName      string                               `json:"database_name"`
	ShardKey          string                               `json:"shard_key"`
	DefaultTtlSeconds int                                  `json:"default_ttl_seconds"`
	Indexes           []AzurermCosmosdbMongoCollectionSpec `json:"indexes"`
}



type AzurermCosmosdbMongoCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCosmosdbMongoCollectionList is a list of AzurermCosmosdbMongoCollections
type AzurermCosmosdbMongoCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbMongoCollection CRD objects
	Items []AzurermCosmosdbMongoCollection `json:"items,omitempty"`
}