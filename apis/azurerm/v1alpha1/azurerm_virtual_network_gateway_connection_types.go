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

type AzurermVirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualNetworkGatewayConnectionSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

type AzurermVirtualNetworkGatewayConnectionSpecIpsecPolicy struct {
	IpsecEncryption string `json:"ipsec_encryption"`
	IpsecIntegrity  string `json:"ipsec_integrity"`
	PfsGroup        string `json:"pfs_group"`
	SaDatasize      int    `json:"sa_datasize"`
	SaLifetime      int    `json:"sa_lifetime"`
	DhGroup         string `json:"dh_group"`
	IkeEncryption   string `json:"ike_encryption"`
	IkeIntegrity    string `json:"ike_integrity"`
}

type AzurermVirtualNetworkGatewayConnectionSpec struct {
	Location                       string                                       `json:"location"`
	Type                           string                                       `json:"type"`
	EnableBgp                      bool                                         `json:"enable_bgp"`
	UsePolicyBasedTrafficSelectors bool                                         `json:"use_policy_based_traffic_selectors"`
	SharedKey                      string                                       `json:"shared_key"`
	Name                           string                                       `json:"name"`
	ResourceGroupName              string                                       `json:"resource_group_name"`
	PeerVirtualNetworkGatewayId    string                                       `json:"peer_virtual_network_gateway_id"`
	VirtualNetworkGatewayId        string                                       `json:"virtual_network_gateway_id"`
	AuthorizationKey               string                                       `json:"authorization_key"`
	RoutingWeight                  int                                          `json:"routing_weight"`
	Tags                           map[string]string                            `json:"tags"`
	ExpressRouteCircuitId          string                                       `json:"express_route_circuit_id"`
	LocalNetworkGatewayId          string                                       `json:"local_network_gateway_id"`
	ExpressRouteGatewayBypass      bool                                         `json:"express_route_gateway_bypass"`
	IpsecPolicy                    []AzurermVirtualNetworkGatewayConnectionSpec `json:"ipsec_policy"`
}



type AzurermVirtualNetworkGatewayConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermVirtualNetworkGatewayConnectionList is a list of AzurermVirtualNetworkGatewayConnections
type AzurermVirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualNetworkGatewayConnection CRD objects
	Items []AzurermVirtualNetworkGatewayConnection `json:"items,omitempty"`
}