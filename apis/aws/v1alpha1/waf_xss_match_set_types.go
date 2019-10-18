package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type WafXssMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafXssMatchSetSpec   `json:"spec,omitempty"`
	Status            WafXssMatchSetStatus `json:"status,omitempty"`
}

type WafXssMatchSetSpecXssMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafXssMatchSetSpecXssMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafXssMatchSetSpecXssMatchTuplesFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	TextTransformation string                                         `json:"textTransformation" tf:"text_transformation"`
}

type WafXssMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	XssMatchTuples []WafXssMatchSetSpecXssMatchTuples `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples,omitempty"`
}

type WafXssMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafXssMatchSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafXssMatchSetList is a list of WafXssMatchSets
type WafXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafXssMatchSet CRD objects
	Items []WafXssMatchSet `json:"items,omitempty"`
}
