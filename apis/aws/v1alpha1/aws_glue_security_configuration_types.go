package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueSecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueSecurityConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsGlueSecurityConfigurationStatus `json:"status,omitempty"`
}

type AwsGlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption struct {
	JobBookmarksEncryptionMode string `json:"job_bookmarks_encryption_mode"`
	KmsKeyArn                  string `json:"kms_key_arn"`
}

type AwsGlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption struct {
	KmsKeyArn        string `json:"kms_key_arn"`
	S3EncryptionMode string `json:"s3_encryption_mode"`
}

type AwsGlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption struct {
	CloudwatchEncryptionMode string `json:"cloudwatch_encryption_mode"`
	KmsKeyArn                string `json:"kms_key_arn"`
}

type AwsGlueSecurityConfigurationSpecEncryptionConfiguration struct {
	JobBookmarksEncryption []AwsGlueSecurityConfigurationSpecEncryptionConfiguration `json:"job_bookmarks_encryption"`
	S3Encryption           []AwsGlueSecurityConfigurationSpecEncryptionConfiguration `json:"s3_encryption"`
	CloudwatchEncryption   []AwsGlueSecurityConfigurationSpecEncryptionConfiguration `json:"cloudwatch_encryption"`
}

type AwsGlueSecurityConfigurationSpec struct {
	EncryptionConfiguration []AwsGlueSecurityConfigurationSpec `json:"encryption_configuration"`
	Name                    string                             `json:"name"`
}

type AwsGlueSecurityConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueSecurityConfigurationList is a list of AwsGlueSecurityConfigurations
type AwsGlueSecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueSecurityConfiguration CRD objects
	Items []AwsGlueSecurityConfiguration `json:"items,omitempty"`
}
