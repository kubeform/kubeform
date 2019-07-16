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

type WafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            WafRateBasedRuleStatus `json:"status,omitempty"`
}

type WafRateBasedRuleSpecPredicates struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

type WafRateBasedRuleSpec struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicates *[]WafRateBasedRuleSpec `json:"predicates,omitempty"`
	RateKey    string                  `json:"rate_key"`
	RateLimit  int                     `json:"rate_limit"`
}

type WafRateBasedRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRateBasedRuleList is a list of WafRateBasedRules
type WafRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRateBasedRule CRD objects
	Items []WafRateBasedRule `json:"items,omitempty"`
}
