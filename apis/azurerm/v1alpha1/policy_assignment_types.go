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

type PolicyAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyAssignmentSpec   `json:"spec,omitempty"`
	Status            PolicyAssignmentStatus `json:"status,omitempty"`
}

type PolicyAssignmentSpecIdentity struct {
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type PolicyAssignmentSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []PolicyAssignmentSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	NotScopes []string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`
	// +optional
	Parameters         string                    `json:"parameters,omitempty" tf:"parameters,omitempty"`
	PolicyDefinitionID string                    `json:"policyDefinitionID" tf:"policy_definition_id"`
	Scope              string                    `json:"scope" tf:"scope"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PolicyAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
