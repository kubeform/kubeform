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

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids,omitempty"`
	// +optional
	FilesystemLabel string `json:"filesystemLabel,omitempty" tf:"filesystem_label,omitempty"`
	// +optional
	// Deprecated
	FilesystemType string `json:"filesystemType,omitempty" tf:"filesystem_type,omitempty"`
	// +optional
	InitialFilesystemLabel string `json:"initialFilesystemLabel,omitempty" tf:"initial_filesystem_label,omitempty"`
	// +optional
	InitialFilesystemType string `json:"initialFilesystemType,omitempty" tf:"initial_filesystem_type,omitempty"`
	Name                  string `json:"name" tf:"name"`
	Region                string `json:"region" tf:"region"`
	Size                  int    `json:"size" tf:"size"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	Urn string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type VolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VolumeSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeList is a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []Volume `json:"items,omitempty"`
}
