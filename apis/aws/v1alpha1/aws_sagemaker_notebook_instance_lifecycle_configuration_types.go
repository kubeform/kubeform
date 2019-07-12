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

type AwsSagemakerNotebookInstanceLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSagemakerNotebookInstanceLifecycleConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsSagemakerNotebookInstanceLifecycleConfigurationStatus `json:"status,omitempty"`
}

type AwsSagemakerNotebookInstanceLifecycleConfigurationSpec struct {
	Arn      string `json:"arn"`
	Name     string `json:"name"`
	OnCreate string `json:"on_create"`
	OnStart  string `json:"on_start"`
}

type AwsSagemakerNotebookInstanceLifecycleConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSagemakerNotebookInstanceLifecycleConfigurationList is a list of AwsSagemakerNotebookInstanceLifecycleConfigurations
type AwsSagemakerNotebookInstanceLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSagemakerNotebookInstanceLifecycleConfiguration CRD objects
	Items []AwsSagemakerNotebookInstanceLifecycleConfiguration `json:"items,omitempty"`
}
