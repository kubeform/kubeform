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

type AwsEc2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteSpec struct {
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
	DestinationCidrBlock       string `json:"destination_cidr_block"`
	Blackhole                  bool   `json:"blackhole"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
}



type AwsEc2TransitGatewayRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayRouteList is a list of AwsEc2TransitGatewayRoutes
type AwsEc2TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRoute CRD objects
	Items []AwsEc2TransitGatewayRoute `json:"items,omitempty"`
}