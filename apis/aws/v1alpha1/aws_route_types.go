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

type AwsRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRouteSpec   `json:"spec,omitempty"`
	Status            AwsRouteStatus `json:"status,omitempty"`
}

type AwsRouteSpec struct {
	DestinationIpv6CidrBlock string `json:"destination_ipv6_cidr_block"`
	NatGatewayId             string `json:"nat_gateway_id"`
	InstanceId               string `json:"instance_id"`
	Origin                   string `json:"origin"`
	State                    string `json:"state"`
	DestinationCidrBlock     string `json:"destination_cidr_block"`
	NetworkInterfaceId       string `json:"network_interface_id"`
	RouteTableId             string `json:"route_table_id"`
	EgressOnlyGatewayId      string `json:"egress_only_gateway_id"`
	GatewayId                string `json:"gateway_id"`
	InstanceOwnerId          string `json:"instance_owner_id"`
	TransitGatewayId         string `json:"transit_gateway_id"`
	VpcPeeringConnectionId   string `json:"vpc_peering_connection_id"`
	DestinationPrefixListId  string `json:"destination_prefix_list_id"`
}



type AwsRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRouteList is a list of AwsRoutes
type AwsRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute CRD objects
	Items []AwsRoute `json:"items,omitempty"`
}