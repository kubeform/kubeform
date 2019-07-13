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
	Policy                               string            `json:"policy"`
	ApplicationFailureFeedbackRoleArn    string            `json:"application_failure_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn           string            `json:"http_success_feedback_role_arn"`
	HttpFailureFeedbackRoleArn           string            `json:"http_failure_feedback_role_arn"`
	LambdaSuccessFeedbackRoleArn         string            `json:"lambda_success_feedback_role_arn"`
	SqsSuccessFeedbackSampleRate         int               `json:"sqs_success_feedback_sample_rate"`
	Name                                 string            `json:"name"`
	DeliveryPolicy                       string            `json:"delivery_policy"`
	ApplicationSuccessFeedbackRoleArn    string            `json:"application_success_feedback_role_arn"`
	LambdaSuccessFeedbackSampleRate      int               `json:"lambda_success_feedback_sample_rate"`
	LambdaFailureFeedbackRoleArn         string            `json:"lambda_failure_feedback_role_arn"`
	ApplicationSuccessFeedbackSampleRate int               `json:"application_success_feedback_sample_rate"`
	HttpSuccessFeedbackSampleRate        int               `json:"http_success_feedback_sample_rate"`
	KmsMasterKeyId                       string            `json:"kms_master_key_id"`
	Arn                                  string            `json:"arn"`
	NamePrefix                           string            `json:"name_prefix"`
	DisplayName                          string            `json:"display_name"`
	SqsSuccessFeedbackRoleArn            string            `json:"sqs_success_feedback_role_arn"`
	SqsFailureFeedbackRoleArn            string            `json:"sqs_failure_feedback_role_arn"`
	Tags                                 map[string]string `json:"tags"`
}



type AwsSnsTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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