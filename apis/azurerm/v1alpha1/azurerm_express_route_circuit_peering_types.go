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
	PrimaryPeerAddressPrefix   string                                  `json:"primary_peer_address_prefix"`
	VlanId                     int                                     `json:"vlan_id"`
	SharedKey                  string                                  `json:"shared_key"`
	PeerAsn                    int                                     `json:"peer_asn"`
	PeeringType                string                                  `json:"peering_type"`
	ResourceGroupName          string                                  `json:"resource_group_name"`
	MicrosoftPeeringConfig     []AzurermExpressRouteCircuitPeeringSpec `json:"microsoft_peering_config"`
	AzureAsn                   int                                     `json:"azure_asn"`
	PrimaryAzurePort           string                                  `json:"primary_azure_port"`
	SecondaryAzurePort         string                                  `json:"secondary_azure_port"`
	ExpressRouteCircuitName    string                                  `json:"express_route_circuit_name"`
	SecondaryPeerAddressPrefix string                                  `json:"secondary_peer_address_prefix"`
}



type AzurermExpressRouteCircuitPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermExpressRouteCircuitPeeringList is a list of AzurermExpressRouteCircuitPeerings
type AzurermExpressRouteCircuitPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermExpressRouteCircuitPeering CRD objects
	Items []AzurermExpressRouteCircuitPeering `json:"items,omitempty"`
}