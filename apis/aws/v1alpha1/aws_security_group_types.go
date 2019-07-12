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

type AwsSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsSecurityGroupStatus `json:"status,omitempty"`
}

type AwsSecurityGroupSpecIngress struct {
	CidrBlocks     []string `json:"cidr_blocks"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	SecurityGroups []string `json:"security_groups"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
}

type AwsSecurityGroupSpecEgress struct {
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	SecurityGroups []string `json:"security_groups"`
	ToPort         int      `json:"to_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
}

type AwsSecurityGroupSpec struct {
	NamePrefix          string                 `json:"name_prefix"`
	Description         string                 `json:"description"`
	Ingress             []AwsSecurityGroupSpec `json:"ingress"`
	Arn                 string                 `json:"arn"`
	RevokeRulesOnDelete bool                   `json:"revoke_rules_on_delete"`
	Name                string                 `json:"name"`
	VpcId               string                 `json:"vpc_id"`
	Egress              []AwsSecurityGroupSpec `json:"egress"`
	OwnerId             string                 `json:"owner_id"`
	Tags                map[string]string      `json:"tags"`
}

type AwsSecurityGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSecurityGroupList is a list of AwsSecurityGroups
type AwsSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityGroup CRD objects
	Items []AwsSecurityGroup `json:"items,omitempty"`
}
