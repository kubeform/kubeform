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

type GoogleComputeRegionDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRegionDiskSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRegionDiskStatus `json:"status,omitempty"`
}

type GoogleComputeRegionDiskSpecDiskEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeRegionDiskSpecSourceSnapshotEncryptionKey struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type GoogleComputeRegionDiskSpec struct {
	Description                 string                        `json:"description"`
	DiskEncryptionKey           []GoogleComputeRegionDiskSpec `json:"disk_encryption_key"`
	Snapshot                    string                        `json:"snapshot"`
	Type                        string                        `json:"type"`
	CreationTimestamp           string                        `json:"creation_timestamp"`
	LabelFingerprint            string                        `json:"label_fingerprint"`
	Name                        string                        `json:"name"`
	ReplicaZones                []string                      `json:"replica_zones"`
	SourceSnapshotEncryptionKey []GoogleComputeRegionDiskSpec `json:"source_snapshot_encryption_key"`
	LastAttachTimestamp         string                        `json:"last_attach_timestamp"`
	LastDetachTimestamp         string                        `json:"last_detach_timestamp"`
	Project                     string                        `json:"project"`
	SelfLink                    string                        `json:"self_link"`
	Labels                      map[string]string             `json:"labels"`
	SourceSnapshotId            string                        `json:"source_snapshot_id"`
	Users                       []string                      `json:"users"`
	Region                      string                        `json:"region"`
	Size                        int                           `json:"size"`
}



type GoogleComputeRegionDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRegionDiskList is a list of GoogleComputeRegionDisks
type GoogleComputeRegionDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionDisk CRD objects
	Items []GoogleComputeRegionDisk `json:"items,omitempty"`
}