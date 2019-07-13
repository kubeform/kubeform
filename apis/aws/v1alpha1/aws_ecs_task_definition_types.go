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

type AwsEcsTaskDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcsTaskDefinitionSpec   `json:"spec,omitempty"`
	Status            AwsEcsTaskDefinitionStatus `json:"status,omitempty"`
}

type AwsEcsTaskDefinitionSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type AwsEcsTaskDefinitionSpecProxyConfiguration struct {
	ContainerName string            `json:"container_name"`
	Properties    map[string]string `json:"properties"`
	Type          string            `json:"type"`
}

type AwsEcsTaskDefinitionSpecVolumeDockerVolumeConfiguration struct {
	Driver        string            `json:"driver"`
	DriverOpts    map[string]string `json:"driver_opts"`
	Labels        map[string]string `json:"labels"`
	Scope         string            `json:"scope"`
	Autoprovision bool              `json:"autoprovision"`
}

type AwsEcsTaskDefinitionSpecVolume struct {
	Name                      string                           `json:"name"`
	HostPath                  string                           `json:"host_path"`
	DockerVolumeConfiguration []AwsEcsTaskDefinitionSpecVolume `json:"docker_volume_configuration"`
}

type AwsEcsTaskDefinitionSpec struct {
	PidMode                 string                     `json:"pid_mode"`
	Arn                     string                     `json:"arn"`
	TaskRoleArn             string                     `json:"task_role_arn"`
	Revision                int                        `json:"revision"`
	ContainerDefinitions    string                     `json:"container_definitions"`
	ExecutionRoleArn        string                     `json:"execution_role_arn"`
	PlacementConstraints    []AwsEcsTaskDefinitionSpec `json:"placement_constraints"`
	RequiresCompatibilities []string                   `json:"requires_compatibilities"`
	IpcMode                 string                     `json:"ipc_mode"`
	ProxyConfiguration      []AwsEcsTaskDefinitionSpec `json:"proxy_configuration"`
	Tags                    map[string]string          `json:"tags"`
	Cpu                     string                     `json:"cpu"`
	Memory                  string                     `json:"memory"`
	NetworkMode             string                     `json:"network_mode"`
	Volume                  []AwsEcsTaskDefinitionSpec `json:"volume"`
	Family                  string                     `json:"family"`
}



type AwsEcsTaskDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEcsTaskDefinitionList is a list of AwsEcsTaskDefinitions
type AwsEcsTaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcsTaskDefinition CRD objects
	Items []AwsEcsTaskDefinition `json:"items,omitempty"`
}