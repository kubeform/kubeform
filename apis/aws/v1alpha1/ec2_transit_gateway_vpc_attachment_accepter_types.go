package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
	// +optional
	Tags                       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TransitGatewayAttachmentID string            `json:"transitGatewayAttachmentID" tf:"transit_gateway_attachment_id"`
	// +optional
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`
	// +optional
	TransitGatewayDefaultRouteTablePropagation bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
}

type Ec2TransitGatewayVpcAttachmentAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
