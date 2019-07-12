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

type AwsSpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotInstanceRequestSpec   `json:"spec,omitempty"`
	Status            AwsSpotInstanceRequestStatus `json:"status,omitempty"`
}

type AwsSpotInstanceRequestSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
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

type AwsSpotInstanceRequestSpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
}

type AwsSpotInstanceRequestSpecNetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
}

type AwsSpotInstanceRequestSpec struct {
	NetworkInterfaceId                string                       `json:"network_interface_id"`
	PrivateDns                        string                       `json:"private_dns"`
	SpotPrice                         string                       `json:"spot_price"`
	LaunchGroup                       string                       `json:"launch_group"`
	SpotRequestState                  string                       `json:"spot_request_state"`
	InstanceType                      string                       `json:"instance_type"`
	SourceDestCheck                   bool                         `json:"source_dest_check"`
	InstanceInitiatedShutdownBehavior string                       `json:"instance_initiated_shutdown_behavior"`
	CreditSpecification               []AwsSpotInstanceRequestSpec `json:"credit_specification"`
	EbsOptimized                      bool                         `json:"ebs_optimized"`
	Ipv6Addresses                     []string                     `json:"ipv6_addresses"`
	CpuCoreCount                      int                          `json:"cpu_core_count"`
	VolumeTags                        map[string]string            `json:"volume_tags"`
	PrivateIp                         string                       `json:"private_ip"`
	SecurityGroups                    []string                     `json:"security_groups"`
	VpcSecurityGroupIds               []string                     `json:"vpc_security_group_ids"`
	PublicIp                          string                       `json:"public_ip"`
	EphemeralBlockDevice              []AwsSpotInstanceRequestSpec `json:"ephemeral_block_device"`
	RootBlockDevice                   []AwsSpotInstanceRequestSpec `json:"root_block_device"`
	SpotBidStatus                     string                       `json:"spot_bid_status"`
	GetPasswordData                   bool                         `json:"get_password_data"`
	Monitoring                        bool                         `json:"monitoring"`
	Tags                              map[string]string            `json:"tags"`
	SpotType                          string                       `json:"spot_type"`
	EbsBlockDevice                    []AwsSpotInstanceRequestSpec `json:"ebs_block_device"`
	Arn                               string                       `json:"arn"`
	AssociatePublicIpAddress          bool                         `json:"associate_public_ip_address"`
	NetworkInterface                  []AwsSpotInstanceRequestSpec `json:"network_interface"`
	CpuThreadsPerCore                 int                          `json:"cpu_threads_per_core"`
	PlacementGroup                    string                       `json:"placement_group"`
	PublicDns                         string                       `json:"public_dns"`
	Ipv6AddressCount                  int                          `json:"ipv6_address_count"`
	SpotInstanceId                    string                       `json:"spot_instance_id"`
	UserDataBase64                    string                       `json:"user_data_base64"`
	Tenancy                           string                       `json:"tenancy"`
	WaitForFulfillment                bool                         `json:"wait_for_fulfillment"`
	ValidFrom                         string                       `json:"valid_from"`
	SubnetId                          string                       `json:"subnet_id"`
	UserData                          string                       `json:"user_data"`
	PrimaryNetworkInterfaceId         string                       `json:"primary_network_interface_id"`
	InstanceState                     string                       `json:"instance_state"`
	Ami                               string                       `json:"ami"`
	AvailabilityZone                  string                       `json:"availability_zone"`
	KeyName                           string                       `json:"key_name"`
	PasswordData                      string                       `json:"password_data"`
	InstanceInterruptionBehaviour     string                       `json:"instance_interruption_behaviour"`
	ValidUntil                        string                       `json:"valid_until"`
	DisableApiTermination             bool                         `json:"disable_api_termination"`
	IamInstanceProfile                string                       `json:"iam_instance_profile"`
	HostId                            string                       `json:"host_id"`
	BlockDurationMinutes              int                          `json:"block_duration_minutes"`
}

type AwsSpotInstanceRequestStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequests
type AwsSpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotInstanceRequest CRD objects
	Items []AwsSpotInstanceRequest `json:"items,omitempty"`
}
