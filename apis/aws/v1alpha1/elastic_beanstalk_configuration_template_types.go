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

type ElasticBeanstalkConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkConfigurationTemplateStatus `json:"status,omitempty"`
}

type ElasticBeanstalkConfigurationTemplateSpecSetting struct {
	Name      string `json:"name" tf:"name"`
	Namespace string `json:"namespace" tf:"namespace"`
	// +optional
	Resource string `json:"resource,omitempty" tf:"resource,omitempty"`
	Value    string `json:"value" tf:"value"`
}

type ElasticBeanstalkConfigurationTemplateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Application string `json:"application" tf:"application"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnvironmentID string `json:"environmentID,omitempty" tf:"environment_id,omitempty"`
	Name          string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Setting []ElasticBeanstalkConfigurationTemplateSpecSetting `json:"setting,omitempty" tf:"setting,omitempty"`
	// +optional
	SolutionStackName string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
}

type ElasticBeanstalkConfigurationTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElasticBeanstalkConfigurationTemplateSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticBeanstalkConfigurationTemplateList is a list of ElasticBeanstalkConfigurationTemplates
type ElasticBeanstalkConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticBeanstalkConfigurationTemplate CRD objects
	Items []ElasticBeanstalkConfigurationTemplate `json:"items,omitempty"`
}
