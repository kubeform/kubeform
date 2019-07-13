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

type AwsEc2TransitGatewayRouteTablePropagation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewayRouteTablePropagationSpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayRouteTablePropagationStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewayRouteTablePropagationSpec struct {
	ResourceId                 string `json:"resource_id"`
	ResourceType               string `json:"resource_type"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}



type AwsEc2TransitGatewayRouteTablePropagationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayRouteTablePropagationList is a list of AwsEc2TransitGatewayRouteTablePropagations
type AwsEc2TransitGatewayRouteTablePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGatewayRouteTablePropagation CRD objects
	Items []AwsEc2TransitGatewayRouteTablePropagation `json:"items,omitempty"`
}