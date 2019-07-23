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

type WafRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRuleSpec   `json:"spec,omitempty"`
	Status            WafRuleStatus `json:"status,omitempty"`
}

type WafRuleSpecPredicates struct {
	DataID  string `json:"dataID" tf:"data_id"`
	Negated bool   `json:"negated" tf:"negated"`
	Type    string `json:"type" tf:"type"`
}

type WafRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	MetricName string `json:"metricName" tf:"metric_name"`
	Name       string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicates []WafRuleSpecPredicates `json:"predicates,omitempty" tf:"predicates,omitempty"`
}

type WafRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRuleList is a list of WafRules
type WafRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRule CRD objects
	Items []WafRule `json:"items,omitempty"`
}
