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

type AwsLaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLaunchConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsLaunchConfigurationStatus `json:"status,omitempty"`
}

type AwsLaunchConfigurationSpecEbsBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	NoDevice            bool   `json:"no_device"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsLaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsLaunchConfigurationSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsLaunchConfigurationSpec struct {
	KeyName                      string                       `json:"key_name"`
	SpotPrice                    string                       `json:"spot_price"`
	EbsBlockDevice               []AwsLaunchConfigurationSpec `json:"ebs_block_device"`
	ImageId                      string                       `json:"image_id"`
	InstanceType                 string                       `json:"instance_type"`
	UserData                     string                       `json:"user_data"`
	VpcClassicLinkSecurityGroups []string                     `json:"vpc_classic_link_security_groups"`
	EphemeralBlockDevice         []AwsLaunchConfigurationSpec `json:"ephemeral_block_device"`
	RootBlockDevice              []AwsLaunchConfigurationSpec `json:"root_block_device"`
	IamInstanceProfile           string                       `json:"iam_instance_profile"`
	EbsOptimized                 bool                         `json:"ebs_optimized"`
	PlacementTenancy             string                       `json:"placement_tenancy"`
	EnableMonitoring             bool                         `json:"enable_monitoring"`
	Name                         string                       `json:"name"`
	NamePrefix                   string                       `json:"name_prefix"`
	UserDataBase64               string                       `json:"user_data_base64"`
	SecurityGroups               []string                     `json:"security_groups"`
	VpcClassicLinkId             string                       `json:"vpc_classic_link_id"`
	AssociatePublicIpAddress     bool                         `json:"associate_public_ip_address"`
}

type AwsLaunchConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLaunchConfigurationList is a list of AwsLaunchConfigurations
type AwsLaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLaunchConfiguration CRD objects
	Items []AwsLaunchConfiguration `json:"items,omitempty"`
}
