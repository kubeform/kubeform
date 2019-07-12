package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEc2TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteTableSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteTableSpec struct {
	TransitGatewayId             string            `json:"transit_gateway_id"`
	DefaultAssociationRouteTable bool              `json:"default_association_route_table"`
	DefaultPropagationRouteTable bool              `json:"default_propagation_route_table"`
	Tags                         map[string]string `json:"tags"`
}

type AwsEc2TransitGatewayRouteTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEc2TransitGatewayRouteTableList is a list of AwsEc2TransitGatewayRouteTables
type AwsEc2TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRouteTable CRD objects
	Items []AwsEc2TransitGatewayRouteTable `json:"items,omitempty"`
}
