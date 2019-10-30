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

type EcsTaskDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsTaskDefinitionSpec   `json:"spec,omitempty"`
	Status            EcsTaskDefinitionStatus `json:"status,omitempty"`
}

type EcsTaskDefinitionSpecPlacementConstraints struct {
	// +optional
	Expression string `json:"expression,omitempty" tf:"expression,omitempty"`
	Type       string `json:"type" tf:"type"`
}

type EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration struct {
	// +optional
	Autoprovision bool `json:"autoprovision,omitempty" tf:"autoprovision,omitempty"`
	// +optional
	Driver string `json:"driver,omitempty" tf:"driver,omitempty"`
	// +optional
	DriverOpts map[string]string `json:"driverOpts,omitempty" tf:"driver_opts,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EcsTaskDefinitionSpecVolume struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DockerVolumeConfiguration []EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration `json:"dockerVolumeConfiguration,omitempty" tf:"docker_volume_configuration,omitempty"`
	// +optional
	HostPath string `json:"hostPath,omitempty" tf:"host_path,omitempty"`
	Name     string `json:"name" tf:"name"`
}

type EcsTaskDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                  string `json:"arn,omitempty" tf:"arn,omitempty"`
	ContainerDefinitions string `json:"containerDefinitions" tf:"container_definitions"`
	// +optional
	Cpu string `json:"cpu,omitempty" tf:"cpu,omitempty"`
	// +optional
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`
	Family           string `json:"family" tf:"family"`
	// +optional
	IpcMode string `json:"ipcMode,omitempty" tf:"ipc_mode,omitempty"`
	// +optional
	Memory string `json:"memory,omitempty" tf:"memory,omitempty"`
	// +optional
	NetworkMode string `json:"networkMode,omitempty" tf:"network_mode,omitempty"`
	// +optional
	PidMode string `json:"pidMode,omitempty" tf:"pid_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PlacementConstraints []EcsTaskDefinitionSpecPlacementConstraints `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`
	// +optional
	RequiresCompatibilities []string `json:"requiresCompatibilities,omitempty" tf:"requires_compatibilities,omitempty"`
	// +optional
	Revision int `json:"revision,omitempty" tf:"revision,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TaskRoleArn string `json:"taskRoleArn,omitempty" tf:"task_role_arn,omitempty"`
	// +optional
	Volume []EcsTaskDefinitionSpecVolume `json:"volume,omitempty" tf:"volume,omitempty"`
}

type EcsTaskDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EcsTaskDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcsTaskDefinitionList is a list of EcsTaskDefinitions
type EcsTaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcsTaskDefinition CRD objects
	Items []EcsTaskDefinition `json:"items,omitempty"`
}
