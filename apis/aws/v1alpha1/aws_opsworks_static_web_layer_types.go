package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksStaticWebLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksStaticWebLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksStaticWebLayerSpecEbsVolume struct {
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	EbsVolume                []AwsOpsworksStaticWebLayerSpec `json:"ebs_volume"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	AutoHealing              bool                            `json:"auto_healing"`
	SystemPackages           []string                        `json:"system_packages"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	CustomJson               string                          `json:"custom_json"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	StackId                  string                          `json:"stack_id"`
	Name                     string                          `json:"name"`
}



type AwsOpsworksStaticWebLayerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayers
type AwsOpsworksStaticWebLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksStaticWebLayer CRD objects
	Items []AwsOpsworksStaticWebLayer `json:"items,omitempty"`
}