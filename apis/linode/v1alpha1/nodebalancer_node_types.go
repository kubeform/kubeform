package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NodebalancerNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerNodeSpec   `json:"spec,omitempty"`
	Status            NodebalancerNodeStatus `json:"status,omitempty"`
}

type NodebalancerNodeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.
	Address string `json:"address" tf:"address"`
	// The ID of the NodeBalancerConfig to access.
	ConfigID int `json:"configID" tf:"config_id"`
	// The label for this node. This is for display purposes only.
	Label string `json:"label" tf:"label"`
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it.
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
	// The ID of the NodeBalancer to access.
	NodebalancerID int `json:"nodebalancerID" tf:"nodebalancer_id"`
	// The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN)
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}

type NodebalancerNodeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NodebalancerNodeSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodebalancerNodeList is a list of NodebalancerNodes
type NodebalancerNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NodebalancerNode CRD objects
	Items []NodebalancerNode `json:"items,omitempty"`
}
