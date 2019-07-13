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

type AwsOpsworksNodejsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksNodejsAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksNodejsAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksNodejsAppLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksNodejsAppLayerSpec struct {
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	Name                     string                          `json:"name"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	EbsVolume                []AwsOpsworksNodejsAppLayerSpec `json:"ebs_volume"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	StackId                  string                          `json:"stack_id"`
	CustomJson               string                          `json:"custom_json"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                        `json:"system_packages"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	AutoHealing              bool                            `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	NodejsVersion            string                          `json:"nodejs_version"`
}



type AwsOpsworksNodejsAppLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksNodejsAppLayerList is a list of AwsOpsworksNodejsAppLayers
type AwsOpsworksNodejsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksNodejsAppLayer CRD objects
	Items []AwsOpsworksNodejsAppLayer `json:"items,omitempty"`
}