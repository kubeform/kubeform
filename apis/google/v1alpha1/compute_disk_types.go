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
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeDiskSpecSourceImageEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeDiskSpecDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	// +optional
	DiskEncryptionKeyRaw string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	// Deprecated
	DiskEncryptionKeySha256 string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LastAttachTimestamp string `json:"lastAttachTimestamp,omitempty" tf:"last_attach_timestamp,omitempty"`
	// +optional
	LastDetachTimestamp string `json:"lastDetachTimestamp,omitempty" tf:"last_detach_timestamp,omitempty"`
	Name                string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Snapshot string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceImageEncryptionKey []ComputeDiskSpecSourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty" tf:"source_image_encryption_key,omitempty"`
	// +optional
	SourceImageID string `json:"sourceImageID,omitempty" tf:"source_image_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey []ComputeDiskSpecSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`
	// +optional
	SourceSnapshotID string `json:"sourceSnapshotID,omitempty" tf:"source_snapshot_id,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Users []string `json:"users,omitempty" tf:"users,omitempty"`
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
