package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafSizeConstraintSetSpec   `json:"spec,omitempty"`
	Status            AwsWafSizeConstraintSetStatus `json:"status,omitempty"`
}

type AwsWafSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafSizeConstraintSetSpecSizeConstraints struct {
	FieldToMatch       []AwsWafSizeConstraintSetSpecSizeConstraints `json:"field_to_match"`
	ComparisonOperator string                                       `json:"comparison_operator"`
	Size               int                                          `json:"size"`
	TextTransformation string                                       `json:"text_transformation"`
}

type AwsWafSizeConstraintSetSpec struct {
	Name            string                        `json:"name"`
	SizeConstraints []AwsWafSizeConstraintSetSpec `json:"size_constraints"`
}

type AwsWafSizeConstraintSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSizeConstraintSetList is a list of AwsWafSizeConstraintSets
type AwsWafSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafSizeConstraintSet CRD objects
	Items []AwsWafSizeConstraintSet `json:"items,omitempty"`
}
