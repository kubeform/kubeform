package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualNetworkGatewayConnectionSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

type AzurermVirtualNetworkGatewayConnectionSpecIpsecPolicy struct {
	DhGroup         string `json:"dh_group"`
	IkeEncryption   string `json:"ike_encryption"`
	IkeIntegrity    string `json:"ike_integrity"`
	IpsecEncryption string `json:"ipsec_encryption"`
	IpsecIntegrity  string `json:"ipsec_integrity"`
	PfsGroup        string `json:"pfs_group"`
	SaDatasize      int    `json:"sa_datasize"`
	SaLifetime      int    `json:"sa_lifetime"`
}

type AzurermVirtualNetworkGatewayConnectionSpec struct {
	ResourceGroupName              string                                       `json:"resource_group_name"`
	Location                       string                                       `json:"location"`
	PeerVirtualNetworkGatewayId    string                                       `json:"peer_virtual_network_gateway_id"`
	Tags                           map[string]string                            `json:"tags"`
	LocalNetworkGatewayId          string                                       `json:"local_network_gateway_id"`
	RoutingWeight                  int                                          `json:"routing_weight"`
	SharedKey                      string                                       `json:"shared_key"`
	Name                           string                                       `json:"name"`
	Type                           string                                       `json:"type"`
	AuthorizationKey               string                                       `json:"authorization_key"`
	ExpressRouteCircuitId          string                                       `json:"express_route_circuit_id"`
	VirtualNetworkGatewayId        string                                       `json:"virtual_network_gateway_id"`
	EnableBgp                      bool                                         `json:"enable_bgp"`
	ExpressRouteGatewayBypass      bool                                         `json:"express_route_gateway_bypass"`
	IpsecPolicy                    []AzurermVirtualNetworkGatewayConnectionSpec `json:"ipsec_policy"`
	UsePolicyBasedTrafficSelectors bool                                         `json:"use_policy_based_traffic_selectors"`
}



type AzurermVirtualNetworkGatewayConnectionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualNetworkGatewayConnectionList is a list of AzurermVirtualNetworkGatewayConnections
type AzurermVirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualNetworkGatewayConnection CRD objects
	Items []AzurermVirtualNetworkGatewayConnection `json:"items,omitempty"`
}