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

type AwsEcsTaskDefinitionSpecVolumeDockerVolumeConfiguration struct {
	Scope         string            `json:"scope"`
	Autoprovision bool              `json:"autoprovision"`
	Driver        string            `json:"driver"`
	DriverOpts    map[string]string `json:"driver_opts"`
	Labels        map[string]string `json:"labels"`
}

type AwsEcsTaskDefinitionSpecVolume struct {
	Name                      string                           `json:"name"`
	HostPath                  string                           `json:"host_path"`
	DockerVolumeConfiguration []AwsEcsTaskDefinitionSpecVolume `json:"docker_volume_configuration"`
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

type AwsEcsTaskDefinitionSpec struct {
	Arn                     string                     `json:"arn"`
	Family                  string                     `json:"family"`
	Memory                  string                     `json:"memory"`
	Volume                  []AwsEcsTaskDefinitionSpec `json:"volume"`
	Revision                int                        `json:"revision"`
	ExecutionRoleArn        string                     `json:"execution_role_arn"`
	PlacementConstraints    []AwsEcsTaskDefinitionSpec `json:"placement_constraints"`
	IpcMode                 string                     `json:"ipc_mode"`
	Tags                    map[string]string          `json:"tags"`
	Cpu                     string                     `json:"cpu"`
	RequiresCompatibilities []string                   `json:"requires_compatibilities"`
	PidMode                 string                     `json:"pid_mode"`
	ContainerDefinitions    string                     `json:"container_definitions"`
	TaskRoleArn             string                     `json:"task_role_arn"`
	NetworkMode             string                     `json:"network_mode"`
	ProxyConfiguration      []AwsEcsTaskDefinitionSpec `json:"proxy_configuration"`
}

type AwsEcsTaskDefinitionStatus struct {
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
