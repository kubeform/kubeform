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

type WafregionalSqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalSqlInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalSqlInjectionMatchSetStatus `json:"status,omitempty"`
}

type WafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"field_to_match"`
	TextTransformation string                                                      `json:"text_transformation"`
}

type WafregionalSqlInjectionMatchSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SqlInjectionMatchTuple *[]WafregionalSqlInjectionMatchSetSpec `json:"sql_injection_match_tuple,omitempty"`
}

type WafregionalSqlInjectionMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalSqlInjectionMatchSetList is a list of WafregionalSqlInjectionMatchSets
type WafregionalSqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalSqlInjectionMatchSet CRD objects
	Items []WafregionalSqlInjectionMatchSet `json:"items,omitempty"`
}
