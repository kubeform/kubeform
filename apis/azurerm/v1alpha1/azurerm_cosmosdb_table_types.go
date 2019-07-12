package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermCosmosdbTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbTableSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbTableStatus `json:"status,omitempty"`
}

type AzurermCosmosdbTableSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	AccountName       string `json:"account_name"`
}



type AzurermCosmosdbTableStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermCosmosdbTableList is a list of AzurermCosmosdbTables
type AzurermCosmosdbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbTable CRD objects
	Items []AzurermCosmosdbTable `json:"items,omitempty"`
}