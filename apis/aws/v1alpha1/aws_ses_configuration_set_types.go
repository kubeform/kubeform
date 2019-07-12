package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesConfigurationSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesConfigurationSetSpec   `json:"spec,omitempty"`
	Status            AwsSesConfigurationSetStatus `json:"status,omitempty"`
}

type AwsSesConfigurationSetSpec struct {
	Name string `json:"name"`
}

type AwsSesConfigurationSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesConfigurationSetList is a list of AwsSesConfigurationSets
type AwsSesConfigurationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesConfigurationSet CRD objects
	Items []AwsSesConfigurationSet `json:"items,omitempty"`
}
