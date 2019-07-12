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
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
}

type AwsOpsworksHaproxyLayerSpec struct {
	ElasticLoadBalancer      string                        `json:"elastic_load_balancer"`
	AutoHealing              bool                          `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                          `json:"install_updates_on_boot"`
	EbsVolume                []AwsOpsworksHaproxyLayerSpec `json:"ebs_volume"`
	StatsPassword            string                        `json:"stats_password"`
	AutoAssignPublicIps      bool                          `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                      `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                      `json:"custom_undeploy_recipes"`
	SystemPackages           []string                      `json:"system_packages"`
	UseEbsOptimizedInstances bool                          `json:"use_ebs_optimized_instances"`
	Name                     string                        `json:"name"`
	StatsUser                string                        `json:"stats_user"`
	StatsUrl                 string                        `json:"stats_url"`
	DrainElbOnShutdown       bool                          `json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps     bool                          `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn string                        `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                      `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                      `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string                      `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   []string                      `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int                           `json:"instance_shutdown_timeout"`
	StatsEnabled             bool                          `json:"stats_enabled"`
	CustomJson               string                        `json:"custom_json"`
	StackId                  string                        `json:"stack_id"`
	HealthcheckUrl           string                        `json:"healthcheck_url"`
	HealthcheckMethod        string                        `json:"healthcheck_method"`
}

type AwsOpsworksHaproxyLayerStatus struct {
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
