package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEmrClusterSpec   `json:"spec,omitempty"`
	Status            AwsEmrClusterStatus `json:"status,omitempty"`
}

type AwsEmrClusterSpecInstanceGroupEbsConfig struct {
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
}

type AwsEmrClusterSpecInstanceGroup struct {
	EbsConfig         []AwsEmrClusterSpecInstanceGroup `json:"ebs_config"`
	InstanceCount     int                              `json:"instance_count"`
	AutoscalingPolicy string                           `json:"autoscaling_policy"`
	InstanceRole      string                           `json:"instance_role"`
	InstanceType      string                           `json:"instance_type"`
	Name              string                           `json:"name"`
	Id                string                           `json:"id"`
	BidPrice          string                           `json:"bid_price"`
}

type AwsEmrClusterSpecStepHadoopJarStep struct {
	Args       []string          `json:"args"`
	Jar        string            `json:"jar"`
	MainClass  string            `json:"main_class"`
	Properties map[string]string `json:"properties"`
}

type AwsEmrClusterSpecStep struct {
	ActionOnFailure string                  `json:"action_on_failure"`
	HadoopJarStep   []AwsEmrClusterSpecStep `json:"hadoop_jar_step"`
	Name            string                  `json:"name"`
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

type AwsEmrClusterSpecKerberosAttributes struct {
	Realm                            string `json:"realm"`
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
}

type AwsEmrClusterSpecBootstrapAction struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type AwsEmrClusterSpec struct {
	EbsRootVolumeSize           int                 `json:"ebs_root_volume_size"`
	AdditionalInfo              string              `json:"additional_info"`
	Applications                []string            `json:"applications"`
	MasterInstanceType          string              `json:"master_instance_type"`
	Tags                        map[string]string   `json:"tags"`
	ServiceRole                 string              `json:"service_role"`
	CustomAmiId                 string              `json:"custom_ami_id"`
	InstanceGroup               []AwsEmrClusterSpec `json:"instance_group"`
	Step                        []AwsEmrClusterSpec `json:"step"`
	CoreInstanceType            string              `json:"core_instance_type"`
	KeepJobFlowAliveWhenNoSteps bool                `json:"keep_job_flow_alive_when_no_steps"`
	Configurations              string              `json:"configurations"`
	SecurityConfiguration       string              `json:"security_configuration"`
	AutoscalingRole             string              `json:"autoscaling_role"`
	Name                        string              `json:"name"`
	MasterPublicDns             string              `json:"master_public_dns"`
	TerminationProtection       bool                `json:"termination_protection"`
	VisibleToAllUsers           bool                `json:"visible_to_all_users"`
	ClusterState                string              `json:"cluster_state"`
	LogUri                      string              `json:"log_uri"`
	Ec2Attributes               []AwsEmrClusterSpec `json:"ec2_attributes"`
	KerberosAttributes          []AwsEmrClusterSpec `json:"kerberos_attributes"`
	ScaleDownBehavior           string              `json:"scale_down_behavior"`
	ReleaseLabel                string              `json:"release_label"`
	BootstrapAction             []AwsEmrClusterSpec `json:"bootstrap_action"`
	ConfigurationsJson          string              `json:"configurations_json"`
	CoreInstanceCount           int                 `json:"core_instance_count"`
}



type AwsEmrClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrClusters
type AwsEmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEmrCluster CRD objects
	Items []AwsEmrCluster `json:"items,omitempty"`
}