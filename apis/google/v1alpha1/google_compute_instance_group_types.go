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
	Name string `json:"name"`
	Port int    `json:"port"`
}

type GoogleComputeInstanceGroupSpec struct {
	Zone        string                           `json:"zone"`
	Description string                           `json:"description"`
	Instances   []string                         `json:"instances"`
	NamedPort   []GoogleComputeInstanceGroupSpec `json:"named_port"`
	Project     string                           `json:"project"`
	Name        string                           `json:"name"`
	SelfLink    string                           `json:"self_link"`
	Size        int                              `json:"size"`
	Network     string                           `json:"network"`
}



type GoogleComputeInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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