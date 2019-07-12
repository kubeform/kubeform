package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	DiskEncryptionKey           []GoogleComputeRegionDiskSpec `json:"disk_encryption_key"`
	Size                        int                           `json:"size"`
	Snapshot                    string                        `json:"snapshot"`
	LabelFingerprint            string                        `json:"label_fingerprint"`
	Users                       []string                      `json:"users"`
	Description                 string                        `json:"description"`
	Labels                      map[string]string             `json:"labels"`
	SourceSnapshotEncryptionKey []GoogleComputeRegionDiskSpec `json:"source_snapshot_encryption_key"`
	CreationTimestamp           string                        `json:"creation_timestamp"`
	LastAttachTimestamp         string                        `json:"last_attach_timestamp"`
	Type                        string                        `json:"type"`
	Project                     string                        `json:"project"`
	SelfLink                    string                        `json:"self_link"`
	Name                        string                        `json:"name"`
	ReplicaZones                []string                      `json:"replica_zones"`
	Region                      string                        `json:"region"`
	LastDetachTimestamp         string                        `json:"last_detach_timestamp"`
	SourceSnapshotId            string                        `json:"source_snapshot_id"`
}



type GoogleComputeRegionDiskStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRegionDiskList is a list of GoogleComputeRegionDisks
type GoogleComputeRegionDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionDisk CRD objects
	Items []GoogleComputeRegionDisk `json:"items,omitempty"`
}