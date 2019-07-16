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

type ComputeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeImageSpec   `json:"spec,omitempty"`
	Status            ComputeImageStatus `json:"status,omitempty"`
}

type ComputeImageSpecRawDisk struct {
	// +optional
	ContainerType string `json:"container_type,omitempty"`
	// +optional
	Sha1   string `json:"sha1,omitempty"`
	Source string `json:"source"`
}

type ComputeImageSpec struct {
	// +optional
	// Deprecated
	CreateTimeout int `json:"create_timeout,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Family string `json:"family,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RawDisk *[]ComputeImageSpec `json:"raw_disk,omitempty"`
	// +optional
	SourceDisk string `json:"source_disk,omitempty"`
}

type ComputeImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeImageList is a list of ComputeImages
type ComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeImage CRD objects
	Items []ComputeImage `json:"items,omitempty"`
}
