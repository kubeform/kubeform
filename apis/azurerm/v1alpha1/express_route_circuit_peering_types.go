package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ExpressRouteCircuitPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitPeeringSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitPeeringStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig struct {
	AdvertisedPublicPrefixes []string `json:"advertisedPublicPrefixes" tf:"advertised_public_prefixes"`
}

type ExpressRouteCircuitPeeringSpec struct {
	ExpressRouteCircuitName string `json:"expressRouteCircuitName" tf:"express_route_circuit_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MicrosoftPeeringConfig []ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig `json:"microsoftPeeringConfig,omitempty" tf:"microsoft_peering_config,omitempty"`
	// +optional
	PeerAsn                    int    `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`
	PeeringType                string `json:"peeringType" tf:"peering_type"`
	PrimaryPeerAddressPrefix   string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	SecondaryPeerAddressPrefix string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	SharedKey   *core.LocalObjectReference `json:"sharedKey,omitempty" tf:"shared_key,omitempty"`
	VlanID      int                        `json:"vlanID" tf:"vlan_id"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type ExpressRouteCircuitPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ExpressRouteCircuitPeeringList is a list of ExpressRouteCircuitPeerings
type ExpressRouteCircuitPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ExpressRouteCircuitPeering CRD objects
	Items []ExpressRouteCircuitPeering `json:"items,omitempty"`
}
