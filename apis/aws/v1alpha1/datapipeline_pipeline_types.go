package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DatapipelinePipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatapipelinePipelineSpec   `json:"spec,omitempty"`
	Status            DatapipelinePipelineStatus `json:"status,omitempty"`
}

type DatapipelinePipelineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DatapipelinePipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DatapipelinePipelineSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatapipelinePipelineList is a list of DatapipelinePipelines
type DatapipelinePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatapipelinePipeline CRD objects
	Items []DatapipelinePipeline `json:"items,omitempty"`
}
