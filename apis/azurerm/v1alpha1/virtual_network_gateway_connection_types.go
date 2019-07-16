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

type VirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewayConnectionSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

type VirtualNetworkGatewayConnectionSpecIpsecPolicy struct {
	DhGroup         string `json:"dh_group"`
	IkeEncryption   string `json:"ike_encryption"`
	IkeIntegrity    string `json:"ike_integrity"`
	IpsecEncryption string `json:"ipsec_encryption"`
	IpsecIntegrity  string `json:"ipsec_integrity"`
	PfsGroup        string `json:"pfs_group"`
}

type VirtualNetworkGatewayConnectionSpec struct {
	// +optional
	AuthorizationKey string `json:"authorization_key,omitempty"`
	// +optional
	ExpressRouteCircuitId string `json:"express_route_circuit_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpsecPolicy *[]VirtualNetworkGatewayConnectionSpec `json:"ipsec_policy,omitempty"`
	// +optional
	LocalNetworkGatewayId string `json:"local_network_gateway_id,omitempty"`
	Location              string `json:"location"`
	Name                  string `json:"name"`
	// +optional
	PeerVirtualNetworkGatewayId string `json:"peer_virtual_network_gateway_id,omitempty"`
	ResourceGroupName           string `json:"resource_group_name"`
	// +optional
	SharedKey               string `json:"shared_key,omitempty"`
	Type                    string `json:"type"`
	VirtualNetworkGatewayId string `json:"virtual_network_gateway_id"`
}

type VirtualNetworkGatewayConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
