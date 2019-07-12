package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermContainerRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermContainerRegistrySpec   `json:"spec,omitempty"`
	Status            AzurermContainerRegistryStatus `json:"status,omitempty"`
}

type AzurermContainerRegistrySpecStorageAccount struct {
	AccessKey string `json:"access_key"`
	Name      string `json:"name"`
}

type AzurermContainerRegistrySpec struct {
	AdminPassword           string                         `json:"admin_password"`
	Name                    string                         `json:"name"`
	ResourceGroupName       string                         `json:"resource_group_name"`
	AdminEnabled            bool                           `json:"admin_enabled"`
	StorageAccountId        string                         `json:"storage_account_id"`
	StorageAccount          []AzurermContainerRegistrySpec `json:"storage_account"`
	LoginServer             string                         `json:"login_server"`
	AdminUsername           string                         `json:"admin_username"`
	Tags                    map[string]string              `json:"tags"`
	Location                string                         `json:"location"`
	Sku                     string                         `json:"sku"`
	GeoreplicationLocations []string                       `json:"georeplication_locations"`
}



type AzurermContainerRegistryStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermContainerRegistryList is a list of AzurermContainerRegistrys
type AzurermContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerRegistry CRD objects
	Items []AzurermContainerRegistry `json:"items,omitempty"`
}