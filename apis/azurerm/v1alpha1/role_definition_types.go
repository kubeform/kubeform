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

type RoleDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleDefinitionSpec   `json:"spec,omitempty"`
	Status            RoleDefinitionStatus `json:"status,omitempty"`
}

type RoleDefinitionSpecPermissions struct {
	// +optional
	Actions []string `json:"actions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DataActions []string `json:"data_actions,omitempty"`
	// +optional
	NotActions []string `json:"not_actions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NotDataActions []string `json:"not_data_actions,omitempty"`
}

type RoleDefinitionSpec struct {
	AssignableScopes []string `json:"assignable_scopes"`
	// +optional
	Description string               `json:"description,omitempty"`
	Name        string               `json:"name"`
	Permissions []RoleDefinitionSpec `json:"permissions"`
	Scope       string               `json:"scope"`
}

type RoleDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RoleDefinitionList is a list of RoleDefinitions
type RoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RoleDefinition CRD objects
	Items []RoleDefinition `json:"items,omitempty"`
}
