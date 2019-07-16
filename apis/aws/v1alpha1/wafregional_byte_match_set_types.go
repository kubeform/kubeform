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

type WafregionalByteMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalByteMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalByteMatchSetStatus `json:"status,omitempty"`
}

type WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafregionalByteMatchSetSpecByteMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch         []WafregionalByteMatchSetSpecByteMatchTuples `json:"field_to_match"`
	PositionalConstraint string                                       `json:"positional_constraint"`
	// +optional
	TargetString       string `json:"target_string,omitempty"`
	TextTransformation string `json:"text_transformation"`
}

type WafregionalByteMatchSetSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ByteMatchTuples *[]WafregionalByteMatchSetSpec `json:"byte_match_tuples,omitempty"`
	Name            string                         `json:"name"`
}

type WafregionalByteMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalByteMatchSetList is a list of WafregionalByteMatchSets
type WafregionalByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalByteMatchSet CRD objects
	Items []WafregionalByteMatchSet `json:"items,omitempty"`
}
