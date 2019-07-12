package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            AzurermVirtualNetworkGatewayStatus `json:"status,omitempty"`
}

type AzurermVirtualNetworkGatewaySpecBgpSettings struct {
	Asn            int    `json:"asn"`
	PeeringAddress string `json:"peering_address"`
	PeerWeight     int    `json:"peer_weight"`
}

type AzurermVirtualNetworkGatewaySpecIpConfiguration struct {
	Name                       string `json:"name"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	SubnetId                   string `json:"subnet_id"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
}

type AzurermVirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate struct {
	Name           string `json:"name"`
	PublicCertData string `json:"public_cert_data"`
}

type AzurermVirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate struct {
	Name       string `json:"name"`
	Thumbprint string `json:"thumbprint"`
}

type AzurermVirtualNetworkGatewaySpecVpnClientConfiguration struct {
	VpnClientProtocols  []string                                                 `json:"vpn_client_protocols"`
	AddressSpace        []string                                                 `json:"address_space"`
	RootCertificate     []AzurermVirtualNetworkGatewaySpecVpnClientConfiguration `json:"root_certificate"`
	RevokedCertificate  []AzurermVirtualNetworkGatewaySpecVpnClientConfiguration `json:"revoked_certificate"`
	RadiusServerAddress string                                                   `json:"radius_server_address"`
	RadiusServerSecret  string                                                   `json:"radius_server_secret"`
}

type AzurermVirtualNetworkGatewaySpec struct {
	ResourceGroupName            string                             `json:"resource_group_name"`
	Type                         string                             `json:"type"`
	ActiveActive                 bool                               `json:"active_active"`
	Sku                          string                             `json:"sku"`
	BgpSettings                  []AzurermVirtualNetworkGatewaySpec `json:"bgp_settings"`
	Name                         string                             `json:"name"`
	VpnType                      string                             `json:"vpn_type"`
	EnableBgp                    bool                               `json:"enable_bgp"`
	IpConfiguration              []AzurermVirtualNetworkGatewaySpec `json:"ip_configuration"`
	VpnClientConfiguration       []AzurermVirtualNetworkGatewaySpec `json:"vpn_client_configuration"`
	DefaultLocalNetworkGatewayId string                             `json:"default_local_network_gateway_id"`
	Tags                         map[string]string                  `json:"tags"`
	Location                     string                             `json:"location"`
}

type AzurermVirtualNetworkGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualNetworkGatewayList is a list of AzurermVirtualNetworkGateways
type AzurermVirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualNetworkGateway CRD objects
	Items []AzurermVirtualNetworkGateway `json:"items,omitempty"`
}
