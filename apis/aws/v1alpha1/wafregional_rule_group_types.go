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

type WafregionalRuleGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRuleGroupSpec   `json:"spec,omitempty"`
	Status            WafregionalRuleGroupStatus `json:"status,omitempty"`
}

type WafregionalRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type" tf:"type"`
}

type WafregionalRuleGroupSpecActivatedRule struct {
	// +kubebuilder:validation:MaxItems=1
	Action   []WafregionalRuleGroupSpecActivatedRuleAction `json:"action" tf:"action"`
	Priority int                                           `json:"priority" tf:"priority"`
	RuleID   string                                        `json:"ruleID" tf:"rule_id"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type WafregionalRuleGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ActivatedRule []WafregionalRuleGroupSpecActivatedRule `json:"activatedRule,omitempty" tf:"activated_rule,omitempty"`
	MetricName    string                                  `json:"metricName" tf:"metric_name"`
	Name          string                                  `json:"name" tf:"name"`
}

type WafregionalRuleGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalRuleGroupList is a list of WafregionalRuleGroups
type WafregionalRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalRuleGroup CRD objects
	Items []WafregionalRuleGroup `json:"items,omitempty"`
}
