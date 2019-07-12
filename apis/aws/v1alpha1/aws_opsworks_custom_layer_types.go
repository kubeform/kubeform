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

type AwsOpsworksCustomLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksCustomLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksCustomLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksCustomLayerSpecEbsVolume struct {
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

type AwsOpsworksCustomLayerSpec struct {
	AutoAssignPublicIps      bool                         `json:"auto_assign_public_ips"`
	CustomShutdownRecipes    []string                     `json:"custom_shutdown_recipes"`
	UseEbsOptimizedInstances bool                         `json:"use_ebs_optimized_instances"`
	EbsVolume                []AwsOpsworksCustomLayerSpec `json:"ebs_volume"`
	ShortName                string                       `json:"short_name"`
	CustomInstanceProfileArn string                       `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                       `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                     `json:"custom_configure_recipes"`
	CustomSecurityGroupIds   []string                     `json:"custom_security_group_ids"`
	CustomJson               string                       `json:"custom_json"`
	InstallUpdatesOnBoot     bool                         `json:"install_updates_on_boot"`
	SystemPackages           []string                     `json:"system_packages"`
	AutoAssignElasticIps     bool                         `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                     `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                     `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                     `json:"custom_undeploy_recipes"`
	StackId                  string                       `json:"stack_id"`
	AutoHealing              bool                         `json:"auto_healing"`
	InstanceShutdownTimeout  int                          `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                         `json:"drain_elb_on_shutdown"`
	Name                     string                       `json:"name"`
}

type AwsOpsworksCustomLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayers
type AwsOpsworksCustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksCustomLayer CRD objects
	Items []AwsOpsworksCustomLayer `json:"items,omitempty"`
}
