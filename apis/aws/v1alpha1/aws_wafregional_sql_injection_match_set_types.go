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

type AwsWafregionalSqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalSqlInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalSqlInjectionMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	FieldToMatch       []AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"field_to_match"`
	TextTransformation string                                                         `json:"text_transformation"`
}

type AwsWafregionalSqlInjectionMatchSetSpec struct {
	Name                   string                                   `json:"name"`
	SqlInjectionMatchTuple []AwsWafregionalSqlInjectionMatchSetSpec `json:"sql_injection_match_tuple"`
}



type AwsWafregionalSqlInjectionMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalSqlInjectionMatchSetList is a list of AwsWafregionalSqlInjectionMatchSets
type AwsWafregionalSqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalSqlInjectionMatchSet CRD objects
	Items []AwsWafregionalSqlInjectionMatchSet `json:"items,omitempty"`
}