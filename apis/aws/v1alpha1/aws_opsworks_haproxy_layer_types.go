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

type AwsOpsworksHaproxyLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksHaproxyLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksHaproxyLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksHaproxyLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

type AwsOpsworksHaproxyLayerSpec struct {
	AutoAssignElasticIps     bool                          `json:"auto_assign_elastic_ips"`
	AutoHealing              bool                          `json:"auto_healing"`
	StackId                  string                        `json:"stack_id"`
	UseEbsOptimizedInstances bool                          `json:"use_ebs_optimized_instances"`
	StatsEnabled             bool                          `json:"stats_enabled"`
	StatsUser                string                        `json:"stats_user"`
	HealthcheckMethod        string                        `json:"healthcheck_method"`
	CustomUndeployRecipes    []string                      `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   []string                      `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int                           `json:"instance_shutdown_timeout"`
	CustomInstanceProfileArn string                        `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                      `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                      `json:"custom_shutdown_recipes"`
	CustomJson               string                        `json:"custom_json"`
	Name                     string                        `json:"name"`
	StatsUrl                 string                        `json:"stats_url"`
	AutoAssignPublicIps      bool                          `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                        `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                      `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                      `json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot     bool                          `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                          `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                      `json:"system_packages"`
	EbsVolume                []AwsOpsworksHaproxyLayerSpec `json:"ebs_volume"`
	StatsPassword            string                        `json:"stats_password"`
	HealthcheckUrl           string                        `json:"healthcheck_url"`
}



type AwsOpsworksHaproxyLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayers
type AwsOpsworksHaproxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksHaproxyLayer CRD objects
	Items []AwsOpsworksHaproxyLayer `json:"items,omitempty"`
}