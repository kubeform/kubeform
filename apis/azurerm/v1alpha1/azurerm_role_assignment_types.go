package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRoleAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRoleAssignmentSpec   `json:"spec,omitempty"`
	Status            AzurermRoleAssignmentStatus `json:"status,omitempty"`
}

type AzurermRoleAssignmentSpec struct {
	Name               string `json:"name"`
	Scope              string `json:"scope"`
	RoleDefinitionId   string `json:"role_definition_id"`
	RoleDefinitionName string `json:"role_definition_name"`
	PrincipalId        string `json:"principal_id"`
}

type AzurermRoleAssignmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRoleAssignmentList is a list of AzurermRoleAssignments
type AzurermRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRoleAssignment CRD objects
	Items []AzurermRoleAssignment `json:"items,omitempty"`
}
