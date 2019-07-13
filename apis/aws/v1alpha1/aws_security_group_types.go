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

type AwsSecurityGroupSpecEgress struct {
	CidrBlocks     []string `json:"cidr_blocks"`
	Self           bool     `json:"self"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups []string `json:"security_groups"`
	Description    string   `json:"description"`
	Protocol       string   `json:"protocol"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
}

type AwsSecurityGroupSpecIngress struct {
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Description    string   `json:"description"`
}

type AwsSecurityGroupSpec struct {
	VpcId               string                 `json:"vpc_id"`
	OwnerId             string                 `json:"owner_id"`
	Egress              []AwsSecurityGroupSpec `json:"egress"`
	Arn                 string                 `json:"arn"`
	Tags                map[string]string      `json:"tags"`
	RevokeRulesOnDelete bool                   `json:"revoke_rules_on_delete"`
	Name                string                 `json:"name"`
	NamePrefix          string                 `json:"name_prefix"`
	Description         string                 `json:"description"`
	Ingress             []AwsSecurityGroupSpec `json:"ingress"`
}



type AwsSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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