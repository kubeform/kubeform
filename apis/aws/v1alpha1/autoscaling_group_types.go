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

type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingGroupSpec   `json:"spec,omitempty"`
	Status            AutoscalingGroupStatus `json:"status,omitempty"`
}

type AutoscalingGroupSpecInitialLifecycleHook struct {
	// +optional
	HeartbeatTimeout    int    `json:"heartbeat_timeout,omitempty"`
	LifecycleTransition string `json:"lifecycle_transition"`
	Name                string `json:"name"`
	// +optional
	NotificationMetadata string `json:"notification_metadata,omitempty"`
	// +optional
	NotificationTargetArn string `json:"notification_target_arn,omitempty"`
	// +optional
	RoleArn string `json:"role_arn,omitempty"`
}

type AutoscalingGroupSpecLaunchTemplate struct {
	// +optional
	Version string `json:"version,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution struct {
	// +optional
	OnDemandAllocationStrategy string `json:"on_demand_allocation_strategy,omitempty"`
	// +optional
	OnDemandBaseCapacity int `json:"on_demand_base_capacity,omitempty"`
	// +optional
	OnDemandPercentageAboveBaseCapacity int `json:"on_demand_percentage_above_base_capacity,omitempty"`
	// +optional
	SpotAllocationStrategy string `json:"spot_allocation_strategy,omitempty"`
	// +optional
	SpotMaxPrice string `json:"spot_max_price,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	// +optional
	Version string `json:"version,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride struct {
	// +optional
	InstanceType string `json:"instance_type,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateSpecification []AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"launch_template_specification"`
	// +optional
	Override *[]AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"override,omitempty"`
}

type AutoscalingGroupSpecMixedInstancesPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InstancesDistribution *[]AutoscalingGroupSpecMixedInstancesPolicy `json:"instances_distribution,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplate []AutoscalingGroupSpecMixedInstancesPolicy `json:"launch_template"`
}

type AutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
	Value             string `json:"value"`
}

type AutoscalingGroupSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledMetrics []string `json:"enabled_metrics,omitempty"`
	// +optional
	ForceDelete bool `json:"force_delete,omitempty"`
	// +optional
	HealthCheckGracePeriod int `json:"health_check_grace_period,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InitialLifecycleHook *[]AutoscalingGroupSpec `json:"initial_lifecycle_hook,omitempty"`
	// +optional
	LaunchConfiguration string `json:"launch_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LaunchTemplate *[]AutoscalingGroupSpec `json:"launch_template,omitempty"`
	MaxSize        int                     `json:"max_size"`
	// +optional
	MetricsGranularity string `json:"metrics_granularity,omitempty"`
	// +optional
	MinElbCapacity int `json:"min_elb_capacity,omitempty"`
	MinSize        int `json:"min_size"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MixedInstancesPolicy *[]AutoscalingGroupSpec `json:"mixed_instances_policy,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	PlacementGroup string `json:"placement_group,omitempty"`
	// +optional
	ProtectFromScaleIn bool `json:"protect_from_scale_in,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SuspendedProcesses []string `json:"suspended_processes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tag *[]AutoscalingGroupSpec `json:"tag,omitempty"`
	// +optional
	// +optional
	TerminationPolicies []string `json:"termination_policies,omitempty"`
	// +optional
	WaitForCapacityTimeout string `json:"wait_for_capacity_timeout,omitempty"`
	// +optional
	WaitForElbCapacity int `json:"wait_for_elb_capacity,omitempty"`
}

type AutoscalingGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
