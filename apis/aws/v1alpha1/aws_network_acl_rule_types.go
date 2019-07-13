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

type AwsNetworkAclRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkAclRuleSpec   `json:"spec,omitempty"`
	Status            AwsNetworkAclRuleStatus `json:"status,omitempty"`
}

type AwsNetworkAclRuleSpec struct {
	ToPort        int    `json:"to_port"`
	IcmpCode      string `json:"icmp_code"`
	Protocol      string `json:"protocol"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	FromPort      int    `json:"from_port"`
	IcmpType      string `json:"icmp_type"`
	NetworkAclId  string `json:"network_acl_id"`
	RuleNumber    int    `json:"rule_number"`
	Egress        bool   `json:"egress"`
	RuleAction    string `json:"rule_action"`
}



type AwsNetworkAclRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNetworkAclRuleList is a list of AwsNetworkAclRules
type AwsNetworkAclRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkAclRule CRD objects
	Items []AwsNetworkAclRule `json:"items,omitempty"`
}