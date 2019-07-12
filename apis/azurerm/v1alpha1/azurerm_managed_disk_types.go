package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermManagedDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermManagedDiskSpec   `json:"spec,omitempty"`
	Status            AzurermManagedDiskStatus `json:"status,omitempty"`
}

type AzurermManagedDiskSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretUrl     string `json:"secret_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type AzurermManagedDiskSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type AzurermManagedDiskSpecEncryptionSettings struct {
	Enabled           bool                                       `json:"enabled"`
	DiskEncryptionKey []AzurermManagedDiskSpecEncryptionSettings `json:"disk_encryption_key"`
	KeyEncryptionKey  []AzurermManagedDiskSpecEncryptionSettings `json:"key_encryption_key"`
}

type AzurermManagedDiskSpec struct {
	StorageAccountType string                   `json:"storage_account_type"`
	SourceResourceId   string                   `json:"source_resource_id"`
	ImageReferenceId   string                   `json:"image_reference_id"`
	Tags               map[string]string        `json:"tags"`
	CreateOption       string                   `json:"create_option"`
	SourceUri          string                   `json:"source_uri"`
	OsType             string                   `json:"os_type"`
	DiskSizeGb         int                      `json:"disk_size_gb"`
	Name               string                   `json:"name"`
	Location           string                   `json:"location"`
	ResourceGroupName  string                   `json:"resource_group_name"`
	Zones              []string                 `json:"zones"`
	EncryptionSettings []AzurermManagedDiskSpec `json:"encryption_settings"`
}

type AzurermManagedDiskStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermManagedDiskList is a list of AzurermManagedDisks
type AzurermManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermManagedDisk CRD objects
	Items []AzurermManagedDisk `json:"items,omitempty"`
}
