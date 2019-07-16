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

type Ec2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayRouteStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayRouteSpec struct {
	// +optional
	Blackhole            bool   `json:"blackhole,omitempty"`
	DestinationCidrBlock string `json:"destination_cidr_block"`
	// +optional
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id,omitempty"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}

type Ec2TransitGatewayRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
