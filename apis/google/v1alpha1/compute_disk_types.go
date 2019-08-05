package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeDiskSpec   `json:"spec,omitempty"`
	Status            ComputeDiskStatus `json:"status,omitempty"`
}

type ComputeDiskSpecDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeDiskSpecSourceImageEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type ComputeDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeDiskSpecDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	// +optional
	DiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Snapshot string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceImageEncryptionKey []ComputeDiskSpecSourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty" tf:"source_image_encryption_key,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey []ComputeDiskSpecSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComputeDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeDiskSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
