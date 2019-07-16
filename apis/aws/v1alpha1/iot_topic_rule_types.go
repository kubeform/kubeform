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

type IotTopicRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotTopicRuleSpec   `json:"spec,omitempty"`
	Status            IotTopicRuleStatus `json:"status,omitempty"`
}

type IotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type IotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	// +optional
	MetricTimestamp string `json:"metric_timestamp,omitempty"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type IotTopicRuleSpecDynamodb struct {
	HashKeyField string `json:"hash_key_field"`
	// +optional
	HashKeyType  string `json:"hash_key_type,omitempty"`
	HashKeyValue string `json:"hash_key_value"`
	// +optional
	PayloadField string `json:"payload_field,omitempty"`
	// +optional
	RangeKeyField string `json:"range_key_field,omitempty"`
	// +optional
	RangeKeyType string `json:"range_key_type,omitempty"`
	// +optional
	RangeKeyValue string `json:"range_key_value,omitempty"`
	RoleArn       string `json:"role_arn"`
	TableName     string `json:"table_name"`
}

type IotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type IotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	// +optional
	Separator string `json:"separator,omitempty"`
}

type IotTopicRuleSpecKinesis struct {
	// +optional
	PartitionKey string `json:"partition_key,omitempty"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type IotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type IotTopicRuleSpecRepublish struct {
	RoleArn string `json:"role_arn"`
	Topic   string `json:"topic"`
}

type IotTopicRuleSpecS3 struct {
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type IotTopicRuleSpecSns struct {
	// +optional
	MessageFormat string `json:"message_format,omitempty"`
	RoleArn       string `json:"role_arn"`
	TargetArn     string `json:"target_arn"`
}

type IotTopicRuleSpecSqs struct {
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
}

type IotTopicRuleSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchAlarm *[]IotTopicRuleSpec `json:"cloudwatch_alarm,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchMetric *[]IotTopicRuleSpec `json:"cloudwatch_metric,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Dynamodb *[]IotTopicRuleSpec `json:"dynamodb,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Elasticsearch *[]IotTopicRuleSpec `json:"elasticsearch,omitempty"`
	Enabled       bool                `json:"enabled"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Firehose *[]IotTopicRuleSpec `json:"firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Kinesis *[]IotTopicRuleSpec `json:"kinesis,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Lambda *[]IotTopicRuleSpec `json:"lambda,omitempty"`
	Name   string              `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Republish *[]IotTopicRuleSpec `json:"republish,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	S3 *[]IotTopicRuleSpec `json:"s3,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Sns        *[]IotTopicRuleSpec `json:"sns,omitempty"`
	Sql        string              `json:"sql"`
	SqlVersion string              `json:"sql_version"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Sqs *[]IotTopicRuleSpec `json:"sqs,omitempty"`
}

type IotTopicRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotTopicRuleList is a list of IotTopicRules
type IotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotTopicRule CRD objects
	Items []IotTopicRule `json:"items,omitempty"`
}
