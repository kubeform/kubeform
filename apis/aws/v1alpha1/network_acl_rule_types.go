package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NetworkACLRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLRuleSpec   `json:"spec,omitempty"`
	Status            NetworkACLRuleStatus `json:"status,omitempty"`
}

type NetworkACLRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	// +optional
	Egress bool `json:"egress,omitempty" tf:"egress,omitempty"`
	// +optional
	FromPort int `json:"fromPort,omitempty" tf:"from_port,omitempty"`
	// +optional
	IcmpCode string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`
	// +optional
	IcmpType string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	NetworkACLID  string `json:"networkACLID" tf:"network_acl_id"`
	Protocol      string `json:"protocol" tf:"protocol"`
	RuleAction    string `json:"ruleAction" tf:"rule_action"`
	RuleNumber    int    `json:"ruleNumber" tf:"rule_number"`
	// +optional
	ToPort int `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type NetworkACLRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkACLRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkACLRuleList is a list of NetworkACLRules
type NetworkACLRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkACLRule CRD objects
	Items []NetworkACLRule `json:"items,omitempty"`
}