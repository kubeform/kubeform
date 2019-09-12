package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type EmrCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrClusterSpec   `json:"spec,omitempty"`
	Status            EmrClusterStatus `json:"status,omitempty"`
}

type EmrClusterSpecBootstrapAction struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	Name string   `json:"name" tf:"name"`
	Path string   `json:"path" tf:"path"`
}

type EmrClusterSpecEc2Attributes struct {
	// +optional
	AdditionalMasterSecurityGroups string `json:"additionalMasterSecurityGroups,omitempty" tf:"additional_master_security_groups,omitempty"`
	// +optional
	AdditionalSlaveSecurityGroups string `json:"additionalSlaveSecurityGroups,omitempty" tf:"additional_slave_security_groups,omitempty"`
	// +optional
	EmrManagedMasterSecurityGroup string `json:"emrManagedMasterSecurityGroup,omitempty" tf:"emr_managed_master_security_group,omitempty"`
	// +optional
	EmrManagedSlaveSecurityGroup string `json:"emrManagedSlaveSecurityGroup,omitempty" tf:"emr_managed_slave_security_group,omitempty"`
	InstanceProfile              string `json:"instanceProfile" tf:"instance_profile"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	ServiceAccessSecurityGroup string `json:"serviceAccessSecurityGroup,omitempty" tf:"service_access_security_group,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type EmrClusterSpecInstanceGroupEbsConfig struct {
	// +optional
	Iops int    `json:"iops,omitempty" tf:"iops,omitempty"`
	Size int    `json:"size" tf:"size"`
	Type string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type EmrClusterSpecInstanceGroup struct {
	// +optional
	AutoscalingPolicy string `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`
	// +optional
	BidPrice string `json:"bidPrice,omitempty" tf:"bid_price,omitempty"`
	// +optional
	EbsConfig []EmrClusterSpecInstanceGroupEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	InstanceCount int    `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`
	InstanceRole  string `json:"instanceRole" tf:"instance_role"`
	InstanceType  string `json:"instanceType" tf:"instance_type"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type EmrClusterSpecKerberosAttributes struct {
	// +optional
	AdDomainJoinPassword string `json:"-" sensitive:"true" tf:"ad_domain_join_password,omitempty"`
	// +optional
	AdDomainJoinUser string `json:"adDomainJoinUser,omitempty" tf:"ad_domain_join_user,omitempty"`
	// +optional
	CrossRealmTrustPrincipalPassword string `json:"-" sensitive:"true" tf:"cross_realm_trust_principal_password,omitempty"`
	KdcAdminPassword                 string `json:"-" sensitive:"true" tf:"kdc_admin_password"`
	Realm                            string `json:"realm" tf:"realm"`
}

type EmrClusterSpecStepHadoopJarStep struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	Jar  string   `json:"jar" tf:"jar"`
	// +optional
	MainClass string `json:"mainClass,omitempty" tf:"main_class,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type EmrClusterSpecStep struct {
	ActionOnFailure string `json:"actionOnFailure" tf:"action_on_failure"`
	// +kubebuilder:validation:MaxItems=1
	HadoopJarStep []EmrClusterSpecStepHadoopJarStep `json:"hadoopJarStep" tf:"hadoop_jar_step"`
	Name          string                            `json:"name" tf:"name"`
}

type EmrClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdditionalInfo string `json:"additionalInfo,omitempty" tf:"additional_info,omitempty"`
	// +optional
	Applications []string `json:"applications,omitempty" tf:"applications,omitempty"`
	// +optional
	AutoscalingRole string `json:"autoscalingRole,omitempty" tf:"autoscaling_role,omitempty"`
	// +optional
	BootstrapAction []EmrClusterSpecBootstrapAction `json:"bootstrapAction,omitempty" tf:"bootstrap_action,omitempty"`
	// +optional
	ClusterState string `json:"clusterState,omitempty" tf:"cluster_state,omitempty"`
	// +optional
	Configurations string `json:"configurations,omitempty" tf:"configurations,omitempty"`
	// +optional
	ConfigurationsJSON string `json:"configurationsJSON,omitempty" tf:"configurations_json,omitempty"`
	// +optional
	CoreInstanceCount int `json:"coreInstanceCount,omitempty" tf:"core_instance_count,omitempty"`
	// +optional
	CoreInstanceType string `json:"coreInstanceType,omitempty" tf:"core_instance_type,omitempty"`
	// +optional
	CustomAmiID string `json:"customAmiID,omitempty" tf:"custom_ami_id,omitempty"`
	// +optional
	EbsRootVolumeSize int `json:"ebsRootVolumeSize,omitempty" tf:"ebs_root_volume_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ec2Attributes []EmrClusterSpecEc2Attributes `json:"ec2Attributes,omitempty" tf:"ec2_attributes,omitempty"`
	// +optional
	InstanceGroup []EmrClusterSpecInstanceGroup `json:"instanceGroup,omitempty" tf:"instance_group,omitempty"`
	// +optional
	KeepJobFlowAliveWhenNoSteps bool `json:"keepJobFlowAliveWhenNoSteps,omitempty" tf:"keep_job_flow_alive_when_no_steps,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KerberosAttributes []EmrClusterSpecKerberosAttributes `json:"kerberosAttributes,omitempty" tf:"kerberos_attributes,omitempty"`
	// +optional
	LogURI string `json:"logURI,omitempty" tf:"log_uri,omitempty"`
	// +optional
	MasterInstanceType string `json:"masterInstanceType,omitempty" tf:"master_instance_type,omitempty"`
	// +optional
	MasterPublicDNS string `json:"masterPublicDNS,omitempty" tf:"master_public_dns,omitempty"`
	Name            string `json:"name" tf:"name"`
	ReleaseLabel    string `json:"releaseLabel" tf:"release_label"`
	// +optional
	ScaleDownBehavior string `json:"scaleDownBehavior,omitempty" tf:"scale_down_behavior,omitempty"`
	// +optional
	SecurityConfiguration string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`
	ServiceRole           string `json:"serviceRole" tf:"service_role"`
	// +optional
	Step []EmrClusterSpecStep `json:"step,omitempty" tf:"step,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TerminationProtection bool `json:"terminationProtection,omitempty" tf:"termination_protection,omitempty"`
	// +optional
	VisibleToAllUsers bool `json:"visibleToAllUsers,omitempty" tf:"visible_to_all_users,omitempty"`
}

type EmrClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EmrClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EmrClusterList is a list of EmrClusters
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmrCluster CRD objects
	Items []EmrCluster `json:"items,omitempty"`
}
