package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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

type VpnConnectionSpecRoutes struct {
	// +optional
	DestinationCIDRBlock string `json:"destinationCIDRBlock,omitempty" tf:"destination_cidr_block,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
}

type VpnConnectionSpecVgwTelemetry struct {
	// +optional
	AcceptedRouteCount int `json:"acceptedRouteCount,omitempty" tf:"accepted_route_count,omitempty"`
	// +optional
	LastStatusChange string `json:"lastStatusChange,omitempty" tf:"last_status_change,omitempty"`
	// +optional
	OutsideIPAddress string `json:"outsideIPAddress,omitempty" tf:"outside_ip_address,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StatusMessage string `json:"statusMessage,omitempty" tf:"status_message,omitempty"`
}

type VpnConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	CustomerGatewayConfiguration string `json:"customerGatewayConfiguration,omitempty" tf:"customer_gateway_configuration,omitempty"`
	CustomerGatewayID            string `json:"customerGatewayID" tf:"customer_gateway_id"`
	// +optional
	Routes []VpnConnectionSpecRoutes `json:"routes,omitempty" tf:"routes,omitempty"`
	// +optional
	StaticRoutesOnly bool `json:"staticRoutesOnly,omitempty" tf:"static_routes_only,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitGatewayAttachmentID string `json:"transitGatewayAttachmentID,omitempty" tf:"transit_gateway_attachment_id,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	Tunnel1Address string `json:"tunnel1Address,omitempty" tf:"tunnel1_address,omitempty"`
	// +optional
	Tunnel1BGPAsn string `json:"tunnel1BGPAsn,omitempty" tf:"tunnel1_bgp_asn,omitempty"`
	// +optional
	Tunnel1BGPHoldtime int `json:"tunnel1BGPHoldtime,omitempty" tf:"tunnel1_bgp_holdtime,omitempty"`
	// +optional
	Tunnel1CgwInsideAddress string `json:"tunnel1CgwInsideAddress,omitempty" tf:"tunnel1_cgw_inside_address,omitempty"`
	// +optional
	Tunnel1InsideCIDR string `json:"tunnel1InsideCIDR,omitempty" tf:"tunnel1_inside_cidr,omitempty"`
	// +optional
	Tunnel1PresharedKey string `json:"-" sensitive:"true" tf:"tunnel1_preshared_key,omitempty"`
	// +optional
	Tunnel1VgwInsideAddress string `json:"tunnel1VgwInsideAddress,omitempty" tf:"tunnel1_vgw_inside_address,omitempty"`
	// +optional
	Tunnel2Address string `json:"tunnel2Address,omitempty" tf:"tunnel2_address,omitempty"`
	// +optional
	Tunnel2BGPAsn string `json:"tunnel2BGPAsn,omitempty" tf:"tunnel2_bgp_asn,omitempty"`
	// +optional
	Tunnel2BGPHoldtime int `json:"tunnel2BGPHoldtime,omitempty" tf:"tunnel2_bgp_holdtime,omitempty"`
	// +optional
	Tunnel2CgwInsideAddress string `json:"tunnel2CgwInsideAddress,omitempty" tf:"tunnel2_cgw_inside_address,omitempty"`
	// +optional
	Tunnel2InsideCIDR string `json:"tunnel2InsideCIDR,omitempty" tf:"tunnel2_inside_cidr,omitempty"`
	// +optional
	Tunnel2PresharedKey string `json:"-" sensitive:"true" tf:"tunnel2_preshared_key,omitempty"`
	// +optional
	Tunnel2VgwInsideAddress string `json:"tunnel2VgwInsideAddress,omitempty" tf:"tunnel2_vgw_inside_address,omitempty"`
	Type                    string `json:"type" tf:"type"`
	// +optional
	VgwTelemetry []VpnConnectionSpecVgwTelemetry `json:"vgwTelemetry,omitempty" tf:"vgw_telemetry,omitempty"`
	// +optional
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type VpnConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpnConnectionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
