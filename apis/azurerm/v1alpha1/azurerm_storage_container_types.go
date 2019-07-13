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

type AzurermStorageContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageContainerSpec   `json:"spec,omitempty"`
	Status            AzurermStorageContainerStatus `json:"status,omitempty"`
}

type AzurermStorageContainerSpec struct {
	Properties          map[string]string `json:"properties"`
	Name                string            `json:"name"`
	ResourceGroupName   string            `json:"resource_group_name"`
	StorageAccountName  string            `json:"storage_account_name"`
	ContainerAccessType string            `json:"container_access_type"`
}



type AzurermStorageContainerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStorageContainerList is a list of AzurermStorageContainers
type AzurermStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageContainer CRD objects
	Items []AzurermStorageContainer `json:"items,omitempty"`
}