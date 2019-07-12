package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksMysqlLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksMysqlLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksMysqlLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksMysqlLayerSpec struct {
	ElasticLoadBalancer        string                      `json:"elastic_load_balancer"`
	CustomSetupRecipes         []string                    `json:"custom_setup_recipes"`
	CustomUndeployRecipes      []string                    `json:"custom_undeploy_recipes"`
	CustomJson                 string                      `json:"custom_json"`
	InstanceShutdownTimeout    int                         `json:"instance_shutdown_timeout"`
	StackId                    string                      `json:"stack_id"`
	AutoAssignElasticIps       bool                        `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps        bool                        `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn   string                      `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds     []string                    `json:"custom_security_group_ids"`
	AutoHealing                bool                        `json:"auto_healing"`
	SystemPackages             []string                    `json:"system_packages"`
	CustomDeployRecipes        []string                    `json:"custom_deploy_recipes"`
	CustomShutdownRecipes      []string                    `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot       bool                        `json:"install_updates_on_boot"`
	DrainElbOnShutdown         bool                        `json:"drain_elb_on_shutdown"`
	EbsVolume                  []AwsOpsworksMysqlLayerSpec `json:"ebs_volume"`
	CustomConfigureRecipes     []string                    `json:"custom_configure_recipes"`
	UseEbsOptimizedInstances   bool                        `json:"use_ebs_optimized_instances"`
	Name                       string                      `json:"name"`
	RootPassword               string                      `json:"root_password"`
	RootPasswordOnAllInstances bool                        `json:"root_password_on_all_instances"`
}



type AwsOpsworksMysqlLayerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayers
type AwsOpsworksMysqlLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksMysqlLayer CRD objects
	Items []AwsOpsworksMysqlLayer `json:"items,omitempty"`
}