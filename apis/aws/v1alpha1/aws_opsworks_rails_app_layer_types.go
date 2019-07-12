package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRailsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksRailsAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksRailsAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksRailsAppLayerSpecEbsVolume struct {
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

type AwsOpsworksRailsAppLayerSpec struct {
	AutoAssignElasticIps     bool                           `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                       `json:"custom_setup_recipes"`
	CustomJson               string                         `json:"custom_json"`
	InstanceShutdownTimeout  int                            `json:"instance_shutdown_timeout"`
	StackId                  string                         `json:"stack_id"`
	EbsVolume                []AwsOpsworksRailsAppLayerSpec `json:"ebs_volume"`
	Name                     string                         `json:"name"`
	ManageBundler            bool                           `json:"manage_bundler"`
	RubyVersion              string                         `json:"ruby_version"`
	AppServer                string                         `json:"app_server"`
	CustomShutdownRecipes    []string                       `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot     bool                           `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                           `json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances bool                           `json:"use_ebs_optimized_instances"`
	CustomInstanceProfileArn string                         `json:"custom_instance_profile_arn"`
	CustomConfigureRecipes   []string                       `json:"custom_configure_recipes"`
	AutoHealing              bool                           `json:"auto_healing"`
	PassengerVersion         string                         `json:"passenger_version"`
	AutoAssignPublicIps      bool                           `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                         `json:"elastic_load_balancer"`
	CustomDeployRecipes      []string                       `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                       `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   []string                       `json:"custom_security_group_ids"`
	SystemPackages           []string                       `json:"system_packages"`
	RubygemsVersion          string                         `json:"rubygems_version"`
	BundlerVersion           string                         `json:"bundler_version"`
}



type AwsOpsworksRailsAppLayerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayers
type AwsOpsworksRailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksRailsAppLayer CRD objects
	Items []AwsOpsworksRailsAppLayer `json:"items,omitempty"`
}