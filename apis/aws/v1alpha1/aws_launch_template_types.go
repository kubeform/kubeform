package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLaunchTemplateSpec   `json:"spec,omitempty"`
	Status            AwsLaunchTemplateStatus `json:"status,omitempty"`
}

type AwsLaunchTemplateSpecIamInstanceProfile struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

type AwsLaunchTemplateSpecInstanceMarketOptionsSpotOptions struct {
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
	BlockDurationMinutes         int    `json:"block_duration_minutes"`
}

type AwsLaunchTemplateSpecInstanceMarketOptions struct {
	MarketType  string                                       `json:"market_type"`
	SpotOptions []AwsLaunchTemplateSpecInstanceMarketOptions `json:"spot_options"`
}

type AwsLaunchTemplateSpecElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type AwsLaunchTemplateSpecLicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type AwsLaunchTemplateSpecMonitoring struct {
	Enabled bool `json:"enabled"`
}

type AwsLaunchTemplateSpecBlockDeviceMappingsEbs struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination string `json:"delete_on_termination"`
	Encrypted           string `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
}

type AwsLaunchTemplateSpecBlockDeviceMappings struct {
	DeviceName  string                                     `json:"device_name"`
	NoDevice    string                                     `json:"no_device"`
	VirtualName string                                     `json:"virtual_name"`
	Ebs         []AwsLaunchTemplateSpecBlockDeviceMappings `json:"ebs"`
}

type AwsLaunchTemplateSpecPlacement struct {
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
}

type AwsLaunchTemplateSpecTagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

type AwsLaunchTemplateSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsLaunchTemplateSpecNetworkInterfaces struct {
	SubnetId                 string   `json:"subnet_id"`
	DeleteOnTermination      bool     `json:"delete_on_termination"`
	DeviceIndex              int      `json:"device_index"`
	Ipv6AddressCount         int      `json:"ipv6_address_count"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	Ipv4AddressCount         int      `json:"ipv4_address_count"`
	AssociatePublicIpAddress bool     `json:"associate_public_ip_address"`
	Description              string   `json:"description"`
	SecurityGroups           []string `json:"security_groups"`
	PrivateIpAddress         string   `json:"private_ip_address"`
}

type AwsLaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	CapacityReservationId string `json:"capacity_reservation_id"`
}

type AwsLaunchTemplateSpecCapacityReservationSpecification struct {
	CapacityReservationPreference string                                                  `json:"capacity_reservation_preference"`
	CapacityReservationTarget     []AwsLaunchTemplateSpecCapacityReservationSpecification `json:"capacity_reservation_target"`
}

type AwsLaunchTemplateSpecElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type AwsLaunchTemplateSpec struct {
	SecurityGroupNames                []string                `json:"security_group_names"`
	Name                              string                  `json:"name"`
	Description                       string                  `json:"description"`
	IamInstanceProfile                []AwsLaunchTemplateSpec `json:"iam_instance_profile"`
	InstanceMarketOptions             []AwsLaunchTemplateSpec `json:"instance_market_options"`
	KernelId                          string                  `json:"kernel_id"`
	RamDiskId                         string                  `json:"ram_disk_id"`
	DefaultVersion                    int                     `json:"default_version"`
	LatestVersion                     int                     `json:"latest_version"`
	ElasticInferenceAccelerator       []AwsLaunchTemplateSpec `json:"elastic_inference_accelerator"`
	ImageId                           string                  `json:"image_id"`
	LicenseSpecification              []AwsLaunchTemplateSpec `json:"license_specification"`
	Monitoring                        []AwsLaunchTemplateSpec `json:"monitoring"`
	VpcSecurityGroupIds               []string                `json:"vpc_security_group_ids"`
	UserData                          string                  `json:"user_data"`
	Arn                               string                  `json:"arn"`
	BlockDeviceMappings               []AwsLaunchTemplateSpec `json:"block_device_mappings"`
	InstanceType                      string                  `json:"instance_type"`
	Placement                         []AwsLaunchTemplateSpec `json:"placement"`
	TagSpecifications                 []AwsLaunchTemplateSpec `json:"tag_specifications"`
	CreditSpecification               []AwsLaunchTemplateSpec `json:"credit_specification"`
	NetworkInterfaces                 []AwsLaunchTemplateSpec `json:"network_interfaces"`
	Tags                              map[string]string       `json:"tags"`
	NamePrefix                        string                  `json:"name_prefix"`
	EbsOptimized                      string                  `json:"ebs_optimized"`
	InstanceInitiatedShutdownBehavior string                  `json:"instance_initiated_shutdown_behavior"`
	KeyName                           string                  `json:"key_name"`
	CapacityReservationSpecification  []AwsLaunchTemplateSpec `json:"capacity_reservation_specification"`
	DisableApiTermination             bool                    `json:"disable_api_termination"`
	ElasticGpuSpecifications          []AwsLaunchTemplateSpec `json:"elastic_gpu_specifications"`
}



type AwsLaunchTemplateStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplates
type AwsLaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLaunchTemplate CRD objects
	Items []AwsLaunchTemplate `json:"items,omitempty"`
}