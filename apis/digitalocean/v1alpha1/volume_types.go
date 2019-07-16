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

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	InitialFilesystemLabel string `json:"initial_filesystem_label,omitempty"`
	// +optional
	InitialFilesystemType string `json:"initial_filesystem_type,omitempty"`
	Name                  string `json:"name"`
	Region                string `json:"region"`
	Size                  int    `json:"size"`
	// +optional
	SnapshotId string `json:"snapshot_id,omitempty"`
}

type VolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
