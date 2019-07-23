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

type DxPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxPrivateVirtualInterfaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AddressFamily string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`
	BgpAsn        int    `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`
	ConnectionID string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`
	// +optional
	DxGatewayID string `json:"dxGatewayID,omitempty" tf:"dx_gateway_id,omitempty"`
	// +optional
	Mtu  int    `json:"mtu,omitempty" tf:"mtu,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Vlan int               `json:"vlan" tf:"vlan"`
	// +optional
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type DxPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxPrivateVirtualInterfaceList is a list of DxPrivateVirtualInterfaces
type DxPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxPrivateVirtualInterface CRD objects
	Items []DxPrivateVirtualInterface `json:"items,omitempty"`
}
