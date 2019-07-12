package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermExpressRouteCircuitPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermExpressRouteCircuitPeeringSpec   `json:"spec,omitempty"`
	Status            AzurermExpressRouteCircuitPeeringStatus `json:"status,omitempty"`
}

type AzurermExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig struct {
	AdvertisedPublicPrefixes []string `json:"advertised_public_prefixes"`
}

type AzurermExpressRouteCircuitPeeringSpec struct {
	ExpressRouteCircuitName    string                                  `json:"express_route_circuit_name"`
	PrimaryPeerAddressPrefix   string                                  `json:"primary_peer_address_prefix"`
	SecondaryPeerAddressPrefix string                                  `json:"secondary_peer_address_prefix"`
	VlanId                     int                                     `json:"vlan_id"`
	PeerAsn                    int                                     `json:"peer_asn"`
	PrimaryAzurePort           string                                  `json:"primary_azure_port"`
	SecondaryAzurePort         string                                  `json:"secondary_azure_port"`
	PeeringType                string                                  `json:"peering_type"`
	SharedKey                  string                                  `json:"shared_key"`
	MicrosoftPeeringConfig     []AzurermExpressRouteCircuitPeeringSpec `json:"microsoft_peering_config"`
	AzureAsn                   int                                     `json:"azure_asn"`
	ResourceGroupName          string                                  `json:"resource_group_name"`
}

type AzurermExpressRouteCircuitPeeringStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermExpressRouteCircuitPeeringList is a list of AzurermExpressRouteCircuitPeerings
type AzurermExpressRouteCircuitPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermExpressRouteCircuitPeering CRD objects
	Items []AzurermExpressRouteCircuitPeering `json:"items,omitempty"`
}
