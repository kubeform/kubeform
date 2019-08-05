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

type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultRouteTableSpec   `json:"spec,omitempty"`
	Status            DefaultRouteTableStatus `json:"status,omitempty"`
}

type DefaultRouteTableSpecRoute struct {
	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	// +optional
	EgressOnlyGatewayID string `json:"egressOnlyGatewayID,omitempty" tf:"egress_only_gateway_id,omitempty"`
	// +optional
	GatewayID string `json:"gatewayID,omitempty" tf:"gateway_id,omitempty"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	// +optional
	NatGatewayID string `json:"natGatewayID,omitempty" tf:"nat_gateway_id,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	VpcPeeringConnectionID string `json:"vpcPeeringConnectionID,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type DefaultRouteTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DefaultRouteTableID string `json:"defaultRouteTableID" tf:"default_route_table_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PropagatingVgws []string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Route []DefaultRouteTableSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultRouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultRouteTableSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultRouteTableList is a list of DefaultRouteTables
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultRouteTable CRD objects
	Items []DefaultRouteTable `json:"items,omitempty"`
}
