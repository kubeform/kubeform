package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AzureAsn                int    `json:"azureAsn,omitempty" tf:"azure_asn,omitempty"`
	ExpressRouteCircuitName string `json:"expressRouteCircuitName" tf:"express_route_circuit_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MicrosoftPeeringConfig []ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig `json:"microsoftPeeringConfig,omitempty" tf:"microsoft_peering_config,omitempty"`
	// +optional
	PeerAsn     int    `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`
	PeeringType string `json:"peeringType" tf:"peering_type"`
	// +optional
	PrimaryAzurePort         string `json:"primaryAzurePort,omitempty" tf:"primary_azure_port,omitempty"`
	PrimaryPeerAddressPrefix string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix"`
	ResourceGroupName        string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAzurePort         string `json:"secondaryAzurePort,omitempty" tf:"secondary_azure_port,omitempty"`
	SecondaryPeerAddressPrefix string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix"`
	// +optional
	SharedKey string `json:"-" sensitive:"true" tf:"shared_key,omitempty"`
	VlanID    int    `json:"vlanID" tf:"vlan_id"`
}

type ExpressRouteCircuitPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ExpressRouteCircuitPeeringSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
