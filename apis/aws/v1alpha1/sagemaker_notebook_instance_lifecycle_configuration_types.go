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

type SagemakerNotebookInstanceLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerNotebookInstanceLifecycleConfigurationSpec   `json:"spec,omitempty"`
	Status            SagemakerNotebookInstanceLifecycleConfigurationStatus `json:"status,omitempty"`
}

type SagemakerNotebookInstanceLifecycleConfigurationSpec struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	OnCreate string `json:"on_create,omitempty"`
	// +optional
	OnStart string `json:"on_start,omitempty"`
}

type SagemakerNotebookInstanceLifecycleConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerNotebookInstanceLifecycleConfigurationList is a list of SagemakerNotebookInstanceLifecycleConfigurations
type SagemakerNotebookInstanceLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerNotebookInstanceLifecycleConfiguration CRD objects
	Items []SagemakerNotebookInstanceLifecycleConfiguration `json:"items,omitempty"`
}
