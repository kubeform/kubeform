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

type DataflowJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataflowJobSpec   `json:"spec,omitempty"`
	Status            DataflowJobStatus `json:"status,omitempty"`
}

type DataflowJobSpec struct {
	// +optional
	MaxWorkers int    `json:"max_workers,omitempty"`
	Name       string `json:"name"`
	// +optional
	OnDelete string `json:"on_delete,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	Project string `json:"project,omitempty"`
	// +optional
	Region          string `json:"region,omitempty"`
	TempGcsLocation string `json:"temp_gcs_location"`
	TemplateGcsPath string `json:"template_gcs_path"`
	// +optional
	Zone string `json:"zone,omitempty"`
}

type DataflowJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
