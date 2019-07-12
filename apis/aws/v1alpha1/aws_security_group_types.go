package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsSecurityGroupStatus `json:"status,omitempty"`
}

type AwsSecurityGroupSpecIngress struct {
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	FromPort       int      `json:"from_port"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	Description    string   `json:"description"`
	ToPort         int      `json:"to_port"`
}

type AwsSecurityGroupSpecEgress struct {
	FromPort       int      `json:"from_port"`
	Protocol       string   `json:"protocol"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	SecurityGroups []string `json:"security_groups"`
	ToPort         int      `json:"to_port"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
}

type AwsSecurityGroupSpec struct {
	NamePrefix          string                 `json:"name_prefix"`
	Ingress             []AwsSecurityGroupSpec `json:"ingress"`
	Egress              []AwsSecurityGroupSpec `json:"egress"`
	Tags                map[string]string      `json:"tags"`
	RevokeRulesOnDelete bool                   `json:"revoke_rules_on_delete"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	VpcId               string                 `json:"vpc_id"`
	Arn                 string                 `json:"arn"`
	OwnerId             string                 `json:"owner_id"`
}



type AwsSecurityGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupList is a list of AwsSecurityGroups
type AwsSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityGroup CRD objects
	Items []AwsSecurityGroup `json:"items,omitempty"`
}