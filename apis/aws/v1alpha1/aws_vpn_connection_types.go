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

type AwsVpnConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnConnectionSpec   `json:"spec,omitempty"`
	Status            AwsVpnConnectionStatus `json:"status,omitempty"`
}

type AwsVpnConnectionSpecVgwTelemetry struct {
	StatusMessage      string `json:"status_message"`
	AcceptedRouteCount int    `json:"accepted_route_count"`
	LastStatusChange   string `json:"last_status_change"`
	OutsideIpAddress   string `json:"outside_ip_address"`
	Status             string `json:"status"`
}

type AwsVpnConnectionSpecRoutes struct {
	DestinationCidrBlock string `json:"destination_cidr_block"`
	Source               string `json:"source"`
	State                string `json:"state"`
}

type AwsVpnConnectionSpec struct {
	VpnGatewayId                 string                 `json:"vpn_gateway_id"`
	Tunnel2InsideCidr            string                 `json:"tunnel2_inside_cidr"`
	Tunnel2PresharedKey          string                 `json:"tunnel2_preshared_key"`
	Tags                         map[string]string      `json:"tags"`
	Tunnel1VgwInsideAddress      string                 `json:"tunnel1_vgw_inside_address"`
	Tunnel2Address               string                 `json:"tunnel2_address"`
	Tunnel2VgwInsideAddress      string                 `json:"tunnel2_vgw_inside_address"`
	CustomerGatewayId            string                 `json:"customer_gateway_id"`
	Type                         string                 `json:"type"`
	CustomerGatewayConfiguration string                 `json:"customer_gateway_configuration"`
	Tunnel1Address               string                 `json:"tunnel1_address"`
	Tunnel1BgpAsn                string                 `json:"tunnel1_bgp_asn"`
	Tunnel1BgpHoldtime           int                    `json:"tunnel1_bgp_holdtime"`
	VgwTelemetry                 []AwsVpnConnectionSpec `json:"vgw_telemetry"`
	TransitGatewayId             string                 `json:"transit_gateway_id"`
	StaticRoutesOnly             bool                   `json:"static_routes_only"`
	Tunnel1CgwInsideAddress      string                 `json:"tunnel1_cgw_inside_address"`
	Tunnel2CgwInsideAddress      string                 `json:"tunnel2_cgw_inside_address"`
	TransitGatewayAttachmentId   string                 `json:"transit_gateway_attachment_id"`
	Tunnel1InsideCidr            string                 `json:"tunnel1_inside_cidr"`
	Tunnel1PresharedKey          string                 `json:"tunnel1_preshared_key"`
	Tunnel2BgpAsn                string                 `json:"tunnel2_bgp_asn"`
	Tunnel2BgpHoldtime           int                    `json:"tunnel2_bgp_holdtime"`
	Routes                       []AwsVpnConnectionSpec `json:"routes"`
}

type AwsVpnConnectionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpnConnectionList is a list of AwsVpnConnections
type AwsVpnConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnConnection CRD objects
	Items []AwsVpnConnection `json:"items,omitempty"`
}
