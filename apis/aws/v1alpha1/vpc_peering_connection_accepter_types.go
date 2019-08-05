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

type VpcPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionAccepterSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionAccepterSpecAccepter struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionAccepterSpecRequester struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionAccepterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AcceptStatus string `json:"acceptStatus" tf:"accept_status"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Accepter []VpcPeeringConnectionAccepterSpecAccepter `json:"accepter,omitempty" tf:"accepter,omitempty"`
	// +optional
	AutoAccept  bool   `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	PeerOwnerID string `json:"peerOwnerID" tf:"peer_owner_id"`
	PeerRegion  string `json:"peerRegion" tf:"peer_region"`
	PeerVpcID   string `json:"peerVpcID" tf:"peer_vpc_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Requester []VpcPeeringConnectionAccepterSpecRequester `json:"requester,omitempty" tf:"requester,omitempty"`
	// +optional
	Tags                   map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID                  string            `json:"vpcID" tf:"vpc_id"`
	VpcPeeringConnectionID string            `json:"vpcPeeringConnectionID" tf:"vpc_peering_connection_id"`
}

type VpcPeeringConnectionAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcPeeringConnectionAccepterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionAccepterList is a list of VpcPeeringConnectionAccepters
type VpcPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnectionAccepter CRD objects
	Items []VpcPeeringConnectionAccepter `json:"items,omitempty"`
}
