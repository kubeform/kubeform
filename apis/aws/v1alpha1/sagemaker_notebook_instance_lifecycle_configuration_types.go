package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OnCreate string `json:"onCreate,omitempty" tf:"on_create,omitempty"`
	// +optional
	OnStart string `json:"onStart,omitempty" tf:"on_start,omitempty"`
}

type SagemakerNotebookInstanceLifecycleConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SagemakerNotebookInstanceLifecycleConfigurationSpec `json:"state,omitempty"`
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
