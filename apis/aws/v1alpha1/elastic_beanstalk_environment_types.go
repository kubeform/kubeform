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

type ElasticBeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkEnvironmentSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkEnvironmentStatus `json:"status,omitempty"`
}

type ElasticBeanstalkEnvironmentSpecAllSettings struct {
	Name      string `json:"name" tf:"name"`
	Namespace string `json:"namespace" tf:"namespace"`
	// +optional
	Resource string `json:"resource,omitempty" tf:"resource,omitempty"`
	Value    string `json:"value" tf:"value"`
}

type ElasticBeanstalkEnvironmentSpecSetting struct {
	Name      string `json:"name" tf:"name"`
	Namespace string `json:"namespace" tf:"namespace"`
	// +optional
	Resource string `json:"resource,omitempty" tf:"resource,omitempty"`
	Value    string `json:"value" tf:"value"`
}

type ElasticBeanstalkEnvironmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:UniqueItems=true
	AllSettings       []ElasticBeanstalkEnvironmentSpecAllSettings `json:"allSettings" tf:"all_settings"`
	Application       string                                       `json:"application" tf:"application"`
	Arn               string                                       `json:"arn" tf:"arn"`
	AutoscalingGroups []string                                     `json:"autoscalingGroups" tf:"autoscaling_groups"`
	Cname             string                                       `json:"cname" tf:"cname"`
	// +optional
	CnamePrefix string `json:"cnamePrefix,omitempty" tf:"cname_prefix,omitempty"`
	// +optional
	Description          string   `json:"description,omitempty" tf:"description,omitempty"`
	Instances            []string `json:"instances" tf:"instances"`
	LaunchConfigurations []string `json:"launchConfigurations" tf:"launch_configurations"`
	LoadBalancers        []string `json:"loadBalancers" tf:"load_balancers"`
	Name                 string   `json:"name" tf:"name"`
	// +optional
	PlatformArn string `json:"platformArn,omitempty" tf:"platform_arn,omitempty"`
	// +optional
	PollInterval string   `json:"pollInterval,omitempty" tf:"poll_interval,omitempty"`
	Queues       []string `json:"queues" tf:"queues"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Setting []ElasticBeanstalkEnvironmentSpecSetting `json:"setting,omitempty" tf:"setting,omitempty"`
	// +optional
	SolutionStackName string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TemplateName string `json:"templateName,omitempty" tf:"template_name,omitempty"`
	// +optional
	Tier     string   `json:"tier,omitempty" tf:"tier,omitempty"`
	Triggers []string `json:"triggers" tf:"triggers"`
	// +optional
	VersionLabel string `json:"versionLabel,omitempty" tf:"version_label,omitempty"`
	// +optional
	WaitForReadyTimeout string `json:"waitForReadyTimeout,omitempty" tf:"wait_for_ready_timeout,omitempty"`
}

type ElasticBeanstalkEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElasticBeanstalkEnvironmentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticBeanstalkEnvironmentList is a list of ElasticBeanstalkEnvironments
type ElasticBeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticBeanstalkEnvironment CRD objects
	Items []ElasticBeanstalkEnvironment `json:"items,omitempty"`
}
