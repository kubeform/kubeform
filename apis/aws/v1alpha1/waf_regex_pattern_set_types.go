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

type WafRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRegexPatternSetSpec   `json:"spec,omitempty"`
	Status            WafRegexPatternSetStatus `json:"status,omitempty"`
}

type WafRegexPatternSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RegexPatternStrings []string `json:"regex_pattern_strings,omitempty"`
}

type WafRegexPatternSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRegexPatternSetList is a list of WafRegexPatternSets
type WafRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRegexPatternSet CRD objects
	Items []WafRegexPatternSet `json:"items,omitempty"`
}
