package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSnapshotSpec   `json:"spec,omitempty"`
	Status            AzurermSnapshotStatus `json:"status,omitempty"`
}

type AzurermSnapshotSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretUrl     string `json:"secret_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type AzurermSnapshotSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type AzurermSnapshotSpecEncryptionSettings struct {
	Enabled           bool                                    `json:"enabled"`
	DiskEncryptionKey []AzurermSnapshotSpecEncryptionSettings `json:"disk_encryption_key"`
	KeyEncryptionKey  []AzurermSnapshotSpecEncryptionSettings `json:"key_encryption_key"`
}

type AzurermSnapshotSpec struct {
	Name               string                `json:"name"`
	Location           string                `json:"location"`
	ResourceGroupName  string                `json:"resource_group_name"`
	CreateOption       string                `json:"create_option"`
	StorageAccountId   string                `json:"storage_account_id"`
	Tags               map[string]string     `json:"tags"`
	SourceUri          string                `json:"source_uri"`
	SourceResourceId   string                `json:"source_resource_id"`
	DiskSizeGb         int                   `json:"disk_size_gb"`
	EncryptionSettings []AzurermSnapshotSpec `json:"encryption_settings"`
}

type AzurermSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSnapshotList is a list of AzurermSnapshots
type AzurermSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSnapshot CRD objects
	Items []AzurermSnapshot `json:"items,omitempty"`
}
