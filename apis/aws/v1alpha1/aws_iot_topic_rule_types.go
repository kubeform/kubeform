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

type AwsIotTopicRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotTopicRuleSpec   `json:"spec,omitempty"`
	Status            AwsIotTopicRuleStatus `json:"status,omitempty"`
}

type AwsIotTopicRuleSpecS3 struct {
	RoleArn    string `json:"role_arn"`
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
}

type AwsIotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	Separator          string `json:"separator"`
}

type AwsIotTopicRuleSpecKinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type AwsIotTopicRuleSpecSns struct {
	MessageFormat string `json:"message_format"`
	TargetArn     string `json:"target_arn"`
	RoleArn       string `json:"role_arn"`
}

type AwsIotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type AwsIotTopicRuleSpecCloudwatchMetric struct {
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
}

type AwsIotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type AwsIotTopicRuleSpecSqs struct {
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
}

type AwsIotTopicRuleSpecDynamodb struct {
	HashKeyType   string `json:"hash_key_type"`
	RangeKeyField string `json:"range_key_field"`
	HashKeyValue  string `json:"hash_key_value"`
	PayloadField  string `json:"payload_field"`
	RangeKeyValue string `json:"range_key_value"`
	RangeKeyType  string `json:"range_key_type"`
	RoleArn       string `json:"role_arn"`
	TableName     string `json:"table_name"`
	HashKeyField  string `json:"hash_key_field"`
}

type AwsIotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type AwsIotTopicRuleSpecRepublish struct {
	RoleArn string `json:"role_arn"`
	Topic   string `json:"topic"`
}

type AwsIotTopicRuleSpec struct {
	S3               []AwsIotTopicRuleSpec `json:"s3"`
	Sql              string                `json:"sql"`
	Firehose         []AwsIotTopicRuleSpec `json:"firehose"`
	Kinesis          []AwsIotTopicRuleSpec `json:"kinesis"`
	Sns              []AwsIotTopicRuleSpec `json:"sns"`
	CloudwatchAlarm  []AwsIotTopicRuleSpec `json:"cloudwatch_alarm"`
	CloudwatchMetric []AwsIotTopicRuleSpec `json:"cloudwatch_metric"`
	Elasticsearch    []AwsIotTopicRuleSpec `json:"elasticsearch"`
	SqlVersion       string                `json:"sql_version"`
	Sqs              []AwsIotTopicRuleSpec `json:"sqs"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Enabled          bool                  `json:"enabled"`
	Arn              string                `json:"arn"`
	Dynamodb         []AwsIotTopicRuleSpec `json:"dynamodb"`
	Lambda           []AwsIotTopicRuleSpec `json:"lambda"`
	Republish        []AwsIotTopicRuleSpec `json:"republish"`
}



type AwsIotTopicRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIotTopicRuleList is a list of AwsIotTopicRules
type AwsIotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotTopicRule CRD objects
	Items []AwsIotTopicRule `json:"items,omitempty"`
}