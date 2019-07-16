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

type WafregionalSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalSizeConstraintSetSpec   `json:"spec,omitempty"`
	Status            WafregionalSizeConstraintSetStatus `json:"status,omitempty"`
}

type WafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafregionalSizeConstraintSetSpecSizeConstraints struct {
	ComparisonOperator string `json:"comparison_operator"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch       []WafregionalSizeConstraintSetSpecSizeConstraints `json:"field_to_match"`
	Size               int                                               `json:"size"`
	TextTransformation string                                            `json:"text_transformation"`
}

type WafregionalSizeConstraintSetSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SizeConstraints *[]WafregionalSizeConstraintSetSpec `json:"size_constraints,omitempty"`
}

type WafregionalSizeConstraintSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
