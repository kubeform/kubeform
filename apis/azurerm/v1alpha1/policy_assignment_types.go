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

type PolicyAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyAssignmentSpec   `json:"spec,omitempty"`
	Status            PolicyAssignmentStatus `json:"status,omitempty"`
}

type PolicyAssignmentSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	Name     string `json:"name"`
	// +optional
	NotScopes []string `json:"not_scopes,omitempty"`
	// +optional
	Parameters         string `json:"parameters,omitempty"`
	PolicyDefinitionId string `json:"policy_definition_id"`
	Scope              string `json:"scope"`
}

type PolicyAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PolicyAssignmentList is a list of PolicyAssignments
type PolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyAssignment CRD objects
	Items []PolicyAssignment `json:"items,omitempty"`
}
