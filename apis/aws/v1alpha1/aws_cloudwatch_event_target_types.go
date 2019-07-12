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

type AwsCloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchEventTargetSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchEventTargetStatus `json:"status,omitempty"`
}

type AwsCloudwatchEventTargetSpecInputTransformer struct {
	InputTemplate string            `json:"input_template"`
	InputPaths    map[string]string `json:"input_paths"`
}

type AwsCloudwatchEventTargetSpecRunCommandTargets struct {
	Values []string `json:"values"`
	Key    string   `json:"key"`
}

type AwsCloudwatchEventTargetSpecEcsTargetNetworkConfiguration struct {
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
	AssignPublicIp bool     `json:"assign_public_ip"`
}

type AwsCloudwatchEventTargetSpecEcsTarget struct {
	Group                string                                  `json:"group"`
	LaunchType           string                                  `json:"launch_type"`
	NetworkConfiguration []AwsCloudwatchEventTargetSpecEcsTarget `json:"network_configuration"`
	PlatformVersion      string                                  `json:"platform_version"`
	TaskCount            int                                     `json:"task_count"`
	TaskDefinitionArn    string                                  `json:"task_definition_arn"`
}

type AwsCloudwatchEventTargetSpecBatchTarget struct {
	JobDefinition string `json:"job_definition"`
	JobName       string `json:"job_name"`
	ArraySize     int    `json:"array_size"`
	JobAttempts   int    `json:"job_attempts"`
}

type AwsCloudwatchEventTargetSpecKinesisTarget struct {
	PartitionKeyPath string `json:"partition_key_path"`
}

type AwsCloudwatchEventTargetSpecSqsTarget struct {
	MessageGroupId string `json:"message_group_id"`
}

type AwsCloudwatchEventTargetSpec struct {
	InputTransformer  []AwsCloudwatchEventTargetSpec `json:"input_transformer"`
	TargetId          string                         `json:"target_id"`
	InputPath         string                         `json:"input_path"`
	RunCommandTargets []AwsCloudwatchEventTargetSpec `json:"run_command_targets"`
	EcsTarget         []AwsCloudwatchEventTargetSpec `json:"ecs_target"`
	BatchTarget       []AwsCloudwatchEventTargetSpec `json:"batch_target"`
	KinesisTarget     []AwsCloudwatchEventTargetSpec `json:"kinesis_target"`
	Rule              string                         `json:"rule"`
	Arn               string                         `json:"arn"`
	Input             string                         `json:"input"`
	RoleArn           string                         `json:"role_arn"`
	SqsTarget         []AwsCloudwatchEventTargetSpec `json:"sqs_target"`
}

type AwsCloudwatchEventTargetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchEventTargetList is a list of AwsCloudwatchEventTargets
type AwsCloudwatchEventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchEventTarget CRD objects
	Items []AwsCloudwatchEventTarget `json:"items,omitempty"`
}
