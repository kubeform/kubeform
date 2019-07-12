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

type GoogleComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInstanceGroupStatus `json:"status,omitempty"`
}

type GoogleComputeInstanceGroupSpecNamedPort struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}

type GoogleComputeInstanceGroupSpec struct {
	Name        string                           `json:"name"`
	Zone        string                           `json:"zone"`
	Description string                           `json:"description"`
	NamedPort   []GoogleComputeInstanceGroupSpec `json:"named_port"`
	Network     string                           `json:"network"`
	SelfLink    string                           `json:"self_link"`
	Instances   []string                         `json:"instances"`
	Project     string                           `json:"project"`
	Size        int                              `json:"size"`
}

type GoogleComputeInstanceGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInstanceGroupList is a list of GoogleComputeInstanceGroups
type GoogleComputeInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInstanceGroup CRD objects
	Items []GoogleComputeInstanceGroup `json:"items,omitempty"`
}
