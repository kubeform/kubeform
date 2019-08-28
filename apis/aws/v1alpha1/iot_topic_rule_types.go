package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	AlarmName   string `json:"alarmName" tf:"alarm_name"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
	StateReason string `json:"stateReason" tf:"state_reason"`
	StateValue  string `json:"stateValue" tf:"state_value"`
}

type IotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metricName" tf:"metric_name"`
	MetricNamespace string `json:"metricNamespace" tf:"metric_namespace"`
	// +optional
	MetricTimestamp string `json:"metricTimestamp,omitempty" tf:"metric_timestamp,omitempty"`
	MetricUnit      string `json:"metricUnit" tf:"metric_unit"`
	MetricValue     string `json:"metricValue" tf:"metric_value"`
	RoleArn         string `json:"roleArn" tf:"role_arn"`
}

type IotTopicRuleSpecDynamodb struct {
	HashKeyField string `json:"hashKeyField" tf:"hash_key_field"`
	// +optional
	HashKeyType  string `json:"hashKeyType,omitempty" tf:"hash_key_type,omitempty"`
	HashKeyValue string `json:"hashKeyValue" tf:"hash_key_value"`
	// +optional
	PayloadField string `json:"payloadField,omitempty" tf:"payload_field,omitempty"`
	// +optional
	RangeKeyField string `json:"rangeKeyField,omitempty" tf:"range_key_field,omitempty"`
	// +optional
	RangeKeyType string `json:"rangeKeyType,omitempty" tf:"range_key_type,omitempty"`
	// +optional
	RangeKeyValue string `json:"rangeKeyValue,omitempty" tf:"range_key_value,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	TableName     string `json:"tableName" tf:"table_name"`
}

type IotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`
	ID       string `json:"ID" tf:"id"`
	Index    string `json:"index" tf:"index"`
	RoleArn  string `json:"roleArn" tf:"role_arn"`
	Type     string `json:"type" tf:"type"`
}

type IotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"deliveryStreamName" tf:"delivery_stream_name"`
	RoleArn            string `json:"roleArn" tf:"role_arn"`
	// +optional
	Separator string `json:"separator,omitempty" tf:"separator,omitempty"`
}

type IotTopicRuleSpecKinesis struct {
	// +optional
	PartitionKey string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`
	RoleArn      string `json:"roleArn" tf:"role_arn"`
	StreamName   string `json:"streamName" tf:"stream_name"`
}

type IotTopicRuleSpecLambda struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

type IotTopicRuleSpecRepublish struct {
	RoleArn string `json:"roleArn" tf:"role_arn"`
	Topic   string `json:"topic" tf:"topic"`
}

type IotTopicRuleSpecS3 struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	Key        string `json:"key" tf:"key"`
	RoleArn    string `json:"roleArn" tf:"role_arn"`
}

type IotTopicRuleSpecSns struct {
	// +optional
	MessageFormat string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	TargetArn     string `json:"targetArn" tf:"target_arn"`
}

type IotTopicRuleSpecSqs struct {
	QueueURL  string `json:"queueURL" tf:"queue_url"`
	RoleArn   string `json:"roleArn" tf:"role_arn"`
	UseBase64 bool   `json:"useBase64" tf:"use_base64"`
}

type IotTopicRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchAlarm []IotTopicRuleSpecCloudwatchAlarm `json:"cloudwatchAlarm,omitempty" tf:"cloudwatch_alarm,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchMetric []IotTopicRuleSpecCloudwatchMetric `json:"cloudwatchMetric,omitempty" tf:"cloudwatch_metric,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Dynamodb []IotTopicRuleSpecDynamodb `json:"dynamodb,omitempty" tf:"dynamodb,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Elasticsearch []IotTopicRuleSpecElasticsearch `json:"elasticsearch,omitempty" tf:"elasticsearch,omitempty"`
	Enabled       bool                            `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Firehose []IotTopicRuleSpecFirehose `json:"firehose,omitempty" tf:"firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Kinesis []IotTopicRuleSpecKinesis `json:"kinesis,omitempty" tf:"kinesis,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Lambda []IotTopicRuleSpecLambda `json:"lambda,omitempty" tf:"lambda,omitempty"`
	Name   string                   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Republish []IotTopicRuleSpecRepublish `json:"republish,omitempty" tf:"republish,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	S3 []IotTopicRuleSpecS3 `json:"s3,omitempty" tf:"s3,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Sns        []IotTopicRuleSpecSns `json:"sns,omitempty" tf:"sns,omitempty"`
	Sql        string                `json:"sql" tf:"sql"`
	SqlVersion string                `json:"sqlVersion" tf:"sql_version"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Sqs []IotTopicRuleSpecSqs `json:"sqs,omitempty" tf:"sqs,omitempty"`
}



type IotTopicRuleStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *IotTopicRuleSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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