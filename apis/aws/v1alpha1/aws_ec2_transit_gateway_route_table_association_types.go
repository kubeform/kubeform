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

type AwsEc2TransitGatewayRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteTableAssociationStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteTableAssociationSpec struct {
	ResourceId                 string `json:"resource_id"`
	ResourceType               string `json:"resource_type"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}

type AwsEc2TransitGatewayRouteTableAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayRouteTableAssociationList is a list of AwsEc2TransitGatewayRouteTableAssociations
type AwsEc2TransitGatewayRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRouteTableAssociation CRD objects
	Items []AwsEc2TransitGatewayRouteTableAssociation `json:"items,omitempty"`
}
