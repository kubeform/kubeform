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

type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec,omitempty"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

type LaunchTemplateSpecBlockDeviceMappingsEbs struct {
	// +optional
	DeleteOnTermination string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Encrypted string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type LaunchTemplateSpecBlockDeviceMappings struct {
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ebs []LaunchTemplateSpecBlockDeviceMappingsEbs `json:"ebs,omitempty" tf:"ebs,omitempty"`
	// +optional
	NoDevice string `json:"noDevice,omitempty" tf:"no_device,omitempty"`
	// +optional
	VirtualName string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	// +optional
	CapacityReservationID string `json:"capacityReservationID,omitempty" tf:"capacity_reservation_id,omitempty"`
}

type LaunchTemplateSpecCapacityReservationSpecification struct {
	// +optional
	CapacityReservationPreference string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CapacityReservationTarget []LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target,omitempty"`
}

type LaunchTemplateSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type LaunchTemplateSpecElasticGpuSpecifications struct {
	Type string `json:"type" tf:"type"`
}

type LaunchTemplateSpecElasticInferenceAccelerator struct {
	Type string `json:"type" tf:"type"`
}

type LaunchTemplateSpecIamInstanceProfile struct {
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type LaunchTemplateSpecInstanceMarketOptionsSpotOptions struct {
	// +optional
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`
	// +optional
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`
	// +optional
	MaxPrice string `json:"maxPrice,omitempty" tf:"max_price,omitempty"`
	// +optional
	SpotInstanceType string `json:"spotInstanceType,omitempty" tf:"spot_instance_type,omitempty"`
	// +optional
	ValidUntil string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type LaunchTemplateSpecInstanceMarketOptions struct {
	// +optional
	MarketType string `json:"marketType,omitempty" tf:"market_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SpotOptions []LaunchTemplateSpecInstanceMarketOptionsSpotOptions `json:"spotOptions,omitempty" tf:"spot_options,omitempty"`
}

type LaunchTemplateSpecLicenseSpecification struct {
	LicenseConfigurationArn string `json:"licenseConfigurationArn" tf:"license_configuration_arn"`
}

type LaunchTemplateSpecMonitoring struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LaunchTemplateSpecNetworkInterfaces struct {
	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DeviceIndex int `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`
	// +optional
	Ipv4AddressCount int `json:"ipv4AddressCount,omitempty" tf:"ipv4_address_count,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty" tf:"ipv4_addresses,omitempty"`
	// +optional
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type LaunchTemplateSpecPlacement struct {
	// +optional
	Affinity string `json:"affinity,omitempty" tf:"affinity,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	GroupName string `json:"groupName,omitempty" tf:"group_name,omitempty"`
	// +optional
	HostID string `json:"hostID,omitempty" tf:"host_id,omitempty"`
	// +optional
	SpreadDomain string `json:"spreadDomain,omitempty" tf:"spread_domain,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type LaunchTemplateSpecTagSpecifications struct {
	// +optional
	ResourceType string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LaunchTemplateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	BlockDeviceMappings []LaunchTemplateSpecBlockDeviceMappings `json:"blockDeviceMappings,omitempty" tf:"block_device_mappings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CapacityReservationSpecification []LaunchTemplateSpecCapacityReservationSpecification `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification []LaunchTemplateSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DisableAPITermination bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination,omitempty"`
	// +optional
	EbsOptimized string `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	ElasticGpuSpecifications []LaunchTemplateSpecElasticGpuSpecifications `json:"elasticGpuSpecifications,omitempty" tf:"elastic_gpu_specifications,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticInferenceAccelerator []LaunchTemplateSpecElasticInferenceAccelerator `json:"elasticInferenceAccelerator,omitempty" tf:"elastic_inference_accelerator,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IamInstanceProfile []LaunchTemplateSpecIamInstanceProfile `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	ImageID string `json:"imageID,omitempty" tf:"image_id,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InstanceMarketOptions []LaunchTemplateSpecInstanceMarketOptions `json:"instanceMarketOptions,omitempty" tf:"instance_market_options,omitempty"`
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	// +optional
	KernelID string `json:"kernelID,omitempty" tf:"kernel_id,omitempty"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LicenseSpecification []LaunchTemplateSpecLicenseSpecification `json:"licenseSpecification,omitempty" tf:"license_specification,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Monitoring []LaunchTemplateSpecMonitoring `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	NetworkInterfaces []LaunchTemplateSpecNetworkInterfaces `json:"networkInterfaces,omitempty" tf:"network_interfaces,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Placement []LaunchTemplateSpecPlacement `json:"placement,omitempty" tf:"placement,omitempty"`
	// +optional
	RamDiskID string `json:"ramDiskID,omitempty" tf:"ram_disk_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
	// +optional
	TagSpecifications []LaunchTemplateSpecTagSpecifications `json:"tagSpecifications,omitempty" tf:"tag_specifications,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type LaunchTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LaunchTemplateList is a list of LaunchTemplates
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LaunchTemplate CRD objects
	Items []LaunchTemplate `json:"items,omitempty"`
}
