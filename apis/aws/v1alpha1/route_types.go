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

type Route struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec,omitempty"`
	Status            RouteStatus `json:"status,omitempty"`
}

type RouteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DestinationCIDRBlock string `json:"destinationCIDRBlock,omitempty" tf:"destination_cidr_block,omitempty"`
	// +optional
	DestinationIpv6CIDRBlock string `json:"destinationIpv6CIDRBlock,omitempty" tf:"destination_ipv6_cidr_block,omitempty"`
	// +optional
	DestinationPrefixListID string `json:"destinationPrefixListID,omitempty" tf:"destination_prefix_list_id,omitempty"`
	// +optional
	EgressOnlyGatewayID string `json:"egressOnlyGatewayID,omitempty" tf:"egress_only_gateway_id,omitempty"`
	// +optional
	GatewayID string `json:"gatewayID,omitempty" tf:"gateway_id,omitempty"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// +optional
	InstanceOwnerID string `json:"instanceOwnerID,omitempty" tf:"instance_owner_id,omitempty"`
	// +optional
	NatGatewayID string `json:"natGatewayID,omitempty" tf:"nat_gateway_id,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	Origin       string `json:"origin,omitempty" tf:"origin,omitempty"`
	RouteTableID string `json:"routeTableID" tf:"route_table_id"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	VpcPeeringConnectionID string `json:"vpcPeeringConnectionID,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type RouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RouteSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RouteList is a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route CRD objects
	Items []Route `json:"items,omitempty"`
}
