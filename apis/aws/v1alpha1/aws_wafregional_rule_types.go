package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalRuleSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalRuleStatus `json:"status,omitempty"`
}

type AwsWafregionalRuleSpecPredicate struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

type AwsWafregionalRuleSpec struct {
	Name       string                   `json:"name"`
	MetricName string                   `json:"metric_name"`
	Predicate  []AwsWafregionalRuleSpec `json:"predicate"`
}



type AwsWafregionalRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRuleList is a list of AwsWafregionalRules
type AwsWafregionalRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalRule CRD objects
	Items []AwsWafregionalRule `json:"items,omitempty"`
}