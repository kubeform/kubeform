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

type ExpressRouteCircuitPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitPeeringSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitPeeringStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig struct {
	AdvertisedPublicPrefixes []string `json:"advertised_public_prefixes"`
}

type ExpressRouteCircuitPeeringSpec struct {
	ExpressRouteCircuitName string `json:"express_route_circuit_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MicrosoftPeeringConfig     *[]ExpressRouteCircuitPeeringSpec `json:"microsoft_peering_config,omitempty"`
	PeeringType                string                            `json:"peering_type"`
	PrimaryPeerAddressPrefix   string                            `json:"primary_peer_address_prefix"`
	ResourceGroupName          string                            `json:"resource_group_name"`
	SecondaryPeerAddressPrefix string                            `json:"secondary_peer_address_prefix"`
	// +optional
	SharedKey string `json:"shared_key,omitempty"`
	VlanId    int    `json:"vlan_id"`
}

type ExpressRouteCircuitPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
