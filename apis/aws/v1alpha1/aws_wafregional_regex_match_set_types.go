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

type AwsWafregionalRegexMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalRegexMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalRegexMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalRegexMatchSetSpecRegexMatchTuple struct {
	TextTransformation string                                           `json:"text_transformation"`
	FieldToMatch       []AwsWafregionalRegexMatchSetSpecRegexMatchTuple `json:"field_to_match"`
	RegexPatternSetId  string                                           `json:"regex_pattern_set_id"`
}

type AwsWafregionalRegexMatchSetSpec struct {
	RegexMatchTuple []AwsWafregionalRegexMatchSetSpec `json:"regex_match_tuple"`
	Name            string                            `json:"name"`
}

type AwsWafregionalRegexMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalRegexMatchSetList is a list of AwsWafregionalRegexMatchSets
type AwsWafregionalRegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalRegexMatchSet CRD objects
	Items []AwsWafregionalRegexMatchSet `json:"items,omitempty"`
}
