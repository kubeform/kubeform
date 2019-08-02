package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`
	// +optional
	S3BucketPrefix string `json:"s3BucketPrefix,omitempty" tf:"s3_bucket_prefix,omitempty"`
	S3Region       string `json:"s3Region" tf:"s3_region"`
}

type SsmMaintenanceWindowTaskSpecTargets struct {
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type SsmMaintenanceWindowTaskSpecTaskParameters struct {
	Name   string   `json:"name" tf:"name"`
	Values []string `json:"values" tf:"values"`
}

type SsmMaintenanceWindowTaskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingInfo    []SsmMaintenanceWindowTaskSpecLoggingInfo `json:"loggingInfo,omitempty" tf:"logging_info,omitempty"`
	MaxConcurrency string                                    `json:"maxConcurrency" tf:"max_concurrency"`
	MaxErrors      string                                    `json:"maxErrors" tf:"max_errors"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Priority       int                                   `json:"priority,omitempty" tf:"priority,omitempty"`
	ServiceRoleArn string                                `json:"serviceRoleArn" tf:"service_role_arn"`
	Targets        []SsmMaintenanceWindowTaskSpecTargets `json:"targets" tf:"targets"`
	TaskArn        string                                `json:"taskArn" tf:"task_arn"`
	// +optional
	TaskParameters []SsmMaintenanceWindowTaskSpecTaskParameters `json:"taskParameters,omitempty" tf:"task_parameters,omitempty"`
	TaskType       string                                       `json:"taskType" tf:"task_type"`
	WindowID       string                                       `json:"windowID" tf:"window_id"`
}

type SsmMaintenanceWindowTaskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SsmMaintenanceWindowTaskSpec `json:"state,omitempty"`
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
