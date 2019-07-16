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

type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec,omitempty"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

type SecurityGroupRuleSpec struct {
	// +optional
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	FromPort    int    `json:"from_port"`
	// +optional
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks,omitempty"`
	// +optional
	PrefixListIds   []string `json:"prefix_list_ids,omitempty"`
	Protocol        string   `json:"protocol"`
	SecurityGroupId string   `json:"security_group_id"`
	// +optional
	Self   bool   `json:"self,omitempty"`
	ToPort int    `json:"to_port"`
	Type   string `json:"type"`
}

type SecurityGroupRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
