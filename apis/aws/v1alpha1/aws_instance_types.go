package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInstanceSpec   `json:"spec,omitempty"`
	Status            AwsInstanceStatus `json:"status,omitempty"`
}

type AwsInstanceSpecEbsBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type AwsInstanceSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsInstanceSpecEphemeralBlockDevice struct {
	NoDevice    bool   `json:"no_device"`
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsInstanceSpecNetworkInterface struct {
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsInstanceSpecRootBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsInstanceSpec struct {
	EbsBlockDevice                    []AwsInstanceSpec `json:"ebs_block_device"`
	NetworkInterfaceId                string            `json:"network_interface_id"`
	InstanceState                     string            `json:"instance_state"`
	Ipv6Addresses                     []string          `json:"ipv6_addresses"`
	Tenancy                           string            `json:"tenancy"`
	PrivateIp                         string            `json:"private_ip"`
	CreditSpecification               []AwsInstanceSpec `json:"credit_specification"`
	InstanceInitiatedShutdownBehavior string            `json:"instance_initiated_shutdown_behavior"`
	Arn                               string            `json:"arn"`
	GetPasswordData                   bool              `json:"get_password_data"`
	SubnetId                          string            `json:"subnet_id"`
	EbsOptimized                      bool              `json:"ebs_optimized"`
	PrimaryNetworkInterfaceId         string            `json:"primary_network_interface_id"`
	HostId                            string            `json:"host_id"`
	UserData                          string            `json:"user_data"`
	UserDataBase64                    string            `json:"user_data_base64"`
	SecurityGroups                    []string          `json:"security_groups"`
	PublicDns                         string            `json:"public_dns"`
	IamInstanceProfile                string            `json:"iam_instance_profile"`
	CpuCoreCount                      int               `json:"cpu_core_count"`
	EphemeralBlockDevice              []AwsInstanceSpec `json:"ephemeral_block_device"`
	Ami                               string            `json:"ami"`
	KeyName                           string            `json:"key_name"`
	SourceDestCheck                   bool              `json:"source_dest_check"`
	PrivateDns                        string            `json:"private_dns"`
	VpcSecurityGroupIds               []string          `json:"vpc_security_group_ids"`
	NetworkInterface                  []AwsInstanceSpec `json:"network_interface"`
	DisableApiTermination             bool              `json:"disable_api_termination"`
	CpuThreadsPerCore                 int               `json:"cpu_threads_per_core"`
	AssociatePublicIpAddress          bool              `json:"associate_public_ip_address"`
	AvailabilityZone                  string            `json:"availability_zone"`
	PlacementGroup                    string            `json:"placement_group"`
	InstanceType                      string            `json:"instance_type"`
	Tags                              map[string]string `json:"tags"`
	VolumeTags                        map[string]string `json:"volume_tags"`
	PasswordData                      string            `json:"password_data"`
	PublicIp                          string            `json:"public_ip"`
	Monitoring                        bool              `json:"monitoring"`
	RootBlockDevice                   []AwsInstanceSpec `json:"root_block_device"`
	Ipv6AddressCount                  int               `json:"ipv6_address_count"`
}



type AwsInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstances
type AwsInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInstance CRD objects
	Items []AwsInstance `json:"items,omitempty"`
}