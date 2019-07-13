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

type AwsWafByteMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafByteMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafByteMatchSetStatus `json:"status,omitempty"`
}

type AwsWafByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafByteMatchSetSpecByteMatchTuples struct {
	FieldToMatch         []AwsWafByteMatchSetSpecByteMatchTuples `json:"field_to_match"`
	PositionalConstraint string                                  `json:"positional_constraint"`
	TargetString         string                                  `json:"target_string"`
	TextTransformation   string                                  `json:"text_transformation"`
}

type AwsWafByteMatchSetSpec struct {
	Name            string                   `json:"name"`
	ByteMatchTuples []AwsWafByteMatchSetSpec `json:"byte_match_tuples"`
}



type AwsWafByteMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafByteMatchSetList is a list of AwsWafByteMatchSets
type AwsWafByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafByteMatchSet CRD objects
	Items []AwsWafByteMatchSet `json:"items,omitempty"`
}