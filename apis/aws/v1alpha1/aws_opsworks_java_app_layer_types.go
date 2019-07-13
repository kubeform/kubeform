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
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type AwsOpsworksJavaAppLayerSpec struct {
	CustomConfigureRecipes   []string                      `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                      `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                      `json:"custom_shutdown_recipes"`
	CustomJson               string                        `json:"custom_json"`
	DrainElbOnShutdown       bool                          `json:"drain_elb_on_shutdown"`
	SystemPackages           []string                      `json:"system_packages"`
	CustomSetupRecipes       []string                      `json:"custom_setup_recipes"`
	CustomInstanceProfileArn string                        `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                        `json:"elastic_load_balancer"`
	CustomSecurityGroupIds   []string                      `json:"custom_security_group_ids"`
	AutoHealing              bool                          `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                          `json:"install_updates_on_boot"`
	EbsVolume                []AwsOpsworksJavaAppLayerSpec `json:"ebs_volume"`
	JvmType                  string                        `json:"jvm_type"`
	AutoAssignElasticIps     bool                          `json:"auto_assign_elastic_ips"`
	AppServer                string                        `json:"app_server"`
	JvmVersion               string                        `json:"jvm_version"`
	UseEbsOptimizedInstances bool                          `json:"use_ebs_optimized_instances"`
	Name                     string                        `json:"name"`
	JvmOptions               string                        `json:"jvm_options"`
	InstanceShutdownTimeout  int                           `json:"instance_shutdown_timeout"`
	CustomUndeployRecipes    []string                      `json:"custom_undeploy_recipes"`
	StackId                  string                        `json:"stack_id"`
	AppServerVersion         string                        `json:"app_server_version"`
	AutoAssignPublicIps      bool                          `json:"auto_assign_public_ips"`
}



type AwsOpsworksJavaAppLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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