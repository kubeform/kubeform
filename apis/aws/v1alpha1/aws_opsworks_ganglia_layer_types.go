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

type AwsOpsworksGangliaLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksGangliaLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksGangliaLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksGangliaLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksGangliaLayerSpec struct {
	CustomConfigureRecipes   []string                      `json:"custom_configure_recipes"`
	AutoHealing              bool                          `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                          `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                           `json:"instance_shutdown_timeout"`
	SystemPackages           []string                      `json:"system_packages"`
	Name                     string                        `json:"name"`
	Password                 string                        `json:"password"`
	AutoAssignPublicIps      bool                          `json:"auto_assign_public_ips"`
	CustomSetupRecipes       []string                      `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                      `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                      `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   []string                      `json:"custom_security_group_ids"`
	CustomJson               string                        `json:"custom_json"`
	DrainElbOnShutdown       bool                          `json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances bool                          `json:"use_ebs_optimized_instances"`
	AutoAssignElasticIps     bool                          `json:"auto_assign_elastic_ips"`
	Url                      string                        `json:"url"`
	EbsVolume                []AwsOpsworksGangliaLayerSpec `json:"ebs_volume"`
	Username                 string                        `json:"username"`
	CustomInstanceProfileArn string                        `json:"custom_instance_profile_arn"`
	CustomShutdownRecipes    []string                      `json:"custom_shutdown_recipes"`
	StackId                  string                        `json:"stack_id"`
	ElasticLoadBalancer      string                        `json:"elastic_load_balancer"`
}



type AwsOpsworksGangliaLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksGangliaLayerList is a list of AwsOpsworksGangliaLayers
type AwsOpsworksGangliaLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksGangliaLayer CRD objects
	Items []AwsOpsworksGangliaLayer `json:"items,omitempty"`
}