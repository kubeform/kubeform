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

type AzurermLocalNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLocalNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            AzurermLocalNetworkGatewayStatus `json:"status,omitempty"`
}

type AzurermLocalNetworkGatewaySpecBgpSettings struct {
	Asn               int    `json:"asn"`
	BgpPeeringAddress string `json:"bgp_peering_address"`
	PeerWeight        int    `json:"peer_weight"`
}

type AzurermLocalNetworkGatewaySpec struct {
	ResourceGroupName string                           `json:"resource_group_name"`
	GatewayAddress    string                           `json:"gateway_address"`
	AddressSpace      []string                         `json:"address_space"`
	BgpSettings       []AzurermLocalNetworkGatewaySpec `json:"bgp_settings"`
	Tags              map[string]string                `json:"tags"`
	Name              string                           `json:"name"`
	Location          string                           `json:"location"`
}

type AzurermLocalNetworkGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLocalNetworkGatewayList is a list of AzurermLocalNetworkGateways
type AzurermLocalNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLocalNetworkGateway CRD objects
	Items []AzurermLocalNetworkGateway `json:"items,omitempty"`
}
