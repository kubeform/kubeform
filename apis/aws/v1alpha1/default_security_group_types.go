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

type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec,omitempty"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

type DefaultSecurityGroupSpecEgress struct {
	// +optional
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	FromPort    int    `json:"from_port"`
	// +optional
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIds []string `json:"prefix_list_ids,omitempty"`
	Protocol      string   `json:"protocol"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty"`
	ToPort int  `json:"to_port"`
}

type DefaultSecurityGroupSpecIngress struct {
	// +optional
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	FromPort    int    `json:"from_port"`
	// +optional
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIds []string `json:"prefix_list_ids,omitempty"`
	Protocol      string   `json:"protocol"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +optional
	Self   bool `json:"self,omitempty"`
	ToPort int  `json:"to_port"`
}

type DefaultSecurityGroupSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Egress *[]DefaultSecurityGroupSpec `json:"egress,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ingress *[]DefaultSecurityGroupSpec `json:"ingress,omitempty"`
	// +optional
	RevokeRulesOnDelete bool `json:"revoke_rules_on_delete,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DefaultSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
