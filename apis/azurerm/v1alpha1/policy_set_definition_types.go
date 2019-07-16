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

type PolicySetDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySetDefinitionSpec   `json:"spec,omitempty"`
	Status            PolicySetDefinitionStatus `json:"status,omitempty"`
}

type PolicySetDefinitionSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +optional
	ManagementGroupId string `json:"management_group_id,omitempty"`
	// +optional
	Metadata string `json:"metadata,omitempty"`
	Name     string `json:"name"`
	// +optional
	Parameters string `json:"parameters,omitempty"`
	// +optional
	PolicyDefinitions string `json:"policy_definitions,omitempty"`
	PolicyType        string `json:"policy_type"`
}

type PolicySetDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PolicySetDefinitionList is a list of PolicySetDefinitions
type PolicySetDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicySetDefinition CRD objects
	Items []PolicySetDefinition `json:"items,omitempty"`
}
