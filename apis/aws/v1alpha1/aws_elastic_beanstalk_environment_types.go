package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticBeanstalkEnvironmentSpec   `json:"spec,omitempty"`
	Status            AwsElasticBeanstalkEnvironmentStatus `json:"status,omitempty"`
}

type AwsElasticBeanstalkEnvironmentSpecAllSettings struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpecSetting struct {
	Value     string `json:"value"`
	Resource  string `json:"resource"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	Application          string                               `json:"application"`
	TemplateName         string                               `json:"template_name"`
	PollInterval         string                               `json:"poll_interval"`
	LaunchConfigurations []string                             `json:"launch_configurations"`
	Queues               []string                             `json:"queues"`
	Arn                  string                               `json:"arn"`
	Name                 string                               `json:"name"`
	VersionLabel         string                               `json:"version_label"`
	Tier                 string                               `json:"tier"`
	AllSettings          []AwsElasticBeanstalkEnvironmentSpec `json:"all_settings"`
	SolutionStackName    string                               `json:"solution_stack_name"`
	Instances            []string                             `json:"instances"`
	Cname                string                               `json:"cname"`
	PlatformArn          string                               `json:"platform_arn"`
	WaitForReadyTimeout  string                               `json:"wait_for_ready_timeout"`
	AutoscalingGroups    []string                             `json:"autoscaling_groups"`
	LoadBalancers        []string                             `json:"load_balancers"`
	Triggers             []string                             `json:"triggers"`
	Tags                 map[string]string                    `json:"tags"`
	Description          string                               `json:"description"`
	CnamePrefix          string                               `json:"cname_prefix"`
	Setting              []AwsElasticBeanstalkEnvironmentSpec `json:"setting"`
}



type AwsElasticBeanstalkEnvironmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkEnvironmentList is a list of AwsElasticBeanstalkEnvironments
type AwsElasticBeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticBeanstalkEnvironment CRD objects
	Items []AwsElasticBeanstalkEnvironment `json:"items,omitempty"`
}