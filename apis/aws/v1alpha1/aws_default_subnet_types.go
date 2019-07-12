package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSubnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultSubnetSpec   `json:"spec,omitempty"`
	Status            AwsDefaultSubnetStatus `json:"status,omitempty"`
}

type AwsDefaultSubnetSpec struct {
	CidrBlock                   string            `json:"cidr_block"`
	Ipv6CidrBlock               string            `json:"ipv6_cidr_block"`
	AvailabilityZone            string            `json:"availability_zone"`
	AvailabilityZoneId          string            `json:"availability_zone_id"`
	VpcId                       string            `json:"vpc_id"`
	MapPublicIpOnLaunch         bool              `json:"map_public_ip_on_launch"`
	AssignIpv6AddressOnCreation bool              `json:"assign_ipv6_address_on_creation"`
	Ipv6CidrBlockAssociationId  string            `json:"ipv6_cidr_block_association_id"`
	Arn                         string            `json:"arn"`
	Tags                        map[string]string `json:"tags"`
	OwnerId                     string            `json:"owner_id"`
}

type AwsDefaultSubnetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSubnetList is a list of AwsDefaultSubnets
type AwsDefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultSubnet CRD objects
	Items []AwsDefaultSubnet `json:"items,omitempty"`
}
