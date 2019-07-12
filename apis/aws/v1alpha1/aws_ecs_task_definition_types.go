package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcsTaskDefinitionSpec   `json:"spec,omitempty"`
	Status            AwsEcsTaskDefinitionStatus `json:"status,omitempty"`
}

type AwsEcsTaskDefinitionSpecPlacementConstraints struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
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

type AwsEcsTaskDefinitionSpec struct {
	Revision                int                        `json:"revision"`
	TaskRoleArn             string                     `json:"task_role_arn"`
	Memory                  string                     `json:"memory"`
	PlacementConstraints    []AwsEcsTaskDefinitionSpec `json:"placement_constraints"`
	RequiresCompatibilities []string                   `json:"requires_compatibilities"`
	IpcMode                 string                     `json:"ipc_mode"`
	Tags                    map[string]string          `json:"tags"`
	Arn                     string                     `json:"arn"`
	Cpu                     string                     `json:"cpu"`
	ContainerDefinitions    string                     `json:"container_definitions"`
	NetworkMode             string                     `json:"network_mode"`
	Volume                  []AwsEcsTaskDefinitionSpec `json:"volume"`
	PidMode                 string                     `json:"pid_mode"`
	Family                  string                     `json:"family"`
	ExecutionRoleArn        string                     `json:"execution_role_arn"`
}

type AwsEcsTaskDefinitionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsTaskDefinitionList is a list of AwsEcsTaskDefinitions
type AwsEcsTaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcsTaskDefinition CRD objects
	Items []AwsEcsTaskDefinition `json:"items,omitempty"`
}
