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

type AwsSnsTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsTopicSpec   `json:"spec,omitempty"`
	Status            AwsSnsTopicStatus `json:"status,omitempty"`
}

type AwsSnsTopicSpec struct {
	SqsSuccessFeedbackRoleArn            string            `json:"sqs_success_feedback_role_arn"`
	Arn                                  string            `json:"arn"`
	DeliveryPolicy                       string            `json:"delivery_policy"`
	HttpSuccessFeedbackSampleRate        int               `json:"http_success_feedback_sample_rate"`
	LambdaSuccessFeedbackRoleArn         string            `json:"lambda_success_feedback_role_arn"`
	HttpFailureFeedbackRoleArn           string            `json:"http_failure_feedback_role_arn"`
	LambdaSuccessFeedbackSampleRate      int               `json:"lambda_success_feedback_sample_rate"`
	LambdaFailureFeedbackRoleArn         string            `json:"lambda_failure_feedback_role_arn"`
	SqsSuccessFeedbackSampleRate         int               `json:"sqs_success_feedback_sample_rate"`
	SqsFailureFeedbackRoleArn            string            `json:"sqs_failure_feedback_role_arn"`
	NamePrefix                           string            `json:"name_prefix"`
	ApplicationSuccessFeedbackSampleRate int               `json:"application_success_feedback_sample_rate"`
	ApplicationFailureFeedbackRoleArn    string            `json:"application_failure_feedback_role_arn"`
	Tags                                 map[string]string `json:"tags"`
	Name                                 string            `json:"name"`
	DisplayName                          string            `json:"display_name"`
	Policy                               string            `json:"policy"`
	ApplicationSuccessFeedbackRoleArn    string            `json:"application_success_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn           string            `json:"http_success_feedback_role_arn"`
	KmsMasterKeyId                       string            `json:"kms_master_key_id"`
}

type AwsSnsTopicStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSnsTopicList is a list of AwsSnsTopics
type AwsSnsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsTopic CRD objects
	Items []AwsSnsTopic `json:"items,omitempty"`
}
