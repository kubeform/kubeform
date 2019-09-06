package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TemplateDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateDeploymentSpec   `json:"spec,omitempty"`
	Status            TemplateDeploymentStatus `json:"status,omitempty"`
}

type TemplateDeploymentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DeploymentMode string `json:"deploymentMode" tf:"deployment_mode"`
	Name           string `json:"name" tf:"name"`
	// +optional
	Outputs map[string]string `json:"outputs,omitempty" tf:"outputs,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	ParametersBody    string `json:"parametersBody,omitempty" tf:"parameters_body,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	TemplateBody string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
}

type TemplateDeploymentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *TemplateDeploymentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TemplateDeploymentList is a list of TemplateDeployments
type TemplateDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TemplateDeployment CRD objects
	Items []TemplateDeployment `json:"items,omitempty"`
}
