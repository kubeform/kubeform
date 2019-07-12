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
	ManyOf  string `json:"many_of"`
	Default string `json:"default"`
	Label   string `json:"label"`
	Name    string `json:"name"`
	Example string `json:"example"`
	OneOf   string `json:"one_of"`
}

type LinodeStackscriptSpec struct {
	Description       string                  `json:"description"`
	DeploymentsTotal  int                     `json:"deployments_total"`
	Username          string                  `json:"username"`
	Created           string                  `json:"created"`
	UserDefinedFields []LinodeStackscriptSpec `json:"user_defined_fields"`
	UserGravatarId    string                  `json:"user_gravatar_id"`
	Updated           string                  `json:"updated"`
	Label             string                  `json:"label"`
	Script            string                  `json:"script"`
	RevNote           string                  `json:"rev_note"`
	IsPublic          bool                    `json:"is_public"`
	Images            []string                `json:"images"`
	DeploymentsActive int                     `json:"deployments_active"`
}

type LinodeStackscriptStatus struct {
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
