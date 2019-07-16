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

type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            VirtualNetworkGatewayStatus `json:"status,omitempty"`
}

type VirtualNetworkGatewaySpecIpConfiguration struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation,omitempty"`
	// +optional
	PublicIpAddressId string `json:"public_ip_address_id,omitempty"`
	SubnetId          string `json:"subnet_id"`
}

type VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate struct {
	Name       string `json:"name"`
	Thumbprint string `json:"thumbprint"`
}

type VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate struct {
	Name           string `json:"name"`
	PublicCertData string `json:"public_cert_data"`
}

type VirtualNetworkGatewaySpecVpnClientConfiguration struct {
	AddressSpace []string `json:"address_space"`
	// +optional
	RadiusServerAddress string `json:"radius_server_address,omitempty"`
	// +optional
	RadiusServerSecret string `json:"radius_server_secret,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RevokedCertificate *[]VirtualNetworkGatewaySpecVpnClientConfiguration `json:"revoked_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RootCertificate *[]VirtualNetworkGatewaySpecVpnClientConfiguration `json:"root_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpnClientProtocols []string `json:"vpn_client_protocols,omitempty"`
}

type VirtualNetworkGatewaySpec struct {
	// +optional
	DefaultLocalNetworkGatewayId string `json:"default_local_network_gateway_id,omitempty"`
	// +kubebuilder:validation:MaxItems=2
	IpConfiguration   []VirtualNetworkGatewaySpec `json:"ip_configuration"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Sku               string                      `json:"sku"`
	Type              string                      `json:"type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpnClientConfiguration *[]VirtualNetworkGatewaySpec `json:"vpn_client_configuration,omitempty"`
	// +optional
	VpnType string `json:"vpn_type,omitempty"`
}

type VirtualNetworkGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
