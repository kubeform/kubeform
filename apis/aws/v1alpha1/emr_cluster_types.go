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

type EmrCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrClusterSpec   `json:"spec,omitempty"`
	Status            EmrClusterStatus `json:"status,omitempty"`
}

type EmrClusterSpecBootstrapAction struct {
	// +optional
	Args []string `json:"args,omitempty"`
	Name string   `json:"name"`
	Path string   `json:"path"`
}

type EmrClusterSpecEc2Attributes struct {
	// +optional
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups,omitempty"`
	// +optional
	AdditionalSlaveSecurityGroups string `json:"additional_slave_security_groups,omitempty"`
	// +optional
	EmrManagedMasterSecurityGroup string `json:"emr_managed_master_security_group,omitempty"`
	// +optional
	EmrManagedSlaveSecurityGroup string `json:"emr_managed_slave_security_group,omitempty"`
	InstanceProfile              string `json:"instance_profile"`
	// +optional
	KeyName string `json:"key_name,omitempty"`
	// +optional
	ServiceAccessSecurityGroup string `json:"service_access_security_group,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
}

type EmrClusterSpecKerberosAttributes struct {
	// +optional
	AdDomainJoinPassword string `json:"ad_domain_join_password,omitempty"`
	// +optional
	AdDomainJoinUser string `json:"ad_domain_join_user,omitempty"`
	// +optional
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password,omitempty"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
	Realm                            string `json:"realm"`
}

type EmrClusterSpec struct {
	// +optional
	AdditionalInfo string `json:"additional_info,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Applications []string `json:"applications,omitempty"`
	// +optional
	AutoscalingRole string `json:"autoscaling_role,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BootstrapAction *[]EmrClusterSpec `json:"bootstrap_action,omitempty"`
	// +optional
	Configurations string `json:"configurations,omitempty"`
	// +optional
	ConfigurationsJson string `json:"configurations_json,omitempty"`
	// +optional
	CustomAmiId string `json:"custom_ami_id,omitempty"`
	// +optional
	EbsRootVolumeSize int `json:"ebs_root_volume_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ec2Attributes *[]EmrClusterSpec `json:"ec2_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KerberosAttributes *[]EmrClusterSpec `json:"kerberos_attributes,omitempty"`
	// +optional
	LogUri       string `json:"log_uri,omitempty"`
	Name         string `json:"name"`
	ReleaseLabel string `json:"release_label"`
	// +optional
	SecurityConfiguration string `json:"security_configuration,omitempty"`
	ServiceRole           string `json:"service_role"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	VisibleToAllUsers bool `json:"visible_to_all_users,omitempty"`
}

type EmrClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
