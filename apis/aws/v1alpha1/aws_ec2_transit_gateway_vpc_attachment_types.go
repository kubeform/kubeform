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

type AwsEc2TransitGatewayVpcAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayVpcAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayVpcAttachmentStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayVpcAttachmentSpec struct {
	Tags                                       map[string]string `json:"tags"`
	TransitGatewayDefaultRouteTableAssociation bool              `json:"transit_gateway_default_route_table_association"`
	TransitGatewayDefaultRouteTablePropagation bool              `json:"transit_gateway_default_route_table_propagation"`
	VpcId                                      string            `json:"vpc_id"`
	DnsSupport                                 string            `json:"dns_support"`
	SubnetIds                                  []string          `json:"subnet_ids"`
	TransitGatewayId                           string            `json:"transit_gateway_id"`
	VpcOwnerId                                 string            `json:"vpc_owner_id"`
	Ipv6Support                                string            `json:"ipv6_support"`
}



type AwsEc2TransitGatewayVpcAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayVpcAttachmentList is a list of AwsEc2TransitGatewayVpcAttachments
type AwsEc2TransitGatewayVpcAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayVpcAttachment CRD objects
	Items []AwsEc2TransitGatewayVpcAttachment `json:"items,omitempty"`
}