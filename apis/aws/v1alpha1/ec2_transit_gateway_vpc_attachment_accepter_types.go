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

type Ec2TransitGatewayVpcAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayVpcAttachmentAccepterSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayVpcAttachmentAccepterStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayVpcAttachmentAccepterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DnsSupport string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`
	// +optional
	Ipv6Support string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	Tags                       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TransitGatewayAttachmentID string            `json:"transitGatewayAttachmentID" tf:"transit_gateway_attachment_id"`
	// +optional
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`
	// +optional
	TransitGatewayDefaultRouteTablePropagation bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	// +optional
	VpcOwnerID string `json:"vpcOwnerID,omitempty" tf:"vpc_owner_id,omitempty"`
}

type Ec2TransitGatewayVpcAttachmentAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Ec2TransitGatewayVpcAttachmentAccepterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachmentAccepterList is a list of Ec2TransitGatewayVpcAttachmentAccepters
type Ec2TransitGatewayVpcAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGatewayVpcAttachmentAccepter CRD objects
	Items []Ec2TransitGatewayVpcAttachmentAccepter `json:"items,omitempty"`
}
