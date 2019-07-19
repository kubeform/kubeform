package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ContainerType string `json:"containerType,omitempty" tf:"container_type,omitempty"`
	// +optional
	Sha1   string `json:"sha1,omitempty" tf:"sha1,omitempty"`
	Source string `json:"source" tf:"source"`
}

type ComputeImageSpec struct {
	// +optional
	// Deprecated
	CreateTimeout int `json:"createTimeout,omitempty" tf:"create_timeout,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Family string `json:"family,omitempty" tf:"family,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Licenses []string `json:"licenses,omitempty" tf:"licenses,omitempty"`
	Name     string   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RawDisk []ComputeImageSpecRawDisk `json:"rawDisk,omitempty" tf:"raw_disk,omitempty"`
	// +optional
	SourceDisk  string                    `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
