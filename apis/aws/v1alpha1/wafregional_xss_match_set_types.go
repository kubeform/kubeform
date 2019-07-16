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

type WafregionalXssMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalXssMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalXssMatchSetStatus `json:"status,omitempty"`
}

type WafregionalXssMatchSetSpecXssMatchTupleFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafregionalXssMatchSetSpecXssMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafregionalXssMatchSetSpecXssMatchTuple `json:"field_to_match"`
	TextTransformation string                                    `json:"text_transformation"`
}

type WafregionalXssMatchSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	XssMatchTuple *[]WafregionalXssMatchSetSpec `json:"xss_match_tuple,omitempty"`
}

type WafregionalXssMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalXssMatchSetList is a list of WafregionalXssMatchSets
type WafregionalXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalXssMatchSet CRD objects
	Items []WafregionalXssMatchSet `json:"items,omitempty"`
}
