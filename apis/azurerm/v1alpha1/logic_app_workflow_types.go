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

type LogicAppWorkflow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppWorkflowSpec   `json:"spec,omitempty"`
	Status            LogicAppWorkflowStatus `json:"status,omitempty"`
}

type LogicAppWorkflowSpec struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty"`
	ResourceGroupName string            `json:"resource_group_name"`
	// +optional
	WorkflowSchema string `json:"workflow_schema,omitempty"`
	// +optional
	WorkflowVersion string `json:"workflow_version,omitempty"`
}

type LogicAppWorkflowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
