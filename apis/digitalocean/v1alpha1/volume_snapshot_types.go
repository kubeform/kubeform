package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VolumeSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSnapshotSpec   `json:"spec,omitempty"`
	Status            VolumeSnapshotStatus `json:"status,omitempty"`
}

type VolumeSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	MinDiskSize int    `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Regions []string `json:"regions,omitempty" tf:"regions,omitempty"`
	// +optional
	Size     json.Number `json:"size,omitempty" tf:"size,omitempty"`
	VolumeID string      `json:"volumeID" tf:"volume_id"`
}



type VolumeSnapshotStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *VolumeSnapshotSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeSnapshotList is a list of VolumeSnapshots
type VolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeSnapshot CRD objects
	Items []VolumeSnapshot `json:"items,omitempty"`
}