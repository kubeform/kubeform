package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPolicyAssignmentSpec   `json:"spec,omitempty"`
	Status            AzurermPolicyAssignmentStatus `json:"status,omitempty"`
}

type AzurermPolicyAssignmentSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermPolicyAssignmentSpec struct {
	Scope              string                        `json:"scope"`
	Description        string                        `json:"description"`
	DisplayName        string                        `json:"display_name"`
	Identity           []AzurermPolicyAssignmentSpec `json:"identity"`
	Name               string                        `json:"name"`
	PolicyDefinitionId string                        `json:"policy_definition_id"`
	Location           string                        `json:"location"`
	Parameters         string                        `json:"parameters"`
	NotScopes          []string                      `json:"not_scopes"`
}

type AzurermPolicyAssignmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPolicyAssignmentList is a list of AzurermPolicyAssignments
type AzurermPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPolicyAssignment CRD objects
	Items []AzurermPolicyAssignment `json:"items,omitempty"`
}
