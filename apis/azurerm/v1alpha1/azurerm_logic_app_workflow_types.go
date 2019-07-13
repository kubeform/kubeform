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

type AzurermLogicAppWorkflow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppWorkflowSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppWorkflowStatus `json:"status,omitempty"`
}

type AzurermLogicAppWorkflowSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	Parameters        map[string]string `json:"parameters"`
	WorkflowSchema    string            `json:"workflow_schema"`
	WorkflowVersion   string            `json:"workflow_version"`
	Tags              map[string]string `json:"tags"`
	AccessEndpoint    string            `json:"access_endpoint"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
}



type AzurermLogicAppWorkflowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogicAppWorkflowList is a list of AzurermLogicAppWorkflows
type AzurermLogicAppWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppWorkflow CRD objects
	Items []AzurermLogicAppWorkflow `json:"items,omitempty"`
}