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

type AwsSpotInstanceRequestSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
}

type AwsSpotInstanceRequestSpecNetworkInterface struct {
	DeviceIndex         int    `json:"device_index"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

type AwsSpotInstanceRequestSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsSpotInstanceRequestSpecEphemeralBlockDevice struct {
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
	DeviceName  string `json:"device_name"`
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

type AwsSpotInstanceRequestSpec struct {
	AssociatePublicIpAddress          bool                         `json:"associate_public_ip_address"`
	InstanceType                      string                       `json:"instance_type"`
	UserData                          string                       `json:"user_data"`
	Monitoring                        bool                         `json:"monitoring"`
	RootBlockDevice                   []AwsSpotInstanceRequestSpec `json:"root_block_device"`
	LaunchGroup                       string                       `json:"launch_group"`
	SpotRequestState                  string                       `json:"spot_request_state"`
	KeyName                           string                       `json:"key_name"`
	PasswordData                      string                       `json:"password_data"`
	SourceDestCheck                   bool                         `json:"source_dest_check"`
	PrivateDns                        string                       `json:"private_dns"`
	IamInstanceProfile                string                       `json:"iam_instance_profile"`
	SpotType                          string                       `json:"spot_type"`
	AvailabilityZone                  string                       `json:"availability_zone"`
	GetPasswordData                   bool                         `json:"get_password_data"`
	VpcSecurityGroupIds               []string                     `json:"vpc_security_group_ids"`
	PublicIp                          string                       `json:"public_ip"`
	Ipv6Addresses                     []string                     `json:"ipv6_addresses"`
	InstanceInterruptionBehaviour     string                       `json:"instance_interruption_behaviour"`
	WaitForFulfillment                bool                         `json:"wait_for_fulfillment"`
	ValidUntil                        string                       `json:"valid_until"`
	PrivateIp                         string                       `json:"private_ip"`
	PublicDns                         string                       `json:"public_dns"`
	NetworkInterface                  []AwsSpotInstanceRequestSpec `json:"network_interface"`
	InstanceInitiatedShutdownBehavior string                       `json:"instance_initiated_shutdown_behavior"`
	HostId                            string                       `json:"host_id"`
	CpuThreadsPerCore                 int                          `json:"cpu_threads_per_core"`
	SpotPrice                         string                       `json:"spot_price"`
	ValidFrom                         string                       `json:"valid_from"`
	SubnetId                          string                       `json:"subnet_id"`
	PrimaryNetworkInterfaceId         string                       `json:"primary_network_interface_id"`
	DisableApiTermination             bool                         `json:"disable_api_termination"`
	Ipv6AddressCount                  int                          `json:"ipv6_address_count"`
	Tags                              map[string]string            `json:"tags"`
	CreditSpecification               []AwsSpotInstanceRequestSpec `json:"credit_specification"`
	Arn                               string                       `json:"arn"`
	Tenancy                           string                       `json:"tenancy"`
	EphemeralBlockDevice              []AwsSpotInstanceRequestSpec `json:"ephemeral_block_device"`
	PlacementGroup                    string                       `json:"placement_group"`
	NetworkInterfaceId                string                       `json:"network_interface_id"`
	CpuCoreCount                      int                          `json:"cpu_core_count"`
	EbsBlockDevice                    []AwsSpotInstanceRequestSpec `json:"ebs_block_device"`
	SpotBidStatus                     string                       `json:"spot_bid_status"`
	Ami                               string                       `json:"ami"`
	UserDataBase64                    string                       `json:"user_data_base64"`
	SecurityGroups                    []string                     `json:"security_groups"`
	InstanceState                     string                       `json:"instance_state"`
	EbsOptimized                      bool                         `json:"ebs_optimized"`
	VolumeTags                        map[string]string            `json:"volume_tags"`
	SpotInstanceId                    string                       `json:"spot_instance_id"`
	BlockDurationMinutes              int                          `json:"block_duration_minutes"`
}



type AwsSpotInstanceRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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