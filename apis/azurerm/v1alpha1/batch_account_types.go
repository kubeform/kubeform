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

type BatchAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccountSpec   `json:"spec,omitempty"`
	Status            BatchAccountStatus `json:"status,omitempty"`
}

type BatchAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Location    string                    `json:"location" tf:"location"`
	Name        string                    `json:"name" tf:"name"`
	// +optional
	PoolAllocationMode string `json:"poolAllocationMode,omitempty" tf:"pool_allocation_mode,omitempty"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BatchAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchAccountList is a list of BatchAccounts
type BatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchAccount CRD objects
	Items []BatchAccount `json:"items,omitempty"`
}
