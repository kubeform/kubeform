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

type AzurermPolicySetDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPolicySetDefinitionSpec   `json:"spec,omitempty"`
	Status            AzurermPolicySetDefinitionStatus `json:"status,omitempty"`
}

type AzurermPolicySetDefinitionSpec struct {
	Description       string `json:"description"`
	Metadata          string `json:"metadata"`
	Parameters        string `json:"parameters"`
	PolicyDefinitions string `json:"policy_definitions"`
	Name              string `json:"name"`
	PolicyType        string `json:"policy_type"`
	ManagementGroupId string `json:"management_group_id"`
	DisplayName       string `json:"display_name"`
}



type AzurermPolicySetDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPolicySetDefinitionList is a list of AzurermPolicySetDefinitions
type AzurermPolicySetDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPolicySetDefinition CRD objects
	Items []AzurermPolicySetDefinition `json:"items,omitempty"`
}