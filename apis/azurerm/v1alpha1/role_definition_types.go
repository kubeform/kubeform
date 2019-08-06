package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Actions []string `json:"actions,omitempty" tf:"actions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DataActions []string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`
	// +optional
	NotActions []string `json:"notActions,omitempty" tf:"not_actions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NotDataActions []string `json:"notDataActions,omitempty" tf:"not_data_actions,omitempty"`
}

type RoleDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AssignableScopes []string `json:"assignableScopes" tf:"assignable_scopes"`
	// +optional
	Description string                          `json:"description,omitempty" tf:"description,omitempty"`
	Name        string                          `json:"name" tf:"name"`
	Permissions []RoleDefinitionSpecPermissions `json:"permissions" tf:"permissions"`
	// +optional
	RoleDefinitionID string `json:"roleDefinitionID,omitempty" tf:"role_definition_id,omitempty"`
	Scope            string `json:"scope" tf:"scope"`
}

type RoleDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RoleDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
