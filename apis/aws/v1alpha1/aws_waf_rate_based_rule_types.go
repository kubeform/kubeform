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

type AwsWafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            AwsWafRateBasedRuleStatus `json:"status,omitempty"`
}

type AwsWafRateBasedRuleSpecPredicates struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

type AwsWafRateBasedRuleSpec struct {
	MetricName string                    `json:"metric_name"`
	Predicates []AwsWafRateBasedRuleSpec `json:"predicates"`
	RateKey    string                    `json:"rate_key"`
	RateLimit  int                       `json:"rate_limit"`
	Name       string                    `json:"name"`
}



type AwsWafRateBasedRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafRateBasedRuleList is a list of AwsWafRateBasedRules
type AwsWafRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRateBasedRule CRD objects
	Items []AwsWafRateBasedRule `json:"items,omitempty"`
}