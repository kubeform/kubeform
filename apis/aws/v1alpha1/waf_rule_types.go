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

type WafRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRuleSpec   `json:"spec,omitempty"`
	Status            WafRuleStatus `json:"status,omitempty"`
}

type WafRuleSpecPredicates struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

type WafRuleSpec struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Predicates *[]WafRuleSpec `json:"predicates,omitempty"`
}

type WafRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
