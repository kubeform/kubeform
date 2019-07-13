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

type LinodeVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeVolumeSpec   `json:"spec,omitempty"`
	Status            LinodeVolumeStatus `json:"status,omitempty"`
}

type LinodeVolumeSpec struct {
	Tags           []string `json:"tags"`
	Label          string   `json:"label"`
	Status         string   `json:"status"`
	Region         string   `json:"region"`
	Size           int      `json:"size"`
	LinodeId       int      `json:"linode_id"`
	FilesystemPath string   `json:"filesystem_path"`
}



type LinodeVolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeVolumeList is a list of LinodeVolumes
type LinodeVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeVolume CRD objects
	Items []LinodeVolume `json:"items,omitempty"`
}