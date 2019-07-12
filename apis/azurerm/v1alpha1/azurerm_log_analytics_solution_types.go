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

type AzurermLogAnalyticsSolution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogAnalyticsSolutionSpec   `json:"spec,omitempty"`
	Status            AzurermLogAnalyticsSolutionStatus `json:"status,omitempty"`
}

type AzurermLogAnalyticsSolutionSpecPlan struct {
	PromotionCode string `json:"promotion_code"`
	Product       string `json:"product"`
	Name          string `json:"name"`
	Publisher     string `json:"publisher"`
}

type AzurermLogAnalyticsSolutionSpec struct {
	WorkspaceName       string                            `json:"workspace_name"`
	WorkspaceResourceId string                            `json:"workspace_resource_id"`
	Location            string                            `json:"location"`
	ResourceGroupName   string                            `json:"resource_group_name"`
	Plan                []AzurermLogAnalyticsSolutionSpec `json:"plan"`
	SolutionName        string                            `json:"solution_name"`
}

type AzurermLogAnalyticsSolutionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogAnalyticsSolutionList is a list of AzurermLogAnalyticsSolutions
type AzurermLogAnalyticsSolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogAnalyticsSolution CRD objects
	Items []AzurermLogAnalyticsSolution `json:"items,omitempty"`
}
