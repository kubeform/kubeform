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

type DxBgpPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxBgpPeerSpec   `json:"spec,omitempty"`
	Status            DxBgpPeerStatus `json:"status,omitempty"`
}

type DxBgpPeerSpec struct {
	AddressFamily      string `json:"address_family"`
	BgpAsn             int    `json:"bgp_asn"`
	VirtualInterfaceId string `json:"virtual_interface_id"`
}

type DxBgpPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxBgpPeerList is a list of DxBgpPeers
type DxBgpPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxBgpPeer CRD objects
	Items []DxBgpPeer `json:"items,omitempty"`
}
