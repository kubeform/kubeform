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

type SnsTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicSpec   `json:"spec,omitempty"`
	Status            SnsTopicStatus `json:"status,omitempty"`
}

type SnsTopicSpec struct {
	// +optional
	ApplicationFailureFeedbackRoleArn string `json:"application_failure_feedback_role_arn,omitempty"`
	// +optional
	ApplicationSuccessFeedbackRoleArn string `json:"application_success_feedback_role_arn,omitempty"`
	// +optional
	ApplicationSuccessFeedbackSampleRate int `json:"application_success_feedback_sample_rate,omitempty"`
	// +optional
	DeliveryPolicy string `json:"delivery_policy,omitempty"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	HttpFailureFeedbackRoleArn string `json:"http_failure_feedback_role_arn,omitempty"`
	// +optional
	HttpSuccessFeedbackRoleArn string `json:"http_success_feedback_role_arn,omitempty"`
	// +optional
	HttpSuccessFeedbackSampleRate int `json:"http_success_feedback_sample_rate,omitempty"`
	// +optional
	KmsMasterKeyId string `json:"kms_master_key_id,omitempty"`
	// +optional
	LambdaFailureFeedbackRoleArn string `json:"lambda_failure_feedback_role_arn,omitempty"`
	// +optional
	LambdaSuccessFeedbackRoleArn string `json:"lambda_success_feedback_role_arn,omitempty"`
	// +optional
	LambdaSuccessFeedbackSampleRate int `json:"lambda_success_feedback_sample_rate,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	SqsFailureFeedbackRoleArn string `json:"sqs_failure_feedback_role_arn,omitempty"`
	// +optional
	SqsSuccessFeedbackRoleArn string `json:"sqs_success_feedback_role_arn,omitempty"`
	// +optional
	SqsSuccessFeedbackSampleRate int `json:"sqs_success_feedback_sample_rate,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SnsTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsTopicList is a list of SnsTopics
type SnsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsTopic CRD objects
	Items []SnsTopic `json:"items,omitempty"`
}
