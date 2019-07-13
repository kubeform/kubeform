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

type AzurermKeyVaultAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKeyVaultAccessPolicySpec   `json:"spec,omitempty"`
	Status            AzurermKeyVaultAccessPolicyStatus `json:"status,omitempty"`
}

type AzurermKeyVaultAccessPolicySpec struct {
	ObjectId               string   `json:"object_id"`
	KeyPermissions         []string `json:"key_permissions"`
	VaultName              string   `json:"vault_name"`
	ResourceGroupName      string   `json:"resource_group_name"`
	TenantId               string   `json:"tenant_id"`
	SecretPermissions      []string `json:"secret_permissions"`
	StoragePermissions     []string `json:"storage_permissions"`
	KeyVaultId             string   `json:"key_vault_id"`
	ApplicationId          string   `json:"application_id"`
	CertificatePermissions []string `json:"certificate_permissions"`
}



type AzurermKeyVaultAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermKeyVaultAccessPolicyList is a list of AzurermKeyVaultAccessPolicys
type AzurermKeyVaultAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKeyVaultAccessPolicy CRD objects
	Items []AzurermKeyVaultAccessPolicy `json:"items,omitempty"`
}