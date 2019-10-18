package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Ec2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayRouteStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayRouteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DestinationCIDRBlock       string `json:"destinationCIDRBlock" tf:"destination_cidr_block"`
	TransitGatewayAttachmentID string `json:"transitGatewayAttachmentID" tf:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableID string `json:"transitGatewayRouteTableID" tf:"transit_gateway_route_table_id"`
}

type Ec2TransitGatewayRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Ec2TransitGatewayRouteSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteList is a list of Ec2TransitGatewayRoutes
type Ec2TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGatewayRoute CRD objects
	Items []Ec2TransitGatewayRoute `json:"items,omitempty"`
}
