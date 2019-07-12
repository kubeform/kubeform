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

type AwsOpsworksMysqlLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksMysqlLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksMysqlLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksMysqlLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

type AwsOpsworksMysqlLayerSpec struct {
	StackId                    string                      `json:"stack_id"`
	RootPassword               string                      `json:"root_password"`
	ElasticLoadBalancer        string                      `json:"elastic_load_balancer"`
	CustomConfigureRecipes     []string                    `json:"custom_configure_recipes"`
	CustomShutdownRecipes      []string                    `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds     []string                    `json:"custom_security_group_ids"`
	CustomJson                 string                      `json:"custom_json"`
	AutoHealing                bool                        `json:"auto_healing"`
	InstanceShutdownTimeout    int                         `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances   bool                        `json:"use_ebs_optimized_instances"`
	EbsVolume                  []AwsOpsworksMysqlLayerSpec `json:"ebs_volume"`
	RootPasswordOnAllInstances bool                        `json:"root_password_on_all_instances"`
	AutoAssignPublicIps        bool                        `json:"auto_assign_public_ips"`
	CustomSetupRecipes         []string                    `json:"custom_setup_recipes"`
	CustomDeployRecipes        []string                    `json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot       bool                        `json:"install_updates_on_boot"`
	DrainElbOnShutdown         bool                        `json:"drain_elb_on_shutdown"`
	SystemPackages             []string                    `json:"system_packages"`
	AutoAssignElasticIps       bool                        `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn   string                      `json:"custom_instance_profile_arn"`
	CustomUndeployRecipes      []string                    `json:"custom_undeploy_recipes"`
	Name                       string                      `json:"name"`
}

type AwsOpsworksMysqlLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayers
type AwsOpsworksMysqlLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksMysqlLayer CRD objects
	Items []AwsOpsworksMysqlLayer `json:"items,omitempty"`
}
