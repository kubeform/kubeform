package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksNodejsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksNodejsAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksNodejsAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksNodejsAppLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

type AwsOpsworksNodejsAppLayerSpec struct {
	EbsVolume                []AwsOpsworksNodejsAppLayerSpec `json:"ebs_volume"`
	AutoAssignPublicIps      bool                            `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                          `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                          `json:"elastic_load_balancer"`
	CustomShutdownRecipes    []string                        `json:"custom_shutdown_recipes"`
	UseEbsOptimizedInstances bool                            `json:"use_ebs_optimized_instances"`
	CustomDeployRecipes      []string                        `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                        `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                            `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                        `json:"system_packages"`
	CustomJson               string                          `json:"custom_json"`
	Name                     string                          `json:"name"`
	NodejsVersion            string                          `json:"nodejs_version"`
	InstallUpdatesOnBoot     bool                            `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                             `json:"instance_shutdown_timeout"`
	StackId                  string                          `json:"stack_id"`
	AutoAssignElasticIps     bool                            `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                        `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                        `json:"custom_configure_recipes"`
	CustomSecurityGroupIds   []string                        `json:"custom_security_group_ids"`
	AutoHealing              bool                            `json:"auto_healing"`
}



type AwsOpsworksNodejsAppLayerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayerList is a list of AwsOpsworksNodejsAppLayers
type AwsOpsworksNodejsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksNodejsAppLayer CRD objects
	Items []AwsOpsworksNodejsAppLayer `json:"items,omitempty"`
}