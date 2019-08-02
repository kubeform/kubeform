package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PolicyDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyDefinitionSpec   `json:"spec,omitempty"`
	Status            PolicyDefinitionStatus `json:"status,omitempty"`
}

type PolicyDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +optional
	ManagementGroupID string `json:"managementGroupID,omitempty" tf:"management_group_id,omitempty"`
	// +optional
	Metadata string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	Mode     string `json:"mode" tf:"mode"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Parameters string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	PolicyRule string `json:"policyRule,omitempty" tf:"policy_rule,omitempty"`
	PolicyType string `json:"policyType" tf:"policy_type"`
}

type PolicyDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *PolicyDefinitionSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PolicyDefinitionList is a list of PolicyDefinitions
type PolicyDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyDefinition CRD objects
	Items []PolicyDefinition `json:"items,omitempty"`
}
