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

type WafByteMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafByteMatchSetSpec   `json:"spec,omitempty"`
	Status            WafByteMatchSetStatus `json:"status,omitempty"`
}

type WafByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafByteMatchSetSpecByteMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch         []WafByteMatchSetSpecByteMatchTuples `json:"field_to_match"`
	PositionalConstraint string                               `json:"positional_constraint"`
	// +optional
	TargetString       string `json:"target_string,omitempty"`
	TextTransformation string `json:"text_transformation"`
}

type WafByteMatchSetSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ByteMatchTuples *[]WafByteMatchSetSpec `json:"byte_match_tuples,omitempty"`
	Name            string                 `json:"name"`
}

type WafByteMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafByteMatchSetList is a list of WafByteMatchSets
type WafByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafByteMatchSet CRD objects
	Items []WafByteMatchSet `json:"items,omitempty"`
}
