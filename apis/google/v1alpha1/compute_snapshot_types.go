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

type ComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSnapshotSpec   `json:"spec,omitempty"`
	Status            ComputeSnapshotStatus `json:"status,omitempty"`
}

type ComputeSnapshotSpecSnapshotEncryptionKey struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	RawKey core.LocalObjectReference `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeSnapshotSpecSourceDiskEncryptionKey struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	RawKey core.LocalObjectReference `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeSnapshotSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotEncryptionKey []ComputeSnapshotSpecSnapshotEncryptionKey `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	SnapshotEncryptionKeyRaw core.LocalObjectReference `json:"snapshotEncryptionKeyRaw,omitempty" tf:"snapshot_encryption_key_raw,omitempty"`
	SourceDisk               string                    `json:"sourceDisk" tf:"source_disk"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceDiskEncryptionKey []ComputeSnapshotSpecSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey,omitempty" tf:"source_disk_encryption_key,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	SourceDiskEncryptionKeyRaw core.LocalObjectReference `json:"sourceDiskEncryptionKeyRaw,omitempty" tf:"source_disk_encryption_key_raw,omitempty"`
	// +optional
	Zone        string                    `json:"zone,omitempty" tf:"zone,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
