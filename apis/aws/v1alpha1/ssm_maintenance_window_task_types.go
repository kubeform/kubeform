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

type SsmMaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowTaskSpec   `json:"spec,omitempty"`
	Status            SsmMaintenanceWindowTaskStatus `json:"status,omitempty"`
}

type SsmMaintenanceWindowTaskSpecLoggingInfo struct {
	S3BucketName string `json:"s3_bucket_name"`
	// +optional
	S3BucketPrefix string `json:"s3_bucket_prefix,omitempty"`
	S3Region       string `json:"s3_region"`
}

type SsmMaintenanceWindowTaskSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SsmMaintenanceWindowTaskSpecTaskParameters struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type SsmMaintenanceWindowTaskSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingInfo    *[]SsmMaintenanceWindowTaskSpec `json:"logging_info,omitempty"`
	MaxConcurrency string                          `json:"max_concurrency"`
	MaxErrors      string                          `json:"max_errors"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Priority       int                            `json:"priority,omitempty"`
	ServiceRoleArn string                         `json:"service_role_arn"`
	Targets        []SsmMaintenanceWindowTaskSpec `json:"targets"`
	TaskArn        string                         `json:"task_arn"`
	// +optional
	TaskParameters *[]SsmMaintenanceWindowTaskSpec `json:"task_parameters,omitempty"`
	TaskType       string                          `json:"task_type"`
	WindowId       string                          `json:"window_id"`
}

type SsmMaintenanceWindowTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmMaintenanceWindowTaskList is a list of SsmMaintenanceWindowTasks
type SsmMaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmMaintenanceWindowTask CRD objects
	Items []SsmMaintenanceWindowTask `json:"items,omitempty"`
}
