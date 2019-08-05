package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Nodebalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerSpec   `json:"spec,omitempty"`
	Status            NodebalancerStatus `json:"status,omitempty"`
}

type NodebalancerSpecTransfer struct {
	In    json.Number `json:"in" tf:"in"`
	Out   json.Number `json:"out" tf:"out"`
	Total json.Number `json:"total" tf:"total"`
}

type NodebalancerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ClientConnThrottle int    `json:"clientConnThrottle,omitempty" tf:"client_conn_throttle,omitempty"`
	Created            string `json:"created" tf:"created"`
	Hostname           string `json:"hostname" tf:"hostname"`
	Ipv4               string `json:"ipv4" tf:"ipv4"`
	Ipv6               string `json:"ipv6" tf:"ipv6"`
	// +optional
	Label  string `json:"label,omitempty" tf:"label,omitempty"`
	Region string `json:"region" tf:"region"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags     []string                            `json:"tags,omitempty" tf:"tags,omitempty"`
	Transfer map[string]NodebalancerSpecTransfer `json:"transfer" tf:"transfer"`
	Updated  string                              `json:"updated" tf:"updated"`
}

type NodebalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NodebalancerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodebalancerList is a list of Nodebalancers
type NodebalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Nodebalancer CRD objects
	Items []Nodebalancer `json:"items,omitempty"`
}
