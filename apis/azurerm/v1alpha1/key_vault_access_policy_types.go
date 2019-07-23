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

type KeyVaultAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultAccessPolicySpec   `json:"spec,omitempty"`
	Status            KeyVaultAccessPolicyStatus `json:"status,omitempty"`
}

type KeyVaultAccessPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ApplicationID string `json:"applicationID,omitempty" tf:"application_id,omitempty"`
	// +optional
	CertificatePermissions []string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`
	// +optional
	KeyPermissions []string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	ObjectID   string `json:"objectID" tf:"object_id"`
	// +optional
	// Deprecated
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	// +optional
	SecretPermissions []string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`
	// +optional
	StoragePermissions []string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`
	TenantID           string   `json:"tenantID" tf:"tenant_id"`
	// +optional
	// Deprecated
	VaultName string `json:"vaultName,omitempty" tf:"vault_name,omitempty"`
}

type KeyVaultAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultAccessPolicyList is a list of KeyVaultAccessPolicys
type KeyVaultAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultAccessPolicy CRD objects
	Items []KeyVaultAccessPolicy `json:"items,omitempty"`
}
