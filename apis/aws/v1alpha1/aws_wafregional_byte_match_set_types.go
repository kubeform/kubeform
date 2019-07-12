package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalByteMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalByteMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalByteMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTuples struct {
	FieldToMatch         []AwsWafregionalByteMatchSetSpecByteMatchTuples `json:"field_to_match"`
	PositionalConstraint string                                          `json:"positional_constraint"`
	TargetString         string                                          `json:"target_string"`
	TextTransformation   string                                          `json:"text_transformation"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTuple struct {
	FieldToMatch         []AwsWafregionalByteMatchSetSpecByteMatchTuple `json:"field_to_match"`
	PositionalConstraint string                                         `json:"positional_constraint"`
	TargetString         string                                         `json:"target_string"`
	TextTransformation   string                                         `json:"text_transformation"`
}

type AwsWafregionalByteMatchSetSpec struct {
	ByteMatchTuples []AwsWafregionalByteMatchSetSpec `json:"byte_match_tuples"`
	Name            string                           `json:"name"`
	ByteMatchTuple  []AwsWafregionalByteMatchSetSpec `json:"byte_match_tuple"`
}

type AwsWafregionalByteMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalByteMatchSetList is a list of AwsWafregionalByteMatchSets
type AwsWafregionalByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalByteMatchSet CRD objects
	Items []AwsWafregionalByteMatchSet `json:"items,omitempty"`
}
