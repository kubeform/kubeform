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

type DxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxHostedPrivateVirtualInterfaceSpec struct {
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
	Mtu            int                       `json:"mtu,omitempty" tf:"mtu,omitempty"`
	Name           string                    `json:"name" tf:"name"`
	OwnerAccountID string                    `json:"ownerAccountID" tf:"owner_account_id"`
	Vlan           int                       `json:"vlan" tf:"vlan"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DxHostedPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
