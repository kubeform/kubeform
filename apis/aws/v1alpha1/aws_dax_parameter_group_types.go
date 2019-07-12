package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDaxParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsDaxParameterGroupStatus `json:"status,omitempty"`
}

type AwsDaxParameterGroupSpecParameters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsDaxParameterGroupSpec struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Parameters  []AwsDaxParameterGroupSpec `json:"parameters"`
}

type AwsDaxParameterGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxParameterGroupList is a list of AwsDaxParameterGroups
type AwsDaxParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDaxParameterGroup CRD objects
	Items []AwsDaxParameterGroup `json:"items,omitempty"`
}
