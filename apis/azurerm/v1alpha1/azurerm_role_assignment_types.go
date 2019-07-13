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

type AzurermRoleAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRoleAssignmentSpec   `json:"spec,omitempty"`
	Status            AzurermRoleAssignmentStatus `json:"status,omitempty"`
}

type AzurermRoleAssignmentSpec struct {
	PrincipalId        string `json:"principal_id"`
	Name               string `json:"name"`
	Scope              string `json:"scope"`
	RoleDefinitionId   string `json:"role_definition_id"`
	RoleDefinitionName string `json:"role_definition_name"`
}



type AzurermRoleAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermRoleAssignmentList is a list of AzurermRoleAssignments
type AzurermRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRoleAssignment CRD objects
	Items []AzurermRoleAssignment `json:"items,omitempty"`
}