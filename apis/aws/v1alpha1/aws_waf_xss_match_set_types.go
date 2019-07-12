package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafXssMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafXssMatchSetStatus `json:"status,omitempty"`
}

type AwsWafXssMatchSetSpecXssMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafXssMatchSetSpecXssMatchTuples struct {
	FieldToMatch       []AwsWafXssMatchSetSpecXssMatchTuples `json:"field_to_match"`
	TextTransformation string                                `json:"text_transformation"`
}

type AwsWafXssMatchSetSpec struct {
	Name           string                  `json:"name"`
	XssMatchTuples []AwsWafXssMatchSetSpec `json:"xss_match_tuples"`
}



type AwsWafXssMatchSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafXssMatchSetList is a list of AwsWafXssMatchSets
type AwsWafXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafXssMatchSet CRD objects
	Items []AwsWafXssMatchSet `json:"items,omitempty"`
}