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

type SesReceiptRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesReceiptRuleSpec   `json:"spec,omitempty"`
	Status            SesReceiptRuleStatus `json:"status,omitempty"`
}

type SesReceiptRuleSpecAddHeaderAction struct {
	HeaderName  string `json:"headerName" tf:"header_name"`
	HeaderValue string `json:"headerValue" tf:"header_value"`
	Position    int    `json:"position" tf:"position"`
}

type SesReceiptRuleSpecBounceAction struct {
	Message       string `json:"message" tf:"message"`
	Position      int    `json:"position" tf:"position"`
	Sender        string `json:"sender" tf:"sender"`
	SmtpReplyCode string `json:"smtpReplyCode" tf:"smtp_reply_code"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecLambdaAction struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`
	// +optional
	InvocationType string `json:"invocationType,omitempty" tf:"invocation_type,omitempty"`
	Position       int    `json:"position" tf:"position"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecS3Action struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	ObjectKeyPrefix string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix,omitempty"`
	Position        int    `json:"position" tf:"position"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecSnsAction struct {
	Position int    `json:"position" tf:"position"`
	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

type SesReceiptRuleSpecStopAction struct {
	Position int    `json:"position" tf:"position"`
	Scope    string `json:"scope" tf:"scope"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecWorkmailAction struct {
	OrganizationArn string `json:"organizationArn" tf:"organization_arn"`
	Position        int    `json:"position" tf:"position"`
	// +optional
	TopicArn string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SesReceiptRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AddHeaderAction []SesReceiptRuleSpecAddHeaderAction `json:"addHeaderAction,omitempty" tf:"add_header_action,omitempty"`
	// +optional
	After string `json:"after,omitempty" tf:"after,omitempty"`
	// +optional
	BounceAction []SesReceiptRuleSpecBounceAction `json:"bounceAction,omitempty" tf:"bounce_action,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LambdaAction []SesReceiptRuleSpecLambdaAction `json:"lambdaAction,omitempty" tf:"lambda_action,omitempty"`
	Name         string                           `json:"name" tf:"name"`
	// +optional
	Recipients  []string `json:"recipients,omitempty" tf:"recipients,omitempty"`
	RuleSetName string   `json:"ruleSetName" tf:"rule_set_name"`
	// +optional
	S3Action []SesReceiptRuleSpecS3Action `json:"s3Action,omitempty" tf:"s3_action,omitempty"`
	// +optional
	ScanEnabled bool `json:"scanEnabled,omitempty" tf:"scan_enabled,omitempty"`
	// +optional
	SnsAction []SesReceiptRuleSpecSnsAction `json:"snsAction,omitempty" tf:"sns_action,omitempty"`
	// +optional
	StopAction []SesReceiptRuleSpecStopAction `json:"stopAction,omitempty" tf:"stop_action,omitempty"`
	// +optional
	TlsPolicy string `json:"tlsPolicy,omitempty" tf:"tls_policy,omitempty"`
	// +optional
	WorkmailAction []SesReceiptRuleSpecWorkmailAction `json:"workmailAction,omitempty" tf:"workmail_action,omitempty"`
}

type SesReceiptRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SesReceiptRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesReceiptRuleList is a list of SesReceiptRules
type SesReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesReceiptRule CRD objects
	Items []SesReceiptRule `json:"items,omitempty"`
}
