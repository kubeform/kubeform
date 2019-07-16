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

type WafSqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafSqlInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            WafSqlInjectionMatchSetStatus `json:"status,omitempty"`
}

type WafSqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafSqlInjectionMatchSetSpecSqlInjectionMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafSqlInjectionMatchSetSpecSqlInjectionMatchTuples `json:"field_to_match"`
	TextTransformation string                                               `json:"text_transformation"`
}

type WafSqlInjectionMatchSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SqlInjectionMatchTuples *[]WafSqlInjectionMatchSetSpec `json:"sql_injection_match_tuples,omitempty"`
}

type WafSqlInjectionMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafSqlInjectionMatchSetList is a list of WafSqlInjectionMatchSets
type WafSqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafSqlInjectionMatchSet CRD objects
	Items []WafSqlInjectionMatchSet `json:"items,omitempty"`
}
