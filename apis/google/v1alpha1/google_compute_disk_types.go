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
	DiskEncryptionKey           []GoogleComputeDiskSpec `json:"disk_encryption_key"`
	SourceImageEncryptionKey    []GoogleComputeDiskSpec `json:"source_image_encryption_key"`
	LastAttachTimestamp         string                  `json:"last_attach_timestamp"`
	LastDetachTimestamp         string                  `json:"last_detach_timestamp"`
	SourceSnapshotId            string                  `json:"source_snapshot_id"`
	SelfLink                    string                  `json:"self_link"`
	Description                 string                  `json:"description"`
	Labels                      map[string]string       `json:"labels"`
	Size                        int                     `json:"size"`
	Snapshot                    string                  `json:"snapshot"`
	Type                        string                  `json:"type"`
	SourceSnapshotEncryptionKey []GoogleComputeDiskSpec `json:"source_snapshot_encryption_key"`
	LabelFingerprint            string                  `json:"label_fingerprint"`
	DiskEncryptionKeyRaw        string                  `json:"disk_encryption_key_raw"`
	Project                     string                  `json:"project"`
	Name                        string                  `json:"name"`
	Image                       string                  `json:"image"`
	Zone                        string                  `json:"zone"`
	CreationTimestamp           string                  `json:"creation_timestamp"`
	SourceImageId               string                  `json:"source_image_id"`
	Users                       []string                `json:"users"`
	DiskEncryptionKeySha256     string                  `json:"disk_encryption_key_sha256"`
}



type GoogleComputeDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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