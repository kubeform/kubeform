package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanFirewallSpec   `json:"spec,omitempty"`
	Status            DigitaloceanFirewallStatus `json:"status,omitempty"`
}

type DigitaloceanFirewallSpecPendingChanges struct {
	Status    string `json:"status"`
	DropletId int    `json:"droplet_id"`
	Removing  bool   `json:"removing"`
}

type DigitaloceanFirewallSpecInboundRule struct {
	Protocol               string   `json:"protocol"`
	PortRange              string   `json:"port_range"`
	SourceAddresses        []string `json:"source_addresses"`
	SourceDropletIds       []int64  `json:"source_droplet_ids"`
	SourceLoadBalancerUids []string `json:"source_load_balancer_uids"`
	SourceTags             []string `json:"source_tags"`
}

type DigitaloceanFirewallSpecOutboundRule struct {
	DestinationDropletIds       []int64  `json:"destination_droplet_ids"`
	DestinationLoadBalancerUids []string `json:"destination_load_balancer_uids"`
	DestinationTags             []string `json:"destination_tags"`
	Protocol                    string   `json:"protocol"`
	PortRange                   string   `json:"port_range"`
	DestinationAddresses        []string `json:"destination_addresses"`
}

type DigitaloceanFirewallSpec struct {
	Status         string                     `json:"status"`
	CreatedAt      string                     `json:"created_at"`
	PendingChanges []DigitaloceanFirewallSpec `json:"pending_changes"`
	Tags           []string                   `json:"tags"`
	Name           string                     `json:"name"`
	DropletIds     []int64                    `json:"droplet_ids"`
	InboundRule    []DigitaloceanFirewallSpec `json:"inbound_rule"`
	OutboundRule   []DigitaloceanFirewallSpec `json:"outbound_rule"`
}



type DigitaloceanFirewallStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanFirewallList is a list of DigitaloceanFirewalls
type DigitaloceanFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanFirewall CRD objects
	Items []DigitaloceanFirewall `json:"items,omitempty"`
}