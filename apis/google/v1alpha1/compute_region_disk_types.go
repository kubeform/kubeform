package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeRegionDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeRegionDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeRegionDiskSpecDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	ReplicaZones []string `json:"replicaZones" tf:"replica_zones"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Snapshot string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey []ComputeRegionDiskSpecSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeRegionDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
