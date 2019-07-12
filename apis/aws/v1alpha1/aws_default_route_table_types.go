package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultRouteTableSpec   `json:"spec,omitempty"`
	Status            AwsDefaultRouteTableStatus `json:"status,omitempty"`
}

type AwsDefaultRouteTableSpecRoute struct {
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	TransitGatewayId       string `json:"transit_gateway_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	CidrBlock              string `json:"cidr_block"`
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	GatewayId              string `json:"gateway_id"`
	InstanceId             string `json:"instance_id"`
	NatGatewayId           string `json:"nat_gateway_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
}

type AwsDefaultRouteTableSpec struct {
	DefaultRouteTableId string                     `json:"default_route_table_id"`
	VpcId               string                     `json:"vpc_id"`
	PropagatingVgws     []string                   `json:"propagating_vgws"`
	Route               []AwsDefaultRouteTableSpec `json:"route"`
	Tags                map[string]string          `json:"tags"`
	OwnerId             string                     `json:"owner_id"`
}

type AwsDefaultRouteTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultRouteTableList is a list of AwsDefaultRouteTables
type AwsDefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultRouteTable CRD objects
	Items []AwsDefaultRouteTable `json:"items,omitempty"`
}
