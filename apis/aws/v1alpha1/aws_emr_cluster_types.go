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

type AwsEmrCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEmrClusterSpec   `json:"spec,omitempty"`
	Status            AwsEmrClusterStatus `json:"status,omitempty"`
}

type AwsEmrClusterSpecInstanceGroupEbsConfig struct {
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
}

type AwsEmrClusterSpecInstanceGroup struct {
	Name              string                           `json:"name"`
	Id                string                           `json:"id"`
	BidPrice          string                           `json:"bid_price"`
	EbsConfig         []AwsEmrClusterSpecInstanceGroup `json:"ebs_config"`
	InstanceCount     int                              `json:"instance_count"`
	AutoscalingPolicy string                           `json:"autoscaling_policy"`
	InstanceRole      string                           `json:"instance_role"`
	InstanceType      string                           `json:"instance_type"`
}

type AwsEmrClusterSpecMasterInstanceGroupEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

type AwsEmrClusterSpecMasterInstanceGroup struct {
	BidPrice     string                                 `json:"bid_price"`
	EbsConfig    []AwsEmrClusterSpecMasterInstanceGroup `json:"ebs_config"`
	Id           string                                 `json:"id"`
	InstanceType string                                 `json:"instance_type"`
	Name         string                                 `json:"name"`
}

type AwsEmrClusterSpecKerberosAttributes struct {
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
	Realm                            string `json:"realm"`
}

type AwsEmrClusterSpecCoreInstanceGroupEbsConfig struct {
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
}

type AwsEmrClusterSpecCoreInstanceGroup struct {
	Id                string                               `json:"id"`
	InstanceCount     int                                  `json:"instance_count"`
	InstanceType      string                               `json:"instance_type"`
	Name              string                               `json:"name"`
	AutoscalingPolicy string                               `json:"autoscaling_policy"`
	BidPrice          string                               `json:"bid_price"`
	EbsConfig         []AwsEmrClusterSpecCoreInstanceGroup `json:"ebs_config"`
}

type AwsEmrClusterSpecBootstrapAction struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type AwsEmrClusterSpecEc2Attributes struct {
	EmrManagedSlaveSecurityGroup   string `json:"emr_managed_slave_security_group"`
	InstanceProfile                string `json:"instance_profile"`
	ServiceAccessSecurityGroup     string `json:"service_access_security_group"`
	KeyName                        string `json:"key_name"`
	SubnetId                       string `json:"subnet_id"`
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups  string `json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup  string `json:"emr_managed_master_security_group"`
}

type AwsEmrClusterSpecStepHadoopJarStep struct {
	Jar        string            `json:"jar"`
	MainClass  string            `json:"main_class"`
	Properties map[string]string `json:"properties"`
	Args       []string          `json:"args"`
}

type AwsEmrClusterSpecStep struct {
	ActionOnFailure string                  `json:"action_on_failure"`
	HadoopJarStep   []AwsEmrClusterSpecStep `json:"hadoop_jar_step"`
	Name            string                  `json:"name"`
}

type AwsEmrClusterSpec struct {
	CoreInstanceType            string              `json:"core_instance_type"`
	CoreInstanceCount           int                 `json:"core_instance_count"`
	Applications                []string            `json:"applications"`
	EbsRootVolumeSize           int                 `json:"ebs_root_volume_size"`
	CustomAmiId                 string              `json:"custom_ami_id"`
	MasterPublicDns             string              `json:"master_public_dns"`
	Name                        string              `json:"name"`
	ReleaseLabel                string              `json:"release_label"`
	MasterInstanceType          string              `json:"master_instance_type"`
	ClusterState                string              `json:"cluster_state"`
	KeepJobFlowAliveWhenNoSteps bool                `json:"keep_job_flow_alive_when_no_steps"`
	SecurityConfiguration       string              `json:"security_configuration"`
	AdditionalInfo              string              `json:"additional_info"`
	InstanceGroup               []AwsEmrClusterSpec `json:"instance_group"`
	ConfigurationsJson          string              `json:"configurations_json"`
	ScaleDownBehavior           string              `json:"scale_down_behavior"`
	MasterInstanceGroup         []AwsEmrClusterSpec `json:"master_instance_group"`
	Tags                        map[string]string   `json:"tags"`
	VisibleToAllUsers           bool                `json:"visible_to_all_users"`
	KerberosAttributes          []AwsEmrClusterSpec `json:"kerberos_attributes"`
	CoreInstanceGroup           []AwsEmrClusterSpec `json:"core_instance_group"`
	BootstrapAction             []AwsEmrClusterSpec `json:"bootstrap_action"`
	AutoscalingRole             string              `json:"autoscaling_role"`
	LogUri                      string              `json:"log_uri"`
	TerminationProtection       bool                `json:"termination_protection"`
	Ec2Attributes               []AwsEmrClusterSpec `json:"ec2_attributes"`
	Step                        []AwsEmrClusterSpec `json:"step"`
	Configurations              string              `json:"configurations"`
	ServiceRole                 string              `json:"service_role"`
}

type AwsEmrClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEmrClusterList is a list of AwsEmrClusters
type AwsEmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEmrCluster CRD objects
	Items []AwsEmrCluster `json:"items,omitempty"`
}
