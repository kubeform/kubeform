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

type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingGroupSpec   `json:"spec,omitempty"`
	Status            AutoscalingGroupStatus `json:"status,omitempty"`
}

type AutoscalingGroupSpecInitialLifecycleHook struct {
	// +optional
	DefaultResult string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`
	// +optional
	HeartbeatTimeout    int    `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`
	LifecycleTransition string `json:"lifecycleTransition" tf:"lifecycle_transition"`
	Name                string `json:"name" tf:"name"`
	// +optional
	NotificationMetadata string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`
	// +optional
	NotificationTargetArn string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type AutoscalingGroupSpecLaunchTemplate struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution struct {
	// +optional
	OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy,omitempty"`
	// +optional
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity,omitempty" tf:"on_demand_base_capacity,omitempty"`
	// +optional
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity,omitempty" tf:"on_demand_percentage_above_base_capacity,omitempty"`
	// +optional
	SpotAllocationStrategy string `json:"spotAllocationStrategy,omitempty" tf:"spot_allocation_strategy,omitempty"`
	// +optional
	SpotInstancePools int `json:"spotInstancePools,omitempty" tf:"spot_instance_pools,omitempty"`
	// +optional
	SpotMaxPrice string `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	// +optional
	LaunchTemplateID string `json:"launchTemplateID,omitempty" tf:"launch_template_id,omitempty"`
	// +optional
	LaunchTemplateName string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride struct {
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateSpecification []AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `json:"launchTemplateSpecification" tf:"launch_template_specification"`
	// +optional
	Override []AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride `json:"override,omitempty" tf:"override,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InstancesDistribution []AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution `json:"instancesDistribution,omitempty" tf:"instances_distribution,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplate []AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"launchTemplate" tf:"launch_template"`
}

type AutoscalingGroupSpecTag struct {
	Key               string `json:"key" tf:"key"`
	PropagateAtLaunch bool   `json:"propagateAtLaunch" tf:"propagate_at_launch"`
	Value             string `json:"value" tf:"value"`
}

type AutoscalingGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	DefaultCooldown int `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`
	// +optional
	DesiredCapacity int `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`
	// +optional
	EnabledMetrics []string `json:"enabledMetrics,omitempty" tf:"enabled_metrics,omitempty"`
	// +optional
	ForceDelete bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`
	// +optional
	HealthCheckGracePeriod int `json:"healthCheckGracePeriod,omitempty" tf:"health_check_grace_period,omitempty"`
	// +optional
	HealthCheckType string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`
	// +optional
	InitialLifecycleHook []AutoscalingGroupSpecInitialLifecycleHook `json:"initialLifecycleHook,omitempty" tf:"initial_lifecycle_hook,omitempty"`
	// +optional
	LaunchConfiguration string `json:"launchConfiguration,omitempty" tf:"launch_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LaunchTemplate []AutoscalingGroupSpecLaunchTemplate `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`
	// +optional
	LoadBalancers []string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`
	MaxSize       int      `json:"maxSize" tf:"max_size"`
	// +optional
	MetricsGranularity string `json:"metricsGranularity,omitempty" tf:"metrics_granularity,omitempty"`
	// +optional
	MinElbCapacity int `json:"minElbCapacity,omitempty" tf:"min_elb_capacity,omitempty"`
	MinSize        int `json:"minSize" tf:"min_size"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MixedInstancesPolicy []AutoscalingGroupSpecMixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty" tf:"mixed_instances_policy,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	PlacementGroup string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`
	// +optional
	ProtectFromScaleIn bool `json:"protectFromScaleIn,omitempty" tf:"protect_from_scale_in,omitempty"`
	// +optional
	ServiceLinkedRoleArn string `json:"serviceLinkedRoleArn,omitempty" tf:"service_linked_role_arn,omitempty"`
	// +optional
	SuspendedProcesses []string `json:"suspendedProcesses,omitempty" tf:"suspended_processes,omitempty"`
	// +optional
	Tag []AutoscalingGroupSpecTag `json:"tag,omitempty" tf:"tag,omitempty"`
	// +optional
	// +optional
	TargetGroupArns []string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`
	// +optional
	TerminationPolicies []string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`
	// +optional
	VpcZoneIdentifier []string `json:"vpcZoneIdentifier,omitempty" tf:"vpc_zone_identifier,omitempty"`
	// +optional
	WaitForCapacityTimeout string `json:"waitForCapacityTimeout,omitempty" tf:"wait_for_capacity_timeout,omitempty"`
	// +optional
	WaitForElbCapacity int `json:"waitForElbCapacity,omitempty" tf:"wait_for_elb_capacity,omitempty"`
}

type AutoscalingGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutoscalingGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingGroupList is a list of AutoscalingGroups
type AutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingGroup CRD objects
	Items []AutoscalingGroup `json:"items,omitempty"`
}
