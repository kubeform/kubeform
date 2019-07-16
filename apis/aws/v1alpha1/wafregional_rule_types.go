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

type WafregionalRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRuleSpec   `json:"spec,omitempty"`
	Status            WafregionalRuleStatus `json:"status,omitempty"`
}

type WafregionalRuleSpecPredicate struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

type WafregionalRuleSpec struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicate *[]WafregionalRuleSpec `json:"predicate,omitempty"`
}

type WafregionalRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
