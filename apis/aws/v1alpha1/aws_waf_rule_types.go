package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRuleSpec   `json:"spec,omitempty"`
	Status            AwsWafRuleStatus `json:"status,omitempty"`
}

type AwsWafRuleSpecPredicates struct {
	Type    string `json:"type"`
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
}

type AwsWafRuleSpec struct {
	MetricName string           `json:"metric_name"`
	Predicates []AwsWafRuleSpec `json:"predicates"`
	Name       string           `json:"name"`
}



type AwsWafRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRuleList is a list of AwsWafRules
type AwsWafRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRule CRD objects
	Items []AwsWafRule `json:"items,omitempty"`
}