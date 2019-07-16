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

type ComputeAttachedDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeAttachedDiskSpec   `json:"spec,omitempty"`
	Status            ComputeAttachedDiskStatus `json:"status,omitempty"`
}

type ComputeAttachedDiskSpec struct {
	Disk     string `json:"disk"`
	Instance string `json:"instance"`
	// +optional
	Mode string `json:"mode,omitempty"`
}

type ComputeAttachedDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeAttachedDiskList is a list of ComputeAttachedDisks
type ComputeAttachedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeAttachedDisk CRD objects
	Items []ComputeAttachedDisk `json:"items,omitempty"`
}
