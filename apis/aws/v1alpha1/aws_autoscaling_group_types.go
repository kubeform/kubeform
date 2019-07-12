package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingGroupSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingGroupStatus `json:"status,omitempty"`
}

type AwsAutoscalingGroupSpecInitialLifecycleHook struct {
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	DefaultResult         string `json:"default_result"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution struct {
	SpotMaxPrice                        string `json:"spot_max_price"`
	OnDemandAllocationStrategy          string `json:"on_demand_allocation_strategy"`
	OnDemandBaseCapacity                int    `json:"on_demand_base_capacity"`
	OnDemandPercentageAboveBaseCapacity int    `json:"on_demand_percentage_above_base_capacity"`
	SpotAllocationStrategy              string `json:"spot_allocation_strategy"`
	SpotInstancePools                   int    `json:"spot_instance_pools"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	LaunchTemplateId   string `json:"launch_template_id"`
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride struct {
	InstanceType string `json:"instance_type"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate struct {
	LaunchTemplateSpecification []AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"launch_template_specification"`
	Override                    []AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"override"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicy struct {
	InstancesDistribution []AwsAutoscalingGroupSpecMixedInstancesPolicy `json:"instances_distribution"`
	LaunchTemplate        []AwsAutoscalingGroupSpecMixedInstancesPolicy `json:"launch_template"`
}

type AwsAutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
}

type AwsAutoscalingGroupSpecLaunchTemplate struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Id      string `json:"id"`
}

type AwsAutoscalingGroupSpec struct {
	MetricsGranularity     string                    `json:"metrics_granularity"`
	ServiceLinkedRoleArn   string                    `json:"service_linked_role_arn"`
	AvailabilityZones      []string                  `json:"availability_zones"`
	LaunchConfiguration    string                    `json:"launch_configuration"`
	WaitForElbCapacity     int                       `json:"wait_for_elb_capacity"`
	TargetGroupArns        []string                  `json:"target_group_arns"`
	Name                   string                    `json:"name"`
	ForceDelete            bool                      `json:"force_delete"`
	LoadBalancers          []string                  `json:"load_balancers"`
	SuspendedProcesses     []string                  `json:"suspended_processes"`
	DefaultCooldown        int                       `json:"default_cooldown"`
	MaxSize                int                       `json:"max_size"`
	InitialLifecycleHook   []AwsAutoscalingGroupSpec `json:"initial_lifecycle_hook"`
	DesiredCapacity        int                       `json:"desired_capacity"`
	MixedInstancesPolicy   []AwsAutoscalingGroupSpec `json:"mixed_instances_policy"`
	MinElbCapacity         int                       `json:"min_elb_capacity"`
	HealthCheckType        string                    `json:"health_check_type"`
	VpcZoneIdentifier      []string                  `json:"vpc_zone_identifier"`
	WaitForCapacityTimeout string                    `json:"wait_for_capacity_timeout"`
	ProtectFromScaleIn     bool                      `json:"protect_from_scale_in"`
	Tag                    []AwsAutoscalingGroupSpec `json:"tag"`
	NamePrefix             string                    `json:"name_prefix"`
	TerminationPolicies    []string                  `json:"termination_policies"`
	Arn                    string                    `json:"arn"`
	LaunchTemplate         []AwsAutoscalingGroupSpec `json:"launch_template"`
	HealthCheckGracePeriod int                       `json:"health_check_grace_period"`
	PlacementGroup         string                    `json:"placement_group"`
	EnabledMetrics         []string                  `json:"enabled_metrics"`
	MinSize                int                       `json:"min_size"`
}



type AwsAutoscalingGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroups
type AwsAutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingGroup CRD objects
	Items []AwsAutoscalingGroup `json:"items,omitempty"`
}