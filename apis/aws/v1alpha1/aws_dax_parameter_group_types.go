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
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDaxParameterGroupList is a list of AwsDaxParameterGroups
type AwsDaxParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDaxParameterGroup CRD objects
	Items []AwsDaxParameterGroup `json:"items,omitempty"`
}