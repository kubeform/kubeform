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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The full filesystem path for the Volume based on the Volume's label. Path is /dev/disk/by-id/scsi-0Linode_Volume_ + Volume label.
	// +optional
	FilesystemPath string `json:"filesystemPath,omitempty" tf:"filesystem_path,omitempty"`
	// The label of the Linode Volume.
	Label string `json:"label" tf:"label"`
	// The Linode ID where the Volume should be attached.
	// +optional
	LinodeID int `json:"linodeID,omitempty" tf:"linode_id,omitempty"`
	// The region where this volume will be deployed.
	Region string `json:"region" tf:"region"`
	// Size of the Volume in GB
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// The status of the volume, indicating the current readiness state.
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
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