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

type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec,omitempty"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

type LaunchTemplateSpecBlockDeviceMappingsEbs struct {
	// +optional
	DeleteOnTermination string `json:"delete_on_termination,omitempty"`
	// +optional
	Encrypted string `json:"encrypted,omitempty"`
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	// +optional
	SnapshotId string `json:"snapshot_id,omitempty"`
}

type LaunchTemplateSpecBlockDeviceMappings struct {
	// +optional
	DeviceName string `json:"device_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ebs *[]LaunchTemplateSpecBlockDeviceMappings `json:"ebs,omitempty"`
	// +optional
	NoDevice string `json:"no_device,omitempty"`
	// +optional
	VirtualName string `json:"virtual_name,omitempty"`
}

type LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	// +optional
	CapacityReservationId string `json:"capacity_reservation_id,omitempty"`
}

type LaunchTemplateSpecCapacityReservationSpecification struct {
	// +optional
	CapacityReservationPreference string `json:"capacity_reservation_preference,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CapacityReservationTarget *[]LaunchTemplateSpecCapacityReservationSpecification `json:"capacity_reservation_target,omitempty"`
}

type LaunchTemplateSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpu_credits,omitempty"`
}

type LaunchTemplateSpecElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type LaunchTemplateSpecElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type LaunchTemplateSpecIamInstanceProfile struct {
	// +optional
	Arn string `json:"arn,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
}

type LaunchTemplateSpecInstanceMarketOptionsSpotOptions struct {
	// +optional
	BlockDurationMinutes int `json:"block_duration_minutes,omitempty"`
	// +optional
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior,omitempty"`
	// +optional
	MaxPrice string `json:"max_price,omitempty"`
	// +optional
	SpotInstanceType string `json:"spot_instance_type,omitempty"`
}

type LaunchTemplateSpecInstanceMarketOptions struct {
	// +optional
	MarketType string `json:"market_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SpotOptions *[]LaunchTemplateSpecInstanceMarketOptions `json:"spot_options,omitempty"`
}

type LaunchTemplateSpecLicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type LaunchTemplateSpecMonitoring struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type LaunchTemplateSpecNetworkInterfaces struct {
	// +optional
	AssociatePublicIpAddress bool `json:"associate_public_ip_address,omitempty"`
	// +optional
	DeleteOnTermination bool `json:"delete_on_termination,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DeviceIndex int `json:"device_index,omitempty"`
	// +optional
	Ipv4AddressCount int `json:"ipv4_address_count,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ipv4Addresses []string `json:"ipv4_addresses,omitempty"`
	// +optional
	Ipv6AddressCount int `json:"ipv6_address_count,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ipv6Addresses []string `json:"ipv6_addresses,omitempty"`
	// +optional
	NetworkInterfaceId string `json:"network_interface_id,omitempty"`
	// +optional
	PrivateIpAddress string `json:"private_ip_address,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
}

type LaunchTemplateSpecPlacement struct {
	// +optional
	Affinity string `json:"affinity,omitempty"`
	// +optional
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// +optional
	GroupName string `json:"group_name,omitempty"`
	// +optional
	HostId string `json:"host_id,omitempty"`
	// +optional
	SpreadDomain string `json:"spread_domain,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty"`
}

type LaunchTemplateSpecTagSpecifications struct {
	// +optional
	ResourceType string `json:"resource_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type LaunchTemplateSpec struct {
	// +optional
	BlockDeviceMappings *[]LaunchTemplateSpec `json:"block_device_mappings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CapacityReservationSpecification *[]LaunchTemplateSpec `json:"capacity_reservation_specification,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification *[]LaunchTemplateSpec `json:"credit_specification,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DisableApiTermination bool `json:"disable_api_termination,omitempty"`
	// +optional
	EbsOptimized string `json:"ebs_optimized,omitempty"`
	// +optional
	ElasticGpuSpecifications *[]LaunchTemplateSpec `json:"elastic_gpu_specifications,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticInferenceAccelerator *[]LaunchTemplateSpec `json:"elastic_inference_accelerator,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IamInstanceProfile *[]LaunchTemplateSpec `json:"iam_instance_profile,omitempty"`
	// +optional
	ImageId string `json:"image_id,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instance_initiated_shutdown_behavior,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InstanceMarketOptions *[]LaunchTemplateSpec `json:"instance_market_options,omitempty"`
	// +optional
	InstanceType string `json:"instance_type,omitempty"`
	// +optional
	KernelId string `json:"kernel_id,omitempty"`
	// +optional
	KeyName string `json:"key_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LicenseSpecification *[]LaunchTemplateSpec `json:"license_specification,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Monitoring *[]LaunchTemplateSpec `json:"monitoring,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	NetworkInterfaces *[]LaunchTemplateSpec `json:"network_interfaces,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Placement *[]LaunchTemplateSpec `json:"placement,omitempty"`
	// +optional
	RamDiskId string `json:"ram_disk_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"security_group_names,omitempty"`
	// +optional
	TagSpecifications *[]LaunchTemplateSpec `json:"tag_specifications,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	UserData string `json:"user_data,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIds []string `json:"vpc_security_group_ids,omitempty"`
}

type LaunchTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
