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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecInboundRule struct {
	// +optional
	PortRange string `json:"port_range,omitempty"`
	Protocol  string `json:"protocol"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceAddresses []string `json:"source_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceDropletIds []int64 `json:"source_droplet_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceLoadBalancerUids []string `json:"source_load_balancer_uids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceTags []string `json:"source_tags,omitempty"`
}

type FirewallSpecOutboundRule struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationAddresses []string `json:"destination_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationDropletIds []int64 `json:"destination_droplet_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationLoadBalancerUids []string `json:"destination_load_balancer_uids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationTags []string `json:"destination_tags,omitempty"`
	// +optional
	PortRange string `json:"port_range,omitempty"`
	Protocol  string `json:"protocol"`
}

type FirewallSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DropletIds []int64 `json:"droplet_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InboundRule *[]FirewallSpec `json:"inbound_rule,omitempty"`
	Name        string          `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OutboundRule *[]FirewallSpec `json:"outbound_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type FirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
