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

type StorageTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageTableSpec   `json:"spec,omitempty"`
	Status            StorageTableStatus `json:"status,omitempty"`
}

type StorageTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name               string `json:"name" tf:"name"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
}

type StorageTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageTableList is a list of StorageTables
type StorageTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageTable CRD objects
	Items []StorageTable `json:"items,omitempty"`
}
