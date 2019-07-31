package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LambdaFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaFunctionSpec   `json:"spec,omitempty"`
	Status            LambdaFunctionStatus `json:"status,omitempty"`
}

type LambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"targetArn" tf:"target_arn"`
}

type LambdaFunctionSpecEnvironment struct {
	// +optional
	Variables map[string]string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type LambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode" tf:"mode"`
}

type LambdaFunctionSpecVpcConfig struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
}

type LambdaFunctionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeadLetterConfig []LambdaFunctionSpecDeadLetterConfig `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Environment []LambdaFunctionSpecEnvironment `json:"environment,omitempty" tf:"environment,omitempty"`
	// +optional
	Filename     string `json:"filename,omitempty" tf:"filename,omitempty"`
	FunctionName string `json:"functionName" tf:"function_name"`
	Handler      string `json:"handler" tf:"handler"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Layers []string `json:"layers,omitempty" tf:"layers,omitempty"`
	// +optional
	MemorySize int `json:"memorySize,omitempty" tf:"memory_size,omitempty"`
	// +optional
	Publish bool `json:"publish,omitempty" tf:"publish,omitempty"`
	// +optional
	ReservedConcurrentExecutions int    `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions,omitempty"`
	Role                         string `json:"role" tf:"role"`
	Runtime                      string `json:"runtime" tf:"runtime"`
	// +optional
	S3Bucket string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`
	// +optional
	S3Key string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`
	// +optional
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`
	// +optional
	SourceCodeHash string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TracingConfig []LambdaFunctionSpecTracingConfig `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig []LambdaFunctionSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type LambdaFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaFunctionList is a list of LambdaFunctions
type LambdaFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaFunction CRD objects
	Items []LambdaFunction `json:"items,omitempty"`
}
