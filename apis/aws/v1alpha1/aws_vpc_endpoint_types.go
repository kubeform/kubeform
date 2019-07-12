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
	VpcEndpointType     string               `json:"vpc_endpoint_type"`
	RequesterManaged    bool                 `json:"requester_managed"`
	State               string               `json:"state"`
	Policy              string               `json:"policy"`
	RouteTableIds       []string             `json:"route_table_ids"`
	SubnetIds           []string             `json:"subnet_ids"`
	NetworkInterfaceIds []string             `json:"network_interface_ids"`
	OwnerId             string               `json:"owner_id"`
	ServiceName         string               `json:"service_name"`
	VpcId               string               `json:"vpc_id"`
	DnsEntry            []AwsVpcEndpointSpec `json:"dns_entry"`
	PrivateDnsEnabled   bool                 `json:"private_dns_enabled"`
	PrefixListId        string               `json:"prefix_list_id"`
	SecurityGroupIds    []string             `json:"security_group_ids"`
	Tags                map[string]string    `json:"tags"`
	AutoAccept          bool                 `json:"auto_accept"`
	CidrBlocks          []string             `json:"cidr_blocks"`
}

type AwsVpcEndpointStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcEndpointList is a list of AwsVpcEndpoints
type AwsVpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpoint CRD objects
	Items []AwsVpcEndpoint `json:"items,omitempty"`
}
