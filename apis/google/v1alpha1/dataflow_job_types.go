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

type DataflowJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataflowJobSpec   `json:"spec,omitempty"`
	Status            DataflowJobStatus `json:"status,omitempty"`
}

type DataflowJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	MaxWorkers int    `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	OnDelete string `json:"onDelete,omitempty" tf:"on_delete,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region          string `json:"region,omitempty" tf:"region,omitempty"`
	TempGcsLocation string `json:"tempGcsLocation" tf:"temp_gcs_location"`
	TemplateGcsPath string `json:"templateGcsPath" tf:"template_gcs_path"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DataflowJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataflowJobList is a list of DataflowJobs
type DataflowJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataflowJob CRD objects
	Items []DataflowJob `json:"items,omitempty"`
}
