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

type AwsOpsworksInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksInstanceSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksInstanceStatus `json:"status,omitempty"`
}

type AwsOpsworksInstanceSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsOpsworksInstanceSpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

type AwsOpsworksInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsOpsworksInstanceSpec struct {
	LayerIds                 []string                  `json:"layer_ids"`
	Platform                 string                    `json:"platform"`
	AutoScalingType          string                    `json:"auto_scaling_type"`
	EbsOptimized             bool                      `json:"ebs_optimized"`
	Os                       string                    `json:"os"`
	VirtualizationType       string                    `json:"virtualization_type"`
	RootBlockDevice          []AwsOpsworksInstanceSpec `json:"root_block_device"`
	AgentVersion             string                    `json:"agent_version"`
	AmiId                    string                    `json:"ami_id"`
	EcsClusterArn            string                    `json:"ecs_cluster_arn"`
	InfrastructureClass      string                    `json:"infrastructure_class"`
	RootDeviceType           string                    `json:"root_device_type"`
	SshKeyName               string                    `json:"ssh_key_name"`
	SubnetId                 string                    `json:"subnet_id"`
	Tenancy                  string                    `json:"tenancy"`
	Architecture             string                    `json:"architecture"`
	Ec2InstanceId            string                    `json:"ec2_instance_id"`
	EbsBlockDevice           []AwsOpsworksInstanceSpec `json:"ebs_block_device"`
	InstanceType             string                    `json:"instance_type"`
	ReportedAgentVersion     string                    `json:"reported_agent_version"`
	ReportedOsFamily         string                    `json:"reported_os_family"`
	ReportedOsVersion        string                    `json:"reported_os_version"`
	RootDeviceVolumeId       string                    `json:"root_device_volume_id"`
	Hostname                 string                    `json:"hostname"`
	InstallUpdatesOnBoot     bool                      `json:"install_updates_on_boot"`
	StackId                  string                    `json:"stack_id"`
	PublicIp                 string                    `json:"public_ip"`
	RegisteredBy             string                    `json:"registered_by"`
	DeleteEip                bool                      `json:"delete_eip"`
	ElasticIp                string                    `json:"elastic_ip"`
	Status                   string                    `json:"status"`
	AvailabilityZone         string                    `json:"availability_zone"`
	CreatedAt                string                    `json:"created_at"`
	PublicDns                string                    `json:"public_dns"`
	ReportedOsName           string                    `json:"reported_os_name"`
	SshHostDsaKeyFingerprint string                    `json:"ssh_host_dsa_key_fingerprint"`
	SshHostRsaKeyFingerprint string                    `json:"ssh_host_rsa_key_fingerprint"`
	State                    string                    `json:"state"`
	EphemeralBlockDevice     []AwsOpsworksInstanceSpec `json:"ephemeral_block_device"`
	InstanceProfileArn       string                    `json:"instance_profile_arn"`
	LastServiceErrorId       string                    `json:"last_service_error_id"`
	PrivateIp                string                    `json:"private_ip"`
	SecurityGroupIds         []string                  `json:"security_group_ids"`
	DeleteEbs                bool                      `json:"delete_ebs"`
	PrivateDns               string                    `json:"private_dns"`
}



type AwsOpsworksInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksInstanceList is a list of AwsOpsworksInstances
type AwsOpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksInstance CRD objects
	Items []AwsOpsworksInstance `json:"items,omitempty"`
}