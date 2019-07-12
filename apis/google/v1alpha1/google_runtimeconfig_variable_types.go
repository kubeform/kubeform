package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleRuntimeconfigVariable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleRuntimeconfigVariableSpec   `json:"spec,omitempty"`
	Status            GoogleRuntimeconfigVariableStatus `json:"status,omitempty"`
}

type GoogleRuntimeconfigVariableSpec struct {
	Text       string `json:"text"`
	UpdateTime string `json:"update_time"`
	Name       string `json:"name"`
	Parent     string `json:"parent"`
	Project    string `json:"project"`
	Value      string `json:"value"`
}



type GoogleRuntimeconfigVariableStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleRuntimeconfigVariableList is a list of GoogleRuntimeconfigVariables
type GoogleRuntimeconfigVariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleRuntimeconfigVariable CRD objects
	Items []GoogleRuntimeconfigVariable `json:"items,omitempty"`
}