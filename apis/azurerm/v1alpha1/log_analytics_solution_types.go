package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	Product string `json:"product" tf:"product"`
	// +optional
	PromotionCode string `json:"promotionCode,omitempty" tf:"promotion_code,omitempty"`
	Publisher     string `json:"publisher" tf:"publisher"`
}

type LogAnalyticsSolutionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Location string `json:"location" tf:"location"`
	// +kubebuilder:validation:MaxItems=1
	Plan                []LogAnalyticsSolutionSpecPlan `json:"plan" tf:"plan"`
	ResourceGroupName   string                         `json:"resourceGroupName" tf:"resource_group_name"`
	SolutionName        string                         `json:"solutionName" tf:"solution_name"`
	WorkspaceName       string                         `json:"workspaceName" tf:"workspace_name"`
	WorkspaceResourceID string                         `json:"workspaceResourceID" tf:"workspace_resource_id"`
}

type LogAnalyticsSolutionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
