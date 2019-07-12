package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalSqlInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalSqlInjectionMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	TextTransformation string                                                         `json:"text_transformation"`
	FieldToMatch       []AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"field_to_match"`
}

type AwsWafregionalSqlInjectionMatchSetSpec struct {
	SqlInjectionMatchTuple []AwsWafregionalSqlInjectionMatchSetSpec `json:"sql_injection_match_tuple"`
	Name                   string                                   `json:"name"`
}



type AwsWafregionalSqlInjectionMatchSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSqlInjectionMatchSetList is a list of AwsWafregionalSqlInjectionMatchSets
type AwsWafregionalSqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalSqlInjectionMatchSet CRD objects
	Items []AwsWafregionalSqlInjectionMatchSet `json:"items,omitempty"`
}