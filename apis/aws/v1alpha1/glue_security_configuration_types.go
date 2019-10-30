/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration []GlueSecurityConfigurationSpecEncryptionConfiguration `json:"encryptionConfiguration" tf:"encryption_configuration"`
	Name                    string                                                 `json:"name" tf:"name"`
}

type GlueSecurityConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueSecurityConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
