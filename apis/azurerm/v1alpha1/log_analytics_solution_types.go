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

type LogAnalyticsSolution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsSolutionSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsSolutionStatus `json:"status,omitempty"`
}

type LogAnalyticsSolutionSpecPlan struct {
	Product string `json:"product"`
	// +optional
	PromotionCode string `json:"promotion_code,omitempty"`
	Publisher     string `json:"publisher"`
}

type LogAnalyticsSolutionSpec struct {
	Location string `json:"location"`
	// +kubebuilder:validation:MaxItems=1
	Plan                []LogAnalyticsSolutionSpec `json:"plan"`
	ResourceGroupName   string                     `json:"resource_group_name"`
	SolutionName        string                     `json:"solution_name"`
	WorkspaceName       string                     `json:"workspace_name"`
	WorkspaceResourceId string                     `json:"workspace_resource_id"`
}

type LogAnalyticsSolutionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsSolutionList is a list of LogAnalyticsSolutions
type LogAnalyticsSolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsSolution CRD objects
	Items []LogAnalyticsSolution `json:"items,omitempty"`
}
