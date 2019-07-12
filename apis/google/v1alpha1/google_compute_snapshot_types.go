package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSnapshotSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSnapshotStatus `json:"status,omitempty"`
}

type GoogleComputeSnapshotSpecSourceDiskEncryptionKey struct {
	RawKey string `json:"raw_key"`
}

type GoogleComputeSnapshotSpecSnapshotEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeSnapshotSpec struct {
	Name                          string                      `json:"name"`
	Zone                          string                      `json:"zone"`
	Licenses                      []string                    `json:"licenses"`
	SnapshotId                    int                         `json:"snapshot_id"`
	SnapshotEncryptionKeySha256   string                      `json:"snapshot_encryption_key_sha256"`
	SourceDiskEncryptionKeyRaw    string                      `json:"source_disk_encryption_key_raw"`
	Description                   string                      `json:"description"`
	CreationTimestamp             string                      `json:"creation_timestamp"`
	DiskSizeGb                    int                         `json:"disk_size_gb"`
	LabelFingerprint              string                      `json:"label_fingerprint"`
	SourceDiskEncryptionKey       []GoogleComputeSnapshotSpec `json:"source_disk_encryption_key"`
	SourceDiskLink                string                      `json:"source_disk_link"`
	SnapshotEncryptionKeyRaw      string                      `json:"snapshot_encryption_key_raw"`
	SourceDiskEncryptionKeySha256 string                      `json:"source_disk_encryption_key_sha256"`
	SelfLink                      string                      `json:"self_link"`
	SourceDisk                    string                      `json:"source_disk"`
	Labels                        map[string]string           `json:"labels"`
	SnapshotEncryptionKey         []GoogleComputeSnapshotSpec `json:"snapshot_encryption_key"`
	StorageBytes                  int                         `json:"storage_bytes"`
	Project                       string                      `json:"project"`
}

type GoogleComputeSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSnapshotList is a list of GoogleComputeSnapshots
type GoogleComputeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSnapshot CRD objects
	Items []GoogleComputeSnapshot `json:"items,omitempty"`
}
