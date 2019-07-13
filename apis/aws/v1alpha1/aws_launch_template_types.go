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

type AwsLaunchTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLaunchTemplateSpec   `json:"spec,omitempty"`
	Status            AwsLaunchTemplateStatus `json:"status,omitempty"`
}

type AwsLaunchTemplateSpecInstanceMarketOptionsSpotOptions struct {
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
	BlockDurationMinutes         int    `json:"block_duration_minutes"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
}

type AwsLaunchTemplateSpecInstanceMarketOptions struct {
	MarketType  string                                       `json:"market_type"`
	SpotOptions []AwsLaunchTemplateSpecInstanceMarketOptions `json:"spot_options"`
}

type AwsLaunchTemplateSpecElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type AwsLaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	CapacityReservationId string `json:"capacity_reservation_id"`
}

type AwsLaunchTemplateSpecCapacityReservationSpecification struct {
	CapacityReservationPreference string                                                  `json:"capacity_reservation_preference"`
	CapacityReservationTarget     []AwsLaunchTemplateSpecCapacityReservationSpecification `json:"capacity_reservation_target"`
}

type AwsLaunchTemplateSpecElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type AwsLaunchTemplateSpecNetworkInterfaces struct {
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	AssociatePublicIpAddress bool     `json:"associate_public_ip_address"`
	DeleteOnTermination      bool     `json:"delete_on_termination"`
	Description              string   `json:"description"`
	Ipv6AddressCount         int      `json:"ipv6_address_count"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	DeviceIndex              int      `json:"device_index"`
	SecurityGroups           []string `json:"security_groups"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	Ipv4AddressCount         int      `json:"ipv4_address_count"`
	SubnetId                 string   `json:"subnet_id"`
}

type AwsLaunchTemplateSpecBlockDeviceMappingsEbs struct {
	DeleteOnTermination string `json:"delete_on_termination"`
	Encrypted           string `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsLaunchTemplateSpecBlockDeviceMappings struct {
	DeviceName  string                                     `json:"device_name"`
	NoDevice    string                                     `json:"no_device"`
	VirtualName string                                     `json:"virtual_name"`
	Ebs         []AwsLaunchTemplateSpecBlockDeviceMappings `json:"ebs"`
}

type AwsLaunchTemplateSpecIamInstanceProfile struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

type AwsLaunchTemplateSpecLicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type AwsLaunchTemplateSpecMonitoring struct {
	Enabled bool `json:"enabled"`
}

type AwsLaunchTemplateSpecTagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

type AwsLaunchTemplateSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsLaunchTemplateSpecPlacement struct {
	Tenancy          string `json:"tenancy"`
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
	SpreadDomain     string `json:"spread_domain"`
}

type AwsLaunchTemplateSpec struct {
	LatestVersion                     int                     `json:"latest_version"`
	UserData                          string                  `json:"user_data"`
	Name                              string                  `json:"name"`
	DefaultVersion                    int                     `json:"default_version"`
	InstanceInitiatedShutdownBehavior string                  `json:"instance_initiated_shutdown_behavior"`
	InstanceMarketOptions             []AwsLaunchTemplateSpec `json:"instance_market_options"`
	InstanceType                      string                  `json:"instance_type"`
	VpcSecurityGroupIds               []string                `json:"vpc_security_group_ids"`
	Arn                               string                  `json:"arn"`
	DisableApiTermination             bool                    `json:"disable_api_termination"`
	ElasticGpuSpecifications          []AwsLaunchTemplateSpec `json:"elastic_gpu_specifications"`
	KernelId                          string                  `json:"kernel_id"`
	CapacityReservationSpecification  []AwsLaunchTemplateSpec `json:"capacity_reservation_specification"`
	EbsOptimized                      string                  `json:"ebs_optimized"`
	ElasticInferenceAccelerator       []AwsLaunchTemplateSpec `json:"elastic_inference_accelerator"`
	ImageId                           string                  `json:"image_id"`
	NetworkInterfaces                 []AwsLaunchTemplateSpec `json:"network_interfaces"`
	Tags                              map[string]string       `json:"tags"`
	Description                       string                  `json:"description"`
	SecurityGroupNames                []string                `json:"security_group_names"`
	BlockDeviceMappings               []AwsLaunchTemplateSpec `json:"block_device_mappings"`
	IamInstanceProfile                []AwsLaunchTemplateSpec `json:"iam_instance_profile"`
	LicenseSpecification              []AwsLaunchTemplateSpec `json:"license_specification"`
	Monitoring                        []AwsLaunchTemplateSpec `json:"monitoring"`
	TagSpecifications                 []AwsLaunchTemplateSpec `json:"tag_specifications"`
	NamePrefix                        string                  `json:"name_prefix"`
	CreditSpecification               []AwsLaunchTemplateSpec `json:"credit_specification"`
	KeyName                           string                  `json:"key_name"`
	Placement                         []AwsLaunchTemplateSpec `json:"placement"`
	RamDiskId                         string                  `json:"ram_disk_id"`
}



type AwsLaunchTemplateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLaunchTemplateList is a list of AwsLaunchTemplates
type AwsLaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLaunchTemplate CRD objects
	Items []AwsLaunchTemplate `json:"items,omitempty"`
}