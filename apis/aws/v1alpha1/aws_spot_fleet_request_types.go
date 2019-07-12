package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotFleetRequestSpec   `json:"spec,omitempty"`
	Status            AwsSpotFleetRequestStatus `json:"status,omitempty"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsSpotFleetRequestSpecLaunchSpecification struct {
	WeightedCapacity         string                                       `json:"weighted_capacity"`
	Tags                     map[string]string                            `json:"tags"`
	AssociatePublicIpAddress bool                                         `json:"associate_public_ip_address"`
	RootBlockDevice          []AwsSpotFleetRequestSpecLaunchSpecification `json:"root_block_device"`
	UserData                 string                                       `json:"user_data"`
	KeyName                  string                                       `json:"key_name"`
	PlacementGroup           string                                       `json:"placement_group"`
	SubnetId                 string                                       `json:"subnet_id"`
	AvailabilityZone         string                                       `json:"availability_zone"`
	EbsBlockDevice           []AwsSpotFleetRequestSpecLaunchSpecification `json:"ebs_block_device"`
	EphemeralBlockDevice     []AwsSpotFleetRequestSpecLaunchSpecification `json:"ephemeral_block_device"`
	EbsOptimized             bool                                         `json:"ebs_optimized"`
	Monitoring               bool                                         `json:"monitoring"`
	PlacementTenancy         string                                       `json:"placement_tenancy"`
	SpotPrice                string                                       `json:"spot_price"`
	VpcSecurityGroupIds      []string                                     `json:"vpc_security_group_ids"`
	IamInstanceProfileArn    string                                       `json:"iam_instance_profile_arn"`
	Ami                      string                                       `json:"ami"`
	IamInstanceProfile       string                                       `json:"iam_instance_profile"`
	InstanceType             string                                       `json:"instance_type"`
}

type AwsSpotFleetRequestSpec struct {
	IamFleetRole                     string                    `json:"iam_fleet_role"`
	TargetCapacity                   int                       `json:"target_capacity"`
	SpotPrice                        string                    `json:"spot_price"`
	ValidFrom                        string                    `json:"valid_from"`
	FleetType                        string                    `json:"fleet_type"`
	TargetGroupArns                  []string                  `json:"target_group_arns"`
	LaunchSpecification              []AwsSpotFleetRequestSpec `json:"launch_specification"`
	AllocationStrategy               string                    `json:"allocation_strategy"`
	InstancePoolsToUseCount          int                       `json:"instance_pools_to_use_count"`
	ExcessCapacityTerminationPolicy  string                    `json:"excess_capacity_termination_policy"`
	InstanceInterruptionBehaviour    string                    `json:"instance_interruption_behaviour"`
	ClientToken                      string                    `json:"client_token"`
	TerminateInstancesWithExpiration bool                      `json:"terminate_instances_with_expiration"`
	ValidUntil                       string                    `json:"valid_until"`
	LoadBalancers                    []string                  `json:"load_balancers"`
	ReplaceUnhealthyInstances        bool                      `json:"replace_unhealthy_instances"`
	WaitForFulfillment               bool                      `json:"wait_for_fulfillment"`
	SpotRequestState                 string                    `json:"spot_request_state"`
}

type AwsSpotFleetRequestStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequests
type AwsSpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotFleetRequest CRD objects
	Items []AwsSpotFleetRequest `json:"items,omitempty"`
}
