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

type AwsGlueJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueJobSpec   `json:"spec,omitempty"`
	Status            AwsGlueJobStatus `json:"status,omitempty"`
}

type AwsGlueJobSpecCommand struct {
	ScriptLocation string `json:"script_location"`
	Name           string `json:"name"`
}

type AwsGlueJobSpecExecutionProperty struct {
	MaxConcurrentRuns int `json:"max_concurrent_runs"`
}

type AwsGlueJobSpec struct {
	AllocatedCapacity     int               `json:"allocated_capacity"`
	Command               []AwsGlueJobSpec  `json:"command"`
	Connections           []string          `json:"connections"`
	Description           string            `json:"description"`
	ExecutionProperty     []AwsGlueJobSpec  `json:"execution_property"`
	MaxCapacity           float64           `json:"max_capacity"`
	Timeout               int               `json:"timeout"`
	SecurityConfiguration string            `json:"security_configuration"`
	DefaultArguments      map[string]string `json:"default_arguments"`
	MaxRetries            int               `json:"max_retries"`
	Name                  string            `json:"name"`
	RoleArn               string            `json:"role_arn"`
}

type AwsGlueJobStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueJobList is a list of AwsGlueJobs
type AwsGlueJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueJob CRD objects
	Items []AwsGlueJob `json:"items,omitempty"`
}
