package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	CloudwatchEncryptionMode string `json:"cloudwatchEncryptionMode,omitempty" tf:"cloudwatch_encryption_mode,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption struct {
	// +optional
	JobBookmarksEncryptionMode string `json:"jobBookmarksEncryptionMode,omitempty" tf:"job_bookmarks_encryption_mode,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption struct {
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	S3EncryptionMode string `json:"s3EncryptionMode,omitempty" tf:"s3_encryption_mode,omitempty"`
}

type GlueSecurityConfigurationSpecEncryptionConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	CloudwatchEncryption []GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption `json:"cloudwatchEncryption" tf:"cloudwatch_encryption"`
	// +kubebuilder:validation:MaxItems=1
	JobBookmarksEncryption []GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption `json:"jobBookmarksEncryption" tf:"job_bookmarks_encryption"`
	// +kubebuilder:validation:MaxItems=1
	S3Encryption []GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption `json:"s3Encryption" tf:"s3_encryption"`
}

type GlueSecurityConfigurationSpec struct {
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration []GlueSecurityConfigurationSpecEncryptionConfiguration `json:"encryptionConfiguration" tf:"encryption_configuration"`
	Name                    string                                                 `json:"name" tf:"name"`
	ProviderRef             core.LocalObjectReference                              `json:"providerRef" tf:"-"`
}

type GlueSecurityConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
