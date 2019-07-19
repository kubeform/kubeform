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

type WafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            WafRateBasedRuleStatus `json:"status,omitempty"`
}

type WafRateBasedRuleSpecPredicates struct {
	DataID  string `json:"dataID" tf:"data_id"`
	Negated bool   `json:"negated" tf:"negated"`
	Type    string `json:"type" tf:"type"`
}

type WafRateBasedRuleSpec struct {
	MetricName string `json:"metricName" tf:"metric_name"`
	Name       string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicates  []WafRateBasedRuleSpecPredicates `json:"predicates,omitempty" tf:"predicates,omitempty"`
	RateKey     string                           `json:"rateKey" tf:"rate_key"`
	RateLimit   int                              `json:"rateLimit" tf:"rate_limit"`
	ProviderRef core.LocalObjectReference        `json:"providerRef" tf:"-"`
}

type WafRateBasedRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
