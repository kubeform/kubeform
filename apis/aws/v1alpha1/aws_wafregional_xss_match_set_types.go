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

type AwsWafregionalXssMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalXssMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalXssMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalXssMatchSetSpecXssMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalXssMatchSetSpecXssMatchTuple struct {
	FieldToMatch       []AwsWafregionalXssMatchSetSpecXssMatchTuple `json:"field_to_match"`
	TextTransformation string                                       `json:"text_transformation"`
}

type AwsWafregionalXssMatchSetSpec struct {
	Name          string                          `json:"name"`
	XssMatchTuple []AwsWafregionalXssMatchSetSpec `json:"xss_match_tuple"`
}

type AwsWafregionalXssMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalXssMatchSetList is a list of AwsWafregionalXssMatchSets
type AwsWafregionalXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalXssMatchSet CRD objects
	Items []AwsWafregionalXssMatchSet `json:"items,omitempty"`
}
