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

type ComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSnapshotSpec   `json:"spec,omitempty"`
	Status            ComputeSnapshotStatus `json:"status,omitempty"`
}

type ComputeSnapshotSpecSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"-" sensitive:"true" tf:"raw_key,omitempty"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeSnapshotSpecSourceDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"-" sensitive:"true" tf:"raw_key,omitempty"`
}

type ComputeSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Licenses []string `json:"licenses,omitempty" tf:"licenses,omitempty"`
	Name     string   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotEncryptionKey []ComputeSnapshotSpecSnapshotEncryptionKey `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`
	// +optional
	SnapshotEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"snapshot_encryption_key_raw,omitempty"`
	// +optional
	// Deprecated
	SnapshotEncryptionKeySha256 string `json:"snapshotEncryptionKeySha256,omitempty" tf:"snapshot_encryption_key_sha256,omitempty"`
	// +optional
	SnapshotID int    `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	SourceDisk string `json:"sourceDisk" tf:"source_disk"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceDiskEncryptionKey []ComputeSnapshotSpecSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey,omitempty" tf:"source_disk_encryption_key,omitempty"`
	// +optional
	SourceDiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"source_disk_encryption_key_raw,omitempty"`
	// +optional
	// Deprecated
	SourceDiskEncryptionKeySha256 string `json:"sourceDiskEncryptionKeySha256,omitempty" tf:"source_disk_encryption_key_sha256,omitempty"`
	// +optional
	SourceDiskLink string `json:"sourceDiskLink,omitempty" tf:"source_disk_link,omitempty"`
	// +optional
	StorageBytes int `json:"storageBytes,omitempty" tf:"storage_bytes,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComputeSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSnapshotList is a list of ComputeSnapshots
type ComputeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSnapshot CRD objects
	Items []ComputeSnapshot `json:"items,omitempty"`
}
