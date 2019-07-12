package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStorageTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageTableSpec   `json:"spec,omitempty"`
	Status            AzurermStorageTableStatus `json:"status,omitempty"`
}

type AzurermStorageTableSpec struct {
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
	StorageAccountName string `json:"storage_account_name"`
}

type AzurermStorageTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStorageTableList is a list of AzurermStorageTables
type AzurermStorageTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageTable CRD objects
	Items []AzurermStorageTable `json:"items,omitempty"`
}
