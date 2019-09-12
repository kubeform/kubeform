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

type CloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventTargetSpec   `json:"spec,omitempty"`
	Status            CloudwatchEventTargetStatus `json:"status,omitempty"`
}

type CloudwatchEventTargetSpecBatchTarget struct {
	// +optional
	ArraySize int `json:"arraySize,omitempty" tf:"array_size,omitempty"`
	// +optional
	JobAttempts   int    `json:"jobAttempts,omitempty" tf:"job_attempts,omitempty"`
	JobDefinition string `json:"jobDefinition" tf:"job_definition"`
	JobName       string `json:"jobName" tf:"job_name"`
}

type CloudwatchEventTargetSpecEcsTargetNetworkConfiguration struct {
	// +optional
	AssignPublicIP bool `json:"assignPublicIP,omitempty" tf:"assign_public_ip,omitempty"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	Subnets        []string `json:"subnets" tf:"subnets"`
}

type CloudwatchEventTargetSpecEcsTarget struct {
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// +optional
	LaunchType string `json:"launchType,omitempty" tf:"launch_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkConfiguration []CloudwatchEventTargetSpecEcsTargetNetworkConfiguration `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`
	// +optional
	PlatformVersion string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`
	// +optional
	TaskCount         int    `json:"taskCount,omitempty" tf:"task_count,omitempty"`
	TaskDefinitionArn string `json:"taskDefinitionArn" tf:"task_definition_arn"`
}

type CloudwatchEventTargetSpecInputTransformer struct {
	// +optional
	InputPaths    map[string]string `json:"inputPaths,omitempty" tf:"input_paths,omitempty"`
	InputTemplate string            `json:"inputTemplate" tf:"input_template"`
}

type CloudwatchEventTargetSpecKinesisTarget struct {
	// +optional
	PartitionKeyPath string `json:"partitionKeyPath,omitempty" tf:"partition_key_path,omitempty"`
}

type CloudwatchEventTargetSpecRunCommandTargets struct {
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type CloudwatchEventTargetSpecSqsTarget struct {
	// +optional
	MessageGroupID string `json:"messageGroupID,omitempty" tf:"message_group_id,omitempty"`
}

type CloudwatchEventTargetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BatchTarget []CloudwatchEventTargetSpecBatchTarget `json:"batchTarget,omitempty" tf:"batch_target,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EcsTarget []CloudwatchEventTargetSpecEcsTarget `json:"ecsTarget,omitempty" tf:"ecs_target,omitempty"`
	// +optional
	Input string `json:"input,omitempty" tf:"input,omitempty"`
	// +optional
	InputPath string `json:"inputPath,omitempty" tf:"input_path,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputTransformer []CloudwatchEventTargetSpecInputTransformer `json:"inputTransformer,omitempty" tf:"input_transformer,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisTarget []CloudwatchEventTargetSpecKinesisTarget `json:"kinesisTarget,omitempty" tf:"kinesis_target,omitempty"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
	Rule    string `json:"rule" tf:"rule"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	RunCommandTargets []CloudwatchEventTargetSpecRunCommandTargets `json:"runCommandTargets,omitempty" tf:"run_command_targets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SqsTarget []CloudwatchEventTargetSpecSqsTarget `json:"sqsTarget,omitempty" tf:"sqs_target,omitempty"`
	// +optional
	TargetID string `json:"targetID,omitempty" tf:"target_id,omitempty"`
}

type CloudwatchEventTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudwatchEventTargetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
