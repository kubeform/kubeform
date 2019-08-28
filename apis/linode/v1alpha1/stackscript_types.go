package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
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

type StackscriptSpecUserDefinedFields struct {
	// +optional
	Default string `json:"default,omitempty" tf:"default,omitempty"`
	// +optional
	Example string `json:"example,omitempty" tf:"example,omitempty"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// +optional
	ManyOf string `json:"manyOf,omitempty" tf:"many_of,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OneOf string `json:"oneOf,omitempty" tf:"one_of,omitempty"`
}

type StackscriptSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Created string `json:"created,omitempty" tf:"created,omitempty"`
	// +optional
	DeploymentsActive int `json:"deploymentsActive,omitempty" tf:"deployments_active,omitempty"`
	// +optional
	DeploymentsTotal int      `json:"deploymentsTotal,omitempty" tf:"deployments_total,omitempty"`
	Description      string   `json:"description" tf:"description"`
	Images           []string `json:"images" tf:"images"`
	// +optional
	IsPublic bool   `json:"isPublic,omitempty" tf:"is_public,omitempty"`
	Label    string `json:"label" tf:"label"`
	// +optional
	RevNote string `json:"revNote,omitempty" tf:"rev_note,omitempty"`
	Script  string `json:"script" tf:"script"`
	// +optional
	Updated string `json:"updated,omitempty" tf:"updated,omitempty"`
	// +optional
	UserDefinedFields []StackscriptSpecUserDefinedFields `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
	// +optional
	UserGravatarID string `json:"userGravatarID,omitempty" tf:"user_gravatar_id,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}



type StackscriptStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *StackscriptSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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