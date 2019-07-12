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

type AwsLambdaFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaFunctionSpec   `json:"spec,omitempty"`
	Status            AwsLambdaFunctionStatus `json:"status,omitempty"`
}

type AwsLambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

type AwsLambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode"`
}

type AwsLambdaFunctionSpecEnvironment struct {
	Variables map[string]string `json:"variables"`
}

type AwsLambdaFunctionSpecVpcConfig struct {
	SubnetIds        []string `json:"subnet_ids"`
	SecurityGroupIds []string `json:"security_group_ids"`
	VpcId            string   `json:"vpc_id"`
}

type AwsLambdaFunctionSpec struct {
	Filename                     string                  `json:"filename"`
	DeadLetterConfig             []AwsLambdaFunctionSpec `json:"dead_letter_config"`
	Runtime                      string                  `json:"runtime"`
	Publish                      bool                    `json:"publish"`
	TracingConfig                []AwsLambdaFunctionSpec `json:"tracing_config"`
	KmsKeyArn                    string                  `json:"kms_key_arn"`
	S3Bucket                     string                  `json:"s3_bucket"`
	Layers                       []string                `json:"layers"`
	ReservedConcurrentExecutions int                     `json:"reserved_concurrent_executions"`
	Timeout                      int                     `json:"timeout"`
	LastModified                 string                  `json:"last_modified"`
	SourceCodeSize               int                     `json:"source_code_size"`
	Handler                      string                  `json:"handler"`
	QualifiedArn                 string                  `json:"qualified_arn"`
	Tags                         map[string]string       `json:"tags"`
	S3ObjectVersion              string                  `json:"s3_object_version"`
	InvokeArn                    string                  `json:"invoke_arn"`
	Environment                  []AwsLambdaFunctionSpec `json:"environment"`
	Description                  string                  `json:"description"`
	VpcConfig                    []AwsLambdaFunctionSpec `json:"vpc_config"`
	Arn                          string                  `json:"arn"`
	MemorySize                   int                     `json:"memory_size"`
	Role                         string                  `json:"role"`
	Version                      string                  `json:"version"`
	S3Key                        string                  `json:"s3_key"`
	FunctionName                 string                  `json:"function_name"`
	SourceCodeHash               string                  `json:"source_code_hash"`
}

type AwsLambdaFunctionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLambdaFunctionList is a list of AwsLambdaFunctions
type AwsLambdaFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaFunction CRD objects
	Items []AwsLambdaFunction `json:"items,omitempty"`
}
