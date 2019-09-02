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

	// The date this StackScript was created.
	// +optional
	Created string `json:"created,omitempty" tf:"created,omitempty"`
	// Count of currently active, deployed Linodes created from this StackScript.
	// +optional
	DeploymentsActive int `json:"deploymentsActive,omitempty" tf:"deployments_active,omitempty"`
	// The total number of times this StackScript has been deployed.
	// +optional
	DeploymentsTotal int `json:"deploymentsTotal,omitempty" tf:"deployments_total,omitempty"`
	// A description for the StackScript.
	Description string `json:"description" tf:"description"`
	// An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
	Images []string `json:"images" tf:"images"`
	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.
	// +optional
	IsPublic bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`
	// The StackScript's label is for display purposes only.
	Label string `json:"label" tf:"label"`
	// This field allows you to add notes for the set of revisions made to this StackScript.
	// +optional
	RevNote string `json:"revNote,omitempty" tf:"rev_note,omitempty"`
	// The script to execute when provisioning a new Linode with this StackScript.
	Script string `json:"script" tf:"script"`
	// The date this StackScript was updated.
	// +optional
	Updated string `json:"updated,omitempty" tf:"updated,omitempty"`
	// This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.
	// +optional
	UserDefinedFields []StackscriptSpecUserDefinedFields `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
	// The Gravatar ID for the User who created the StackScript.
	// +optional
	UserGravatarID string `json:"userGravatarID,omitempty" tf:"user_gravatar_id,omitempty"`
	// The User who created the StackScript.
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