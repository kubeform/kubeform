package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcPeeringConnectionAccepterSpec   `json:"spec,omitempty"`
	Status            AwsVpcPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

type AwsVpcPeeringConnectionAccepterSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpec struct {
	Requester              []AwsVpcPeeringConnectionAccepterSpec `json:"requester"`
	AutoAccept             bool                                  `json:"auto_accept"`
	AcceptStatus           string                                `json:"accept_status"`
	PeerVpcId              string                                `json:"peer_vpc_id"`
	PeerRegion             string                                `json:"peer_region"`
	Accepter               []AwsVpcPeeringConnectionAccepterSpec `json:"accepter"`
	Tags                   map[string]string                     `json:"tags"`
	VpcPeeringConnectionId string                                `json:"vpc_peering_connection_id"`
	VpcId                  string                                `json:"vpc_id"`
	PeerOwnerId            string                                `json:"peer_owner_id"`
}

type AwsVpcPeeringConnectionAccepterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcPeeringConnectionAccepterList is a list of AwsVpcPeeringConnectionAccepters
type AwsVpcPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcPeeringConnectionAccepter CRD objects
	Items []AwsVpcPeeringConnectionAccepter `json:"items,omitempty"`
}
