package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsDefaultSecurityGroupStatus `json:"status,omitempty"`
}

type AwsDefaultSecurityGroupSpecIngress struct {
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	CidrBlocks     []string `json:"cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
}

type AwsDefaultSecurityGroupSpecEgress struct {
	Protocol       string   `json:"protocol"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
}

type AwsDefaultSecurityGroupSpec struct {
	Name                string                        `json:"name"`
	Ingress             []AwsDefaultSecurityGroupSpec `json:"ingress"`
	Egress              []AwsDefaultSecurityGroupSpec `json:"egress"`
	OwnerId             string                        `json:"owner_id"`
	Tags                map[string]string             `json:"tags"`
	RevokeRulesOnDelete bool                          `json:"revoke_rules_on_delete"`
	VpcId               string                        `json:"vpc_id"`
	Arn                 string                        `json:"arn"`
}

type AwsDefaultSecurityGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSecurityGroupList is a list of AwsDefaultSecurityGroups
type AwsDefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultSecurityGroup CRD objects
	Items []AwsDefaultSecurityGroup `json:"items,omitempty"`
}
