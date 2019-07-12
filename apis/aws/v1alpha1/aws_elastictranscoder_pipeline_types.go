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

type AwsElastictranscoderPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElastictranscoderPipelineSpec   `json:"spec,omitempty"`
	Status            AwsElastictranscoderPipelineStatus `json:"status,omitempty"`
}

type AwsElastictranscoderPipelineSpecThumbnailConfigPermissions struct {
	GranteeType string   `json:"grantee_type"`
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
}

type AwsElastictranscoderPipelineSpecContentConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
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

type AwsElastictranscoderPipelineSpecThumbnailConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type AwsElastictranscoderPipelineSpec struct {
	ThumbnailConfigPermissions []AwsElastictranscoderPipelineSpec `json:"thumbnail_config_permissions"`
	AwsKmsKeyArn               string                             `json:"aws_kms_key_arn"`
	ContentConfig              []AwsElastictranscoderPipelineSpec `json:"content_config"`
	InputBucket                string                             `json:"input_bucket"`
	Name                       string                             `json:"name"`
	OutputBucket               string                             `json:"output_bucket"`
	Arn                        string                             `json:"arn"`
	ContentConfigPermissions   []AwsElastictranscoderPipelineSpec `json:"content_config_permissions"`
	Notifications              []AwsElastictranscoderPipelineSpec `json:"notifications"`
	Role                       string                             `json:"role"`
	ThumbnailConfig            []AwsElastictranscoderPipelineSpec `json:"thumbnail_config"`
}

type AwsElastictranscoderPipelineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipelines
type AwsElastictranscoderPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElastictranscoderPipeline CRD objects
	Items []AwsElastictranscoderPipeline `json:"items,omitempty"`
}
