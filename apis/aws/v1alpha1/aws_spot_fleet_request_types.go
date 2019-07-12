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

type AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsSpotFleetRequestSpecLaunchSpecification struct {
	AvailabilityZone         string                                       `json:"availability_zone"`
	AssociatePublicIpAddress bool                                         `json:"associate_public_ip_address"`
	EbsBlockDevice           []AwsSpotFleetRequestSpecLaunchSpecification `json:"ebs_block_device"`
	RootBlockDevice          []AwsSpotFleetRequestSpecLaunchSpecification `json:"root_block_device"`
	SpotPrice                string                                       `json:"spot_price"`
	EphemeralBlockDevice     []AwsSpotFleetRequestSpecLaunchSpecification `json:"ephemeral_block_device"`
	EbsOptimized             bool                                         `json:"ebs_optimized"`
	WeightedCapacity         string                                       `json:"weighted_capacity"`
	PlacementTenancy         string                                       `json:"placement_tenancy"`
	SubnetId                 string                                       `json:"subnet_id"`
	Ami                      string                                       `json:"ami"`
	KeyName                  string                                       `json:"key_name"`
	Monitoring               bool                                         `json:"monitoring"`
	PlacementGroup           string                                       `json:"placement_group"`
	UserData                 string                                       `json:"user_data"`
	Tags                     map[string]string                            `json:"tags"`
	VpcSecurityGroupIds      []string                                     `json:"vpc_security_group_ids"`
	IamInstanceProfile       string                                       `json:"iam_instance_profile"`
	IamInstanceProfileArn    string                                       `json:"iam_instance_profile_arn"`
	InstanceType             string                                       `json:"instance_type"`
}

type AwsSpotFleetRequestSpec struct {
	TerminateInstancesWithExpiration bool                      `json:"terminate_instances_with_expiration"`
	LoadBalancers                    []string                  `json:"load_balancers"`
	InstanceInterruptionBehaviour    string                    `json:"instance_interruption_behaviour"`
	SpotPrice                        string                    `json:"spot_price"`
	IamFleetRole                     string                    `json:"iam_fleet_role"`
	LaunchSpecification              []AwsSpotFleetRequestSpec `json:"launch_specification"`
	AllocationStrategy               string                    `json:"allocation_strategy"`
	ValidFrom                        string                    `json:"valid_from"`
	ValidUntil                       string                    `json:"valid_until"`
	ClientToken                      string                    `json:"client_token"`
	WaitForFulfillment               bool                      `json:"wait_for_fulfillment"`
	TargetCapacity                   int                       `json:"target_capacity"`
	InstancePoolsToUseCount          int                       `json:"instance_pools_to_use_count"`
	ExcessCapacityTerminationPolicy  string                    `json:"excess_capacity_termination_policy"`
	FleetType                        string                    `json:"fleet_type"`
	SpotRequestState                 string                    `json:"spot_request_state"`
	TargetGroupArns                  []string                  `json:"target_group_arns"`
	ReplaceUnhealthyInstances        bool                      `json:"replace_unhealthy_instances"`
}

type AwsSpotFleetRequestStatus struct {
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
