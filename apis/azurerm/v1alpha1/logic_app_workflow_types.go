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

type LogicAppWorkflow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppWorkflowSpec   `json:"spec,omitempty"`
	Status            LogicAppWorkflowStatus `json:"status,omitempty"`
}

type LogicAppWorkflowSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	WorkflowSchema string `json:"workflowSchema,omitempty" tf:"workflow_schema,omitempty"`
	// +optional
	WorkflowVersion string `json:"workflowVersion,omitempty" tf:"workflow_version,omitempty"`
}

type LogicAppWorkflowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppWorkflowList is a list of LogicAppWorkflows
type LogicAppWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppWorkflow CRD objects
	Items []LogicAppWorkflow `json:"items,omitempty"`
}
