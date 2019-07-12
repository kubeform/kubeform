package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermCosmosdbSqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbSqlDatabaseSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbSqlDatabaseStatus `json:"status,omitempty"`
}

type AzurermCosmosdbSqlDatabaseSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	AccountName       string `json:"account_name"`
}



type AzurermCosmosdbSqlDatabaseStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermCosmosdbSqlDatabaseList is a list of AzurermCosmosdbSqlDatabases
type AzurermCosmosdbSqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbSqlDatabase CRD objects
	Items []AzurermCosmosdbSqlDatabase `json:"items,omitempty"`
}