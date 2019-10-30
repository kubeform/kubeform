/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type CodedeployDeploymentConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentConfigSpec   `json:"spec,omitempty"`
	Status            CodedeployDeploymentConfigStatus `json:"status,omitempty"`
}

type CodedeployDeploymentConfigSpecMinimumHealthyHosts struct {
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value int `json:"value,omitempty" tf:"value,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary struct {
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	Percentage int `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear struct {
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	Percentage int `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type CodedeployDeploymentConfigSpecTrafficRoutingConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimeBasedCanary []CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary `json:"timeBasedCanary,omitempty" tf:"time_based_canary,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimeBasedLinear []CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear `json:"timeBasedLinear,omitempty" tf:"time_based_linear,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type CodedeployDeploymentConfigSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ComputePlatform string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`
	// +optional
	DeploymentConfigID   string `json:"deploymentConfigID,omitempty" tf:"deployment_config_id,omitempty"`
	DeploymentConfigName string `json:"deploymentConfigName" tf:"deployment_config_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MinimumHealthyHosts []CodedeployDeploymentConfigSpecMinimumHealthyHosts `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TrafficRoutingConfig []CodedeployDeploymentConfigSpecTrafficRoutingConfig `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config,omitempty"`
}

type CodedeployDeploymentConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodedeployDeploymentConfigSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
