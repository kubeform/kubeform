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

type Snapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

type SnapshotSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretUrl     string `json:"secret_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type SnapshotSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type SnapshotSpecEncryptionSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey *[]SnapshotSpecEncryptionSettings `json:"disk_encryption_key,omitempty"`
	Enabled           bool                              `json:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyEncryptionKey *[]SnapshotSpecEncryptionSettings `json:"key_encryption_key,omitempty"`
}

type SnapshotSpec struct {
	CreateOption string `json:"create_option"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionSettings *[]SnapshotSpec `json:"encryption_settings,omitempty"`
	Location           string          `json:"location"`
	Name               string          `json:"name"`
	ResourceGroupName  string          `json:"resource_group_name"`
	// +optional
	SourceResourceId string `json:"source_resource_id,omitempty"`
	// +optional
	SourceUri string `json:"source_uri,omitempty"`
	// +optional
	StorageAccountId string `json:"storage_account_id,omitempty"`
}

type SnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnapshotList is a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Snapshot CRD objects
	Items []Snapshot `json:"items,omitempty"`
}
