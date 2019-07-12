package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	IncludeGlobalServiceEvents bool                `json:"include_global_service_events"`
	SnsTopicName               string              `json:"sns_topic_name"`
	KmsKeyId                   string              `json:"kms_key_id"`
	Arn                        string              `json:"arn"`
	Name                       string              `json:"name"`
	EnableLogging              bool                `json:"enable_logging"`
	IsMultiRegionTrail         bool                `json:"is_multi_region_trail"`
	Tags                       map[string]string   `json:"tags"`
	S3BucketName               string              `json:"s3_bucket_name"`
	CloudWatchLogsGroupArn     string              `json:"cloud_watch_logs_group_arn"`
	IsOrganizationTrail        bool                `json:"is_organization_trail"`
	EnableLogFileValidation    bool                `json:"enable_log_file_validation"`
	EventSelector              []AwsCloudtrailSpec `json:"event_selector"`
	S3KeyPrefix                string              `json:"s3_key_prefix"`
	CloudWatchLogsRoleArn      string              `json:"cloud_watch_logs_role_arn"`
	HomeRegion                 string              `json:"home_region"`
}



type AwsCloudtrailStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrailList is a list of AwsCloudtrails
type AwsCloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudtrail CRD objects
	Items []AwsCloudtrail `json:"items,omitempty"`
}