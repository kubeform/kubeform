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

type AwsSesReceiptRuleSpecLambdaAction struct {
	InvocationType string `json:"invocation_type"`
	TopicArn       string `json:"topic_arn"`
	Position       int    `json:"position"`
	FunctionArn    string `json:"function_arn"`
}

type AwsSesReceiptRuleSpecS3Action struct {
	BucketName      string `json:"bucket_name"`
	KmsKeyArn       string `json:"kms_key_arn"`
	ObjectKeyPrefix string `json:"object_key_prefix"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpecSnsAction struct {
	Position int    `json:"position"`
	TopicArn string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecAddHeaderAction struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Position    int    `json:"position"`
}

type AwsSesReceiptRuleSpecBounceAction struct {
	Position      int    `json:"position"`
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	StatusCode    string `json:"status_code"`
	TopicArn      string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecStopAction struct {
	Scope    string `json:"scope"`
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecWorkmailAction struct {
	OrganizationArn string `json:"organization_arn"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpec struct {
	LambdaAction    []AwsSesReceiptRuleSpec `json:"lambda_action"`
	S3Action        []AwsSesReceiptRuleSpec `json:"s3_action"`
	SnsAction       []AwsSesReceiptRuleSpec `json:"sns_action"`
	RuleSetName     string                  `json:"rule_set_name"`
	Enabled         bool                    `json:"enabled"`
	Recipients      []string                `json:"recipients"`
	TlsPolicy       string                  `json:"tls_policy"`
	Name            string                  `json:"name"`
	AddHeaderAction []AwsSesReceiptRuleSpec `json:"add_header_action"`
	BounceAction    []AwsSesReceiptRuleSpec `json:"bounce_action"`
	After           string                  `json:"after"`
	ScanEnabled     bool                    `json:"scan_enabled"`
	StopAction      []AwsSesReceiptRuleSpec `json:"stop_action"`
	WorkmailAction  []AwsSesReceiptRuleSpec `json:"workmail_action"`
}



type AwsSesReceiptRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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