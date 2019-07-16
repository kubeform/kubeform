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

type EbsVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsVolumeSpec   `json:"spec,omitempty"`
	Status            EbsVolumeStatus `json:"status,omitempty"`
}

type EbsVolumeSpec struct {
	AvailabilityZone string `json:"availability_zone"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type EbsVolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EbsVolumeList is a list of EbsVolumes
type EbsVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EbsVolume CRD objects
	Items []EbsVolume `json:"items,omitempty"`
}
