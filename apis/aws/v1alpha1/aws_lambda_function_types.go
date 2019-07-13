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

type AwsLambdaFunctionSpecVpcConfig struct {
	SubnetIds        []string `json:"subnet_ids"`
	SecurityGroupIds []string `json:"security_group_ids"`
	VpcId            string   `json:"vpc_id"`
}

type AwsLambdaFunctionSpecEnvironment struct {
	Variables map[string]string `json:"variables"`
}

type AwsLambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode"`
}

type AwsLambdaFunctionSpec struct {
	Description                  string                  `json:"description"`
	FunctionName                 string                  `json:"function_name"`
	Role                         string                  `json:"role"`
	Version                      string                  `json:"version"`
	QualifiedArn                 string                  `json:"qualified_arn"`
	ReservedConcurrentExecutions int                     `json:"reserved_concurrent_executions"`
	Timeout                      int                     `json:"timeout"`
	KmsKeyArn                    string                  `json:"kms_key_arn"`
	Filename                     string                  `json:"filename"`
	DeadLetterConfig             []AwsLambdaFunctionSpec `json:"dead_letter_config"`
	MemorySize                   int                     `json:"memory_size"`
	Layers                       []string                `json:"layers"`
	S3Bucket                     string                  `json:"s3_bucket"`
	S3ObjectVersion              string                  `json:"s3_object_version"`
	Handler                      string                  `json:"handler"`
	Runtime                      string                  `json:"runtime"`
	LastModified                 string                  `json:"last_modified"`
	Tags                         map[string]string       `json:"tags"`
	SourceCodeHash               string                  `json:"source_code_hash"`
	S3Key                        string                  `json:"s3_key"`
	Publish                      bool                    `json:"publish"`
	Arn                          string                  `json:"arn"`
	VpcConfig                    []AwsLambdaFunctionSpec `json:"vpc_config"`
	InvokeArn                    string                  `json:"invoke_arn"`
	Environment                  []AwsLambdaFunctionSpec `json:"environment"`
	SourceCodeSize               int                     `json:"source_code_size"`
	TracingConfig                []AwsLambdaFunctionSpec `json:"tracing_config"`
}



type AwsLambdaFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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