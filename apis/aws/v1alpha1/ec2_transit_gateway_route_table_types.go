package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Ec2TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteTableSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayRouteTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DefaultAssociationRouteTable bool `json:"defaultAssociationRouteTable,omitempty" tf:"default_association_route_table,omitempty"`
	// +optional
	DefaultPropagationRouteTable bool `json:"defaultPropagationRouteTable,omitempty" tf:"default_propagation_route_table,omitempty"`
	// +optional
	Tags             map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TransitGatewayID string            `json:"transitGatewayID" tf:"transit_gateway_id"`
}



type Ec2TransitGatewayRouteTableStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *Ec2TransitGatewayRouteTableSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableList is a list of Ec2TransitGatewayRouteTables
type Ec2TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGatewayRouteTable CRD objects
	Items []Ec2TransitGatewayRouteTable `json:"items,omitempty"`
}