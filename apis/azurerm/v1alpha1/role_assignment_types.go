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

type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec,omitempty"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

type RoleAssignmentSpec struct {
	// +optional
	Name        string `json:"name,omitempty" tf:"name,omitempty"`
	PrincipalID string `json:"principalID" tf:"principal_id"`
	// +optional
	RoleDefinitionID string `json:"roleDefinitionID,omitempty" tf:"role_definition_id,omitempty"`
	// +optional
	RoleDefinitionName string                    `json:"roleDefinitionName,omitempty" tf:"role_definition_name,omitempty"`
	Scope              string                    `json:"scope" tf:"scope"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RoleAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RoleAssignmentList is a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RoleAssignment CRD objects
	Items []RoleAssignment `json:"items,omitempty"`
}
