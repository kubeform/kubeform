package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEc2TransitGatewayRouteTablePropagation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteTablePropagationSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteTablePropagationStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteTablePropagationSpec struct {
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
	ResourceId                 string `json:"resource_id"`
	ResourceType               string `json:"resource_type"`
}



type AwsEc2TransitGatewayRouteTablePropagationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEc2TransitGatewayRouteTablePropagationList is a list of AwsEc2TransitGatewayRouteTablePropagations
type AwsEc2TransitGatewayRouteTablePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRouteTablePropagation CRD objects
	Items []AwsEc2TransitGatewayRouteTablePropagation `json:"items,omitempty"`
}