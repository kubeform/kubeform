package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type InstanceSpecEbsBlockDevice struct {
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

type InstanceSpecEphemeralBlockDevice struct {
	DeviceName string `json:"deviceName" tf:"device_name"`
	// +optional
	NoDevice bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`
	// +optional
	VirtualName string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type InstanceSpecNetworkInterface struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceIndex         int    `json:"deviceIndex" tf:"device_index"`
	NetworkInterfaceID  string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type InstanceSpecRootBlockDevice struct {
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type InstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Ami         string                    `json:"ami" tf:"ami"`
	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	CpuCoreCount int `json:"cpuCoreCount,omitempty" tf:"cpu_core_count,omitempty"`
	// +optional
	CpuThreadsPerCore int `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification []InstanceSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`
	// +optional
	DisableAPITermination bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []InstanceSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []InstanceSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	GetPasswordData bool `json:"getPasswordData,omitempty" tf:"get_password_data,omitempty"`
	// +optional
	HostID string `json:"hostID,omitempty" tf:"host_id,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`
	InstanceType                      string `json:"instanceType" tf:"instance_type"`
	// +optional
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`
	// +optional
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NetworkInterface []InstanceSpecNetworkInterface `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`
	// +optional
	PlacementGroup string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RootBlockDevice []InstanceSpecRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SourceDestCheck bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	UserDataBase64 string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`
	// +optional
	VolumeTags map[string]string `json:"volumeTags,omitempty" tf:"volume_tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
