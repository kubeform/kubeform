package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualNetworkPeeringStatus `json:"status,omitempty"`
}

type AzurermVirtualNetworkPeeringSpec struct {
	UseRemoteGateways         bool   `json:"use_remote_gateways"`
	Name                      string `json:"name"`
	ResourceGroupName         string `json:"resource_group_name"`
	VirtualNetworkName        string `json:"virtual_network_name"`
	RemoteVirtualNetworkId    string `json:"remote_virtual_network_id"`
	AllowVirtualNetworkAccess bool   `json:"allow_virtual_network_access"`
	AllowForwardedTraffic     bool   `json:"allow_forwarded_traffic"`
	AllowGatewayTransit       bool   `json:"allow_gateway_transit"`
}

type AzurermVirtualNetworkPeeringStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualNetworkPeeringList is a list of AzurermVirtualNetworkPeerings
type AzurermVirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualNetworkPeering CRD objects
	Items []AzurermVirtualNetworkPeering `json:"items,omitempty"`
}
