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

type AwsElasticBeanstalkEnvironmentSpecSetting struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpecAllSettings struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	Instances            []string                             `json:"instances"`
	Application          string                               `json:"application"`
	Cname                string                               `json:"cname"`
	Tier                 string                               `json:"tier"`
	Setting              []AwsElasticBeanstalkEnvironmentSpec `json:"setting"`
	PollInterval         string                               `json:"poll_interval"`
	Tags                 map[string]string                    `json:"tags"`
	CnamePrefix          string                               `json:"cname_prefix"`
	PlatformArn          string                               `json:"platform_arn"`
	TemplateName         string                               `json:"template_name"`
	AutoscalingGroups    []string                             `json:"autoscaling_groups"`
	Queues               []string                             `json:"queues"`
	LoadBalancers        []string                             `json:"load_balancers"`
	Triggers             []string                             `json:"triggers"`
	Arn                  string                               `json:"arn"`
	Name                 string                               `json:"name"`
	VersionLabel         string                               `json:"version_label"`
	SolutionStackName    string                               `json:"solution_stack_name"`
	LaunchConfigurations []string                             `json:"launch_configurations"`
	Description          string                               `json:"description"`
	AllSettings          []AwsElasticBeanstalkEnvironmentSpec `json:"all_settings"`
	WaitForReadyTimeout  string                               `json:"wait_for_ready_timeout"`
}

type AwsElasticBeanstalkEnvironmentStatus struct {
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
