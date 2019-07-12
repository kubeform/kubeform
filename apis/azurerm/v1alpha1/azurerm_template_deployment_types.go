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

type AzurermTemplateDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermTemplateDeploymentSpec   `json:"spec,omitempty"`
	Status            AzurermTemplateDeploymentStatus `json:"status,omitempty"`
}

type AzurermTemplateDeploymentSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	TemplateBody      string            `json:"template_body"`
	Parameters        map[string]string `json:"parameters"`
	ParametersBody    string            `json:"parameters_body"`
	DeploymentMode    string            `json:"deployment_mode"`
	Outputs           map[string]string `json:"outputs"`
	Name              string            `json:"name"`
}

type AzurermTemplateDeploymentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermTemplateDeploymentList is a list of AzurermTemplateDeployments
type AzurermTemplateDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermTemplateDeployment CRD objects
	Items []AzurermTemplateDeployment `json:"items,omitempty"`
}
