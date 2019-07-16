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

type GlueSecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueSecurityConfigurationSpec   `json:"spec,omitempty"`
	Status            GlueSecurityConfigurationStatus `json:"status,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption struct {
	// +optional
	CloudwatchEncryptionMode string `json:"cloudwatch_encryption_mode,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption struct {
	// +optional
	JobBookmarksEncryptionMode string `json:"job_bookmarks_encryption_mode,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption struct {
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	S3EncryptionMode string `json:"s3_encryption_mode,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	CloudwatchEncryption []GlueSecurityConfigurationSpecEncryptionConfiguration `json:"cloudwatch_encryption"`
	// +kubebuilder:validation:MaxItems=1
	JobBookmarksEncryption []GlueSecurityConfigurationSpecEncryptionConfiguration `json:"job_bookmarks_encryption"`
	// +kubebuilder:validation:MaxItems=1
	S3Encryption []GlueSecurityConfigurationSpecEncryptionConfiguration `json:"s3_encryption"`
}

type GlueSecurityConfigurationSpec struct {
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration []GlueSecurityConfigurationSpec `json:"encryption_configuration"`
	Name                    string                          `json:"name"`
}

type GlueSecurityConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueSecurityConfigurationList is a list of GlueSecurityConfigurations
type GlueSecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueSecurityConfiguration CRD objects
	Items []GlueSecurityConfiguration `json:"items,omitempty"`
}
