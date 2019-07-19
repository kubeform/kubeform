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

type WafSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafSizeConstraintSetSpec   `json:"spec,omitempty"`
	Status            WafSizeConstraintSetStatus `json:"status,omitempty"`
}

type WafSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafSizeConstraintSetSpecSizeConstraints struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafSizeConstraintSetSpecSizeConstraintsFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	Size               int                                                   `json:"size" tf:"size"`
	TextTransformation string                                                `json:"textTransformation" tf:"text_transformation"`
}

type WafSizeConstraintSetSpec struct {
	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SizeConstraints []WafSizeConstraintSetSpecSizeConstraints `json:"sizeConstraints,omitempty" tf:"size_constraints,omitempty"`
	ProviderRef     core.LocalObjectReference                 `json:"providerRef" tf:"-"`
}

type WafSizeConstraintSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafSizeConstraintSetList is a list of WafSizeConstraintSets
type WafSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafSizeConstraintSet CRD objects
	Items []WafSizeConstraintSet `json:"items,omitempty"`
}
