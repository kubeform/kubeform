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

type Stackscript struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackscriptSpec   `json:"spec,omitempty"`
	Status            StackscriptStatus `json:"status,omitempty"`
}

type StackscriptSpecUserDefinedFields struct{}

type StackscriptSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Description string                    `json:"description" tf:"description"`
	Images      []string                  `json:"images" tf:"images"`
	// +optional
	IsPublic bool   `json:"isPublic,omitempty" tf:"is_public,omitempty"`
	Label    string `json:"label" tf:"label"`
	// +optional
	RevNote string `json:"revNote,omitempty" tf:"rev_note,omitempty"`
	Script  string `json:"script" tf:"script"`
	// +optional
	UserDefinedFields []StackscriptSpecUserDefinedFields `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
}

type StackscriptStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StackscriptList is a list of Stackscripts
type StackscriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Stackscript CRD objects
	Items []Stackscript `json:"items,omitempty"`
}
