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

type AzurermDevTestLab struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestLabSpec   `json:"spec,omitempty"`
	Status            AzurermDevTestLabStatus `json:"status,omitempty"`
}

type AzurermDevTestLabSpec struct {
	ArtifactsStorageAccountId       string            `json:"artifacts_storage_account_id"`
	DefaultStorageAccountId         string            `json:"default_storage_account_id"`
	KeyVaultId                      string            `json:"key_vault_id"`
	PremiumDataDiskStorageAccountId string            `json:"premium_data_disk_storage_account_id"`
	Name                            string            `json:"name"`
	ResourceGroupName               string            `json:"resource_group_name"`
	StorageType                     string            `json:"storage_type"`
	Tags                            map[string]string `json:"tags"`
	Location                        string            `json:"location"`
	DefaultPremiumStorageAccountId  string            `json:"default_premium_storage_account_id"`
	UniqueIdentifier                string            `json:"unique_identifier"`
}

type AzurermDevTestLabStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDevTestLabList is a list of AzurermDevTestLabs
type AzurermDevTestLabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestLab CRD objects
	Items []AzurermDevTestLab `json:"items,omitempty"`
}
