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

type AzurermPolicyDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPolicyDefinitionSpec   `json:"spec,omitempty"`
	Status            AzurermPolicyDefinitionStatus `json:"status,omitempty"`
}

type AzurermPolicyDefinitionSpec struct {
	DisplayName       string `json:"display_name"`
	Description       string `json:"description"`
	PolicyRule        string `json:"policy_rule"`
	Metadata          string `json:"metadata"`
	Parameters        string `json:"parameters"`
	PolicyType        string `json:"policy_type"`
	ManagementGroupId string `json:"management_group_id"`
	Name              string `json:"name"`
	Mode              string `json:"mode"`
}



type AzurermPolicyDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPolicyDefinitionList is a list of AzurermPolicyDefinitions
type AzurermPolicyDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPolicyDefinition CRD objects
	Items []AzurermPolicyDefinition `json:"items,omitempty"`
}