package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEc2Fleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2FleetSpec   `json:"spec,omitempty"`
	Status            AwsEc2FleetStatus `json:"status,omitempty"`
}

type AwsEc2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification struct {
	LaunchTemplateId   string `json:"launch_template_id"`
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
}

type AwsEc2FleetSpecLaunchTemplateConfigOverride struct {
	MaxPrice         string  `json:"max_price"`
	Priority         float64 `json:"priority"`
	SubnetId         string  `json:"subnet_id"`
	WeightedCapacity float64 `json:"weighted_capacity"`
	AvailabilityZone string  `json:"availability_zone"`
	InstanceType     string  `json:"instance_type"`
}

type AwsEc2FleetSpecLaunchTemplateConfig struct {
	LaunchTemplateSpecification []AwsEc2FleetSpecLaunchTemplateConfig `json:"launch_template_specification"`
	Override                    []AwsEc2FleetSpecLaunchTemplateConfig `json:"override"`
}

type AwsEc2FleetSpecOnDemandOptions struct {
	AllocationStrategy string `json:"allocation_strategy"`
}

type AwsEc2FleetSpecSpotOptions struct {
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	InstancePoolsToUseCount      int    `json:"instance_pools_to_use_count"`
	AllocationStrategy           string `json:"allocation_strategy"`
}

type AwsEc2FleetSpecTargetCapacitySpecification struct {
	DefaultTargetCapacityType string `json:"default_target_capacity_type"`
	OnDemandTargetCapacity    int    `json:"on_demand_target_capacity"`
	SpotTargetCapacity        int    `json:"spot_target_capacity"`
	TotalTargetCapacity       int    `json:"total_target_capacity"`
}

type AwsEc2FleetSpec struct {
	TerminateInstancesWithExpiration bool              `json:"terminate_instances_with_expiration"`
	LaunchTemplateConfig             []AwsEc2FleetSpec `json:"launch_template_config"`
	OnDemandOptions                  []AwsEc2FleetSpec `json:"on_demand_options"`
	SpotOptions                      []AwsEc2FleetSpec `json:"spot_options"`
	TargetCapacitySpecification      []AwsEc2FleetSpec `json:"target_capacity_specification"`
	Type                             string            `json:"type"`
	ExcessCapacityTerminationPolicy  string            `json:"excess_capacity_termination_policy"`
	ReplaceUnhealthyInstances        bool              `json:"replace_unhealthy_instances"`
	Tags                             map[string]string `json:"tags"`
	TerminateInstances               bool              `json:"terminate_instances"`
}



type AwsEc2FleetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEc2FleetList is a list of AwsEc2Fleets
type AwsEc2FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2Fleet CRD objects
	Items []AwsEc2Fleet `json:"items,omitempty"`
}