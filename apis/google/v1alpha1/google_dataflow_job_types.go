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

type GoogleDataflowJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDataflowJobSpec   `json:"spec,omitempty"`
	Status            GoogleDataflowJobStatus `json:"status,omitempty"`
}

type GoogleDataflowJobSpec struct {
	Zone            string            `json:"zone"`
	Region          string            `json:"region"`
	MaxWorkers      int               `json:"max_workers"`
	OnDelete        string            `json:"on_delete"`
	State           string            `json:"state"`
	Name            string            `json:"name"`
	TempGcsLocation string            `json:"temp_gcs_location"`
	Parameters      map[string]string `json:"parameters"`
	Project         string            `json:"project"`
	TemplateGcsPath string            `json:"template_gcs_path"`
}



type GoogleDataflowJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleDataflowJobList is a list of GoogleDataflowJobs
type GoogleDataflowJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDataflowJob CRD objects
	Items []GoogleDataflowJob `json:"items,omitempty"`
}