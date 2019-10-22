package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllSettings []ElasticBeanstalkEnvironmentSpecAllSettings `json:"allSettings,omitempty" tf:"all_settings,omitempty"`
	Application string                                       `json:"application" tf:"application"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AutoscalingGroups []string `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`
	// +optional
	Cname string `json:"cname,omitempty" tf:"cname,omitempty"`
	// +optional
	CnamePrefix string `json:"cnamePrefix,omitempty" tf:"cname_prefix,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances,omitempty"`
	// +optional
	LaunchConfigurations []string `json:"launchConfigurations,omitempty" tf:"launch_configurations,omitempty"`
	// +optional
	LoadBalancers []string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`
	Name          string   `json:"name" tf:"name"`
	// +optional
	PlatformArn string `json:"platformArn,omitempty" tf:"platform_arn,omitempty"`
	// +optional
	PollInterval string `json:"pollInterval,omitempty" tf:"poll_interval,omitempty"`
	// +optional
	Queues []string `json:"queues,omitempty" tf:"queues,omitempty"`
	// +optional
	Setting []ElasticBeanstalkEnvironmentSpecSetting `json:"setting,omitempty" tf:"setting,omitempty"`
	// +optional
	SolutionStackName string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TemplateName string `json:"templateName,omitempty" tf:"template_name,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
	// +optional
	Triggers []string `json:"triggers,omitempty" tf:"triggers,omitempty"`
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
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
