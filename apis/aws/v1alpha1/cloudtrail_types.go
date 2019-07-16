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

type Cloudtrail struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudtrailSpec   `json:"spec,omitempty"`
	Status            CloudtrailStatus `json:"status,omitempty"`
}

type CloudtrailSpecEventSelectorDataResource struct {
	Type string `json:"type"`
	// +kubebuilder:validation:MaxItems=250
	Values []string `json:"values"`
}

type CloudtrailSpecEventSelector struct {
	// +optional
	DataResource *[]CloudtrailSpecEventSelector `json:"data_resource,omitempty"`
	// +optional
	IncludeManagementEvents bool `json:"include_management_events,omitempty"`
	// +optional
	ReadWriteType string `json:"read_write_type,omitempty"`
}

type CloudtrailSpec struct {
	// +optional
	CloudWatchLogsGroupArn string `json:"cloud_watch_logs_group_arn,omitempty"`
	// +optional
	CloudWatchLogsRoleArn string `json:"cloud_watch_logs_role_arn,omitempty"`
	// +optional
	EnableLogFileValidation bool `json:"enable_log_file_validation,omitempty"`
	// +optional
	EnableLogging bool `json:"enable_logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	EventSelector *[]CloudtrailSpec `json:"event_selector,omitempty"`
	// +optional
	IncludeGlobalServiceEvents bool `json:"include_global_service_events,omitempty"`
	// +optional
	IsMultiRegionTrail bool `json:"is_multi_region_trail,omitempty"`
	// +optional
	IsOrganizationTrail bool `json:"is_organization_trail,omitempty"`
	// +optional
	KmsKeyId     string `json:"kms_key_id,omitempty"`
	Name         string `json:"name"`
	S3BucketName string `json:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3_key_prefix,omitempty"`
	// +optional
	SnsTopicName string `json:"sns_topic_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type CloudtrailStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
