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

type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec,omitempty"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

type DefaultSecurityGroupSpecEgress struct {
	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	FromPort    int    `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIDS []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids,omitempty"`
	Protocol      string   `json:"protocol" tf:"protocol"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty" tf:"self,omitempty"`
	ToPort int  `json:"toPort" tf:"to_port"`
}

type DefaultSecurityGroupSpecIngress struct {
	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	FromPort    int    `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIDS []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids,omitempty"`
	Protocol      string   `json:"protocol" tf:"protocol"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty" tf:"self,omitempty"`
	ToPort int  `json:"toPort" tf:"to_port"`
}

type DefaultSecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Egress []DefaultSecurityGroupSpecEgress `json:"egress,omitempty" tf:"egress,omitempty"`
	// +optional
	Ingress []DefaultSecurityGroupSpecIngress `json:"ingress,omitempty" tf:"ingress,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	RevokeRulesOnDelete bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultSecurityGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultSecurityGroupList is a list of DefaultSecurityGroups
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultSecurityGroup CRD objects
	Items []DefaultSecurityGroup `json:"items,omitempty"`
}