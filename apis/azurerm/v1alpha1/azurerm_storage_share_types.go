package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStorageShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageShareSpec   `json:"spec,omitempty"`
	Status            AzurermStorageShareStatus `json:"status,omitempty"`
}

type AzurermStorageShareSpec struct {
	Url                string `json:"url"`
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
	StorageAccountName string `json:"storage_account_name"`
	Quota              int    `json:"quota"`
}



type AzurermStorageShareStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStorageShareList is a list of AzurermStorageShares
type AzurermStorageShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageShare CRD objects
	Items []AzurermStorageShare `json:"items,omitempty"`
}