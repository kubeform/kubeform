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

type ComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status            ComputeRouterPeerStatus `json:"status,omitempty"`
}

type ComputeRouterPeerSpec struct {
	// +optional
	AdvertisedRoutePriority int    `json:"advertisedRoutePriority,omitempty" tf:"advertised_route_priority,omitempty"`
	Interface               string `json:"interface" tf:"interface"`
	Name                    string `json:"name" tf:"name"`
	PeerAsn                 int    `json:"peerAsn" tf:"peer_asn"`
	// +optional
	PeerIPAddress string `json:"peerIPAddress,omitempty" tf:"peer_ip_address,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region      string                    `json:"region,omitempty" tf:"region,omitempty"`
	Router      string                    `json:"router" tf:"router"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeRouterPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterPeerList is a list of ComputeRouterPeers
type ComputeRouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterPeer CRD objects
	Items []ComputeRouterPeer `json:"items,omitempty"`
}
