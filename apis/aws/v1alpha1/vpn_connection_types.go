package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpnConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnConnectionSpec   `json:"spec,omitempty"`
	Status            VpnConnectionStatus `json:"status,omitempty"`
}

type VpnConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	CustomerGatewayID string `json:"customerGatewayID" tf:"customer_gateway_id"`
	// +optional
	StaticRoutesOnly bool `json:"staticRoutesOnly,omitempty" tf:"static_routes_only,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	Tunnel1InsideCIDR string `json:"tunnel1InsideCIDR,omitempty" tf:"tunnel1_inside_cidr,omitempty"`
	// +optional
	Tunnel1PresharedKey string `json:"-" sensitive:"true" tf:"tunnel1_preshared_key,omitempty"`
	// +optional
	Tunnel2InsideCIDR string `json:"tunnel2InsideCIDR,omitempty" tf:"tunnel2_inside_cidr,omitempty"`
	// +optional
	Tunnel2PresharedKey string `json:"-" sensitive:"true" tf:"tunnel2_preshared_key,omitempty"`
	Type                string `json:"type" tf:"type"`
	// +optional
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type VpnConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnConnectionList is a list of VpnConnections
type VpnConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnConnection CRD objects
	Items []VpnConnection `json:"items,omitempty"`
}
