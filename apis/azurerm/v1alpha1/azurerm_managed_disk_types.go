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
	SourceVaultId string `json:"source_vault_id"`
	KeyUrl        string `json:"key_url"`
}

type AzurermManagedDiskSpecEncryptionSettings struct {
	DiskEncryptionKey []AzurermManagedDiskSpecEncryptionSettings `json:"disk_encryption_key"`
	KeyEncryptionKey  []AzurermManagedDiskSpecEncryptionSettings `json:"key_encryption_key"`
	Enabled           bool                                       `json:"enabled"`
}

type AzurermManagedDiskSpec struct {
	ImageReferenceId   string                   `json:"image_reference_id"`
	OsType             string                   `json:"os_type"`
	DiskSizeGb         int                      `json:"disk_size_gb"`
	ResourceGroupName  string                   `json:"resource_group_name"`
	Zones              []string                 `json:"zones"`
	StorageAccountType string                   `json:"storage_account_type"`
	CreateOption       string                   `json:"create_option"`
	SourceResourceId   string                   `json:"source_resource_id"`
	EncryptionSettings []AzurermManagedDiskSpec `json:"encryption_settings"`
	Tags               map[string]string        `json:"tags"`
	Name               string                   `json:"name"`
	Location           string                   `json:"location"`
	SourceUri          string                   `json:"source_uri"`
}

type AzurermManagedDiskStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermManagedDiskList is a list of AzurermManagedDisks
type AzurermManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermManagedDisk CRD objects
	Items []AzurermManagedDisk `json:"items,omitempty"`
}
