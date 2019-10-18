package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Ec2Fleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2FleetSpec   `json:"spec,omitempty"`
	Status            Ec2FleetStatus `json:"status,omitempty"`
}

type Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification struct {
	// +optional
	LaunchTemplateID string `json:"launchTemplateID,omitempty" tf:"launch_template_id,omitempty"`
	// +optional
	LaunchTemplateName string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`
	Version            string `json:"version" tf:"version"`
}

type Ec2FleetSpecLaunchTemplateConfigOverride struct {
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	// +optional
	MaxPrice string `json:"maxPrice,omitempty" tf:"max_price,omitempty"`
	// +optional
	Priority json.Number `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	WeightedCapacity json.Number `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type Ec2FleetSpecLaunchTemplateConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateSpecification []Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification `json:"launchTemplateSpecification" tf:"launch_template_specification"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Override []Ec2FleetSpecLaunchTemplateConfigOverride `json:"override,omitempty" tf:"override,omitempty"`
}

type Ec2FleetSpecOnDemandOptions struct {
	// +optional
	AllocationStrategy string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`
}

type Ec2FleetSpecSpotOptions struct {
	// +optional
	AllocationStrategy string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`
	// +optional
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`
	// +optional
	InstancePoolsToUseCount int `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type Ec2FleetSpecTargetCapacitySpecification struct {
	DefaultTargetCapacityType string `json:"defaultTargetCapacityType" tf:"default_target_capacity_type"`
	// +optional
	OnDemandTargetCapacity int `json:"onDemandTargetCapacity,omitempty" tf:"on_demand_target_capacity,omitempty"`
	// +optional
	SpotTargetCapacity  int `json:"spotTargetCapacity,omitempty" tf:"spot_target_capacity,omitempty"`
	TotalTargetCapacity int `json:"totalTargetCapacity" tf:"total_target_capacity"`
}

type Ec2FleetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ExcessCapacityTerminationPolicy string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	LaunchTemplateConfig []Ec2FleetSpecLaunchTemplateConfig `json:"launchTemplateConfig" tf:"launch_template_config"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OnDemandOptions []Ec2FleetSpecOnDemandOptions `json:"onDemandOptions,omitempty" tf:"on_demand_options,omitempty"`
	// +optional
	ReplaceUnhealthyInstances bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SpotOptions []Ec2FleetSpecSpotOptions `json:"spotOptions,omitempty" tf:"spot_options,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	TargetCapacitySpecification []Ec2FleetSpecTargetCapacitySpecification `json:"targetCapacitySpecification" tf:"target_capacity_specification"`
	// +optional
	TerminateInstances bool `json:"terminateInstances,omitempty" tf:"terminate_instances,omitempty"`
	// +optional
	TerminateInstancesWithExpiration bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type Ec2FleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Ec2FleetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
