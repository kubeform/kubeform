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

type ContainerRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerRegistrySpec   `json:"spec,omitempty"`
	Status            ContainerRegistryStatus `json:"status,omitempty"`
}

type ContainerRegistrySpecStorageAccount struct {
	AccessKey string `json:"access_key"`
	Name      string `json:"name"`
}

type ContainerRegistrySpec struct {
	// +optional
	AdminEnabled bool `json:"admin_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	GeoreplicationLocations []string `json:"georeplication_locations,omitempty"`
	Location                string   `json:"location"`
	Name                    string   `json:"name"`
	ResourceGroupName       string   `json:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	StorageAccount *[]ContainerRegistrySpec `json:"storage_account,omitempty"`
	// +optional
	StorageAccountId string `json:"storage_account_id,omitempty"`
}

type ContainerRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerRegistryList is a list of ContainerRegistrys
type ContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerRegistry CRD objects
	Items []ContainerRegistry `json:"items,omitempty"`
}
