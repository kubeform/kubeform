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
	Tags                    map[string]string              `json:"tags"`
	ResourceGroupName       string                         `json:"resource_group_name"`
	Sku                     string                         `json:"sku"`
	AdminEnabled            bool                           `json:"admin_enabled"`
	LoginServer             string                         `json:"login_server"`
	AdminUsername           string                         `json:"admin_username"`
	Name                    string                         `json:"name"`
	Location                string                         `json:"location"`
	GeoreplicationLocations []string                       `json:"georeplication_locations"`
	StorageAccountId        string                         `json:"storage_account_id"`
	StorageAccount          []AzurermContainerRegistrySpec `json:"storage_account"`
}



type AzurermContainerRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermContainerRegistryList is a list of AzurermContainerRegistrys
type AzurermContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerRegistry CRD objects
	Items []AzurermContainerRegistry `json:"items,omitempty"`
}