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

type AwsDxBgpPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxBgpPeerSpec   `json:"spec,omitempty"`
	Status            AwsDxBgpPeerStatus `json:"status,omitempty"`
}

type AwsDxBgpPeerSpec struct {
	BgpAuthKey         string `json:"bgp_auth_key"`
	CustomerAddress    string `json:"customer_address"`
	BgpStatus          string `json:"bgp_status"`
	AwsDevice          string `json:"aws_device"`
	AddressFamily      string `json:"address_family"`
	VirtualInterfaceId string `json:"virtual_interface_id"`
	AmazonAddress      string `json:"amazon_address"`
	BgpPeerId          string `json:"bgp_peer_id"`
	BgpAsn             int    `json:"bgp_asn"`
}



type AwsDxBgpPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxBgpPeerList is a list of AwsDxBgpPeers
type AwsDxBgpPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxBgpPeer CRD objects
	Items []AwsDxBgpPeer `json:"items,omitempty"`
}