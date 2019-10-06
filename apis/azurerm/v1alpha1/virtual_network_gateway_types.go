package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            VirtualNetworkGatewayStatus `json:"status,omitempty"`
}

type VirtualNetworkGatewaySpecBgpSettings struct {
	// +optional
	Asn int `json:"asn,omitempty" tf:"asn,omitempty"`
	// +optional
	PeerWeight int `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
	// +optional
	PeeringAddress string `json:"peeringAddress,omitempty" tf:"peering_address,omitempty"`
}

type VirtualNetworkGatewaySpecIpConfiguration struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	PrivateIPAddressAllocation string `json:"privateIPAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	SubnetID          string `json:"subnetID" tf:"subnet_id"`
}

type VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate struct {
	Name       string `json:"name" tf:"name"`
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate struct {
	Name           string `json:"name" tf:"name"`
	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type VirtualNetworkGatewaySpecVpnClientConfiguration struct {
	AddressSpace []string `json:"addressSpace" tf:"address_space"`
	// +optional
	RadiusServerAddress string `json:"radiusServerAddress,omitempty" tf:"radius_server_address,omitempty"`
	// +optional
	RadiusServerSecret string `json:"radiusServerSecret,omitempty" tf:"radius_server_secret,omitempty"`
	// +optional
	RevokedCertificate []VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate `json:"revokedCertificate,omitempty" tf:"revoked_certificate,omitempty"`
	// +optional
	RootCertificate []VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate `json:"rootCertificate,omitempty" tf:"root_certificate,omitempty"`
	// +optional
	VpnClientProtocols []string `json:"vpnClientProtocols,omitempty" tf:"vpn_client_protocols,omitempty"`
}

type VirtualNetworkGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActiveActive bool `json:"activeActive,omitempty" tf:"active_active,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BgpSettings []VirtualNetworkGatewaySpecBgpSettings `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`
	// +optional
	DefaultLocalNetworkGatewayID string `json:"defaultLocalNetworkGatewayID,omitempty" tf:"default_local_network_gateway_id,omitempty"`
	// +optional
	EnableBGP bool `json:"enableBGP,omitempty" tf:"enable_bgp,omitempty"`
	// +kubebuilder:validation:MaxItems=2
	IpConfiguration   []VirtualNetworkGatewaySpecIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Location          string                                     `json:"location" tf:"location"`
	Name              string                                     `json:"name" tf:"name"`
	ResourceGroupName string                                     `json:"resourceGroupName" tf:"resource_group_name"`
	Sku               string                                     `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Type string            `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpnClientConfiguration []VirtualNetworkGatewaySpecVpnClientConfiguration `json:"vpnClientConfiguration,omitempty" tf:"vpn_client_configuration,omitempty"`
	// +optional
	VpnType string `json:"vpnType,omitempty" tf:"vpn_type,omitempty"`
}

type VirtualNetworkGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualNetworkGatewaySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkGatewayList is a list of VirtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetworkGateway CRD objects
	Items []VirtualNetworkGateway `json:"items,omitempty"`
}
