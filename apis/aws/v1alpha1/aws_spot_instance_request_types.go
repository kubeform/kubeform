package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotInstanceRequestSpec   `json:"spec,omitempty"`
	Status            AwsSpotInstanceRequestStatus `json:"status,omitempty"`
}

type AwsSpotInstanceRequestSpecNetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
}

type AwsSpotInstanceRequestSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsSpotInstanceRequestSpecEbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
}

type AwsSpotInstanceRequestSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsSpotInstanceRequestSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
}

type AwsSpotInstanceRequestSpec struct {
	InstanceState                     string                       `json:"instance_state"`
	Ipv6AddressCount                  int                          `json:"ipv6_address_count"`
	CpuThreadsPerCore                 int                          `json:"cpu_threads_per_core"`
	Tags                              map[string]string            `json:"tags"`
	LaunchGroup                       string                       `json:"launch_group"`
	Ami                               string                       `json:"ami"`
	SubnetId                          string                       `json:"subnet_id"`
	VpcSecurityGroupIds               []string                     `json:"vpc_security_group_ids"`
	NetworkInterface                  []AwsSpotInstanceRequestSpec `json:"network_interface"`
	InstanceInitiatedShutdownBehavior string                       `json:"instance_initiated_shutdown_behavior"`
	CreditSpecification               []AwsSpotInstanceRequestSpec `json:"credit_specification"`
	InstanceType                      string                       `json:"instance_type"`
	SecurityGroups                    []string                     `json:"security_groups"`
	PublicDns                         string                       `json:"public_dns"`
	Monitoring                        bool                         `json:"monitoring"`
	VolumeTags                        map[string]string            `json:"volume_tags"`
	SpotInstanceId                    string                       `json:"spot_instance_id"`
	KeyName                           string                       `json:"key_name"`
	UserData                          string                       `json:"user_data"`
	Ipv6Addresses                     []string                     `json:"ipv6_addresses"`
	Tenancy                           string                       `json:"tenancy"`
	EbsBlockDevice                    []AwsSpotInstanceRequestSpec `json:"ebs_block_device"`
	SpotType                          string                       `json:"spot_type"`
	PrivateIp                         string                       `json:"private_ip"`
	PasswordData                      string                       `json:"password_data"`
	SourceDestCheck                   bool                         `json:"source_dest_check"`
	PrimaryNetworkInterfaceId         string                       `json:"primary_network_interface_id"`
	WaitForFulfillment                bool                         `json:"wait_for_fulfillment"`
	ValidFrom                         string                       `json:"valid_from"`
	GetPasswordData                   bool                         `json:"get_password_data"`
	PublicIp                          string                       `json:"public_ip"`
	PrivateDns                        string                       `json:"private_dns"`
	EbsOptimized                      bool                         `json:"ebs_optimized"`
	CpuCoreCount                      int                          `json:"cpu_core_count"`
	EphemeralBlockDevice              []AwsSpotInstanceRequestSpec `json:"ephemeral_block_device"`
	SpotPrice                         string                       `json:"spot_price"`
	AssociatePublicIpAddress          bool                         `json:"associate_public_ip_address"`
	UserDataBase64                    string                       `json:"user_data_base64"`
	IamInstanceProfile                string                       `json:"iam_instance_profile"`
	HostId                            string                       `json:"host_id"`
	SpotBidStatus                     string                       `json:"spot_bid_status"`
	BlockDurationMinutes              int                          `json:"block_duration_minutes"`
	InstanceInterruptionBehaviour     string                       `json:"instance_interruption_behaviour"`
	ValidUntil                        string                       `json:"valid_until"`
	PlacementGroup                    string                       `json:"placement_group"`
	AvailabilityZone                  string                       `json:"availability_zone"`
	NetworkInterfaceId                string                       `json:"network_interface_id"`
	DisableApiTermination             bool                         `json:"disable_api_termination"`
	RootBlockDevice                   []AwsSpotInstanceRequestSpec `json:"root_block_device"`
	SpotRequestState                  string                       `json:"spot_request_state"`
	Arn                               string                       `json:"arn"`
}



type AwsSpotInstanceRequestStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequests
type AwsSpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotInstanceRequest CRD objects
	Items []AwsSpotInstanceRequest `json:"items,omitempty"`
}