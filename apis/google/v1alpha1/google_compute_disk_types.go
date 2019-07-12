package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeDiskSpec   `json:"spec,omitempty"`
	Status            GoogleComputeDiskStatus `json:"status,omitempty"`
}

type GoogleComputeDiskSpecSourceImageEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeDiskSpecSourceSnapshotEncryptionKey struct {
	Sha256 string `json:"sha256"`
	RawKey string `json:"raw_key"`
}

type GoogleComputeDiskSpecDiskEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeDiskSpec struct {
	Project                     string                  `json:"project"`
	Labels                      map[string]string       `json:"labels"`
	LastDetachTimestamp         string                  `json:"last_detach_timestamp"`
	SourceImageId               string                  `json:"source_image_id"`
	DiskEncryptionKeyRaw        string                  `json:"disk_encryption_key_raw"`
	CreationTimestamp           string                  `json:"creation_timestamp"`
	Size                        int                     `json:"size"`
	Snapshot                    string                  `json:"snapshot"`
	SourceImageEncryptionKey    []GoogleComputeDiskSpec `json:"source_image_encryption_key"`
	Zone                        string                  `json:"zone"`
	SelfLink                    string                  `json:"self_link"`
	Description                 string                  `json:"description"`
	SourceSnapshotEncryptionKey []GoogleComputeDiskSpec `json:"source_snapshot_encryption_key"`
	LastAttachTimestamp         string                  `json:"last_attach_timestamp"`
	DiskEncryptionKeySha256     string                  `json:"disk_encryption_key_sha256"`
	LabelFingerprint            string                  `json:"label_fingerprint"`
	SourceSnapshotId            string                  `json:"source_snapshot_id"`
	Users                       []string                `json:"users"`
	Name                        string                  `json:"name"`
	DiskEncryptionKey           []GoogleComputeDiskSpec `json:"disk_encryption_key"`
	Image                       string                  `json:"image"`
	Type                        string                  `json:"type"`
}

type GoogleComputeDiskStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeDiskList is a list of GoogleComputeDisks
type GoogleComputeDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeDisk CRD objects
	Items []GoogleComputeDisk `json:"items,omitempty"`
}
