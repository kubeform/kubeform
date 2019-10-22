package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeRegionDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRegionDiskSpec   `json:"spec,omitempty"`
	Status            ComputeRegionDiskStatus `json:"status,omitempty"`
}

type ComputeRegionDiskSpecDiskEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeRegionDiskSpecSourceSnapshotEncryptionKey struct {
	// +optional
	RawKey string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
	// +optional
	Sha256 string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type ComputeRegionDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ComputeRegionDiskSpecDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
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
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	ReplicaZones []string `json:"replicaZones" tf:"replica_zones"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Snapshot string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceSnapshotEncryptionKey []ComputeRegionDiskSpecSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`
	// +optional
	SourceSnapshotID string `json:"sourceSnapshotID,omitempty" tf:"source_snapshot_id,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Users []string `json:"users,omitempty" tf:"users,omitempty"`
}

type ComputeRegionDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRegionDiskSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
