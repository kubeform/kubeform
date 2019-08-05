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
	DestinationCIDRBlock string `json:"destinationCIDRBlock" tf:"destination_cidr_block"`
	Source               string `json:"source" tf:"source"`
	State                string `json:"state" tf:"state"`
}

type VpnConnectionSpecVgwTelemetry struct {
	AcceptedRouteCount int    `json:"acceptedRouteCount" tf:"accepted_route_count"`
	LastStatusChange   string `json:"lastStatusChange" tf:"last_status_change"`
	OutsideIPAddress   string `json:"outsideIPAddress" tf:"outside_ip_address"`
	Status             string `json:"status" tf:"status"`
	StatusMessage      string `json:"statusMessage" tf:"status_message"`
}

type VpnConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	CustomerGatewayConfiguration string `json:"customerGatewayConfiguration" tf:"customer_gateway_configuration"`
	CustomerGatewayID            string `json:"customerGatewayID" tf:"customer_gateway_id"`
	// +kubebuilder:validation:UniqueItems=true
	Routes []VpnConnectionSpecRoutes `json:"routes" tf:"routes"`
	// +optional
	StaticRoutesOnly bool `json:"staticRoutesOnly,omitempty" tf:"static_routes_only,omitempty"`
	// +optional
	Tags                       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TransitGatewayAttachmentID string            `json:"transitGatewayAttachmentID" tf:"transit_gateway_attachment_id"`
	// +optional
	TransitGatewayID        string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	Tunnel1Address          string `json:"tunnel1Address" tf:"tunnel1_address"`
	Tunnel1BGPAsn           string `json:"tunnel1BGPAsn" tf:"tunnel1_bgp_asn"`
	Tunnel1BGPHoldtime      int    `json:"tunnel1BGPHoldtime" tf:"tunnel1_bgp_holdtime"`
	Tunnel1CgwInsideAddress string `json:"tunnel1CgwInsideAddress" tf:"tunnel1_cgw_inside_address"`
	// +optional
	Tunnel1InsideCIDR string `json:"tunnel1InsideCIDR,omitempty" tf:"tunnel1_inside_cidr,omitempty"`
	// +optional
	Tunnel1PresharedKey     string `json:"-" sensitive:"true" tf:"tunnel1_preshared_key,omitempty"`
	Tunnel1VgwInsideAddress string `json:"tunnel1VgwInsideAddress" tf:"tunnel1_vgw_inside_address"`
	Tunnel2Address          string `json:"tunnel2Address" tf:"tunnel2_address"`
	Tunnel2BGPAsn           string `json:"tunnel2BGPAsn" tf:"tunnel2_bgp_asn"`
	Tunnel2BGPHoldtime      int    `json:"tunnel2BGPHoldtime" tf:"tunnel2_bgp_holdtime"`
	Tunnel2CgwInsideAddress string `json:"tunnel2CgwInsideAddress" tf:"tunnel2_cgw_inside_address"`
	// +optional
	Tunnel2InsideCIDR string `json:"tunnel2InsideCIDR,omitempty" tf:"tunnel2_inside_cidr,omitempty"`
	// +optional
	Tunnel2PresharedKey     string `json:"-" sensitive:"true" tf:"tunnel2_preshared_key,omitempty"`
	Tunnel2VgwInsideAddress string `json:"tunnel2VgwInsideAddress" tf:"tunnel2_vgw_inside_address"`
	Type                    string `json:"type" tf:"type"`
	// +kubebuilder:validation:UniqueItems=true
	VgwTelemetry []VpnConnectionSpecVgwTelemetry `json:"vgwTelemetry" tf:"vgw_telemetry"`
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
