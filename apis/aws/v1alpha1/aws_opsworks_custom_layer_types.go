package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksCustomLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksCustomLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksCustomLayerSpecEbsVolume struct {
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
}

type AwsOpsworksCustomLayerSpec struct {
	CustomSecurityGroupIds   []string                     `json:"custom_security_group_ids"`
	CustomJson               string                       `json:"custom_json"`
	StackId                  string                       `json:"stack_id"`
	EbsVolume                []AwsOpsworksCustomLayerSpec `json:"ebs_volume"`
	AutoAssignElasticIps     bool                         `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn string                       `json:"custom_instance_profile_arn"`
	CustomUndeployRecipes    []string                     `json:"custom_undeploy_recipes"`
	AutoHealing              bool                         `json:"auto_healing"`
	CustomDeployRecipes      []string                     `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                     `json:"custom_shutdown_recipes"`
	UseEbsOptimizedInstances bool                         `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool                         `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                       `json:"elastic_load_balancer"`
	CustomSetupRecipes       []string                     `json:"custom_setup_recipes"`
	DrainElbOnShutdown       bool                         `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                     `json:"system_packages"`
	ShortName                string                       `json:"short_name"`
	Name                     string                       `json:"name"`
	CustomConfigureRecipes   []string                     `json:"custom_configure_recipes"`
	InstallUpdatesOnBoot     bool                         `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                          `json:"instance_shutdown_timeout"`
}

type AwsOpsworksCustomLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayers
type AwsOpsworksCustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksCustomLayer CRD objects
	Items []AwsOpsworksCustomLayer `json:"items,omitempty"`
}
