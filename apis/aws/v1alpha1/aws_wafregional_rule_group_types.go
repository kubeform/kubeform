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

type AwsWafregionalRuleGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalRuleGroupSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalRuleGroupStatus `json:"status,omitempty"`
}

type AwsWafregionalRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

type AwsWafregionalRuleGroupSpecActivatedRule struct {
	Action   []AwsWafregionalRuleGroupSpecActivatedRule `json:"action"`
	Priority int                                        `json:"priority"`
	RuleId   string                                     `json:"rule_id"`
	Type     string                                     `json:"type"`
}

type AwsWafregionalRuleGroupSpec struct {
	Name          string                        `json:"name"`
	MetricName    string                        `json:"metric_name"`
	ActivatedRule []AwsWafregionalRuleGroupSpec `json:"activated_rule"`
}

type AwsWafregionalRuleGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalRuleGroupList is a list of AwsWafregionalRuleGroups
type AwsWafregionalRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalRuleGroup CRD objects
	Items []AwsWafregionalRuleGroup `json:"items,omitempty"`
}
