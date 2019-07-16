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

type VpnGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayRoutePropagationSpec   `json:"spec,omitempty"`
	Status            VpnGatewayRoutePropagationStatus `json:"status,omitempty"`
}

type VpnGatewayRoutePropagationSpec struct {
	RouteTableId string `json:"route_table_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}

type VpnGatewayRoutePropagationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnGatewayRoutePropagationList is a list of VpnGatewayRoutePropagations
type VpnGatewayRoutePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnGatewayRoutePropagation CRD objects
	Items []VpnGatewayRoutePropagation `json:"items,omitempty"`
}
