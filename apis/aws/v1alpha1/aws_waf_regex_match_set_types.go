package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRegexMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafRegexMatchSetStatus `json:"status,omitempty"`
}

type AwsWafRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafRegexMatchSetSpecRegexMatchTuple struct {
	FieldToMatch       []AwsWafRegexMatchSetSpecRegexMatchTuple `json:"field_to_match"`
	RegexPatternSetId  string                                   `json:"regex_pattern_set_id"`
	TextTransformation string                                   `json:"text_transformation"`
}

type AwsWafRegexMatchSetSpec struct {
	Name            string                    `json:"name"`
	RegexMatchTuple []AwsWafRegexMatchSetSpec `json:"regex_match_tuple"`
}

type AwsWafRegexMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexMatchSetList is a list of AwsWafRegexMatchSets
type AwsWafRegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRegexMatchSet CRD objects
	Items []AwsWafRegexMatchSet `json:"items,omitempty"`
}
