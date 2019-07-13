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

type AwsVpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcPeeringConnectionSpec   `json:"spec,omitempty"`
	Status            AwsVpcPeeringConnectionStatus `json:"status,omitempty"`
}

type AwsVpcPeeringConnectionSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionSpec struct {
	Accepter     []AwsVpcPeeringConnectionSpec `json:"accepter"`
	PeerOwnerId  string                        `json:"peer_owner_id"`
	AutoAccept   bool                          `json:"auto_accept"`
	AcceptStatus string                        `json:"accept_status"`
	PeerRegion   string                        `json:"peer_region"`
	Requester    []AwsVpcPeeringConnectionSpec `json:"requester"`
	Tags         map[string]string             `json:"tags"`
	PeerVpcId    string                        `json:"peer_vpc_id"`
	VpcId        string                        `json:"vpc_id"`
}



type AwsVpcPeeringConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcPeeringConnectionList is a list of AwsVpcPeeringConnections
type AwsVpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcPeeringConnection CRD objects
	Items []AwsVpcPeeringConnection `json:"items,omitempty"`
}