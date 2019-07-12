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

type AwsOpsworksJavaAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksJavaAppLayerSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksJavaAppLayerStatus `json:"status,omitempty"`
}

type AwsOpsworksJavaAppLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

type AwsOpsworksJavaAppLayerSpec struct {
	AutoHealing              bool                          `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                          `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                          `json:"drain_elb_on_shutdown"`
	StackId                  string                        `json:"stack_id"`
	JvmOptions               string                        `json:"jvm_options"`
	AppServer                string                        `json:"app_server"`
	CustomShutdownRecipes    []string                      `json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout  int                           `json:"instance_shutdown_timeout"`
	SystemPackages           []string                      `json:"system_packages"`
	CustomUndeployRecipes    []string                      `json:"custom_undeploy_recipes"`
	ElasticLoadBalancer      string                        `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                      `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                      `json:"custom_deploy_recipes"`
	CustomSecurityGroupIds   []string                      `json:"custom_security_group_ids"`
	EbsVolume                []AwsOpsworksJavaAppLayerSpec `json:"ebs_volume"`
	AppServerVersion         string                        `json:"app_server_version"`
	JvmType                  string                        `json:"jvm_type"`
	CustomInstanceProfileArn string                        `json:"custom_instance_profile_arn"`
	AutoAssignPublicIps      bool                          `json:"auto_assign_public_ips"`
	CustomSetupRecipes       []string                      `json:"custom_setup_recipes"`
	CustomJson               string                        `json:"custom_json"`
	UseEbsOptimizedInstances bool                          `json:"use_ebs_optimized_instances"`
	Name                     string                        `json:"name"`
	JvmVersion               string                        `json:"jvm_version"`
	AutoAssignElasticIps     bool                          `json:"auto_assign_elastic_ips"`
}

type AwsOpsworksJavaAppLayerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksJavaAppLayerList is a list of AwsOpsworksJavaAppLayers
type AwsOpsworksJavaAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksJavaAppLayer CRD objects
	Items []AwsOpsworksJavaAppLayer `json:"items,omitempty"`
}
