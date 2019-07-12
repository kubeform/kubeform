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

type GoogleComputeDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeDiskSpec   `json:"spec,omitempty"`
	Status            GoogleComputeDiskStatus `json:"status,omitempty"`
}

type GoogleComputeDiskSpecDiskEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeDiskSpecSourceImageEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeDiskSpecSourceSnapshotEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeDiskSpec struct {
	Description                 string                  `json:"description"`
	Zone                        string                  `json:"zone"`
	SourceImageId               string                  `json:"source_image_id"`
	Users                       []string                `json:"users"`
	DiskEncryptionKey           []GoogleComputeDiskSpec `json:"disk_encryption_key"`
	Image                       string                  `json:"image"`
	Snapshot                    string                  `json:"snapshot"`
	LabelFingerprint            string                  `json:"label_fingerprint"`
	LastDetachTimestamp         string                  `json:"last_detach_timestamp"`
	Size                        int                     `json:"size"`
	SourceImageEncryptionKey    []GoogleComputeDiskSpec `json:"source_image_encryption_key"`
	CreationTimestamp           string                  `json:"creation_timestamp"`
	LastAttachTimestamp         string                  `json:"last_attach_timestamp"`
	SelfLink                    string                  `json:"self_link"`
	DiskEncryptionKeyRaw        string                  `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256     string                  `json:"disk_encryption_key_sha256"`
	Project                     string                  `json:"project"`
	Name                        string                  `json:"name"`
	Labels                      map[string]string       `json:"labels"`
	SourceSnapshotEncryptionKey []GoogleComputeDiskSpec `json:"source_snapshot_encryption_key"`
	Type                        string                  `json:"type"`
	SourceSnapshotId            string                  `json:"source_snapshot_id"`
}

type GoogleComputeDiskStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeDiskList is a list of GoogleComputeDisks
type GoogleComputeDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeDisk CRD objects
	Items []GoogleComputeDisk `json:"items,omitempty"`
}
