package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type WafregionalRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRuleSpec   `json:"spec,omitempty"`
	Status            WafregionalRuleStatus `json:"status,omitempty"`
}

type WafregionalRuleSpecPredicate struct {
	DataID  string `json:"dataID" tf:"data_id"`
	Negated bool   `json:"negated" tf:"negated"`
	Type    string `json:"type" tf:"type"`
}

type WafregionalRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	MetricName string `json:"metricName" tf:"metric_name"`
	Name       string `json:"name" tf:"name"`
	// +optional
	Predicate []WafregionalRuleSpecPredicate `json:"predicate,omitempty" tf:"predicate,omitempty"`
}

type WafregionalRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafregionalRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalRuleList is a list of WafregionalRules
type WafregionalRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalRule CRD objects
	Items []WafregionalRule `json:"items,omitempty"`
}
