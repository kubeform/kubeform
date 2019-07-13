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

type AwsElasticBeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticBeanstalkEnvironmentSpec   `json:"spec,omitempty"`
	Status            AwsElasticBeanstalkEnvironmentStatus `json:"status,omitempty"`
}

type AwsElasticBeanstalkEnvironmentSpecAllSettings struct {
	Value     string `json:"value"`
	Resource  string `json:"resource"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type AwsElasticBeanstalkEnvironmentSpecSetting struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	PollInterval         string                               `json:"poll_interval"`
	Instances            []string                             `json:"instances"`
	Arn                  string                               `json:"arn"`
	Name                 string                               `json:"name"`
	Application          string                               `json:"application"`
	VersionLabel         string                               `json:"version_label"`
	CnamePrefix          string                               `json:"cname_prefix"`
	PlatformArn          string                               `json:"platform_arn"`
	Triggers             []string                             `json:"triggers"`
	LoadBalancers        []string                             `json:"load_balancers"`
	Description          string                               `json:"description"`
	Cname                string                               `json:"cname"`
	AllSettings          []AwsElasticBeanstalkEnvironmentSpec `json:"all_settings"`
	TemplateName         string                               `json:"template_name"`
	WaitForReadyTimeout  string                               `json:"wait_for_ready_timeout"`
	LaunchConfigurations []string                             `json:"launch_configurations"`
	SolutionStackName    string                               `json:"solution_stack_name"`
	Tags                 map[string]string                    `json:"tags"`
	Tier                 string                               `json:"tier"`
	Setting              []AwsElasticBeanstalkEnvironmentSpec `json:"setting"`
	AutoscalingGroups    []string                             `json:"autoscaling_groups"`
	Queues               []string                             `json:"queues"`
}



type AwsElasticBeanstalkEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticBeanstalkEnvironmentList is a list of AwsElasticBeanstalkEnvironments
type AwsElasticBeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticBeanstalkEnvironment CRD objects
	Items []AwsElasticBeanstalkEnvironment `json:"items,omitempty"`
}