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

type AwsWafregionalSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalSizeConstraintSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalSizeConstraintSetStatus `json:"status,omitempty"`
}

type AwsWafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalSizeConstraintSetSpecSizeConstraints struct {
	TextTransformation string                                               `json:"text_transformation"`
	FieldToMatch       []AwsWafregionalSizeConstraintSetSpecSizeConstraints `json:"field_to_match"`
	ComparisonOperator string                                               `json:"comparison_operator"`
	Size               int                                                  `json:"size"`
}

type AwsWafregionalSizeConstraintSetSpec struct {
	Name            string                                `json:"name"`
	SizeConstraints []AwsWafregionalSizeConstraintSetSpec `json:"size_constraints"`
}

type AwsWafregionalSizeConstraintSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalSizeConstraintSetList is a list of AwsWafregionalSizeConstraintSets
type AwsWafregionalSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalSizeConstraintSet CRD objects
	Items []AwsWafregionalSizeConstraintSet `json:"items,omitempty"`
}
