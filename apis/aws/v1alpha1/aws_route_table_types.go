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

type AwsRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRouteTableSpec   `json:"spec,omitempty"`
	Status            AwsRouteTableStatus `json:"status,omitempty"`
}

type AwsRouteTableSpecRoute struct {
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	GatewayId              string `json:"gateway_id"`
	InstanceId             string `json:"instance_id"`
	TransitGatewayId       string `json:"transit_gateway_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
	CidrBlock              string `json:"cidr_block"`
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	NatGatewayId           string `json:"nat_gateway_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
}

type AwsRouteTableSpec struct {
	VpcId           string              `json:"vpc_id"`
	Tags            map[string]string   `json:"tags"`
	PropagatingVgws []string            `json:"propagating_vgws"`
	Route           []AwsRouteTableSpec `json:"route"`
	OwnerId         string              `json:"owner_id"`
}

type AwsRouteTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRouteTableList is a list of AwsRouteTables
type AwsRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRouteTable CRD objects
	Items []AwsRouteTable `json:"items,omitempty"`
}
