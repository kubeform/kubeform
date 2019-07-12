package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElastictranscoderPipelineSpec   `json:"spec,omitempty"`
	Status            AwsElastictranscoderPipelineStatus `json:"status,omitempty"`
}

type AwsElastictranscoderPipelineSpecContentConfig struct {
	StorageClass string `json:"storage_class"`
	Bucket       string `json:"bucket"`
}

type AwsElastictranscoderPipelineSpecContentConfigPermissions struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type AwsElastictranscoderPipelineSpecNotifications struct {
	Completed   string `json:"completed"`
	Error       string `json:"error"`
	Progressing string `json:"progressing"`
	Warning     string `json:"warning"`
}

type AwsElastictranscoderPipelineSpecThumbnailConfigPermissions struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type AwsElastictranscoderPipelineSpecThumbnailConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type AwsElastictranscoderPipelineSpec struct {
	Arn                        string                             `json:"arn"`
	ContentConfig              []AwsElastictranscoderPipelineSpec `json:"content_config"`
	ContentConfigPermissions   []AwsElastictranscoderPipelineSpec `json:"content_config_permissions"`
	InputBucket                string                             `json:"input_bucket"`
	Notifications              []AwsElastictranscoderPipelineSpec `json:"notifications"`
	ThumbnailConfigPermissions []AwsElastictranscoderPipelineSpec `json:"thumbnail_config_permissions"`
	AwsKmsKeyArn               string                             `json:"aws_kms_key_arn"`
	Name                       string                             `json:"name"`
	OutputBucket               string                             `json:"output_bucket"`
	Role                       string                             `json:"role"`
	ThumbnailConfig            []AwsElastictranscoderPipelineSpec `json:"thumbnail_config"`
}

type AwsElastictranscoderPipelineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipelines
type AwsElastictranscoderPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElastictranscoderPipeline CRD objects
	Items []AwsElastictranscoderPipeline `json:"items,omitempty"`
}
