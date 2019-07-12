package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaFunctionSpec   `json:"spec,omitempty"`
	Status            AwsLambdaFunctionStatus `json:"status,omitempty"`
}

type AwsLambdaFunctionSpecEnvironment struct {
	Variables map[string]string `json:"variables"`
}

type AwsLambdaFunctionSpecVpcConfig struct {
	VpcId            string   `json:"vpc_id"`
	SubnetIds        []string `json:"subnet_ids"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

type AwsLambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode"`
}

type AwsLambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

type AwsLambdaFunctionSpec struct {
	Filename                     string                  `json:"filename"`
	S3ObjectVersion              string                  `json:"s3_object_version"`
	InvokeArn                    string                  `json:"invoke_arn"`
	Environment                  []AwsLambdaFunctionSpec `json:"environment"`
	S3Bucket                     string                  `json:"s3_bucket"`
	VpcConfig                    []AwsLambdaFunctionSpec `json:"vpc_config"`
	QualifiedArn                 string                  `json:"qualified_arn"`
	FunctionName                 string                  `json:"function_name"`
	Handler                      string                  `json:"handler"`
	Layers                       []string                `json:"layers"`
	Role                         string                  `json:"role"`
	Publish                      bool                    `json:"publish"`
	SourceCodeSize               int                     `json:"source_code_size"`
	KmsKeyArn                    string                  `json:"kms_key_arn"`
	Tags                         map[string]string       `json:"tags"`
	S3Key                        string                  `json:"s3_key"`
	MemorySize                   int                     `json:"memory_size"`
	Runtime                      string                  `json:"runtime"`
	Timeout                      int                     `json:"timeout"`
	TracingConfig                []AwsLambdaFunctionSpec `json:"tracing_config"`
	Description                  string                  `json:"description"`
	DeadLetterConfig             []AwsLambdaFunctionSpec `json:"dead_letter_config"`
	ReservedConcurrentExecutions int                     `json:"reserved_concurrent_executions"`
	Version                      string                  `json:"version"`
	SourceCodeHash               string                  `json:"source_code_hash"`
	Arn                          string                  `json:"arn"`
	LastModified                 string                  `json:"last_modified"`
}



type AwsLambdaFunctionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunctions
type AwsLambdaFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaFunction CRD objects
	Items []AwsLambdaFunction `json:"items,omitempty"`
}