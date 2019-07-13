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

type AwsCodedeployDeploymentConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodedeployDeploymentConfigSpec   `json:"spec,omitempty"`
	Status            AwsCodedeployDeploymentConfigStatus `json:"status,omitempty"`
}

type AwsCodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary struct {
	Interval   int `json:"interval"`
	Percentage int `json:"percentage"`
}

type AwsCodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear struct {
	Interval   int `json:"interval"`
	Percentage int `json:"percentage"`
}

type AwsCodedeployDeploymentConfigSpecTrafficRoutingConfig struct {
	Type            string                                                  `json:"type"`
	TimeBasedCanary []AwsCodedeployDeploymentConfigSpecTrafficRoutingConfig `json:"time_based_canary"`
	TimeBasedLinear []AwsCodedeployDeploymentConfigSpecTrafficRoutingConfig `json:"time_based_linear"`
}

type AwsCodedeployDeploymentConfigSpecMinimumHealthyHosts struct {
	Value int    `json:"value"`
	Type  string `json:"type"`
}

type AwsCodedeployDeploymentConfigSpec struct {
	TrafficRoutingConfig []AwsCodedeployDeploymentConfigSpec `json:"traffic_routing_config"`
	DeploymentConfigId   string                              `json:"deployment_config_id"`
	DeploymentConfigName string                              `json:"deployment_config_name"`
	ComputePlatform      string                              `json:"compute_platform"`
	MinimumHealthyHosts  []AwsCodedeployDeploymentConfigSpec `json:"minimum_healthy_hosts"`
}



type AwsCodedeployDeploymentConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodedeployDeploymentConfigList is a list of AwsCodedeployDeploymentConfigs
type AwsCodedeployDeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodedeployDeploymentConfig CRD objects
	Items []AwsCodedeployDeploymentConfig `json:"items,omitempty"`
}