package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec,omitempty"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

type SecurityGroupRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	FromPort    int    `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIDS   []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids,omitempty"`
	Protocol        string   `json:"protocol" tf:"protocol"`
	SecurityGroupID string   `json:"securityGroupID" tf:"security_group_id"`
	// +optional
	Self bool `json:"self,omitempty" tf:"self,omitempty"`
	// +optional
	SourceSecurityGroupID string `json:"sourceSecurityGroupID,omitempty" tf:"source_security_group_id,omitempty"`
	ToPort                int    `json:"toPort" tf:"to_port"`
	// Type of rule, ingress (inbound) or egress (outbound).
	Type string `json:"type" tf:"type"`
}

type SecurityGroupRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SecurityGroupRuleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityGroupRuleList is a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityGroupRule CRD objects
	Items []SecurityGroupRule `json:"items,omitempty"`
}
