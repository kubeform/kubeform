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

type VirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewayConnectionSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

type VirtualNetworkGatewayConnectionSpecIpsecPolicy struct {
	DhGroup         string `json:"dhGroup" tf:"dh_group"`
	IkeEncryption   string `json:"ikeEncryption" tf:"ike_encryption"`
	IkeIntegrity    string `json:"ikeIntegrity" tf:"ike_integrity"`
	IpsecEncryption string `json:"ipsecEncryption" tf:"ipsec_encryption"`
	IpsecIntegrity  string `json:"ipsecIntegrity" tf:"ipsec_integrity"`
	PfsGroup        string `json:"pfsGroup" tf:"pfs_group"`
	// +optional
	SaDatasize int `json:"saDatasize,omitempty" tf:"sa_datasize,omitempty"`
	// +optional
	SaLifetime int `json:"saLifetime,omitempty" tf:"sa_lifetime,omitempty"`
}

type VirtualNetworkGatewayConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"kubeFormSecret,omitempty" tf:"-"`

	// +optional
	AuthorizationKey string `json:"-" sensitive:"true" tf:"authorization_key,omitempty"`
	// +optional
	EnableBGP bool `json:"enableBGP,omitempty" tf:"enable_bgp,omitempty"`
	// +optional
	ExpressRouteCircuitID string `json:"expressRouteCircuitID,omitempty" tf:"express_route_circuit_id,omitempty"`
	// +optional
	ExpressRouteGatewayBypass bool `json:"expressRouteGatewayBypass,omitempty" tf:"express_route_gateway_bypass,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpsecPolicy []VirtualNetworkGatewayConnectionSpecIpsecPolicy `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`
	// +optional
	LocalNetworkGatewayID string `json:"localNetworkGatewayID,omitempty" tf:"local_network_gateway_id,omitempty"`
	Location              string `json:"location" tf:"location"`
	Name                  string `json:"name" tf:"name"`
	// +optional
	PeerVirtualNetworkGatewayID string `json:"peerVirtualNetworkGatewayID,omitempty" tf:"peer_virtual_network_gateway_id,omitempty"`
	ResourceGroupName           string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RoutingWeight int `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`
	// +optional
	SharedKey string `json:"-" sensitive:"true" tf:"shared_key,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Type string            `json:"type" tf:"type"`
	// +optional
	UsePolicyBasedTrafficSelectors bool   `json:"usePolicyBasedTrafficSelectors,omitempty" tf:"use_policy_based_traffic_selectors,omitempty"`
	VirtualNetworkGatewayID        string `json:"virtualNetworkGatewayID" tf:"virtual_network_gateway_id"`
}

type VirtualNetworkGatewayConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualNetworkGatewayConnectionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnectionList is a list of VirtualNetworkGatewayConnections
type VirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetworkGatewayConnection CRD objects
	Items []VirtualNetworkGatewayConnection `json:"items,omitempty"`
}
