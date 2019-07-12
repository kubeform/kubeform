package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMemcachedLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksMemcachedLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksMemcachedLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksMemcachedLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksMemcachedLayerSpec struct {
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	StackId                  string                          `json:"stack_id"`
	EbsVolume                []AwsOpsworksMemcachedLayerSpec `json:"ebs_volume"`
	AutoHealing              bool                            `json:"auto_healing"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	SystemPackages           []string                        `json:"system_packages"`
	Name                     string                          `json:"name"`
	AllocatedMemory          int                             `json:"allocated_memory"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	CustomJson               string                          `json:"custom_json"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
}

type AwsOpsworksMemcachedLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayers
type AwsOpsworksMemcachedLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksMemcachedLayer CRD objects
	Items []AwsOpsworksMemcachedLayer `json:"items,omitempty"`
}
