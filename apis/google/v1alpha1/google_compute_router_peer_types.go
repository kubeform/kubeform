package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouterPeerStatus `json:"status,omitempty"`
}

type GoogleComputeRouterPeerSpec struct {
	Name                    string `json:"name"`
	Interface               string `json:"interface"`
	PeerIpAddress           string `json:"peer_ip_address"`
	Project                 string `json:"project"`
	Router                  string `json:"router"`
	PeerAsn                 int    `json:"peer_asn"`
	AdvertisedRoutePriority int    `json:"advertised_route_priority"`
	IpAddress               string `json:"ip_address"`
	Region                  string `json:"region"`
}



type GoogleComputeRouterPeerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRouterPeerList is a list of GoogleComputeRouterPeers
type GoogleComputeRouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRouterPeer CRD objects
	Items []GoogleComputeRouterPeer `json:"items,omitempty"`
}