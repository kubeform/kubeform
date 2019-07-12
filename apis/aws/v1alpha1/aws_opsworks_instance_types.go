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
	ElasticIp                string                    `json:"elastic_ip"`
	PublicDns                string                    `json:"public_dns"`
	ReportedOsFamily         string                    `json:"reported_os_family"`
	Architecture             string                    `json:"architecture"`
	DeleteEip                bool                      `json:"delete_eip"`
	InstallUpdatesOnBoot     bool                      `json:"install_updates_on_boot"`
	InstanceProfileArn       string                    `json:"instance_profile_arn"`
	LayerIds                 []string                  `json:"layer_ids"`
	PrivateIp                string                    `json:"private_ip"`
	RegisteredBy             string                    `json:"registered_by"`
	StackId                  string                    `json:"stack_id"`
	Ec2InstanceId            string                    `json:"ec2_instance_id"`
	InfrastructureClass      string                    `json:"infrastructure_class"`
	CreatedAt                string                    `json:"created_at"`
	ReportedAgentVersion     string                    `json:"reported_agent_version"`
	PrivateDns               string                    `json:"private_dns"`
	DeleteEbs                bool                      `json:"delete_ebs"`
	EbsOptimized             bool                      `json:"ebs_optimized"`
	ReportedOsVersion        string                    `json:"reported_os_version"`
	SshHostRsaKeyFingerprint string                    `json:"ssh_host_rsa_key_fingerprint"`
	VirtualizationType       string                    `json:"virtualization_type"`
	AmiId                    string                    `json:"ami_id"`
	AvailabilityZone         string                    `json:"availability_zone"`
	Os                       string                    `json:"os"`
	PublicIp                 string                    `json:"public_ip"`
	Status                   string                    `json:"status"`
	AutoScalingType          string                    `json:"auto_scaling_type"`
	LastServiceErrorId       string                    `json:"last_service_error_id"`
	InstanceType             string                    `json:"instance_type"`
	Platform                 string                    `json:"platform"`
	ReportedOsName           string                    `json:"reported_os_name"`
	RootDeviceVolumeId       string                    `json:"root_device_volume_id"`
	SecurityGroupIds         []string                  `json:"security_group_ids"`
	SshHostDsaKeyFingerprint string                    `json:"ssh_host_dsa_key_fingerprint"`
	EcsClusterArn            string                    `json:"ecs_cluster_arn"`
	Hostname                 string                    `json:"hostname"`
	RootBlockDevice          []AwsOpsworksInstanceSpec `json:"root_block_device"`
	SubnetId                 string                    `json:"subnet_id"`
	Tenancy                  string                    `json:"tenancy"`
	SshKeyName               string                    `json:"ssh_key_name"`
	State                    string                    `json:"state"`
	EbsBlockDevice           []AwsOpsworksInstanceSpec `json:"ebs_block_device"`
	EphemeralBlockDevice     []AwsOpsworksInstanceSpec `json:"ephemeral_block_device"`
	AgentVersion             string                    `json:"agent_version"`
	RootDeviceType           string                    `json:"root_device_type"`
}

type AwsOpsworksInstanceStatus struct {
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
