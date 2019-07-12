package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsTopicSpec   `json:"spec,omitempty"`
	Status            AwsSnsTopicStatus `json:"status,omitempty"`
}

type AwsSnsTopicSpec struct {
	SqsSuccessFeedbackRoleArn            string `json:"sqs_success_feedback_role_arn"`
	Arn                                  string `json:"arn"`
	DisplayName                          string `json:"display_name"`
	ApplicationSuccessFeedbackRoleArn    string `json:"application_success_feedback_role_arn"`
	ApplicationSuccessFeedbackSampleRate int    `json:"application_success_feedback_sample_rate"`
	LambdaSuccessFeedbackSampleRate      int    `json:"lambda_success_feedback_sample_rate"`
	SqsFailureFeedbackRoleArn            string `json:"sqs_failure_feedback_role_arn"`
	Policy                               string `json:"policy"`
	DeliveryPolicy                       string `json:"delivery_policy"`
	KmsMasterKeyId                       string `json:"kms_master_key_id"`
	LambdaSuccessFeedbackRoleArn         string `json:"lambda_success_feedback_role_arn"`
	NamePrefix                           string `json:"name_prefix"`
	HttpSuccessFeedbackSampleRate        int    `json:"http_success_feedback_sample_rate"`
	HttpFailureFeedbackRoleArn           string `json:"http_failure_feedback_role_arn"`
	LambdaFailureFeedbackRoleArn         string `json:"lambda_failure_feedback_role_arn"`
	SqsSuccessFeedbackSampleRate         int    `json:"sqs_success_feedback_sample_rate"`
	Name                                 string `json:"name"`
	ApplicationFailureFeedbackRoleArn    string `json:"application_failure_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn           string `json:"http_success_feedback_role_arn"`
}

type AwsSnsTopicStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicList is a list of AwsSnsTopics
type AwsSnsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsTopic CRD objects
	Items []AwsSnsTopic `json:"items,omitempty"`
}
