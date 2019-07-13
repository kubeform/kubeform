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

type GoogleComputeAttachedDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeAttachedDiskSpec   `json:"spec,omitempty"`
	Status            GoogleComputeAttachedDiskStatus `json:"status,omitempty"`
}

type GoogleComputeAttachedDiskSpec struct {
	Disk       string `json:"disk"`
	Instance   string `json:"instance"`
	Project    string `json:"project"`
	Zone       string `json:"zone"`
	DeviceName string `json:"device_name"`
	Mode       string `json:"mode"`
}



type GoogleComputeAttachedDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeAttachedDiskList is a list of GoogleComputeAttachedDisks
type GoogleComputeAttachedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeAttachedDisk CRD objects
	Items []GoogleComputeAttachedDisk `json:"items,omitempty"`
}