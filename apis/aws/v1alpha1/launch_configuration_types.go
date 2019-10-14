package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type LaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchConfigurationSpec   `json:"spec,omitempty"`
	Status            LaunchConfigurationStatus `json:"status,omitempty"`
}

type LaunchConfigurationSpecEbsBlockDevice struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	NoDevice bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type LaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"deviceName" tf:"device_name"`
	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type LaunchConfigurationSpecRootBlockDevice struct {
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type LaunchConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	EbsBlockDevice []LaunchConfigurationSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	EnableMonitoring bool `json:"enableMonitoring,omitempty" tf:"enable_monitoring,omitempty"`
	// +optional
	EphemeralBlockDevice []LaunchConfigurationSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	ImageID            string `json:"imageID" tf:"image_id"`
	InstanceType       string `json:"instanceType" tf:"instance_type"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	PlacementTenancy string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RootBlockDevice []LaunchConfigurationSpecRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SpotPrice string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	UserDataBase64 string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`
	// +optional
	VpcClassicLinkID string `json:"vpcClassicLinkID,omitempty" tf:"vpc_classic_link_id,omitempty"`
	// +optional
	VpcClassicLinkSecurityGroups []string `json:"vpcClassicLinkSecurityGroups,omitempty" tf:"vpc_classic_link_security_groups,omitempty"`
}

type LaunchConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LaunchConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LaunchConfigurationList is a list of LaunchConfigurations
type LaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LaunchConfiguration CRD objects
	Items []LaunchConfiguration `json:"items,omitempty"`
}
