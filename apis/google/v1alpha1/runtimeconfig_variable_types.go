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

type RuntimeconfigVariable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuntimeconfigVariableSpec   `json:"spec,omitempty"`
	Status            RuntimeconfigVariableStatus `json:"status,omitempty"`
}

type RuntimeconfigVariableSpec struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
	// +optional
	Text string `json:"text,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type RuntimeconfigVariableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
