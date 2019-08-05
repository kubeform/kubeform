package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

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

type EcsTaskDefinitionSpecProxyConfiguration struct {
	ContainerName string `json:"containerName" tf:"container_name"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
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

	Arn                  string `json:"arn" tf:"arn"`
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
	// +kubebuilder:validation:UniqueItems=true
	PlacementConstraints []EcsTaskDefinitionSpecPlacementConstraints `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProxyConfiguration []EcsTaskDefinitionSpecProxyConfiguration `json:"proxyConfiguration,omitempty" tf:"proxy_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RequiresCompatibilities []string `json:"requiresCompatibilities,omitempty" tf:"requires_compatibilities,omitempty"`
	Revision                int      `json:"revision" tf:"revision"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TaskRoleArn string `json:"taskRoleArn,omitempty" tf:"task_role_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Volume []EcsTaskDefinitionSpecVolume `json:"volume,omitempty" tf:"volume,omitempty"`
}

type EcsTaskDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EcsTaskDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
