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

type NetworkAclRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkAclRuleSpec   `json:"spec,omitempty"`
	Status            NetworkAclRuleStatus `json:"status,omitempty"`
}

type NetworkAclRuleSpec struct {
	// +optional
	CidrBlock string `json:"cidr_block,omitempty"`
	// +optional
	Egress bool `json:"egress,omitempty"`
	// +optional
	FromPort int `json:"from_port,omitempty"`
	// +optional
	IcmpCode string `json:"icmp_code,omitempty"`
	// +optional
	IcmpType string `json:"icmp_type,omitempty"`
	// +optional
	Ipv6CidrBlock string `json:"ipv6_cidr_block,omitempty"`
	NetworkAclId  string `json:"network_acl_id"`
	Protocol      string `json:"protocol"`
	RuleAction    string `json:"rule_action"`
	RuleNumber    int    `json:"rule_number"`
	// +optional
	ToPort int `json:"to_port,omitempty"`
}

type NetworkAclRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkAclRuleList is a list of NetworkAclRules
type NetworkAclRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkAclRule CRD objects
	Items []NetworkAclRule `json:"items,omitempty"`
}
