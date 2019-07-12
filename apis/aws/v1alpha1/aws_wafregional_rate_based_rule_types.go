package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalRateBasedRuleStatus `json:"status,omitempty"`
}

type AwsWafregionalRateBasedRuleSpecPredicate struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

type AwsWafregionalRateBasedRuleSpec struct {
	RateLimit  int                               `json:"rate_limit"`
	Name       string                            `json:"name"`
	MetricName string                            `json:"metric_name"`
	Predicate  []AwsWafregionalRateBasedRuleSpec `json:"predicate"`
	RateKey    string                            `json:"rate_key"`
}



type AwsWafregionalRateBasedRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRateBasedRuleList is a list of AwsWafregionalRateBasedRules
type AwsWafregionalRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalRateBasedRule CRD objects
	Items []AwsWafregionalRateBasedRule `json:"items,omitempty"`
}