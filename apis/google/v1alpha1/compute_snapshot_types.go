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
	Sha256 string `json:"sha256" tf:"sha256"`
}

type ComputeSnapshotSpecSourceDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"-" sensitive:"true" tf:"raw_key,omitempty"`
}

type ComputeSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	CreationTimestamp string `json:"creationTimestamp" tf:"creation_timestamp"`
	// +optional
	Description      string `json:"description,omitempty" tf:"description,omitempty"`
	DiskSizeGb       int    `json:"diskSizeGb" tf:"disk_size_gb"`
	LabelFingerprint string `json:"labelFingerprint" tf:"label_fingerprint"`
	// +optional
	Labels   map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Licenses []string          `json:"licenses" tf:"licenses"`
	Name     string            `json:"name" tf:"name"`
	// +optional
	Project  string `json:"project,omitempty" tf:"project,omitempty"`
	SelfLink string `json:"selfLink" tf:"self_link"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotEncryptionKey []ComputeSnapshotSpecSnapshotEncryptionKey `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`
	// +optional
	SnapshotEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"snapshot_encryption_key_raw,omitempty"`
	// Deprecated
	SnapshotEncryptionKeySha256 string `json:"snapshotEncryptionKeySha256" tf:"snapshot_encryption_key_sha256"`
	SnapshotID                  int    `json:"snapshotID" tf:"snapshot_id"`
	SourceDisk                  string `json:"sourceDisk" tf:"source_disk"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceDiskEncryptionKey []ComputeSnapshotSpecSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey,omitempty" tf:"source_disk_encryption_key,omitempty"`
	// +optional
	SourceDiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"source_disk_encryption_key_raw,omitempty"`
	// Deprecated
	SourceDiskEncryptionKeySha256 string `json:"sourceDiskEncryptionKeySha256" tf:"source_disk_encryption_key_sha256"`
	SourceDiskLink                string `json:"sourceDiskLink" tf:"source_disk_link"`
	StorageBytes                  int    `json:"storageBytes" tf:"storage_bytes"`
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
