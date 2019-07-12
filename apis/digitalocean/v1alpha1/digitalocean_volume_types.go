package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanVolumeSpec   `json:"spec,omitempty"`
	Status            DigitaloceanVolumeStatus `json:"status,omitempty"`
}

type DigitaloceanVolumeSpec struct {
	InitialFilesystemLabel string  `json:"initial_filesystem_label"`
	FilesystemLabel        string  `json:"filesystem_label"`
	Name                   string  `json:"name"`
	Urn                    string  `json:"urn"`
	Size                   int     `json:"size"`
	Description            string  `json:"description"`
	FilesystemType         string  `json:"filesystem_type"`
	Region                 string  `json:"region"`
	SnapshotId             string  `json:"snapshot_id"`
	InitialFilesystemType  string  `json:"initial_filesystem_type"`
	DropletIds             []int64 `json:"droplet_ids"`
}



type DigitaloceanVolumeStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanVolumeList is a list of DigitaloceanVolumes
type DigitaloceanVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanVolume CRD objects
	Items []DigitaloceanVolume `json:"items,omitempty"`
}