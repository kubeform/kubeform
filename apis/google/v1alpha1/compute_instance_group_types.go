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

type ComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceGroupStatus `json:"status,omitempty"`
}

type ComputeInstanceGroupSpecNamedPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type ComputeInstanceGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	NamedPort *[]ComputeInstanceGroupSpec `json:"named_port,omitempty"`
}

type ComputeInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceGroupList is a list of ComputeInstanceGroups
type ComputeInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceGroup CRD objects
	Items []ComputeInstanceGroup `json:"items,omitempty"`
}
