package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MediaServicesAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaServicesAccountSpec   `json:"spec,omitempty"`
	Status            MediaServicesAccountStatus `json:"status,omitempty"`
}

type MediaServicesAccountSpecStorageAccount struct {
	ID string `json:"ID" tf:"id"`
	// +optional
	IsPrimary bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`
}

type MediaServicesAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:UniqueItems=true
	StorageAccount []MediaServicesAccountSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
}

type MediaServicesAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MediaServicesAccountList is a list of MediaServicesAccounts
type MediaServicesAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MediaServicesAccount CRD objects
	Items []MediaServicesAccount `json:"items,omitempty"`
}
