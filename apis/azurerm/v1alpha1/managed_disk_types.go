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

type ManagedDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDiskSpec   `json:"spec,omitempty"`
	Status            ManagedDiskStatus `json:"status,omitempty"`
}

type ManagedDiskSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretUrl     string `json:"secret_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type ManagedDiskSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type ManagedDiskSpecEncryptionSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey *[]ManagedDiskSpecEncryptionSettings `json:"disk_encryption_key,omitempty"`
	Enabled           bool                                 `json:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyEncryptionKey *[]ManagedDiskSpecEncryptionSettings `json:"key_encryption_key,omitempty"`
}

type ManagedDiskSpec struct {
	CreateOption string `json:"create_option"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionSettings *[]ManagedDiskSpec `json:"encryption_settings,omitempty"`
	// +optional
	ImageReferenceId string `json:"image_reference_id,omitempty"`
	Location         string `json:"location"`
	Name             string `json:"name"`
	// +optional
	OsType            string `json:"os_type,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	SourceResourceId   string `json:"source_resource_id,omitempty"`
	StorageAccountType string `json:"storage_account_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type ManagedDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ManagedDiskList is a list of ManagedDisks
type ManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedDisk CRD objects
	Items []ManagedDisk `json:"items,omitempty"`
}
