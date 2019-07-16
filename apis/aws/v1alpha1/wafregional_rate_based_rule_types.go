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

type WafregionalRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            WafregionalRateBasedRuleStatus `json:"status,omitempty"`
}

type WafregionalRateBasedRuleSpecPredicate struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

type WafregionalRateBasedRuleSpec struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicate *[]WafregionalRateBasedRuleSpec `json:"predicate,omitempty"`
	RateKey   string                          `json:"rate_key"`
	RateLimit int                             `json:"rate_limit"`
}

type WafregionalRateBasedRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalRateBasedRuleList is a list of WafregionalRateBasedRules
type WafregionalRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalRateBasedRule CRD objects
	Items []WafregionalRateBasedRule `json:"items,omitempty"`
}
