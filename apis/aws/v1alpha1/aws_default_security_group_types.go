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

type AwsDefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDefaultSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsDefaultSecurityGroupStatus `json:"status,omitempty"`
}

type AwsDefaultSecurityGroupSpecEgress struct {
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	Protocol       string   `json:"protocol"`
	ToPort         int      `json:"to_port"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups []string `json:"security_groups"`
	FromPort       int      `json:"from_port"`
}

type AwsDefaultSecurityGroupSpecIngress struct {
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
}

type AwsDefaultSecurityGroupSpec struct {
	VpcId               string                        `json:"vpc_id"`
	Egress              []AwsDefaultSecurityGroupSpec `json:"egress"`
	Arn                 string                        `json:"arn"`
	Tags                map[string]string             `json:"tags"`
	RevokeRulesOnDelete bool                          `json:"revoke_rules_on_delete"`
	Name                string                        `json:"name"`
	OwnerId             string                        `json:"owner_id"`
	Ingress             []AwsDefaultSecurityGroupSpec `json:"ingress"`
}



type AwsDefaultSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDefaultSecurityGroupList is a list of AwsDefaultSecurityGroups
type AwsDefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDefaultSecurityGroup CRD objects
	Items []AwsDefaultSecurityGroup `json:"items,omitempty"`
}