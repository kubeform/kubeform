package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DropletSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DropletSnapshotSpec   `json:"spec,omitempty"`
	Status            DropletSnapshotStatus `json:"status,omitempty"`
}

type DropletSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	DropletID string `json:"dropletID" tf:"droplet_id"`
	// +optional
	MinDiskSize int    `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Regions []string `json:"regions,omitempty" tf:"regions,omitempty"`
	// +optional
	Size json.Number `json:"size,omitempty" tf:"size,omitempty"`
}

type DropletSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DropletSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DropletSnapshotList is a list of DropletSnapshots
type DropletSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DropletSnapshot CRD objects
	Items []DropletSnapshot `json:"items,omitempty"`
}
