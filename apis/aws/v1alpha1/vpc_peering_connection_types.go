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

type VpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionSpecAccepter struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionSpecRequester struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AcceptStatus string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Accepter []VpcPeeringConnectionSpecAccepter `json:"accepter,omitempty" tf:"accepter,omitempty"`
	// +optional
	AutoAccept bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	// +optional
	PeerOwnerID string `json:"peerOwnerID,omitempty" tf:"peer_owner_id,omitempty"`
	// +optional
	PeerRegion string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`
	PeerVpcID  string `json:"peerVpcID" tf:"peer_vpc_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Requester []VpcPeeringConnectionSpecRequester `json:"requester,omitempty" tf:"requester,omitempty"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID string            `json:"vpcID" tf:"vpc_id"`
}

type VpcPeeringConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcPeeringConnectionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionList is a list of VpcPeeringConnections
type VpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnection CRD objects
	Items []VpcPeeringConnection `json:"items,omitempty"`
}
