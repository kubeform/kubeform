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

type LambdaFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaFunctionSpec   `json:"spec,omitempty"`
	Status            LambdaFunctionStatus `json:"status,omitempty"`
}

type LambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

type LambdaFunctionSpecEnvironment struct {
	// +optional
	Variables map[string]string `json:"variables,omitempty"`
}

type LambdaFunctionSpecVpcConfig struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
}

type LambdaFunctionSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeadLetterConfig *[]LambdaFunctionSpec `json:"dead_letter_config,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Environment *[]LambdaFunctionSpec `json:"environment,omitempty"`
	// +optional
	Filename     string `json:"filename,omitempty"`
	FunctionName string `json:"function_name"`
	Handler      string `json:"handler"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Layers []string `json:"layers,omitempty"`
	// +optional
	MemorySize int `json:"memory_size,omitempty"`
	// +optional
	Publish bool `json:"publish,omitempty"`
	// +optional
	ReservedConcurrentExecutions int    `json:"reserved_concurrent_executions,omitempty"`
	Role                         string `json:"role"`
	Runtime                      string `json:"runtime"`
	// +optional
	S3Bucket string `json:"s3_bucket,omitempty"`
	// +optional
	S3Key string `json:"s3_key,omitempty"`
	// +optional
	S3ObjectVersion string `json:"s3_object_version,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig *[]LambdaFunctionSpec `json:"vpc_config,omitempty"`
}

type LambdaFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
