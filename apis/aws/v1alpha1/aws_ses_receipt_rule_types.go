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

type AwsSesReceiptRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesReceiptRuleSpec   `json:"spec,omitempty"`
	Status            AwsSesReceiptRuleStatus `json:"status,omitempty"`
}

type AwsSesReceiptRuleSpecStopAction struct {
	Scope    string `json:"scope"`
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecAddHeaderAction struct {
	HeaderValue string `json:"header_value"`
	Position    int    `json:"position"`
	HeaderName  string `json:"header_name"`
}

type AwsSesReceiptRuleSpecS3Action struct {
	BucketName      string `json:"bucket_name"`
	KmsKeyArn       string `json:"kms_key_arn"`
	ObjectKeyPrefix string `json:"object_key_prefix"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpecBounceAction struct {
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	StatusCode    string `json:"status_code"`
	TopicArn      string `json:"topic_arn"`
	Position      int    `json:"position"`
}

type AwsSesReceiptRuleSpecLambdaAction struct {
	InvocationType string `json:"invocation_type"`
	TopicArn       string `json:"topic_arn"`
	Position       int    `json:"position"`
	FunctionArn    string `json:"function_arn"`
}

type AwsSesReceiptRuleSpecSnsAction struct {
	Position int    `json:"position"`
	TopicArn string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecWorkmailAction struct {
	OrganizationArn string `json:"organization_arn"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpec struct {
	Name            string                  `json:"name"`
	ScanEnabled     bool                    `json:"scan_enabled"`
	StopAction      []AwsSesReceiptRuleSpec `json:"stop_action"`
	Recipients      []string                `json:"recipients"`
	AddHeaderAction []AwsSesReceiptRuleSpec `json:"add_header_action"`
	S3Action        []AwsSesReceiptRuleSpec `json:"s3_action"`
	RuleSetName     string                  `json:"rule_set_name"`
	After           string                  `json:"after"`
	Enabled         bool                    `json:"enabled"`
	BounceAction    []AwsSesReceiptRuleSpec `json:"bounce_action"`
	TlsPolicy       string                  `json:"tls_policy"`
	LambdaAction    []AwsSesReceiptRuleSpec `json:"lambda_action"`
	SnsAction       []AwsSesReceiptRuleSpec `json:"sns_action"`
	WorkmailAction  []AwsSesReceiptRuleSpec `json:"workmail_action"`
}

type AwsSesReceiptRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesReceiptRuleList is a list of AwsSesReceiptRules
type AwsSesReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesReceiptRule CRD objects
	Items []AwsSesReceiptRule `json:"items,omitempty"`
}
