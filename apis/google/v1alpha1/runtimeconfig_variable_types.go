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

type RuntimeconfigVariable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuntimeconfigVariableSpec   `json:"spec,omitempty"`
	Status            RuntimeconfigVariableStatus `json:"status,omitempty"`
}

type RuntimeconfigVariableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name   string `json:"name" tf:"name"`
	Parent string `json:"parent" tf:"parent"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Text string `json:"text,omitempty" tf:"text,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuntimeconfigVariableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RuntimeconfigVariableSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RuntimeconfigVariableList is a list of RuntimeconfigVariables
type RuntimeconfigVariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RuntimeconfigVariable CRD objects
	Items []RuntimeconfigVariable `json:"items,omitempty"`
}
