package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Cloudtrail struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudtrailSpec   `json:"spec,omitempty"`
	Status            CloudtrailStatus `json:"status,omitempty"`
}

type CloudtrailSpecEventSelectorDataResource struct {
	Type string `json:"type" tf:"type"`
	// +kubebuilder:validation:MaxItems=250
	Values []string `json:"values" tf:"values"`
}

type CloudtrailSpecEventSelector struct {
	// +optional
	DataResource []CloudtrailSpecEventSelectorDataResource `json:"dataResource,omitempty" tf:"data_resource,omitempty"`
	// +optional
	IncludeManagementEvents bool `json:"includeManagementEvents,omitempty" tf:"include_management_events,omitempty"`
	// +optional
	ReadWriteType string `json:"readWriteType,omitempty" tf:"read_write_type,omitempty"`
}

type CloudtrailSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CloudWatchLogsGroupArn string `json:"cloudWatchLogsGroupArn,omitempty" tf:"cloud_watch_logs_group_arn,omitempty"`
	// +optional
	CloudWatchLogsRoleArn string `json:"cloudWatchLogsRoleArn,omitempty" tf:"cloud_watch_logs_role_arn,omitempty"`
	// +optional
	EnableLogFileValidation bool `json:"enableLogFileValidation,omitempty" tf:"enable_log_file_validation,omitempty"`
	// +optional
	EnableLogging bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	EventSelector []CloudtrailSpecEventSelector `json:"eventSelector,omitempty" tf:"event_selector,omitempty"`
	// +optional
	IncludeGlobalServiceEvents bool `json:"includeGlobalServiceEvents,omitempty" tf:"include_global_service_events,omitempty"`
	// +optional
	IsMultiRegionTrail bool `json:"isMultiRegionTrail,omitempty" tf:"is_multi_region_trail,omitempty"`
	// +optional
	IsOrganizationTrail bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail,omitempty"`
	// +optional
	KmsKeyID     string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	Name         string `json:"name" tf:"name"`
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
	// +optional
	SnsTopicName string `json:"snsTopicName,omitempty" tf:"sns_topic_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudtrailStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *CloudtrailSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudtrailList is a list of Cloudtrails
type CloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cloudtrail CRD objects
	Items []Cloudtrail `json:"items,omitempty"`
}
