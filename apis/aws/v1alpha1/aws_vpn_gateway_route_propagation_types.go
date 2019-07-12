package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnGatewayRoutePropagationSpec   `json:"spec,omitempty"`
	Status            AwsVpnGatewayRoutePropagationStatus `json:"status,omitempty"`
}

type AwsVpnGatewayRoutePropagationSpec struct {
	VpnGatewayId string `json:"vpn_gateway_id"`
	RouteTableId string `json:"route_table_id"`
}

type AwsVpnGatewayRoutePropagationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayRoutePropagationList is a list of AwsVpnGatewayRoutePropagations
type AwsVpnGatewayRoutePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnGatewayRoutePropagation CRD objects
	Items []AwsVpnGatewayRoutePropagation `json:"items,omitempty"`
}
