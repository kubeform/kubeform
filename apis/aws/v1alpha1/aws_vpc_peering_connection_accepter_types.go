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

type AwsVpcPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcPeeringConnectionAccepterSpec   `json:"spec,omitempty"`
	Status            AwsVpcPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

type AwsVpcPeeringConnectionAccepterSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpec struct {
	PeerVpcId              string                                `json:"peer_vpc_id"`
	PeerOwnerId            string                                `json:"peer_owner_id"`
	PeerRegion             string                                `json:"peer_region"`
	Tags                   map[string]string                     `json:"tags"`
	AutoAccept             bool                                  `json:"auto_accept"`
	AcceptStatus           string                                `json:"accept_status"`
	VpcId                  string                                `json:"vpc_id"`
	VpcPeeringConnectionId string                                `json:"vpc_peering_connection_id"`
	Accepter               []AwsVpcPeeringConnectionAccepterSpec `json:"accepter"`
	Requester              []AwsVpcPeeringConnectionAccepterSpec `json:"requester"`
}



type AwsVpcPeeringConnectionAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcPeeringConnectionAccepterList is a list of AwsVpcPeeringConnectionAccepters
type AwsVpcPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcPeeringConnectionAccepter CRD objects
	Items []AwsVpcPeeringConnectionAccepter `json:"items,omitempty"`
}