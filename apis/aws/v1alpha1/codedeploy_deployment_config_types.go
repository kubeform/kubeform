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

type CodedeployDeploymentConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentConfigSpec   `json:"spec,omitempty"`
	Status            CodedeployDeploymentConfigStatus `json:"status,omitempty"`
}

type CodedeployDeploymentConfigSpecMinimumHealthyHosts struct {
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Value int `json:"value,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary struct {
	// +optional
	Interval int `json:"interval,omitempty"`
	// +optional
	Percentage int `json:"percentage,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear struct {
	// +optional
	Interval int `json:"interval,omitempty"`
	// +optional
	Percentage int `json:"percentage,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimeBasedCanary *[]CodedeployDeploymentConfigSpecTrafficRoutingConfig `json:"time_based_canary,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimeBasedLinear *[]CodedeployDeploymentConfigSpecTrafficRoutingConfig `json:"time_based_linear,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type CodedeployDeploymentConfigSpec struct {
	// +optional
	ComputePlatform      string `json:"compute_platform,omitempty"`
	DeploymentConfigName string `json:"deployment_config_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MinimumHealthyHosts *[]CodedeployDeploymentConfigSpec `json:"minimum_healthy_hosts,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TrafficRoutingConfig *[]CodedeployDeploymentConfigSpec `json:"traffic_routing_config,omitempty"`
}

type CodedeployDeploymentConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodedeployDeploymentConfigList is a list of CodedeployDeploymentConfigs
type CodedeployDeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodedeployDeploymentConfig CRD objects
	Items []CodedeployDeploymentConfig `json:"items,omitempty"`
}
