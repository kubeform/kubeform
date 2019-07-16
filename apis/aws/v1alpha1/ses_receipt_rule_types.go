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

type SesReceiptRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesReceiptRuleSpec   `json:"spec,omitempty"`
	Status            SesReceiptRuleStatus `json:"status,omitempty"`
}

type SesReceiptRuleSpecAddHeaderAction struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Position    int    `json:"position"`
}

type SesReceiptRuleSpecBounceAction struct {
	Message       string `json:"message"`
	Position      int    `json:"position"`
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	// +optional
	StatusCode string `json:"status_code,omitempty"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecLambdaAction struct {
	FunctionArn string `json:"function_arn"`
	Position    int    `json:"position"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecS3Action struct {
	BucketName string `json:"bucket_name"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	ObjectKeyPrefix string `json:"object_key_prefix,omitempty"`
	Position        int    `json:"position"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecSnsAction struct {
	Position int    `json:"position"`
	TopicArn string `json:"topic_arn"`
}

type SesReceiptRuleSpecStopAction struct {
	Position int    `json:"position"`
	Scope    string `json:"scope"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesReceiptRuleSpecWorkmailAction struct {
	OrganizationArn string `json:"organization_arn"`
	Position        int    `json:"position"`
	// +optional
	TopicArn string `json:"topic_arn,omitempty"`
}

type SesReceiptRuleSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AddHeaderAction *[]SesReceiptRuleSpec `json:"add_header_action,omitempty"`
	// +optional
	After string `json:"after,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BounceAction *[]SesReceiptRuleSpec `json:"bounce_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LambdaAction *[]SesReceiptRuleSpec `json:"lambda_action,omitempty"`
	Name         string                `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Recipients  []string `json:"recipients,omitempty"`
	RuleSetName string   `json:"rule_set_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	S3Action *[]SesReceiptRuleSpec `json:"s3_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnsAction *[]SesReceiptRuleSpec `json:"sns_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StopAction *[]SesReceiptRuleSpec `json:"stop_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WorkmailAction *[]SesReceiptRuleSpec `json:"workmail_action,omitempty"`
}

type SesReceiptRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
