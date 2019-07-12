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

type AwsAutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingGroupSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingGroupStatus `json:"status,omitempty"`
}

type AwsAutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution struct {
	OnDemandAllocationStrategy          string `json:"on_demand_allocation_strategy"`
	OnDemandBaseCapacity                int    `json:"on_demand_base_capacity"`
	OnDemandPercentageAboveBaseCapacity int    `json:"on_demand_percentage_above_base_capacity"`
	SpotAllocationStrategy              string `json:"spot_allocation_strategy"`
	SpotInstancePools                   int    `json:"spot_instance_pools"`
	SpotMaxPrice                        string `json:"spot_max_price"`
}

type AwsAutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
	LaunchTemplateId   string `json:"launch_template_id"`
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

type AwsAutoscalingGroupSpecLaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AwsAutoscalingGroupSpecInitialLifecycleHook struct {
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
}

type AwsAutoscalingGroupSpec struct {
	SuspendedProcesses     []string                  `json:"suspended_processes"`
	TargetGroupArns        []string                  `json:"target_group_arns"`
	ServiceLinkedRoleArn   string                    `json:"service_linked_role_arn"`
	MinElbCapacity         int                       `json:"min_elb_capacity"`
	MinSize                int                       `json:"min_size"`
	HealthCheckType        string                    `json:"health_check_type"`
	DesiredCapacity        int                       `json:"desired_capacity"`
	TerminationPolicies    []string                  `json:"termination_policies"`
	Tag                    []AwsAutoscalingGroupSpec `json:"tag"`
	HealthCheckGracePeriod int                       `json:"health_check_grace_period"`
	PlacementGroup         string                    `json:"placement_group"`
	WaitForCapacityTimeout string                    `json:"wait_for_capacity_timeout"`
	WaitForElbCapacity     int                       `json:"wait_for_elb_capacity"`
	ProtectFromScaleIn     bool                      `json:"protect_from_scale_in"`
	NamePrefix             string                    `json:"name_prefix"`
	MixedInstancesPolicy   []AwsAutoscalingGroupSpec `json:"mixed_instances_policy"`
	DefaultCooldown        int                       `json:"default_cooldown"`
	LaunchConfiguration    string                    `json:"launch_configuration"`
	LaunchTemplate         []AwsAutoscalingGroupSpec `json:"launch_template"`
	InitialLifecycleHook   []AwsAutoscalingGroupSpec `json:"initial_lifecycle_hook"`
	Arn                    string                    `json:"arn"`
	ForceDelete            bool                      `json:"force_delete"`
	LoadBalancers          []string                  `json:"load_balancers"`
	VpcZoneIdentifier      []string                  `json:"vpc_zone_identifier"`
	MaxSize                int                       `json:"max_size"`
	AvailabilityZones      []string                  `json:"availability_zones"`
	EnabledMetrics         []string                  `json:"enabled_metrics"`
	Name                   string                    `json:"name"`
	MetricsGranularity     string                    `json:"metrics_granularity"`
}

type AwsAutoscalingGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroups
type AwsAutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingGroup CRD objects
	Items []AwsAutoscalingGroup `json:"items,omitempty"`
}
