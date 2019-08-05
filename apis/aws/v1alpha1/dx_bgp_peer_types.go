package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DxBGPPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxBGPPeerSpec   `json:"spec,omitempty"`
	Status            DxBGPPeerStatus `json:"status,omitempty"`
}

type DxBGPPeerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AddressFamily string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`
	AwsDevice     string `json:"awsDevice" tf:"aws_device"`
	BgpAsn        int    `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`
	BgpPeerID  string `json:"bgpPeerID" tf:"bgp_peer_id"`
	BgpStatus  string `json:"bgpStatus" tf:"bgp_status"`
	// +optional
	CustomerAddress    string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`
	VirtualInterfaceID string `json:"virtualInterfaceID" tf:"virtual_interface_id"`
}

type DxBGPPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxBGPPeerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxBGPPeerList is a list of DxBGPPeers
type DxBGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxBGPPeer CRD objects
	Items []DxBGPPeer `json:"items,omitempty"`
}
