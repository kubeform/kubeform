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

type WafSQLInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafSQLInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            WafSQLInjectionMatchSetStatus `json:"status,omitempty"`
}

type WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	TextTransformation string                                                           `json:"textTransformation" tf:"text_transformation"`
}

type WafSQLInjectionMatchSetSpec struct {
	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SqlInjectionMatchTuples []WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples `json:"sqlInjectionMatchTuples,omitempty" tf:"sql_injection_match_tuples,omitempty"`
	ProviderRef             core.LocalObjectReference                            `json:"providerRef" tf:"-"`
}

type WafSQLInjectionMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafSQLInjectionMatchSetList is a list of WafSQLInjectionMatchSets
type WafSQLInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafSQLInjectionMatchSet CRD objects
	Items []WafSQLInjectionMatchSet `json:"items,omitempty"`
}
