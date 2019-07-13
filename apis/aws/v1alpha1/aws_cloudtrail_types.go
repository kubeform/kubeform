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

type AwsCloudtrail struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudtrailSpec   `json:"spec,omitempty"`
	Status            AwsCloudtrailStatus `json:"status,omitempty"`
}

type AwsCloudtrailSpecEventSelectorDataResource struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type AwsCloudtrailSpecEventSelector struct {
	ReadWriteType           string                           `json:"read_write_type"`
	IncludeManagementEvents bool                             `json:"include_management_events"`
	DataResource            []AwsCloudtrailSpecEventSelector `json:"data_resource"`
}

type AwsCloudtrailSpec struct {
	KmsKeyId                   string              `json:"kms_key_id"`
	EventSelector              []AwsCloudtrailSpec `json:"event_selector"`
	HomeRegion                 string              `json:"home_region"`
	Name                       string              `json:"name"`
	IncludeGlobalServiceEvents bool                `json:"include_global_service_events"`
	CloudWatchLogsGroupArn     string              `json:"cloud_watch_logs_group_arn"`
	IsMultiRegionTrail         bool                `json:"is_multi_region_trail"`
	IsOrganizationTrail        bool                `json:"is_organization_trail"`
	SnsTopicName               string              `json:"sns_topic_name"`
	S3BucketName               string              `json:"s3_bucket_name"`
	S3KeyPrefix                string              `json:"s3_key_prefix"`
	CloudWatchLogsRoleArn      string              `json:"cloud_watch_logs_role_arn"`
	EnableLogFileValidation    bool                `json:"enable_log_file_validation"`
	Arn                        string              `json:"arn"`
	Tags                       map[string]string   `json:"tags"`
	EnableLogging              bool                `json:"enable_logging"`
}



type AwsCloudtrailStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudtrailList is a list of AwsCloudtrails
type AwsCloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudtrail CRD objects
	Items []AwsCloudtrail `json:"items,omitempty"`
}