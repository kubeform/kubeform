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

type CloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventTargetSpec   `json:"spec,omitempty"`
	Status            CloudwatchEventTargetStatus `json:"status,omitempty"`
}

type CloudwatchEventTargetSpecBatchTarget struct {
	// +optional
	ArraySize int `json:"array_size,omitempty"`
	// +optional
	JobAttempts   int    `json:"job_attempts,omitempty"`
	JobDefinition string `json:"job_definition"`
	JobName       string `json:"job_name"`
}

type CloudwatchEventTargetSpecEcsTargetNetworkConfiguration struct {
	// +optional
	AssignPublicIp bool `json:"assign_public_ip,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets"`
}

type CloudwatchEventTargetSpecEcsTarget struct {
	// +optional
	Group string `json:"group,omitempty"`
	// +optional
	LaunchType string `json:"launch_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkConfiguration *[]CloudwatchEventTargetSpecEcsTarget `json:"network_configuration,omitempty"`
	// +optional
	PlatformVersion string `json:"platform_version,omitempty"`
	// +optional
	TaskCount         int    `json:"task_count,omitempty"`
	TaskDefinitionArn string `json:"task_definition_arn"`
}

type CloudwatchEventTargetSpecInputTransformer struct {
	// +optional
	InputPaths    map[string]string `json:"input_paths,omitempty"`
	InputTemplate string            `json:"input_template"`
}

type CloudwatchEventTargetSpecKinesisTarget struct {
	// +optional
	PartitionKeyPath string `json:"partition_key_path,omitempty"`
}

type CloudwatchEventTargetSpecRunCommandTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type CloudwatchEventTargetSpecSqsTarget struct {
	// +optional
	MessageGroupId string `json:"message_group_id,omitempty"`
}

type CloudwatchEventTargetSpec struct {
	Arn string `json:"arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BatchTarget *[]CloudwatchEventTargetSpec `json:"batch_target,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EcsTarget *[]CloudwatchEventTargetSpec `json:"ecs_target,omitempty"`
	// +optional
	Input string `json:"input,omitempty"`
	// +optional
	InputPath string `json:"input_path,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputTransformer *[]CloudwatchEventTargetSpec `json:"input_transformer,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisTarget *[]CloudwatchEventTargetSpec `json:"kinesis_target,omitempty"`
	// +optional
	RoleArn string `json:"role_arn,omitempty"`
	Rule    string `json:"rule"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	RunCommandTargets *[]CloudwatchEventTargetSpec `json:"run_command_targets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SqsTarget *[]CloudwatchEventTargetSpec `json:"sqs_target,omitempty"`
}

type CloudwatchEventTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchEventTargetList is a list of CloudwatchEventTargets
type CloudwatchEventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchEventTarget CRD objects
	Items []CloudwatchEventTarget `json:"items,omitempty"`
}
