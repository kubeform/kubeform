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

type AwsWafRuleGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRuleGroupSpec   `json:"spec,omitempty"`
	Status            AwsWafRuleGroupStatus `json:"status,omitempty"`
}

type AwsWafRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

type AwsWafRuleGroupSpecActivatedRule struct {
	RuleId   string                             `json:"rule_id"`
	Type     string                             `json:"type"`
	Action   []AwsWafRuleGroupSpecActivatedRule `json:"action"`
	Priority int                                `json:"priority"`
}

type AwsWafRuleGroupSpec struct {
	ActivatedRule []AwsWafRuleGroupSpec `json:"activated_rule"`
	Name          string                `json:"name"`
	MetricName    string                `json:"metric_name"`
}



type AwsWafRuleGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafRuleGroupList is a list of AwsWafRuleGroups
type AwsWafRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRuleGroup CRD objects
	Items []AwsWafRuleGroup `json:"items,omitempty"`
}