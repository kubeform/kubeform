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

type AwsOpsworksRailsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksRailsAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksRailsAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksRailsAppLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

type AwsOpsworksRailsAppLayerSpec struct {
	AutoAssignElasticIps     bool                           `json:"auto_assign_elastic_ips"`
	CustomShutdownRecipes    []string                       `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   []string                       `json:"custom_security_group_ids"`
	StackId                  string                         `json:"stack_id"`
	UseEbsOptimizedInstances bool                           `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool                           `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                         `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                       `json:"custom_setup_recipes"`
	AutoHealing              bool                           `json:"auto_healing"`
	Name                     string                         `json:"name"`
	PassengerVersion         string                         `json:"passenger_version"`
	BundlerVersion           string                         `json:"bundler_version"`
	ElasticLoadBalancer      string                         `json:"elastic_load_balancer"`
	DrainElbOnShutdown       bool                           `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                       `json:"system_packages"`
	EbsVolume                []AwsOpsworksRailsAppLayerSpec `json:"ebs_volume"`
	RubygemsVersion          string                         `json:"rubygems_version"`
	CustomConfigureRecipes   []string                       `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                       `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                       `json:"custom_undeploy_recipes"`
	CustomJson               string                         `json:"custom_json"`
	InstallUpdatesOnBoot     bool                           `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                            `json:"instance_shutdown_timeout"`
	RubyVersion              string                         `json:"ruby_version"`
	AppServer                string                         `json:"app_server"`
	ManageBundler            bool                           `json:"manage_bundler"`
}

type AwsOpsworksRailsAppLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayers
type AwsOpsworksRailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksRailsAppLayer CRD objects
	Items []AwsOpsworksRailsAppLayer `json:"items,omitempty"`
}
