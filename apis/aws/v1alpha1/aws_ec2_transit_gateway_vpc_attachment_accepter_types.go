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

type AwsEc2TransitGatewayVpcAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayVpcAttachmentAccepterSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayVpcAttachmentAccepterStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayVpcAttachmentAccepterSpec struct {
	VpcId                                      string            `json:"vpc_id"`
	VpcOwnerId                                 string            `json:"vpc_owner_id"`
	Tags                                       map[string]string `json:"tags"`
	TransitGatewayDefaultRouteTablePropagation bool              `json:"transit_gateway_default_route_table_propagation"`
	SubnetIds                                  []string          `json:"subnet_ids"`
	TransitGatewayAttachmentId                 string            `json:"transit_gateway_attachment_id"`
	TransitGatewayDefaultRouteTableAssociation bool              `json:"transit_gateway_default_route_table_association"`
	TransitGatewayId                           string            `json:"transit_gateway_id"`
	DnsSupport                                 string            `json:"dns_support"`
	Ipv6Support                                string            `json:"ipv6_support"`
}



type AwsEc2TransitGatewayVpcAttachmentAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayVpcAttachmentAccepterList is a list of AwsEc2TransitGatewayVpcAttachmentAccepters
type AwsEc2TransitGatewayVpcAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayVpcAttachmentAccepter CRD objects
	Items []AwsEc2TransitGatewayVpcAttachmentAccepter `json:"items,omitempty"`
}