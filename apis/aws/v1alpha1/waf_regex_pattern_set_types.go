package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RegexPatternStrings []string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings,omitempty"`
}

type WafRegexPatternSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
