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

type AwsInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsInstanceSpecNetworkInterface struct {
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsInstanceSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsInstanceSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
}

type AwsInstanceSpec struct {
	CpuThreadsPerCore                 int               `json:"cpu_threads_per_core"`
	UserDataBase64                    string            `json:"user_data_base64"`
	PrivateDns                        string            `json:"private_dns"`
	InstanceInitiatedShutdownBehavior string            `json:"instance_initiated_shutdown_behavior"`
	Monitoring                        bool              `json:"monitoring"`
	CpuCoreCount                      int               `json:"cpu_core_count"`
	Tags                              map[string]string `json:"tags"`
	EbsBlockDevice                    []AwsInstanceSpec `json:"ebs_block_device"`
	Arn                               string            `json:"arn"`
	SubnetId                          string            `json:"subnet_id"`
	NetworkInterfaceId                string            `json:"network_interface_id"`
	EbsOptimized                      bool              `json:"ebs_optimized"`
	HostId                            string            `json:"host_id"`
	EphemeralBlockDevice              []AwsInstanceSpec `json:"ephemeral_block_device"`
	AvailabilityZone                  string            `json:"availability_zone"`
	UserData                          string            `json:"user_data"`
	PublicDns                         string            `json:"public_dns"`
	PublicIp                          string            `json:"public_ip"`
	SecurityGroups                    []string          `json:"security_groups"`
	VpcSecurityGroupIds               []string          `json:"vpc_security_group_ids"`
	NetworkInterface                  []AwsInstanceSpec `json:"network_interface"`
	Ami                               string            `json:"ami"`
	AssociatePublicIpAddress          bool              `json:"associate_public_ip_address"`
	PrimaryNetworkInterfaceId         string            `json:"primary_network_interface_id"`
	InstanceState                     string            `json:"instance_state"`
	Ipv6AddressCount                  int               `json:"ipv6_address_count"`
	KeyName                           string            `json:"key_name"`
	PasswordData                      string            `json:"password_data"`
	SourceDestCheck                   bool              `json:"source_dest_check"`
	PrivateIp                         string            `json:"private_ip"`
	DisableApiTermination             bool              `json:"disable_api_termination"`
	IamInstanceProfile                string            `json:"iam_instance_profile"`
	VolumeTags                        map[string]string `json:"volume_tags"`
	CreditSpecification               []AwsInstanceSpec `json:"credit_specification"`
	PlacementGroup                    string            `json:"placement_group"`
	InstanceType                      string            `json:"instance_type"`
	GetPasswordData                   bool              `json:"get_password_data"`
	Ipv6Addresses                     []string          `json:"ipv6_addresses"`
	Tenancy                           string            `json:"tenancy"`
	RootBlockDevice                   []AwsInstanceSpec `json:"root_block_device"`
}

type AwsInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsInstanceList is a list of AwsInstances
type AwsInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInstance CRD objects
	Items []AwsInstance `json:"items,omitempty"`
}
