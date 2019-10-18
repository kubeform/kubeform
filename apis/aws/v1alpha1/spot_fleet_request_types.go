package v1alpha1

import (
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

type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotFleetRequestSpec   `json:"spec,omitempty"`
	Status            SpotFleetRequestStatus `json:"status,omitempty"`
}

type SpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"deviceName" tf:"device_name"`
	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type SpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotFleetRequestSpecLaunchSpecification struct {
	Ami string `json:"ami" tf:"ami"`
	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	EbsBlockDevice []SpotFleetRequestSpecLaunchSpecificationEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	EphemeralBlockDevice []SpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	IamInstanceProfileArn string `json:"iamInstanceProfileArn,omitempty" tf:"iam_instance_profile_arn,omitempty"`
	InstanceType          string `json:"instanceType" tf:"instance_type"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	PlacementGroup string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`
	// +optional
	PlacementTenancy string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`
	// +optional
	RootBlockDevice []SpotFleetRequestSpecLaunchSpecificationRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`
	// +optional
	SpotPrice string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
	// +optional
	WeightedCapacity string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type SpotFleetRequestSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllocationStrategy string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`
	// +optional
	ClientToken string `json:"clientToken,omitempty" tf:"client_token,omitempty"`
	// +optional
	ExcessCapacityTerminationPolicy string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy,omitempty"`
	// +optional
	FleetType    string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`
	IamFleetRole string `json:"iamFleetRole" tf:"iam_fleet_role"`
	// +optional
	InstanceInterruptionBehaviour string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`
	// +optional
	InstancePoolsToUseCount int                                       `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
	LaunchSpecification     []SpotFleetRequestSpecLaunchSpecification `json:"launchSpecification" tf:"launch_specification"`
	// +optional
	LoadBalancers []string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`
	// +optional
	ReplaceUnhealthyInstances bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances,omitempty"`
	// +optional
	SpotPrice string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	// +optional
	SpotRequestState string `json:"spotRequestState,omitempty" tf:"spot_request_state,omitempty"`
	TargetCapacity   int    `json:"targetCapacity" tf:"target_capacity"`
	// +optional
	TargetGroupArns []string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`
	// +optional
	TerminateInstancesWithExpiration bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration,omitempty"`
	// +optional
	ValidFrom string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`
	// +optional
	ValidUntil string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
	// +optional
	WaitForFulfillment bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
}

type SpotFleetRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SpotFleetRequestSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
