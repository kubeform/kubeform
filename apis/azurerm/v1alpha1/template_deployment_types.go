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

type TemplateDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateDeploymentSpec   `json:"spec,omitempty"`
	Status            TemplateDeploymentStatus `json:"status,omitempty"`
}

type TemplateDeploymentSpec struct {
	DeploymentMode string `json:"deployment_mode"`
	Name           string `json:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	ParametersBody    string `json:"parameters_body,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
}

type TemplateDeploymentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
