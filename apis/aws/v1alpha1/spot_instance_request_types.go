package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotInstanceRequestSpec   `json:"spec,omitempty"`
	Status            SpotInstanceRequestStatus `json:"status,omitempty"`
}

type SpotInstanceRequestSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type SpotInstanceRequestSpecEbsBlockDevice struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	VolumeID   string `json:"volumeID" tf:"volume_id"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotInstanceRequestSpecEphemeralBlockDevice struct {
	DeviceName string `json:"deviceName" tf:"device_name"`
	// +optional
	NoDevice bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`
	// +optional
	VirtualName string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type SpotInstanceRequestSpecNetworkInterface struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceIndex         int    `json:"deviceIndex" tf:"device_index"`
	NetworkInterfaceID  string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type SpotInstanceRequestSpecRootBlockDevice struct {
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Iops     int    `json:"iops,omitempty" tf:"iops,omitempty"`
	VolumeID string `json:"volumeID" tf:"volume_id"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotInstanceRequestSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Ami string `json:"ami" tf:"ami"`
	Arn string `json:"arn" tf:"arn"`
	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`
	// +optional
	CpuCoreCount int `json:"cpuCoreCount,omitempty" tf:"cpu_core_count,omitempty"`
	// +optional
	CpuThreadsPerCore int `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification []SpotInstanceRequestSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`
	// +optional
	DisableAPITermination bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []SpotInstanceRequestSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []SpotInstanceRequestSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	GetPasswordData bool `json:"getPasswordData,omitempty" tf:"get_password_data,omitempty"`
	// +optional
	HostID string `json:"hostID,omitempty" tf:"host_id,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`
	// +optional
	InstanceInterruptionBehaviour string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`
	InstanceState                 string `json:"instanceState" tf:"instance_state"`
	InstanceType                  string `json:"instanceType" tf:"instance_type"`
	// +optional
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`
	// +optional
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	LaunchGroup string `json:"launchGroup,omitempty" tf:"launch_group,omitempty"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NetworkInterface []SpotInstanceRequestSpecNetworkInterface `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`
	PasswordData     string                                    `json:"passwordData" tf:"password_data"`
	// +optional
	PlacementGroup            string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`
	PrimaryNetworkInterfaceID string `json:"primaryNetworkInterfaceID" tf:"primary_network_interface_id"`
	PrivateDNS                string `json:"privateDNS" tf:"private_dns"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	PublicDNS string `json:"publicDNS" tf:"public_dns"`
	PublicIP  string `json:"publicIP" tf:"public_ip"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RootBlockDevice []SpotInstanceRequestSpecRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SourceDestCheck bool   `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`
	SpotBidStatus   string `json:"spotBidStatus" tf:"spot_bid_status"`
	SpotInstanceID  string `json:"spotInstanceID" tf:"spot_instance_id"`
	// +optional
	SpotPrice        string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	SpotRequestState string `json:"spotRequestState" tf:"spot_request_state"`
	// +optional
	SpotType string `json:"spotType,omitempty" tf:"spot_type,omitempty"`
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
	ValidFrom string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`
	// +optional
	ValidUntil string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
	// +optional
	VolumeTags map[string]string `json:"volumeTags,omitempty" tf:"volume_tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
	// +optional
	WaitForFulfillment bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
}

type SpotInstanceRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SpotInstanceRequestSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpotInstanceRequestList is a list of SpotInstanceRequests
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpotInstanceRequest CRD objects
	Items []SpotInstanceRequest `json:"items,omitempty"`
}
