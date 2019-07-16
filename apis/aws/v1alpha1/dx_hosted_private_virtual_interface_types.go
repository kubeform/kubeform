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

type DxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxHostedPrivateVirtualInterfaceSpec struct {
	AddressFamily string `json:"address_family"`
	BgpAsn        int    `json:"bgp_asn"`
	ConnectionId  string `json:"connection_id"`
	// +optional
	Mtu            int    `json:"mtu,omitempty"`
	Name           string `json:"name"`
	OwnerAccountId string `json:"owner_account_id"`
	Vlan           int    `json:"vlan"`
}

type DxHostedPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterfaceList is a list of DxHostedPrivateVirtualInterfaces
type DxHostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxHostedPrivateVirtualInterface CRD objects
	Items []DxHostedPrivateVirtualInterface `json:"items,omitempty"`
}
