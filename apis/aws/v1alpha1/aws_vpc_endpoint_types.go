package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointStatus `json:"status,omitempty"`
}

type AwsVpcEndpointSpecDnsEntry struct {
	DnsName      string `json:"dns_name"`
	HostedZoneId string `json:"hosted_zone_id"`
}

type AwsVpcEndpointSpec struct {
	PrivateDnsEnabled   bool                 `json:"private_dns_enabled"`
	VpcId               string               `json:"vpc_id"`
	VpcEndpointType     string               `json:"vpc_endpoint_type"`
	DnsEntry            []AwsVpcEndpointSpec `json:"dns_entry"`
	SecurityGroupIds    []string             `json:"security_group_ids"`
	CidrBlocks          []string             `json:"cidr_blocks"`
	RouteTableIds       []string             `json:"route_table_ids"`
	SubnetIds           []string             `json:"subnet_ids"`
	State               string               `json:"state"`
	NetworkInterfaceIds []string             `json:"network_interface_ids"`
	ServiceName         string               `json:"service_name"`
	Policy              string               `json:"policy"`
	PrefixListId        string               `json:"prefix_list_id"`
	AutoAccept          bool                 `json:"auto_accept"`
}



type AwsVpcEndpointStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointList is a list of AwsVpcEndpoints
type AwsVpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpoint CRD objects
	Items []AwsVpcEndpoint `json:"items,omitempty"`
}