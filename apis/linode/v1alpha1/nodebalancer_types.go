package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type Nodebalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerSpec   `json:"spec,omitempty"`
	Status            NodebalancerStatus `json:"status,omitempty"`
}

type NodebalancerSpecTransfer struct {
	// The total transfer, in MB, used by this NodeBalancer this month
	// +optional
	In json.Number `json:"in,omitempty" tf:"in,omitempty"`
	// The total inbound transfer, in MB, used for this NodeBalancer this month
	// +optional
	Out json.Number `json:"out,omitempty" tf:"out,omitempty"`
	// The total outbound transfer, in MB, used for this NodeBalancer this month
	// +optional
	Total json.Number `json:"total,omitempty" tf:"total,omitempty"`
}

type NodebalancerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Throttle connections per second (0-20). Set to 0 (zero) to disable throttling.
	// +optional
	ClientConnThrottle int `json:"clientConnThrottle,omitempty" tf:"client_conn_throttle,omitempty"`
	// +optional
	Created string `json:"created,omitempty" tf:"created,omitempty"`
	// This NodeBalancer's hostname, ending with .nodebalancer.linode.com
	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	// The Public IPv4 Address of this NodeBalancer
	// +optional
	Ipv4 string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`
	// The Public IPv6 Address of this NodeBalancer
	// +optional
	Ipv6 string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
	// The label of the Linode NodeBalancer.
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// The region where this NodeBalancer will be deployed.
	Region string `json:"region" tf:"region"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Transfer map[string]NodebalancerSpecTransfer `json:"transfer,omitempty" tf:"transfer,omitempty"`
	// +optional
	Updated string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type NodebalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NodebalancerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
