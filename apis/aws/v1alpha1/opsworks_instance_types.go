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

type OpsworksInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksInstanceSpec   `json:"spec,omitempty"`
	Status            OpsworksInstanceStatus `json:"status,omitempty"`
}

type OpsworksInstanceSpecEbsBlockDevice struct {
	// +optional
	DeleteOnTermination bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type OpsworksInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"deviceName" tf:"device_name"`
	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type OpsworksInstanceSpecRootBlockDevice struct {
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type OpsworksInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AgentVersion string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`
	// +optional
	AmiID string `json:"amiID,omitempty" tf:"ami_id,omitempty"`
	// +optional
	Architecture string `json:"architecture,omitempty" tf:"architecture,omitempty"`
	// +optional
	AutoScalingType string `json:"autoScalingType,omitempty" tf:"auto_scaling_type,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	DeleteEbs bool `json:"deleteEbs,omitempty" tf:"delete_ebs,omitempty"`
	// +optional
	DeleteEip bool `json:"deleteEip,omitempty" tf:"delete_eip,omitempty"`
	// +optional
	EbsBlockDevice []OpsworksInstanceSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	Ec2InstanceID string `json:"ec2InstanceID,omitempty" tf:"ec2_instance_id,omitempty"`
	// +optional
	EcsClusterArn string `json:"ecsClusterArn,omitempty" tf:"ecs_cluster_arn,omitempty"`
	// +optional
	ElasticIP string `json:"elasticIP,omitempty" tf:"elastic_ip,omitempty"`
	// +optional
	EphemeralBlockDevice []OpsworksInstanceSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	// +optional
	InfrastructureClass string `json:"infrastructureClass,omitempty" tf:"infrastructure_class,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceProfileArn string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	// +optional
	LastServiceErrorID string   `json:"lastServiceErrorID,omitempty" tf:"last_service_error_id,omitempty"`
	LayerIDS           []string `json:"layerIDS" tf:"layer_ids"`
	// +optional
	Os string `json:"os,omitempty" tf:"os,omitempty"`
	// +optional
	Platform string `json:"platform,omitempty" tf:"platform,omitempty"`
	// +optional
	PrivateDNS string `json:"privateDNS,omitempty" tf:"private_dns,omitempty"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// +optional
	PublicDNS string `json:"publicDNS,omitempty" tf:"public_dns,omitempty"`
	// +optional
	PublicIP string `json:"publicIP,omitempty" tf:"public_ip,omitempty"`
	// +optional
	RegisteredBy string `json:"registeredBy,omitempty" tf:"registered_by,omitempty"`
	// +optional
	ReportedAgentVersion string `json:"reportedAgentVersion,omitempty" tf:"reported_agent_version,omitempty"`
	// +optional
	ReportedOsFamily string `json:"reportedOsFamily,omitempty" tf:"reported_os_family,omitempty"`
	// +optional
	ReportedOsName string `json:"reportedOsName,omitempty" tf:"reported_os_name,omitempty"`
	// +optional
	ReportedOsVersion string `json:"reportedOsVersion,omitempty" tf:"reported_os_version,omitempty"`
	// +optional
	RootBlockDevice []OpsworksInstanceSpecRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`
	// +optional
	RootDeviceType string `json:"rootDeviceType,omitempty" tf:"root_device_type,omitempty"`
	// +optional
	RootDeviceVolumeID string `json:"rootDeviceVolumeID,omitempty" tf:"root_device_volume_id,omitempty"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	SshHostDsaKeyFingerprint string `json:"sshHostDsaKeyFingerprint,omitempty" tf:"ssh_host_dsa_key_fingerprint,omitempty"`
	// +optional
	SshHostRsaKeyFingerprint string `json:"sshHostRsaKeyFingerprint,omitempty" tf:"ssh_host_rsa_key_fingerprint,omitempty"`
	// +optional
	SshKeyName string `json:"sshKeyName,omitempty" tf:"ssh_key_name,omitempty"`
	StackID    string `json:"stackID" tf:"stack_id"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
	// +optional
	VirtualizationType string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type OpsworksInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OpsworksInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksInstanceList is a list of OpsworksInstances
type OpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksInstance CRD objects
	Items []OpsworksInstance `json:"items,omitempty"`
}
