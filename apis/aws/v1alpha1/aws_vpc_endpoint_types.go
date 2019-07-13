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
	HostedZoneId string `json:"hosted_zone_id"`
	DnsName      string `json:"dns_name"`
}

type AwsVpcEndpointSpec struct {
	DnsEntry            []AwsVpcEndpointSpec `json:"dns_entry"`
	NetworkInterfaceIds []string             `json:"network_interface_ids"`
	PrefixListId        string               `json:"prefix_list_id"`
	PrivateDnsEnabled   bool                 `json:"private_dns_enabled"`
	VpcEndpointType     string               `json:"vpc_endpoint_type"`
	AutoAccept          bool                 `json:"auto_accept"`
	CidrBlocks          []string             `json:"cidr_blocks"`
	VpcId               string               `json:"vpc_id"`
	State               string               `json:"state"`
	Tags                map[string]string    `json:"tags"`
	SecurityGroupIds    []string             `json:"security_group_ids"`
	ServiceName         string               `json:"service_name"`
	SubnetIds           []string             `json:"subnet_ids"`
	OwnerId             string               `json:"owner_id"`
	RequesterManaged    bool                 `json:"requester_managed"`
	Policy              string               `json:"policy"`
	RouteTableIds       []string             `json:"route_table_ids"`
}



type AwsVpcEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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