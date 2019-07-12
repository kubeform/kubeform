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

type AwsWafSqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafSqlInjectionMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafSqlInjectionMatchSetStatus `json:"status,omitempty"`
}

type AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuples struct {
	FieldToMatch       []AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuples `json:"field_to_match"`
	TextTransformation string                                                  `json:"text_transformation"`
}

type AwsWafSqlInjectionMatchSetSpec struct {
	Name                    string                           `json:"name"`
	SqlInjectionMatchTuples []AwsWafSqlInjectionMatchSetSpec `json:"sql_injection_match_tuples"`
}

type AwsWafSqlInjectionMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafSqlInjectionMatchSetList is a list of AwsWafSqlInjectionMatchSets
type AwsWafSqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafSqlInjectionMatchSet CRD objects
	Items []AwsWafSqlInjectionMatchSet `json:"items,omitempty"`
}
