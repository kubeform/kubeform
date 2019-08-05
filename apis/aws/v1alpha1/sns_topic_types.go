package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ApplicationFailureFeedbackRoleArn string `json:"applicationFailureFeedbackRoleArn,omitempty" tf:"application_failure_feedback_role_arn,omitempty"`
	// +optional
	ApplicationSuccessFeedbackRoleArn string `json:"applicationSuccessFeedbackRoleArn,omitempty" tf:"application_success_feedback_role_arn,omitempty"`
	// +optional
	ApplicationSuccessFeedbackSampleRate int `json:"applicationSuccessFeedbackSampleRate,omitempty" tf:"application_success_feedback_sample_rate,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	DeliveryPolicy string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	HttpFailureFeedbackRoleArn string `json:"httpFailureFeedbackRoleArn,omitempty" tf:"http_failure_feedback_role_arn,omitempty"`
	// +optional
	HttpSuccessFeedbackRoleArn string `json:"httpSuccessFeedbackRoleArn,omitempty" tf:"http_success_feedback_role_arn,omitempty"`
	// +optional
	HttpSuccessFeedbackSampleRate int `json:"httpSuccessFeedbackSampleRate,omitempty" tf:"http_success_feedback_sample_rate,omitempty"`
	// +optional
	KmsMasterKeyID string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id,omitempty"`
	// +optional
	LambdaFailureFeedbackRoleArn string `json:"lambdaFailureFeedbackRoleArn,omitempty" tf:"lambda_failure_feedback_role_arn,omitempty"`
	// +optional
	LambdaSuccessFeedbackRoleArn string `json:"lambdaSuccessFeedbackRoleArn,omitempty" tf:"lambda_success_feedback_role_arn,omitempty"`
	// +optional
	LambdaSuccessFeedbackSampleRate int `json:"lambdaSuccessFeedbackSampleRate,omitempty" tf:"lambda_success_feedback_sample_rate,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	SqsFailureFeedbackRoleArn string `json:"sqsFailureFeedbackRoleArn,omitempty" tf:"sqs_failure_feedback_role_arn,omitempty"`
	// +optional
	SqsSuccessFeedbackRoleArn string `json:"sqsSuccessFeedbackRoleArn,omitempty" tf:"sqs_success_feedback_role_arn,omitempty"`
	// +optional
	SqsSuccessFeedbackSampleRate int `json:"sqsSuccessFeedbackSampleRate,omitempty" tf:"sqs_success_feedback_sample_rate,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnsTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnsTopicSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
