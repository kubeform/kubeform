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

type DigitaloceanVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanVolumeSpec   `json:"spec,omitempty"`
	Status            DigitaloceanVolumeStatus `json:"status,omitempty"`
}

type DigitaloceanVolumeSpec struct {
	FilesystemLabel        string  `json:"filesystem_label"`
	Region                 string  `json:"region"`
	Name                   string  `json:"name"`
	Urn                    string  `json:"urn"`
	Description            string  `json:"description"`
	InitialFilesystemType  string  `json:"initial_filesystem_type"`
	InitialFilesystemLabel string  `json:"initial_filesystem_label"`
	FilesystemType         string  `json:"filesystem_type"`
	Size                   int     `json:"size"`
	SnapshotId             string  `json:"snapshot_id"`
	DropletIds             []int64 `json:"droplet_ids"`
}

type DigitaloceanVolumeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanVolumeList is a list of DigitaloceanVolumes
type DigitaloceanVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanVolume CRD objects
	Items []DigitaloceanVolume `json:"items,omitempty"`
}
