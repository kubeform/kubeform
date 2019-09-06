package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OpsworksRailsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksRailsAppLayerSpec   `json:"spec,omitempty"`
	Status            OpsworksRailsAppLayerStatus `json:"status,omitempty"`
}

type OpsworksRailsAppLayerSpecEbsVolume struct {
	// +optional
	Iops          int    `json:"iops,omitempty" tf:"iops,omitempty"`
	MountPoint    string `json:"mountPoint" tf:"mount_point"`
	NumberOfDisks int    `json:"numberOfDisks" tf:"number_of_disks"`
	// +optional
	RaidLevel string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`
	Size      int    `json:"size" tf:"size"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type OpsworksRailsAppLayerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppServer string `json:"appServer,omitempty" tf:"app_server,omitempty"`
	// +optional
	AutoAssignElasticIPS bool `json:"autoAssignElasticIPS,omitempty" tf:"auto_assign_elastic_ips,omitempty"`
	// +optional
	AutoAssignPublicIPS bool `json:"autoAssignPublicIPS,omitempty" tf:"auto_assign_public_ips,omitempty"`
	// +optional
	AutoHealing bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`
	// +optional
	BundlerVersion string `json:"bundlerVersion,omitempty" tf:"bundler_version,omitempty"`
	// +optional
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`
	// +optional
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`
	// +optional
	CustomInstanceProfileArn string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`
	// +optional
	CustomJSON string `json:"customJSON,omitempty" tf:"custom_json,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomSecurityGroupIDS []string `json:"customSecurityGroupIDS,omitempty" tf:"custom_security_group_ids,omitempty"`
	// +optional
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`
	// +optional
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`
	// +optional
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`
	// +optional
	DrainElbOnShutdown bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsVolume []OpsworksRailsAppLayerSpecEbsVolume `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`
	// +optional
	ElasticLoadBalancer string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceShutdownTimeout int `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`
	// +optional
	ManageBundler bool `json:"manageBundler,omitempty" tf:"manage_bundler,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	PassengerVersion string `json:"passengerVersion,omitempty" tf:"passenger_version,omitempty"`
	// +optional
	RubyVersion string `json:"rubyVersion,omitempty" tf:"ruby_version,omitempty"`
	// +optional
	RubygemsVersion string `json:"rubygemsVersion,omitempty" tf:"rubygems_version,omitempty"`
	StackID         string `json:"stackID" tf:"stack_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`
	// +optional
	UseEbsOptimizedInstances bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type OpsworksRailsAppLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OpsworksRailsAppLayerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksRailsAppLayerList is a list of OpsworksRailsAppLayers
type OpsworksRailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksRailsAppLayer CRD objects
	Items []OpsworksRailsAppLayer `json:"items,omitempty"`
}
