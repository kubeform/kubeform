package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafregionalXssMatchSetSpecXssMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafregionalXssMatchSetSpecXssMatchTupleFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	TextTransformation string                                                `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalXssMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	XssMatchTuple []WafregionalXssMatchSetSpecXssMatchTuple `json:"xssMatchTuple,omitempty" tf:"xss_match_tuple,omitempty"`
}



type WafregionalXssMatchSetStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *WafregionalXssMatchSetSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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