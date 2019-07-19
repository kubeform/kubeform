package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GlueJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueJobSpec   `json:"spec,omitempty"`
	Status            GlueJobStatus `json:"status,omitempty"`
}

type GlueJobSpecCommand struct {
	// +optional
	Name           string `json:"name,omitempty" tf:"name,omitempty"`
	ScriptLocation string `json:"scriptLocation" tf:"script_location"`
}

type GlueJobSpecExecutionProperty struct {
	// +optional
	MaxConcurrentRuns int `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`
}

type GlueJobSpec struct {
	// +optional
	// Deprecated
	AllocatedCapacity int `json:"allocatedCapacity,omitempty" tf:"allocated_capacity,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Command []GlueJobSpecCommand `json:"command" tf:"command"`
	// +optional
	Connections []string `json:"connections,omitempty" tf:"connections,omitempty"`
	// +optional
	DefaultArguments map[string]string `json:"defaultArguments,omitempty" tf:"default_arguments,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExecutionProperty []GlueJobSpecExecutionProperty `json:"executionProperty,omitempty" tf:"execution_property,omitempty"`
	// +optional
	MaxCapacity json.Number `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	// +optional
	MaxRetries int    `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`
	Name       string `json:"name" tf:"name"`
	RoleArn    string `json:"roleArn" tf:"role_arn"`
	// +optional
	SecurityConfiguration string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`
	// +optional
	Timeout     int                       `json:"timeout,omitempty" tf:"timeout,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GlueJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueJobList is a list of GlueJobs
type GlueJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueJob CRD objects
	Items []GlueJob `json:"items,omitempty"`
}
