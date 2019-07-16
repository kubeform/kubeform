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

type LocalNetworkGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalNetworkGatewaySpec   `json:"spec,omitempty"`
	Status            LocalNetworkGatewayStatus `json:"status,omitempty"`
}

type LocalNetworkGatewaySpecBgpSettings struct {
	Asn               int    `json:"asn"`
	BgpPeeringAddress string `json:"bgp_peering_address"`
}

type LocalNetworkGatewaySpec struct {
	AddressSpace []string `json:"address_space"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BgpSettings       *[]LocalNetworkGatewaySpec `json:"bgp_settings,omitempty"`
	GatewayAddress    string                     `json:"gateway_address"`
	Location          string                     `json:"location"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
}

type LocalNetworkGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
