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
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type AzurermManagedDiskSpecEncryptionSettings struct {
	Enabled           bool                                       `json:"enabled"`
	DiskEncryptionKey []AzurermManagedDiskSpecEncryptionSettings `json:"disk_encryption_key"`
	KeyEncryptionKey  []AzurermManagedDiskSpecEncryptionSettings `json:"key_encryption_key"`
}

type AzurermManagedDiskSpec struct {
	DiskSizeGb         int                      `json:"disk_size_gb"`
	Tags               map[string]string        `json:"tags"`
	Name               string                   `json:"name"`
	Location           string                   `json:"location"`
	Zones              []string                 `json:"zones"`
	CreateOption       string                   `json:"create_option"`
	ImageReferenceId   string                   `json:"image_reference_id"`
	EncryptionSettings []AzurermManagedDiskSpec `json:"encryption_settings"`
	ResourceGroupName  string                   `json:"resource_group_name"`
	StorageAccountType string                   `json:"storage_account_type"`
	SourceUri          string                   `json:"source_uri"`
	SourceResourceId   string                   `json:"source_resource_id"`
	OsType             string                   `json:"os_type"`
}



type AzurermManagedDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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