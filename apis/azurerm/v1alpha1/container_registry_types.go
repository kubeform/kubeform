package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	// Sensitive Data. Provide secret name which contains one value only
	AccessKey core.LocalObjectReference `json:"accessKey" tf:"access_key"`
	Name      string                    `json:"name" tf:"name"`
}

type ContainerRegistrySpec struct {
	// +optional
	AdminEnabled bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	GeoreplicationLocations []string `json:"georeplicationLocations,omitempty" tf:"georeplication_locations,omitempty"`
	Location                string   `json:"location" tf:"location"`
	Name                    string   `json:"name" tf:"name"`
	ResourceGroupName       string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	StorageAccount []ContainerRegistrySpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ContainerRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
