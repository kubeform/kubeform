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

type GlueJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueJobSpec   `json:"spec,omitempty"`
	Status            GlueJobStatus `json:"status,omitempty"`
}

type GlueJobSpecCommand struct {
	// +optional
	Name           string `json:"name,omitempty"`
	ScriptLocation string `json:"script_location"`
}

type GlueJobSpec struct {
	// +kubebuilder:validation:MaxItems=1
	Command []GlueJobSpec `json:"command"`
	// +optional
	Connections []string `json:"connections,omitempty"`
	// +optional
	DefaultArguments map[string]string `json:"default_arguments,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	MaxRetries int    `json:"max_retries,omitempty"`
	Name       string `json:"name"`
	RoleArn    string `json:"role_arn"`
	// +optional
	SecurityConfiguration string `json:"security_configuration,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty"`
}

type GlueJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
