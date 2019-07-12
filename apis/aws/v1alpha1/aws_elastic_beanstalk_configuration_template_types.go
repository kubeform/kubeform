package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticBeanstalkConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status            AwsElasticBeanstalkConfigurationTemplateStatus `json:"status,omitempty"`
}

type AwsElasticBeanstalkConfigurationTemplateSpecSetting struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
	Namespace string `json:"namespace"`
}

type AwsElasticBeanstalkConfigurationTemplateSpec struct {
	Name              string                                         `json:"name"`
	Application       string                                         `json:"application"`
	Description       string                                         `json:"description"`
	EnvironmentId     string                                         `json:"environment_id"`
	Setting           []AwsElasticBeanstalkConfigurationTemplateSpec `json:"setting"`
	SolutionStackName string                                         `json:"solution_stack_name"`
}



type AwsElasticBeanstalkConfigurationTemplateStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkConfigurationTemplateList is a list of AwsElasticBeanstalkConfigurationTemplates
type AwsElasticBeanstalkConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticBeanstalkConfigurationTemplate CRD objects
	Items []AwsElasticBeanstalkConfigurationTemplate `json:"items,omitempty"`
}