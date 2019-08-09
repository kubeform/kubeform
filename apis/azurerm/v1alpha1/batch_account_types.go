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

type BatchAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccountSpec   `json:"spec,omitempty"`
	Status            BatchAccountStatus `json:"status,omitempty"`
}

type BatchAccountSpecKeyVaultReference struct {
	ID  string `json:"ID" tf:"id"`
	Url string `json:"url" tf:"url"`
}

type BatchAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AccountEndpoint string `json:"accountEndpoint,omitempty" tf:"account_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyVaultReference []BatchAccountSpecKeyVaultReference `json:"keyVaultReference,omitempty" tf:"key_vault_reference,omitempty"`
	Location          string                              `json:"location" tf:"location"`
	Name              string                              `json:"name" tf:"name"`
	// +optional
	PoolAllocationMode string `json:"poolAllocationMode,omitempty" tf:"pool_allocation_mode,omitempty"`
	// +optional
	PrimaryAccessKey  string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BatchAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BatchAccountSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
