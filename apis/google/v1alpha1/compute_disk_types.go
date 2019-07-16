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

type ComputeDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeDiskSpec   `json:"spec,omitempty"`
	Status            ComputeDiskStatus `json:"status,omitempty"`
}

type ComputeDiskSpecSourceImageEncryptionKey struct {
	// +optional
	RawKey string `json:"raw_key,omitempty"`
}

type ComputeDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"raw_key,omitempty"`
}

type ComputeDiskSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// Deprecated
	DiskEncryptionKeyRaw string `json:"disk_encryption_key_raw,omitempty"`
	// +optional
	Image string `json:"image,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	Snapshot string `json:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceImageEncryptionKey *[]ComputeDiskSpec `json:"source_image_encryption_key,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey *[]ComputeDiskSpec `json:"source_snapshot_encryption_key,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ComputeDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeDiskList is a list of ComputeDisks
type ComputeDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeDisk CRD objects
	Items []ComputeDisk `json:"items,omitempty"`
}
