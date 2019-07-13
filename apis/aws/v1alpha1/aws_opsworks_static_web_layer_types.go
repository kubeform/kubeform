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
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	Name                     string                          `json:"name"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomJson               string                          `json:"custom_json"`
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	AutoHealing              bool                            `json:"auto_healing"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	StackId                  string                          `json:"stack_id"`
	EbsVolume                []AwsOpsworksStaticWebLayerSpec `json:"ebs_volume"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	SystemPackages           []string                        `json:"system_packages"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
}



type AwsOpsworksStaticWebLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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