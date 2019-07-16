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

type ElastictranscoderPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElastictranscoderPipelineSpec   `json:"spec,omitempty"`
	Status            ElastictranscoderPipelineStatus `json:"status,omitempty"`
}

type ElastictranscoderPipelineSpecContentConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty"`
	// +optional
	Grantee string `json:"grantee,omitempty"`
	// +optional
	GranteeType string `json:"grantee_type,omitempty"`
}

type ElastictranscoderPipelineSpecNotifications struct {
	// +optional
	Completed string `json:"completed,omitempty"`
	// +optional
	Error string `json:"error,omitempty"`
	// +optional
	Progressing string `json:"progressing,omitempty"`
	// +optional
	Warning string `json:"warning,omitempty"`
}

type ElastictranscoderPipelineSpecThumbnailConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty"`
	// +optional
	Grantee string `json:"grantee,omitempty"`
	// +optional
	GranteeType string `json:"grantee_type,omitempty"`
}

type ElastictranscoderPipelineSpec struct {
	// +optional
	AwsKmsKeyArn string `json:"aws_kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ContentConfigPermissions *[]ElastictranscoderPipelineSpec `json:"content_config_permissions,omitempty"`
	InputBucket              string                           `json:"input_bucket"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Notifications *[]ElastictranscoderPipelineSpec `json:"notifications,omitempty"`
	Role          string                           `json:"role"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ThumbnailConfigPermissions *[]ElastictranscoderPipelineSpec `json:"thumbnail_config_permissions,omitempty"`
}

type ElastictranscoderPipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
