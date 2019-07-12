package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEc2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteSpec struct {
	DestinationCidrBlock       string `json:"destination_cidr_block"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}



type AwsEc2TransitGatewayRouteStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEc2TransitGatewayRouteList is a list of AwsEc2TransitGatewayRoutes
type AwsEc2TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRoute CRD objects
	Items []AwsEc2TransitGatewayRoute `json:"items,omitempty"`
}