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

type AwsOpsworksMemcachedLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksMemcachedLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksMemcachedLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksMemcachedLayerSpecEbsVolume struct {
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

type AwsOpsworksMemcachedLayerSpec struct {
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
	Name                     string                          `json:"name"`
	AutoHealing              bool                            `json:"auto_healing"`
	SystemPackages           []string                        `json:"system_packages"`
	StackId                  string                          `json:"stack_id"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	AllocatedMemory          int                             `json:"allocated_memory"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomJson               string                          `json:"custom_json"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	EbsVolume                []AwsOpsworksMemcachedLayerSpec `json:"ebs_volume"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
}

type AwsOpsworksMemcachedLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayers
type AwsOpsworksMemcachedLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksMemcachedLayer CRD objects
	Items []AwsOpsworksMemcachedLayer `json:"items,omitempty"`
}
