package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Route struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec,omitempty"`
	Status            RouteStatus `json:"status,omitempty"`
}

type RouteSpec struct {
	// +optional
	DestinationCIDRBlock string `json:"destinationCIDRBlock,omitempty" tf:"destination_cidr_block,omitempty"`
	// +optional
	DestinationIpv6CIDRBlock string `json:"destinationIpv6CIDRBlock,omitempty" tf:"destination_ipv6_cidr_block,omitempty"`
	// +optional
	EgressOnlyGatewayID string `json:"egressOnlyGatewayID,omitempty" tf:"egress_only_gateway_id,omitempty"`
	// +optional
	GatewayID string `json:"gatewayID,omitempty" tf:"gateway_id,omitempty"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// +optional
	NatGatewayID string `json:"natGatewayID,omitempty" tf:"nat_gateway_id,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	RouteTableID       string `json:"routeTableID" tf:"route_table_id"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	VpcPeeringConnectionID string                    `json:"vpcPeeringConnectionID,omitempty" tf:"vpc_peering_connection_id,omitempty"`
	ProviderRef            core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
