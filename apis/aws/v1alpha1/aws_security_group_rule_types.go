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

type AwsSecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityGroupRuleSpec   `json:"spec,omitempty"`
	Status            AwsSecurityGroupRuleStatus `json:"status,omitempty"`
}

type AwsSecurityGroupRuleSpec struct {
	Type                  string   `json:"type"`
	ToPort                int      `json:"to_port"`
	CidrBlocks            []string `json:"cidr_blocks"`
	SourceSecurityGroupId string   `json:"source_security_group_id"`
	Self                  bool     `json:"self"`
	Description           string   `json:"description"`
	FromPort              int      `json:"from_port"`
	Protocol              string   `json:"protocol"`
	Ipv6CidrBlocks        []string `json:"ipv6_cidr_blocks"`
	PrefixListIds         []string `json:"prefix_list_ids"`
	SecurityGroupId       string   `json:"security_group_id"`
}

type AwsSecurityGroupRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSecurityGroupRuleList is a list of AwsSecurityGroupRules
type AwsSecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityGroupRule CRD objects
	Items []AwsSecurityGroupRule `json:"items,omitempty"`
}
