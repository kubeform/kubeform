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

type DxPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxPrivateVirtualInterfaceSpec struct {
	AddressFamily string `json:"address_family"`
	BgpAsn        int    `json:"bgp_asn"`
	ConnectionId  string `json:"connection_id"`
	// +optional
	DxGatewayId string `json:"dx_gateway_id,omitempty"`
	// +optional
	Mtu  int    `json:"mtu,omitempty"`
	Name string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	Vlan int               `json:"vlan"`
	// +optional
	VpnGatewayId string `json:"vpn_gateway_id,omitempty"`
}

type DxPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
