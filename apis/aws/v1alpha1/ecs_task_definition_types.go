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

type EcsTaskDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsTaskDefinitionSpec   `json:"spec,omitempty"`
	Status            EcsTaskDefinitionStatus `json:"status,omitempty"`
}

type EcsTaskDefinitionSpecPlacementConstraints struct {
	// +optional
	Expression string `json:"expression,omitempty"`
	Type       string `json:"type"`
}

type EcsTaskDefinitionSpecProxyConfiguration struct {
	ContainerName string `json:"container_name"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration struct {
	// +optional
	Autoprovision bool `json:"autoprovision,omitempty"`
	// +optional
	Driver string `json:"driver,omitempty"`
	// +optional
	DriverOpts map[string]string `json:"driver_opts,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type EcsTaskDefinitionSpecVolume struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DockerVolumeConfiguration *[]EcsTaskDefinitionSpecVolume `json:"docker_volume_configuration,omitempty"`
	// +optional
	HostPath string `json:"host_path,omitempty"`
	Name     string `json:"name"`
}

type EcsTaskDefinitionSpec struct {
	ContainerDefinitions string `json:"container_definitions"`
	// +optional
	Cpu string `json:"cpu,omitempty"`
	// +optional
	ExecutionRoleArn string `json:"execution_role_arn,omitempty"`
	Family           string `json:"family"`
	// +optional
	IpcMode string `json:"ipc_mode,omitempty"`
	// +optional
	Memory string `json:"memory,omitempty"`
	// +optional
	PidMode string `json:"pid_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	PlacementConstraints *[]EcsTaskDefinitionSpec `json:"placement_constraints,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProxyConfiguration *[]EcsTaskDefinitionSpec `json:"proxy_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RequiresCompatibilities []string `json:"requires_compatibilities,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TaskRoleArn string `json:"task_role_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Volume *[]EcsTaskDefinitionSpec `json:"volume,omitempty"`
}

type EcsTaskDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
