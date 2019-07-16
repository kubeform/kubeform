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

type WafRegexMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRegexMatchSetSpec   `json:"spec,omitempty"`
	Status            WafRegexMatchSetStatus `json:"status,omitempty"`
}

type WafRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafRegexMatchSetSpecRegexMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafRegexMatchSetSpecRegexMatchTuple `json:"field_to_match"`
	RegexPatternSetId  string                                `json:"regex_pattern_set_id"`
	TextTransformation string                                `json:"text_transformation"`
}

type WafRegexMatchSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RegexMatchTuple *[]WafRegexMatchSetSpec `json:"regex_match_tuple,omitempty"`
}

type WafRegexMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRegexMatchSetList is a list of WafRegexMatchSets
type WafRegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRegexMatchSet CRD objects
	Items []WafRegexMatchSet `json:"items,omitempty"`
}
