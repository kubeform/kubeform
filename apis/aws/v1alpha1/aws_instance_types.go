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

type AwsInstanceSpecRootBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
}

type AwsInstanceSpecNetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
}

type AwsInstanceSpecEbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
}

type AwsInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsInstanceSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsInstanceSpec struct {
	NetworkInterfaceId                string            `json:"network_interface_id"`
	PrivateDns                        string            `json:"private_dns"`
	Monitoring                        bool              `json:"monitoring"`
	RootBlockDevice                   []AwsInstanceSpec `json:"root_block_device"`
	PrivateIp                         string            `json:"private_ip"`
	InstanceType                      string            `json:"instance_type"`
	SecurityGroups                    []string          `json:"security_groups"`
	Tenancy                           string            `json:"tenancy"`
	VolumeTags                        map[string]string `json:"volume_tags"`
	AvailabilityZone                  string            `json:"availability_zone"`
	UserData                          string            `json:"user_data"`
	NetworkInterface                  []AwsInstanceSpec `json:"network_interface"`
	PublicIp                          string            `json:"public_ip"`
	EbsBlockDevice                    []AwsInstanceSpec `json:"ebs_block_device"`
	EphemeralBlockDevice              []AwsInstanceSpec `json:"ephemeral_block_device"`
	Arn                               string            `json:"arn"`
	PasswordData                      string            `json:"password_data"`
	PublicDns                         string            `json:"public_dns"`
	EbsOptimized                      bool              `json:"ebs_optimized"`
	DisableApiTermination             bool              `json:"disable_api_termination"`
	Ipv6Addresses                     []string          `json:"ipv6_addresses"`
	CpuThreadsPerCore                 int               `json:"cpu_threads_per_core"`
	KeyName                           string            `json:"key_name"`
	PlacementGroup                    string            `json:"placement_group"`
	SourceDestCheck                   bool              `json:"source_dest_check"`
	UserDataBase64                    string            `json:"user_data_base64"`
	InstanceState                     string            `json:"instance_state"`
	Tags                              map[string]string `json:"tags"`
	Ami                               string            `json:"ami"`
	VpcSecurityGroupIds               []string          `json:"vpc_security_group_ids"`
	Ipv6AddressCount                  int               `json:"ipv6_address_count"`
	HostId                            string            `json:"host_id"`
	GetPasswordData                   bool              `json:"get_password_data"`
	SubnetId                          string            `json:"subnet_id"`
	InstanceInitiatedShutdownBehavior string            `json:"instance_initiated_shutdown_behavior"`
	IamInstanceProfile                string            `json:"iam_instance_profile"`
	CreditSpecification               []AwsInstanceSpec `json:"credit_specification"`
	AssociatePublicIpAddress          bool              `json:"associate_public_ip_address"`
	CpuCoreCount                      int               `json:"cpu_core_count"`
	PrimaryNetworkInterfaceId         string            `json:"primary_network_interface_id"`
}



type AwsInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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