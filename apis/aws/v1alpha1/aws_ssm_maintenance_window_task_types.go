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

type AwsSsmMaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmMaintenanceWindowTaskSpec   `json:"spec,omitempty"`
	Status            AwsSsmMaintenanceWindowTaskStatus `json:"status,omitempty"`
}

type AwsSsmMaintenanceWindowTaskSpecTaskParameters struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTaskSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTaskSpecLoggingInfo struct {
	S3BucketName   string `json:"s3_bucket_name"`
	S3Region       string `json:"s3_region"`
	S3BucketPrefix string `json:"s3_bucket_prefix"`
}

type AwsSsmMaintenanceWindowTaskSpec struct {
	Priority       int                               `json:"priority"`
	TaskParameters []AwsSsmMaintenanceWindowTaskSpec `json:"task_parameters"`
	WindowId       string                            `json:"window_id"`
	MaxConcurrency string                            `json:"max_concurrency"`
	ServiceRoleArn string                            `json:"service_role_arn"`
	Name           string                            `json:"name"`
	Description    string                            `json:"description"`
	MaxErrors      string                            `json:"max_errors"`
	TaskType       string                            `json:"task_type"`
	TaskArn        string                            `json:"task_arn"`
	Targets        []AwsSsmMaintenanceWindowTaskSpec `json:"targets"`
	LoggingInfo    []AwsSsmMaintenanceWindowTaskSpec `json:"logging_info"`
}

type AwsSsmMaintenanceWindowTaskStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTasks
type AwsSsmMaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmMaintenanceWindowTask CRD objects
	Items []AwsSsmMaintenanceWindowTask `json:"items,omitempty"`
}
