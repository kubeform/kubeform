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

type AwsVpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcSpec   `json:"spec,omitempty"`
	Status            AwsVpcStatus `json:"status,omitempty"`
}

type AwsVpcSpec struct {
	EnableClassiclink            bool              `json:"enable_classiclink"`
	DhcpOptionsId                string            `json:"dhcp_options_id"`
	DefaultRouteTableId          string            `json:"default_route_table_id"`
	Ipv6AssociationId            string            `json:"ipv6_association_id"`
	Ipv6CidrBlock                string            `json:"ipv6_cidr_block"`
	Tags                         map[string]string `json:"tags"`
	CidrBlock                    string            `json:"cidr_block"`
	InstanceTenancy              string            `json:"instance_tenancy"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	DefaultSecurityGroupId       string            `json:"default_security_group_id"`
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	MainRouteTableId             string            `json:"main_route_table_id"`
	Arn                          string            `json:"arn"`
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	DefaultNetworkAclId          string            `json:"default_network_acl_id"`
	OwnerId                      string            `json:"owner_id"`
}



type AwsVpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcList is a list of AwsVpcs
type AwsVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpc CRD objects
	Items []AwsVpc `json:"items,omitempty"`
}