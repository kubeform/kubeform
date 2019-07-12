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

type AwsSubnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSubnetSpec   `json:"spec,omitempty"`
	Status            AwsSubnetStatus `json:"status,omitempty"`
}

type AwsSubnetSpec struct {
	VpcId                       string            `json:"vpc_id"`
	CidrBlock                   string            `json:"cidr_block"`
	Ipv6CidrBlock               string            `json:"ipv6_cidr_block"`
	AvailabilityZone            string            `json:"availability_zone"`
	MapPublicIpOnLaunch         bool              `json:"map_public_ip_on_launch"`
	AssignIpv6AddressOnCreation bool              `json:"assign_ipv6_address_on_creation"`
	Arn                         string            `json:"arn"`
	AvailabilityZoneId          string            `json:"availability_zone_id"`
	Ipv6CidrBlockAssociationId  string            `json:"ipv6_cidr_block_association_id"`
	Tags                        map[string]string `json:"tags"`
	OwnerId                     string            `json:"owner_id"`
}

type AwsSubnetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSubnetList is a list of AwsSubnets
type AwsSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSubnet CRD objects
	Items []AwsSubnet `json:"items,omitempty"`
}
