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

type GoogleComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSnapshotSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSnapshotStatus `json:"status,omitempty"`
}

type GoogleComputeSnapshotSpecSnapshotEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeSnapshotSpecSourceDiskEncryptionKey struct {
	RawKey string `json:"raw_key"`
}

type GoogleComputeSnapshotSpec struct {
	StorageBytes                  int                         `json:"storage_bytes"`
	SourceDiskLink                string                      `json:"source_disk_link"`
	SnapshotEncryptionKeySha256   string                      `json:"snapshot_encryption_key_sha256"`
	SourceDiskEncryptionKeyRaw    string                      `json:"source_disk_encryption_key_raw"`
	SnapshotEncryptionKey         []GoogleComputeSnapshotSpec `json:"snapshot_encryption_key"`
	CreationTimestamp             string                      `json:"creation_timestamp"`
	SnapshotId                    int                         `json:"snapshot_id"`
	SourceDisk                    string                      `json:"source_disk"`
	SourceDiskEncryptionKey       []GoogleComputeSnapshotSpec `json:"source_disk_encryption_key"`
	Zone                          string                      `json:"zone"`
	SourceDiskEncryptionKeySha256 string                      `json:"source_disk_encryption_key_sha256"`
	Project                       string                      `json:"project"`
	SelfLink                      string                      `json:"self_link"`
	Description                   string                      `json:"description"`
	Labels                        map[string]string           `json:"labels"`
	LabelFingerprint              string                      `json:"label_fingerprint"`
	Licenses                      []string                    `json:"licenses"`
	SnapshotEncryptionKeyRaw      string                      `json:"snapshot_encryption_key_raw"`
	Name                          string                      `json:"name"`
	DiskSizeGb                    int                         `json:"disk_size_gb"`
}



type GoogleComputeSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeSnapshotList is a list of GoogleComputeSnapshots
type GoogleComputeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSnapshot CRD objects
	Items []GoogleComputeSnapshot `json:"items,omitempty"`
}