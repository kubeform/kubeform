package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultNetworkAclSpec   `json:"spec,omitempty"`
	Status            AwsDefaultNetworkAclStatus `json:"status,omitempty"`
}

type AwsDefaultNetworkAclSpecIngress struct {
	ToPort        int    `json:"to_port"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
}

type AwsDefaultNetworkAclSpecEgress struct {
	ToPort        int    `json:"to_port"`
	RuleNo        int    `json:"rule_no"`
	Protocol      string `json:"protocol"`
	FromPort      int    `json:"from_port"`
	Action        string `json:"action"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
}

type AwsDefaultNetworkAclSpec struct {
	DefaultNetworkAclId string                     `json:"default_network_acl_id"`
	SubnetIds           []string                   `json:"subnet_ids"`
	Ingress             []AwsDefaultNetworkAclSpec `json:"ingress"`
	Egress              []AwsDefaultNetworkAclSpec `json:"egress"`
	Tags                map[string]string          `json:"tags"`
	OwnerId             string                     `json:"owner_id"`
	VpcId               string                     `json:"vpc_id"`
}

type AwsDefaultNetworkAclStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAclList is a list of AwsDefaultNetworkAcls
type AwsDefaultNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultNetworkAcl CRD objects
	Items []AwsDefaultNetworkAcl `json:"items,omitempty"`
}
