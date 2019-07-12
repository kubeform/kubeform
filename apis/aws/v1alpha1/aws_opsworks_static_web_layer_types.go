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

type AwsOpsworksStaticWebLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksStaticWebLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksStaticWebLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksStaticWebLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	SystemPackages           []string                        `json:"system_packages"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	AutoHealing              bool                            `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	EbsVolume                []AwsOpsworksStaticWebLayerSpec `json:"ebs_volume"`
	Name                     string                          `json:"name"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	CustomJson               string                          `json:"custom_json"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	StackId                  string                          `json:"stack_id"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
}

type AwsOpsworksStaticWebLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayers
type AwsOpsworksStaticWebLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksStaticWebLayer CRD objects
	Items []AwsOpsworksStaticWebLayer `json:"items,omitempty"`
}
