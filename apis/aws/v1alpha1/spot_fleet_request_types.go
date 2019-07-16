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

type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotFleetRequestSpec   `json:"spec,omitempty"`
	Status            SpotFleetRequestStatus `json:"status,omitempty"`
}

type SpotFleetRequestSpecLaunchSpecification struct {
	Ami string `json:"ami"`
	// +optional
	AssociatePublicIpAddress bool `json:"associate_public_ip_address,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebs_optimized,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iam_instance_profile,omitempty"`
	// +optional
	IamInstanceProfileArn string `json:"iam_instance_profile_arn,omitempty"`
	InstanceType          string `json:"instance_type"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty"`
	// +optional
	PlacementTenancy string `json:"placement_tenancy,omitempty"`
	// +optional
	SpotPrice string `json:"spot_price,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	UserData string `json:"user_data,omitempty"`
	// +optional
	WeightedCapacity string `json:"weighted_capacity,omitempty"`
}

type SpotFleetRequestSpec struct {
	// +optional
	AllocationStrategy string `json:"allocation_strategy,omitempty"`
	// +optional
	ExcessCapacityTerminationPolicy string `json:"excess_capacity_termination_policy,omitempty"`
	// +optional
	FleetType    string `json:"fleet_type,omitempty"`
	IamFleetRole string `json:"iam_fleet_role"`
	// +optional
	InstanceInterruptionBehaviour string `json:"instance_interruption_behaviour,omitempty"`
	// +optional
	InstancePoolsToUseCount int `json:"instance_pools_to_use_count,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	LaunchSpecification []SpotFleetRequestSpec `json:"launch_specification"`
	// +optional
	ReplaceUnhealthyInstances bool `json:"replace_unhealthy_instances,omitempty"`
	// +optional
	SpotPrice      string `json:"spot_price,omitempty"`
	TargetCapacity int    `json:"target_capacity"`
	// +optional
	TerminateInstancesWithExpiration bool `json:"terminate_instances_with_expiration,omitempty"`
	// +optional
	ValidFrom string `json:"valid_from,omitempty"`
	// +optional
	ValidUntil string `json:"valid_until,omitempty"`
	// +optional
	WaitForFulfillment bool `json:"wait_for_fulfillment,omitempty"`
}

type SpotFleetRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpotFleetRequestList is a list of SpotFleetRequests
type SpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpotFleetRequest CRD objects
	Items []SpotFleetRequest `json:"items,omitempty"`
}
