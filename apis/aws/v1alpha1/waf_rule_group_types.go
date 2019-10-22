package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type WafRuleGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRuleGroupSpec   `json:"spec,omitempty"`
	Status            WafRuleGroupStatus `json:"status,omitempty"`
}

type WafRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type" tf:"type"`
}

type WafRuleGroupSpecActivatedRule struct {
	// +kubebuilder:validation:MaxItems=1
	Action   []WafRuleGroupSpecActivatedRuleAction `json:"action" tf:"action"`
	Priority int                                   `json:"priority" tf:"priority"`
	RuleID   string                                `json:"ruleID" tf:"rule_id"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type WafRuleGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActivatedRule []WafRuleGroupSpecActivatedRule `json:"activatedRule,omitempty" tf:"activated_rule,omitempty"`
	MetricName    string                          `json:"metricName" tf:"metric_name"`
	Name          string                          `json:"name" tf:"name"`
}

type WafRuleGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafRuleGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
