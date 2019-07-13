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

type AwsDefaultVpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultVpcSpec   `json:"spec,omitempty"`
	Status            AwsDefaultVpcStatus `json:"status,omitempty"`
}

type AwsDefaultVpcSpec struct {
	MainRouteTableId             string            `json:"main_route_table_id"`
	CidrBlock                    string            `json:"cidr_block"`
	InstanceTenancy              string            `json:"instance_tenancy"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	DhcpOptionsId                string            `json:"dhcp_options_id"`
	DefaultRouteTableId          string            `json:"default_route_table_id"`
	Ipv6AssociationId            string            `json:"ipv6_association_id"`
	EnableClassiclink            bool              `json:"enable_classiclink"`
	DefaultNetworkAclId          string            `json:"default_network_acl_id"`
	DefaultSecurityGroupId       string            `json:"default_security_group_id"`
	Ipv6CidrBlock                string            `json:"ipv6_cidr_block"`
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	Arn                          string            `json:"arn"`
	Tags                         map[string]string `json:"tags"`
	OwnerId                      string            `json:"owner_id"`
}



type AwsDefaultVpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDefaultVpcList is a list of AwsDefaultVpcs
type AwsDefaultVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultVpc CRD objects
	Items []AwsDefaultVpc `json:"items,omitempty"`
}