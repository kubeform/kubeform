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

type Ec2TransitGatewayRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayRouteTableAssociationStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayRouteTableAssociationSpec struct {
	TransitGatewayAttachmentID string                    `json:"transitGatewayAttachmentID" tf:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableID string                    `json:"transitGatewayRouteTableID" tf:"transit_gateway_route_table_id"`
	ProviderRef                core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type Ec2TransitGatewayRouteTableAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableAssociationList is a list of Ec2TransitGatewayRouteTableAssociations
type Ec2TransitGatewayRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGatewayRouteTableAssociation CRD objects
	Items []Ec2TransitGatewayRouteTableAssociation `json:"items,omitempty"`
}
