package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec,omitempty"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

type SecurityGroupSpecEgress struct {
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
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty" tf:"self,omitempty"`
	ToPort int  `json:"toPort" tf:"to_port"`
}

type SecurityGroupSpecIngress struct {
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
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty" tf:"self,omitempty"`
	ToPort int  `json:"toPort" tf:"to_port"`
}

type SecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Egress []SecurityGroupSpecEgress `json:"egress,omitempty" tf:"egress,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ingress []SecurityGroupSpecIngress `json:"ingress,omitempty" tf:"ingress,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	RevokeRulesOnDelete bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type SecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityGroupList is a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityGroup CRD objects
	Items []SecurityGroup `json:"items,omitempty"`
}
