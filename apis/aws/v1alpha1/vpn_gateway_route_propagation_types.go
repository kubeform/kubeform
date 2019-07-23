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

type VpnGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayRoutePropagationSpec   `json:"spec,omitempty"`
	Status            VpnGatewayRoutePropagationStatus `json:"status,omitempty"`
}

type VpnGatewayRoutePropagationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	RouteTableID string `json:"routeTableID" tf:"route_table_id"`
	VpnGatewayID string `json:"vpnGatewayID" tf:"vpn_gateway_id"`
}

type VpnGatewayRoutePropagationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
