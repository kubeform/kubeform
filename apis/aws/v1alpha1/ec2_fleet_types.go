package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Ec2Fleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2FleetSpec   `json:"spec,omitempty"`
	Status            Ec2FleetStatus `json:"status,omitempty"`
}

type Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification struct {
	// +optional
	LaunchTemplateId string `json:"launch_template_id,omitempty"`
	// +optional
	LaunchTemplateName string `json:"launch_template_name,omitempty"`
	Version            string `json:"version"`
}

type Ec2FleetSpecLaunchTemplateConfigOverride struct {
	// +optional
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// +optional
	InstanceType string `json:"instance_type,omitempty"`
	// +optional
	MaxPrice string `json:"max_price,omitempty"`
	// +optional
	Priority json.Number `json:"priority,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
	// +optional
	WeightedCapacity json.Number `json:"weighted_capacity,omitempty"`
}

type Ec2FleetSpecLaunchTemplateConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateSpecification []Ec2FleetSpecLaunchTemplateConfig `json:"launch_template_specification"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Override *[]Ec2FleetSpecLaunchTemplateConfig `json:"override,omitempty"`
}

type Ec2FleetSpecOnDemandOptions struct {
	// +optional
	AllocationStrategy string `json:"allocation_strategy,omitempty"`
}

type Ec2FleetSpecSpotOptions struct {
	// +optional
	AllocationStrategy string `json:"allocation_strategy,omitempty"`
	// +optional
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior,omitempty"`
	// +optional
	InstancePoolsToUseCount int `json:"instance_pools_to_use_count,omitempty"`
}

type Ec2FleetSpecTargetCapacitySpecification struct {
	DefaultTargetCapacityType string `json:"default_target_capacity_type"`
	// +optional
	OnDemandTargetCapacity int `json:"on_demand_target_capacity,omitempty"`
	// +optional
	SpotTargetCapacity  int `json:"spot_target_capacity,omitempty"`
	TotalTargetCapacity int `json:"total_target_capacity"`
}

type Ec2FleetSpec struct {
	// +optional
	ExcessCapacityTerminationPolicy string `json:"excess_capacity_termination_policy,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateConfig []Ec2FleetSpec `json:"launch_template_config"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OnDemandOptions *[]Ec2FleetSpec `json:"on_demand_options,omitempty"`
	// +optional
	ReplaceUnhealthyInstances bool `json:"replace_unhealthy_instances,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SpotOptions *[]Ec2FleetSpec `json:"spot_options,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	TargetCapacitySpecification []Ec2FleetSpec `json:"target_capacity_specification"`
	// +optional
	TerminateInstances bool `json:"terminate_instances,omitempty"`
	// +optional
	TerminateInstancesWithExpiration bool `json:"terminate_instances_with_expiration,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type Ec2FleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2FleetList is a list of Ec2Fleets
type Ec2FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2Fleet CRD objects
	Items []Ec2Fleet `json:"items,omitempty"`
}
