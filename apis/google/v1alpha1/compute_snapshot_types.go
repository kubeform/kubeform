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

type ComputeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSnapshotSpec   `json:"spec,omitempty"`
	Status            ComputeSnapshotStatus `json:"status,omitempty"`
}

type ComputeSnapshotSpecSourceDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"raw_key,omitempty"`
}

type ComputeSnapshotSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	// Deprecated
	SnapshotEncryptionKeyRaw string `json:"snapshot_encryption_key_raw,omitempty"`
	SourceDisk               string `json:"source_disk"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceDiskEncryptionKey *[]ComputeSnapshotSpec `json:"source_disk_encryption_key,omitempty"`
	// +optional
	// Deprecated
	SourceDiskEncryptionKeyRaw string `json:"source_disk_encryption_key_raw,omitempty"`
}

type ComputeSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
