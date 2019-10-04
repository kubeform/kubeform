package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	PortRange string `json:"portRange,omitempty" tf:"port_range,omitempty"`
	Protocol  string `json:"protocol" tf:"protocol"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`
	// +optional
	SourceDropletIDS []int64 `json:"sourceDropletIDS,omitempty" tf:"source_droplet_ids,omitempty"`
	// +optional
	SourceLoadBalancerUids []string `json:"sourceLoadBalancerUids,omitempty" tf:"source_load_balancer_uids,omitempty"`
	// +optional
	SourceTags []string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`
}

type FirewallSpecOutboundRule struct {
	// +optional
	DestinationAddresses []string `json:"destinationAddresses,omitempty" tf:"destination_addresses,omitempty"`
	// +optional
	DestinationDropletIDS []int64 `json:"destinationDropletIDS,omitempty" tf:"destination_droplet_ids,omitempty"`
	// +optional
	DestinationLoadBalancerUids []string `json:"destinationLoadBalancerUids,omitempty" tf:"destination_load_balancer_uids,omitempty"`
	// +optional
	DestinationTags []string `json:"destinationTags,omitempty" tf:"destination_tags,omitempty"`
	// +optional
	PortRange string `json:"portRange,omitempty" tf:"port_range,omitempty"`
	Protocol  string `json:"protocol" tf:"protocol"`
}

type FirewallSpecPendingChanges struct {
	// +optional
	DropletID int `json:"dropletID,omitempty" tf:"droplet_id,omitempty"`
	// +optional
	Removing bool `json:"removing,omitempty" tf:"removing,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
}

type FirewallSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids,omitempty"`
	// +optional
	InboundRule []FirewallSpecInboundRule `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`
	Name        string                    `json:"name" tf:"name"`
	// +optional
	OutboundRule []FirewallSpecOutboundRule `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`
	// +optional
	PendingChanges []FirewallSpecPendingChanges `json:"pendingChanges,omitempty" tf:"pending_changes,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FirewallSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
