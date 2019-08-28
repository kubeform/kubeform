package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeSharedVpcHostProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSharedVpcHostProjectSpec   `json:"spec,omitempty"`
	Status            ComputeSharedVpcHostProjectStatus `json:"status,omitempty"`
}

type ComputeSharedVpcHostProjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Project string `json:"project" tf:"project"`
}



type ComputeSharedVpcHostProjectStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ComputeSharedVpcHostProjectSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSharedVpcHostProjectList is a list of ComputeSharedVpcHostProjects
type ComputeSharedVpcHostProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSharedVpcHostProject CRD objects
	Items []ComputeSharedVpcHostProject `json:"items,omitempty"`
}