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

type ElasticBeanstalkConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkConfigurationTemplateStatus `json:"status,omitempty"`
}

type ElasticBeanstalkConfigurationTemplateSpec struct {
	Application string `json:"application"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnvironmentId string `json:"environment_id,omitempty"`
	Name          string `json:"name"`
	// +optional
	SolutionStackName string `json:"solution_stack_name,omitempty"`
}

type ElasticBeanstalkConfigurationTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
