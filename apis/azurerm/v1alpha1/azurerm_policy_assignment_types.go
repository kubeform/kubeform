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

type AzurermPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPolicyAssignmentSpec   `json:"spec,omitempty"`
	Status            AzurermPolicyAssignmentStatus `json:"status,omitempty"`
}

type AzurermPolicyAssignmentSpecIdentity struct {
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
}

type AzurermPolicyAssignmentSpec struct {
	Scope              string                        `json:"scope"`
	PolicyDefinitionId string                        `json:"policy_definition_id"`
	Description        string                        `json:"description"`
	Location           string                        `json:"location"`
	NotScopes          []string                      `json:"not_scopes"`
	Name               string                        `json:"name"`
	DisplayName        string                        `json:"display_name"`
	Identity           []AzurermPolicyAssignmentSpec `json:"identity"`
	Parameters         string                        `json:"parameters"`
}



type AzurermPolicyAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPolicyAssignmentList is a list of AzurermPolicyAssignments
type AzurermPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPolicyAssignment CRD objects
	Items []AzurermPolicyAssignment `json:"items,omitempty"`
}