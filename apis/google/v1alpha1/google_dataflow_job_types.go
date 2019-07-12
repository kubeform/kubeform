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
	TemplateGcsPath string            `json:"template_gcs_path"`
	Zone            string            `json:"zone"`
	Region          string            `json:"region"`
	Project         string            `json:"project"`
	Name            string            `json:"name"`
	TempGcsLocation string            `json:"temp_gcs_location"`
	MaxWorkers      int               `json:"max_workers"`
	Parameters      map[string]string `json:"parameters"`
	OnDelete        string            `json:"on_delete"`
	State           string            `json:"state"`
}

type GoogleDataflowJobStatus struct {
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
