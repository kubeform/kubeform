package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceGroupStatus `json:"status,omitempty"`
}

type ComputeInstanceGroupSpecNamedPort struct {
	Name string `json:"name" tf:"name"`
	Port int    `json:"port" tf:"port"`
}

type ComputeInstanceGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances,omitempty"`
	Name      string   `json:"name" tf:"name"`
	// +optional
	NamedPort []ComputeInstanceGroupSpecNamedPort `json:"namedPort,omitempty" tf:"named_port,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComputeInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeInstanceGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
