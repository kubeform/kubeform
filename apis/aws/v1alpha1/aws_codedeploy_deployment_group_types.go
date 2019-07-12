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

type AwsCodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodedeployDeploymentGroupSpec   `json:"spec,omitempty"`
	Status            AwsCodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

type AwsCodedeployDeploymentGroupSpecAlarmConfiguration struct {
	Alarms                 []string `json:"alarms"`
	Enabled                bool     `json:"enabled"`
	IgnorePollAlarmFailure bool     `json:"ignore_poll_alarm_failure"`
}

type AwsCodedeployDeploymentGroupSpecAutoRollbackConfiguration struct {
	Events  []string `json:"events"`
	Enabled bool     `json:"enabled"`
}

type AwsCodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCodedeployDeploymentGroupSpecTriggerConfiguration struct {
	TriggerEvents    []string `json:"trigger_events"`
	TriggerName      string   `json:"trigger_name"`
	TriggerTargetArn string   `json:"trigger_target_arn"`
}

type AwsCodedeployDeploymentGroupSpecEcsService struct {
	ClusterName string `json:"cluster_name"`
	ServiceName string `json:"service_name"`
}

type AwsCodedeployDeploymentGroupSpecDeploymentStyle struct {
	DeploymentOption string `json:"deployment_option"`
	DeploymentType   string `json:"deployment_type"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption struct {
	ActionOnTimeout   string `json:"action_on_timeout"`
	WaitTimeInMinutes int    `json:"wait_time_in_minutes"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption struct {
	Action string `json:"action"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	Action                       string `json:"action"`
	TerminationWaitTimeInMinutes int    `json:"termination_wait_time_in_minutes"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig struct {
	DeploymentReadyOption                     []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig `json:"deployment_ready_option"`
	GreenFleetProvisioningOption              []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig `json:"green_fleet_provisioning_option"`
	TerminateBlueInstancesOnDeploymentSuccess []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig `json:"terminate_blue_instances_on_deployment_success"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo struct {
	Name string `json:"name"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo struct {
	Name string `json:"name"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute struct {
	ListenerArns []string `json:"listener_arns"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute struct {
	ListenerArns []string `json:"listener_arns"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup struct {
	Name string `json:"name"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo struct {
	TestTrafficRoute []AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo `json:"test_traffic_route"`
	ProdTrafficRoute []AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo `json:"prod_traffic_route"`
	TargetGroup      []AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo `json:"target_group"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfo struct {
	ElbInfo             []AwsCodedeployDeploymentGroupSpecLoadBalancerInfo `json:"elb_info"`
	TargetGroupInfo     []AwsCodedeployDeploymentGroupSpecLoadBalancerInfo `json:"target_group_info"`
	TargetGroupPairInfo []AwsCodedeployDeploymentGroupSpecLoadBalancerInfo `json:"target_group_pair_info"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagSet struct {
	Ec2TagFilter []AwsCodedeployDeploymentGroupSpecEc2TagSet `json:"ec2_tag_filter"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCodedeployDeploymentGroupSpec struct {
	ServiceRoleArn              string                             `json:"service_role_arn"`
	AlarmConfiguration          []AwsCodedeployDeploymentGroupSpec `json:"alarm_configuration"`
	AutoRollbackConfiguration   []AwsCodedeployDeploymentGroupSpec `json:"auto_rollback_configuration"`
	AutoscalingGroups           []string                           `json:"autoscaling_groups"`
	OnPremisesInstanceTagFilter []AwsCodedeployDeploymentGroupSpec `json:"on_premises_instance_tag_filter"`
	DeploymentConfigName        string                             `json:"deployment_config_name"`
	TriggerConfiguration        []AwsCodedeployDeploymentGroupSpec `json:"trigger_configuration"`
	AppName                     string                             `json:"app_name"`
	EcsService                  []AwsCodedeployDeploymentGroupSpec `json:"ecs_service"`
	DeploymentGroupName         string                             `json:"deployment_group_name"`
	DeploymentStyle             []AwsCodedeployDeploymentGroupSpec `json:"deployment_style"`
	BlueGreenDeploymentConfig   []AwsCodedeployDeploymentGroupSpec `json:"blue_green_deployment_config"`
	LoadBalancerInfo            []AwsCodedeployDeploymentGroupSpec `json:"load_balancer_info"`
	Ec2TagSet                   []AwsCodedeployDeploymentGroupSpec `json:"ec2_tag_set"`
	Ec2TagFilter                []AwsCodedeployDeploymentGroupSpec `json:"ec2_tag_filter"`
}

type AwsCodedeployDeploymentGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroups
type AwsCodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodedeployDeploymentGroup CRD objects
	Items []AwsCodedeployDeploymentGroup `json:"items,omitempty"`
}
