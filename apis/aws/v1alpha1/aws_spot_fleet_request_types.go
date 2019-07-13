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

type AwsSpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotFleetRequestSpec   `json:"spec,omitempty"`
	Status            AwsSpotFleetRequestStatus `json:"status,omitempty"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsSpotFleetRequestSpecLaunchSpecification struct {
	Monitoring               bool                                         `json:"monitoring"`
	EphemeralBlockDevice     []AwsSpotFleetRequestSpecLaunchSpecification `json:"ephemeral_block_device"`
	EbsOptimized             bool                                         `json:"ebs_optimized"`
	InstanceType             string                                       `json:"instance_type"`
	AvailabilityZone         string                                       `json:"availability_zone"`
	Tags                     map[string]string                            `json:"tags"`
	VpcSecurityGroupIds      []string                                     `json:"vpc_security_group_ids"`
	PlacementGroup           string                                       `json:"placement_group"`
	PlacementTenancy         string                                       `json:"placement_tenancy"`
	SpotPrice                string                                       `json:"spot_price"`
	SubnetId                 string                                       `json:"subnet_id"`
	RootBlockDevice          []AwsSpotFleetRequestSpecLaunchSpecification `json:"root_block_device"`
	Ami                      string                                       `json:"ami"`
	KeyName                  string                                       `json:"key_name"`
	IamInstanceProfileArn    string                                       `json:"iam_instance_profile_arn"`
	UserData                 string                                       `json:"user_data"`
	WeightedCapacity         string                                       `json:"weighted_capacity"`
	AssociatePublicIpAddress bool                                         `json:"associate_public_ip_address"`
	EbsBlockDevice           []AwsSpotFleetRequestSpecLaunchSpecification `json:"ebs_block_device"`
	IamInstanceProfile       string                                       `json:"iam_instance_profile"`
}

type AwsSpotFleetRequestSpec struct {
	ExcessCapacityTerminationPolicy  string                    `json:"excess_capacity_termination_policy"`
	ValidUntil                       string                    `json:"valid_until"`
	FleetType                        string                    `json:"fleet_type"`
	TargetGroupArns                  []string                  `json:"target_group_arns"`
	WaitForFulfillment               bool                      `json:"wait_for_fulfillment"`
	LaunchSpecification              []AwsSpotFleetRequestSpec `json:"launch_specification"`
	SpotPrice                        string                    `json:"spot_price"`
	LoadBalancers                    []string                  `json:"load_balancers"`
	ReplaceUnhealthyInstances        bool                      `json:"replace_unhealthy_instances"`
	AllocationStrategy               string                    `json:"allocation_strategy"`
	TerminateInstancesWithExpiration bool                      `json:"terminate_instances_with_expiration"`
	SpotRequestState                 string                    `json:"spot_request_state"`
	IamFleetRole                     string                    `json:"iam_fleet_role"`
	InstanceInterruptionBehaviour    string                    `json:"instance_interruption_behaviour"`
	ValidFrom                        string                    `json:"valid_from"`
	ClientToken                      string                    `json:"client_token"`
	TargetCapacity                   int                       `json:"target_capacity"`
	InstancePoolsToUseCount          int                       `json:"instance_pools_to_use_count"`
}



type AwsSpotFleetRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequests
type AwsSpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotFleetRequest CRD objects
	Items []AwsSpotFleetRequest `json:"items,omitempty"`
}