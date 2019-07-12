package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksPhpAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksPhpAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksPhpAppLayerSpecEbsVolume struct {
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
}

type AwsOpsworksPhpAppLayerSpec struct {
	UseEbsOptimizedInstances bool                         `json:"use_ebs_optimized_instances"`
	Name                     string                       `json:"name"`
	AutoAssignElasticIps     bool                         `json:"auto_assign_elastic_ips"`
	CustomUndeployRecipes    []string                     `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   []string                     `json:"custom_security_group_ids"`
	CustomJson               string                       `json:"custom_json"`
	SystemPackages           []string                     `json:"system_packages"`
	EbsVolume                []AwsOpsworksPhpAppLayerSpec `json:"ebs_volume"`
	CustomInstanceProfileArn string                       `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                       `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                     `json:"custom_configure_recipes"`
	AutoHealing              bool                         `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                         `json:"install_updates_on_boot"`
	CustomSetupRecipes       []string                     `json:"custom_setup_recipes"`
	InstanceShutdownTimeout  int                          `json:"instance_shutdown_timeout"`
	AutoAssignPublicIps      bool                         `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                     `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                     `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool                         `json:"drain_elb_on_shutdown"`
	StackId                  string                       `json:"stack_id"`
}



type AwsOpsworksPhpAppLayerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayerList is a list of AwsOpsworksPhpAppLayers
type AwsOpsworksPhpAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksPhpAppLayer CRD objects
	Items []AwsOpsworksPhpAppLayer `json:"items,omitempty"`
}