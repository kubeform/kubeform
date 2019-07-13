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

type AwsVpcPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcPeeringConnectionOptionsSpec   `json:"spec,omitempty"`
	Status            AwsVpcPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

type AwsVpcPeeringConnectionOptionsSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionOptionsSpecRequester struct {
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
}

type AwsVpcPeeringConnectionOptionsSpec struct {
	Accepter               []AwsVpcPeeringConnectionOptionsSpec `json:"accepter"`
	Requester              []AwsVpcPeeringConnectionOptionsSpec `json:"requester"`
	VpcPeeringConnectionId string                               `json:"vpc_peering_connection_id"`
}



type AwsVpcPeeringConnectionOptionsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcPeeringConnectionOptionsList is a list of AwsVpcPeeringConnectionOptionss
type AwsVpcPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcPeeringConnectionOptions CRD objects
	Items []AwsVpcPeeringConnectionOptions `json:"items,omitempty"`
}