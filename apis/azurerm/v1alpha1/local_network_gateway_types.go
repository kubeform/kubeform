package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LocalNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            LocalNetworkGatewayStatus `json:"status,omitempty"`
}

type LocalNetworkGatewaySpecBgpSettings struct {
	Asn               int    `json:"asn" tf:"asn"`
	BgpPeeringAddress string `json:"bgpPeeringAddress" tf:"bgp_peering_address"`
	// +optional
	PeerWeight int `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
}

type LocalNetworkGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AddressSpace []string `json:"addressSpace" tf:"address_space"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BgpSettings       []LocalNetworkGatewaySpecBgpSettings `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`
	GatewayAddress    string                               `json:"gatewayAddress" tf:"gateway_address"`
	Location          string                               `json:"location" tf:"location"`
	Name              string                               `json:"name" tf:"name"`
	ResourceGroupName string                               `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LocalNetworkGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LocalNetworkGatewayList is a list of LocalNetworkGateways
type LocalNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LocalNetworkGateway CRD objects
	Items []LocalNetworkGateway `json:"items,omitempty"`
}
