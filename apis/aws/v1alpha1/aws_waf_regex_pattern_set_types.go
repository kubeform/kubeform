package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRegexPatternSetSpec   `json:"spec,omitempty"`
	Status            AwsWafRegexPatternSetStatus `json:"status,omitempty"`
}

type AwsWafRegexPatternSetSpec struct {
	Name                string   `json:"name"`
	RegexPatternStrings []string `json:"regex_pattern_strings"`
}



type AwsWafRegexPatternSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexPatternSetList is a list of AwsWafRegexPatternSets
type AwsWafRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRegexPatternSet CRD objects
	Items []AwsWafRegexPatternSet `json:"items,omitempty"`
}