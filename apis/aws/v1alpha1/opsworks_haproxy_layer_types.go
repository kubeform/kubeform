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

type OpsworksHaproxyLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksHaproxyLayerSpec   `json:"spec,omitempty"`
	Status            OpsworksHaproxyLayerStatus `json:"status,omitempty"`
}

type OpsworksHaproxyLayerSpecEbsVolume struct {
	// +optional
	Iops          int    `json:"iops,omitempty"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	// +optional
	RaidLevel string `json:"raid_level,omitempty"`
	Size      int    `json:"size"`
	// +optional
	Type string `json:"type,omitempty"`
}

type OpsworksHaproxyLayerSpec struct {
	// +optional
	AutoAssignElasticIps bool `json:"auto_assign_elastic_ips,omitempty"`
	// +optional
	AutoAssignPublicIps bool `json:"auto_assign_public_ips,omitempty"`
	// +optional
	AutoHealing bool `json:"auto_healing,omitempty"`
	// +optional
	CustomConfigureRecipes []string `json:"custom_configure_recipes,omitempty"`
	// +optional
	CustomDeployRecipes []string `json:"custom_deploy_recipes,omitempty"`
	// +optional
	CustomInstanceProfileArn string `json:"custom_instance_profile_arn,omitempty"`
	// +optional
	CustomJson string `json:"custom_json,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomSecurityGroupIds []string `json:"custom_security_group_ids,omitempty"`
	// +optional
	CustomSetupRecipes []string `json:"custom_setup_recipes,omitempty"`
	// +optional
	CustomShutdownRecipes []string `json:"custom_shutdown_recipes,omitempty"`
	// +optional
	CustomUndeployRecipes []string `json:"custom_undeploy_recipes,omitempty"`
	// +optional
	DrainElbOnShutdown bool `json:"drain_elb_on_shutdown,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsVolume *[]OpsworksHaproxyLayerSpec `json:"ebs_volume,omitempty"`
	// +optional
	ElasticLoadBalancer string `json:"elastic_load_balancer,omitempty"`
	// +optional
	HealthcheckMethod string `json:"healthcheck_method,omitempty"`
	// +optional
	HealthcheckUrl string `json:"healthcheck_url,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceShutdownTimeout int `json:"instance_shutdown_timeout,omitempty"`
	// +optional
	Name    string `json:"name,omitempty"`
	StackId string `json:"stack_id"`
	// +optional
	StatsEnabled  bool   `json:"stats_enabled,omitempty"`
	StatsPassword string `json:"stats_password"`
	// +optional
	StatsUrl string `json:"stats_url,omitempty"`
	// +optional
	StatsUser string `json:"stats_user,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SystemPackages []string `json:"system_packages,omitempty"`
	// +optional
	UseEbsOptimizedInstances bool `json:"use_ebs_optimized_instances,omitempty"`
}

type OpsworksHaproxyLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksHaproxyLayerList is a list of OpsworksHaproxyLayers
type OpsworksHaproxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksHaproxyLayer CRD objects
	Items []OpsworksHaproxyLayer `json:"items,omitempty"`
}
