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

type AwsWafregionalRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalRegexPatternSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalRegexPatternSetStatus `json:"status,omitempty"`
}

type AwsWafregionalRegexPatternSetSpec struct {
	Name                string   `json:"name"`
	RegexPatternStrings []string `json:"regex_pattern_strings"`
}

type AwsWafregionalRegexPatternSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalRegexPatternSetList is a list of AwsWafregionalRegexPatternSets
type AwsWafregionalRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalRegexPatternSet CRD objects
	Items []AwsWafregionalRegexPatternSet `json:"items,omitempty"`
}
