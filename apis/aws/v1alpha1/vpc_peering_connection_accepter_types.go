package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AcceptStatus string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Accepter []VpcPeeringConnectionAccepterSpecAccepter `json:"accepter,omitempty" tf:"accepter,omitempty"`
	// +optional
	AutoAccept bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	// +optional
	PeerOwnerID string `json:"peerOwnerID,omitempty" tf:"peer_owner_id,omitempty"`
	// +optional
	PeerRegion string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`
	// +optional
	PeerVpcID string `json:"peerVpcID,omitempty" tf:"peer_vpc_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Requester []VpcPeeringConnectionAccepterSpecRequester `json:"requester,omitempty" tf:"requester,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID                  string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	VpcPeeringConnectionID string `json:"vpcPeeringConnectionID" tf:"vpc_peering_connection_id"`
}



type VpcPeeringConnectionAccepterStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
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