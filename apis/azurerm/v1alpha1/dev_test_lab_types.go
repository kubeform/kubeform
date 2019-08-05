package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DevTestLab struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestLabSpec   `json:"spec,omitempty"`
	Status            DevTestLabStatus `json:"status,omitempty"`
}

type DevTestLabSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ArtifactsStorageAccountID string `json:"artifactsStorageAccountID,omitempty" tf:"artifacts_storage_account_id,omitempty"`
	// +optional
	DefaultPremiumStorageAccountID string `json:"defaultPremiumStorageAccountID,omitempty" tf:"default_premium_storage_account_id,omitempty"`
	// +optional
	DefaultStorageAccountID string `json:"defaultStorageAccountID,omitempty" tf:"default_storage_account_id,omitempty"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	Location   string `json:"location" tf:"location"`
	Name       string `json:"name" tf:"name"`
	// +optional
	PremiumDataDiskStorageAccountID string `json:"premiumDataDiskStorageAccountID,omitempty" tf:"premium_data_disk_storage_account_id,omitempty"`
	ResourceGroupName               string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type DevTestLabStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DevTestLabSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestLabList is a list of DevTestLabs
type DevTestLabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestLab CRD objects
	Items []DevTestLab `json:"items,omitempty"`
}
