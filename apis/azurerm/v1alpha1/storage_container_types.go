package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StorageContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageContainerSpec   `json:"spec,omitempty"`
	Status            StorageContainerStatus `json:"status,omitempty"`
}

type StorageContainerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ContainerAccessType string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`
	Name                string `json:"name" tf:"name"`
	ResourceGroupName   string `json:"resourceGroupName" tf:"resource_group_name"`
	StorageAccountName  string `json:"storageAccountName" tf:"storage_account_name"`
}

type StorageContainerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageContainerList is a list of StorageContainers
type StorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageContainer CRD objects
	Items []StorageContainer `json:"items,omitempty"`
}
