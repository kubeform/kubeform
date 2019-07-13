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
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

type AwsOpsworksRailsAppLayerSpec struct {
	CustomSecurityGroupIds   []string                       `json:"custom_security_group_ids"`
	AutoHealing              bool                           `json:"auto_healing"`
	InstanceShutdownTimeout  int                            `json:"instance_shutdown_timeout"`
	RubygemsVersion          string                         `json:"rubygems_version"`
	ManageBundler            bool                           `json:"manage_bundler"`
	BundlerVersion           string                         `json:"bundler_version"`
	CustomSetupRecipes       []string                       `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                       `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                       `json:"custom_shutdown_recipes"`
	CustomJson               string                         `json:"custom_json"`
	EbsVolume                []AwsOpsworksRailsAppLayerSpec `json:"ebs_volume"`
	PassengerVersion         string                         `json:"passenger_version"`
	Name                     string                         `json:"name"`
	RubyVersion              string                         `json:"ruby_version"`
	CustomInstanceProfileArn string                         `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                         `json:"elastic_load_balancer"`
	CustomUndeployRecipes    []string                       `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                           `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                       `json:"system_packages"`
	UseEbsOptimizedInstances bool                           `json:"use_ebs_optimized_instances"`
	AppServer                string                         `json:"app_server"`
	AutoAssignElasticIps     bool                           `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                           `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string                       `json:"custom_configure_recipes"`
	InstallUpdatesOnBoot     bool                           `json:"install_updates_on_boot"`
	StackId                  string                         `json:"stack_id"`
}



type AwsOpsworksRailsAppLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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