package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ElastictranscoderPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElastictranscoderPipelineSpec   `json:"spec,omitempty"`
	Status            ElastictranscoderPipelineStatus `json:"status,omitempty"`
}

type ElastictranscoderPipelineSpecContentConfig struct {
	// +optional
	Bucket string `json:"bucket,omitempty" tf:"bucket,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ElastictranscoderPipelineSpecContentConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty" tf:"access,omitempty"`
	// +optional
	Grantee string `json:"grantee,omitempty" tf:"grantee,omitempty"`
	// +optional
	GranteeType string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ElastictranscoderPipelineSpecNotifications struct {
	// +optional
	Completed string `json:"completed,omitempty" tf:"completed,omitempty"`
	// +optional
	Error string `json:"error,omitempty" tf:"error,omitempty"`
	// +optional
	Progressing string `json:"progressing,omitempty" tf:"progressing,omitempty"`
	// +optional
	Warning string `json:"warning,omitempty" tf:"warning,omitempty"`
}

type ElastictranscoderPipelineSpecThumbnailConfig struct {
	// +optional
	Bucket string `json:"bucket,omitempty" tf:"bucket,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ElastictranscoderPipelineSpecThumbnailConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty" tf:"access,omitempty"`
	// +optional
	Grantee string `json:"grantee,omitempty" tf:"grantee,omitempty"`
	// +optional
	GranteeType string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ElastictranscoderPipelineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AwsKmsKeyArn string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ContentConfig []ElastictranscoderPipelineSpecContentConfig `json:"contentConfig,omitempty" tf:"content_config,omitempty"`
	// +optional
	ContentConfigPermissions []ElastictranscoderPipelineSpecContentConfigPermissions `json:"contentConfigPermissions,omitempty" tf:"content_config_permissions,omitempty"`
	InputBucket              string                                                  `json:"inputBucket" tf:"input_bucket"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Notifications []ElastictranscoderPipelineSpecNotifications `json:"notifications,omitempty" tf:"notifications,omitempty"`
	// +optional
	OutputBucket string `json:"outputBucket,omitempty" tf:"output_bucket,omitempty"`
	Role         string `json:"role" tf:"role"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ThumbnailConfig []ElastictranscoderPipelineSpecThumbnailConfig `json:"thumbnailConfig,omitempty" tf:"thumbnail_config,omitempty"`
	// +optional
	ThumbnailConfigPermissions []ElastictranscoderPipelineSpecThumbnailConfigPermissions `json:"thumbnailConfigPermissions,omitempty" tf:"thumbnail_config_permissions,omitempty"`
}

type ElastictranscoderPipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElastictranscoderPipelineSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElastictranscoderPipelineList is a list of ElastictranscoderPipelines
type ElastictranscoderPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElastictranscoderPipeline CRD objects
	Items []ElastictranscoderPipeline `json:"items,omitempty"`
}
