package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermCosmosdbMongoDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbMongoDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbMongoDatabaseStatus `json:"status,omitempty"`
}

type AzurermCosmosdbMongoDatabaseSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	AccountName       string `json:"account_name"`
}

type AzurermCosmosdbMongoDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermCosmosdbMongoDatabaseList is a list of AzurermCosmosdbMongoDatabases
type AzurermCosmosdbMongoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbMongoDatabase CRD objects
	Items []AzurermCosmosdbMongoDatabase `json:"items,omitempty"`
}
