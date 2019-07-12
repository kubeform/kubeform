package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksInstanceSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksInstanceStatus `json:"status,omitempty"`
}

type AwsOpsworksInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsOpsworksInstanceSpecRootBlockDevice struct {
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsOpsworksInstanceSpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsOpsworksInstanceSpec struct {
	InstallUpdatesOnBoot     bool                      `json:"install_updates_on_boot"`
	SshHostDsaKeyFingerprint string                    `json:"ssh_host_dsa_key_fingerprint"`
	ElasticIp                string                    `json:"elastic_ip"`
	EcsClusterArn            string                    `json:"ecs_cluster_arn"`
	InfrastructureClass      string                    `json:"infrastructure_class"`
	VirtualizationType       string                    `json:"virtualization_type"`
	AutoScalingType          string                    `json:"auto_scaling_type"`
	ReportedOsVersion        string                    `json:"reported_os_version"`
	SecurityGroupIds         []string                  `json:"security_group_ids"`
	SshKeyName               string                    `json:"ssh_key_name"`
	StackId                  string                    `json:"stack_id"`
	Ec2InstanceId            string                    `json:"ec2_instance_id"`
	AmiId                    string                    `json:"ami_id"`
	Architecture             string                    `json:"architecture"`
	LastServiceErrorId       string                    `json:"last_service_error_id"`
	PrivateDns               string                    `json:"private_dns"`
	PrivateIp                string                    `json:"private_ip"`
	PublicIp                 string                    `json:"public_ip"`
	ReportedOsFamily         string                    `json:"reported_os_family"`
	AgentVersion             string                    `json:"agent_version"`
	EphemeralBlockDevice     []AwsOpsworksInstanceSpec `json:"ephemeral_block_device"`
	State                    string                    `json:"state"`
	Os                       string                    `json:"os"`
	Platform                 string                    `json:"platform"`
	PublicDns                string                    `json:"public_dns"`
	RootDeviceVolumeId       string                    `json:"root_device_volume_id"`
	SubnetId                 string                    `json:"subnet_id"`
	RootBlockDevice          []AwsOpsworksInstanceSpec `json:"root_block_device"`
	InstanceType             string                    `json:"instance_type"`
	EbsOptimized             bool                      `json:"ebs_optimized"`
	Hostname                 string                    `json:"hostname"`
	SshHostRsaKeyFingerprint string                    `json:"ssh_host_rsa_key_fingerprint"`
	Tenancy                  string                    `json:"tenancy"`
	EbsBlockDevice           []AwsOpsworksInstanceSpec `json:"ebs_block_device"`
	DeleteEbs                bool                      `json:"delete_ebs"`
	CreatedAt                string                    `json:"created_at"`
	DeleteEip                bool                      `json:"delete_eip"`
	LayerIds                 []string                  `json:"layer_ids"`
	RegisteredBy             string                    `json:"registered_by"`
	ReportedAgentVersion     string                    `json:"reported_agent_version"`
	ReportedOsName           string                    `json:"reported_os_name"`
	RootDeviceType           string                    `json:"root_device_type"`
	AvailabilityZone         string                    `json:"availability_zone"`
	Status                   string                    `json:"status"`
	InstanceProfileArn       string                    `json:"instance_profile_arn"`
}



type AwsOpsworksInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstanceList is a list of AwsOpsworksInstances
type AwsOpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksInstance CRD objects
	Items []AwsOpsworksInstance `json:"items,omitempty"`
}