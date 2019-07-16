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

type WafRuleGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRuleGroupSpec   `json:"spec,omitempty"`
	Status            WafRuleGroupStatus `json:"status,omitempty"`
}

type WafRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

type WafRuleGroupSpecActivatedRule struct {
	// +kubebuilder:validation:MaxItems=1
	Action   []WafRuleGroupSpecActivatedRule `json:"action"`
	Priority int                             `json:"priority"`
	RuleId   string                          `json:"rule_id"`
	// +optional
	Type string `json:"type,omitempty"`
}

type WafRuleGroupSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ActivatedRule *[]WafRuleGroupSpec `json:"activated_rule,omitempty"`
	MetricName    string              `json:"metric_name"`
	Name          string              `json:"name"`
}

type WafRuleGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRuleGroupList is a list of WafRuleGroups
type WafRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRuleGroup CRD objects
	Items []WafRuleGroup `json:"items,omitempty"`
}
