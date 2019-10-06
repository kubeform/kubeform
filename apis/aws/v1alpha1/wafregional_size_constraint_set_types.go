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

type WafregionalSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalSizeConstraintSetSpec   `json:"spec,omitempty"`
	Status            WafregionalSizeConstraintSetStatus `json:"status,omitempty"`
}

type WafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafregionalSizeConstraintSetSpecSizeConstraints struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	Size               int                                                           `json:"size" tf:"size"`
	TextTransformation string                                                        `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalSizeConstraintSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	SizeConstraints []WafregionalSizeConstraintSetSpecSizeConstraints `json:"sizeConstraints,omitempty" tf:"size_constraints,omitempty"`
}

type WafregionalSizeConstraintSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafregionalSizeConstraintSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalSizeConstraintSetList is a list of WafregionalSizeConstraintSets
type WafregionalSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalSizeConstraintSet CRD objects
	Items []WafregionalSizeConstraintSet `json:"items,omitempty"`
}
