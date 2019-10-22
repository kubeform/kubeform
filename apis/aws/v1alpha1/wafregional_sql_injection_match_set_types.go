package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type WafregionalSQLInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalSQLInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalSQLInjectionMatchSetStatus `json:"status,omitempty"`
}

type WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	TextTransformation string                                                                  `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalSQLInjectionMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	SqlInjectionMatchTuple []WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple `json:"sqlInjectionMatchTuple,omitempty" tf:"sql_injection_match_tuple,omitempty"`
}

type WafregionalSQLInjectionMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafregionalSQLInjectionMatchSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalSQLInjectionMatchSetList is a list of WafregionalSQLInjectionMatchSets
type WafregionalSQLInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalSQLInjectionMatchSet CRD objects
	Items []WafregionalSQLInjectionMatchSet `json:"items,omitempty"`
}
