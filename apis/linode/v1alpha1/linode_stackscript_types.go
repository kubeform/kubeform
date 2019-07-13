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

type LinodeStackscript struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeStackscriptSpec   `json:"spec,omitempty"`
	Status            LinodeStackscriptStatus `json:"status,omitempty"`
}

type LinodeStackscriptSpecUserDefinedFields struct {
	Example string `json:"example"`
	OneOf   string `json:"one_of"`
	ManyOf  string `json:"many_of"`
	Default string `json:"default"`
	Label   string `json:"label"`
	Name    string `json:"name"`
}

type LinodeStackscriptSpec struct {
	Updated           string                  `json:"updated"`
	UserDefinedFields []LinodeStackscriptSpec `json:"user_defined_fields"`
	Script            string                  `json:"script"`
	RevNote           string                  `json:"rev_note"`
	Username          string                  `json:"username"`
	Images            []string                `json:"images"`
	DeploymentsActive int                     `json:"deployments_active"`
	UserGravatarId    string                  `json:"user_gravatar_id"`
	DeploymentsTotal  int                     `json:"deployments_total"`
	Created           string                  `json:"created"`
	Label             string                  `json:"label"`
	Description       string                  `json:"description"`
	IsPublic          bool                    `json:"is_public"`
}



type LinodeStackscriptStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeStackscriptList is a list of LinodeStackscripts
type LinodeStackscriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeStackscript CRD objects
	Items []LinodeStackscript `json:"items,omitempty"`
}