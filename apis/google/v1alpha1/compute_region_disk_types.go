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

type ComputeRegionDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRegionDiskSpec   `json:"spec,omitempty"`
	Status            ComputeRegionDiskStatus `json:"status,omitempty"`
}

type ComputeRegionDiskSpecDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"raw_key,omitempty"`
}

type ComputeRegionDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"raw_key,omitempty"`
}

type ComputeRegionDiskSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey *[]ComputeRegionDiskSpec `json:"disk_encryption_key,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	ReplicaZones []string `json:"replica_zones"`
	// +optional
	Snapshot string `json:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey *[]ComputeRegionDiskSpec `json:"source_snapshot_encryption_key,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ComputeRegionDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRegionDiskList is a list of ComputeRegionDisks
type ComputeRegionDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRegionDisk CRD objects
	Items []ComputeRegionDisk `json:"items,omitempty"`
}
