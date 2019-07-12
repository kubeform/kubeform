package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotTopicRuleSpec   `json:"spec,omitempty"`
	Status            AwsIotTopicRuleStatus `json:"status,omitempty"`
}

type AwsIotTopicRuleSpecSns struct {
	MessageFormat string `json:"message_format"`
	TargetArn     string `json:"target_arn"`
	RoleArn       string `json:"role_arn"`
}

type AwsIotTopicRuleSpecDynamodb struct {
	HashKeyField  string `json:"hash_key_field"`
	HashKeyType   string `json:"hash_key_type"`
	RangeKeyValue string `json:"range_key_value"`
	RangeKeyType  string `json:"range_key_type"`
	TableName     string `json:"table_name"`
	HashKeyValue  string `json:"hash_key_value"`
	PayloadField  string `json:"payload_field"`
	RangeKeyField string `json:"range_key_field"`
	RoleArn       string `json:"role_arn"`
}

type AwsIotTopicRuleSpecKinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type AwsIotTopicRuleSpecS3 struct {
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type AwsIotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	Separator          string `json:"separator"`
}

type AwsIotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type AwsIotTopicRuleSpecRepublish struct {
	RoleArn string `json:"role_arn"`
	Topic   string `json:"topic"`
}

type AwsIotTopicRuleSpecSqs struct {
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
	QueueUrl  string `json:"queue_url"`
}

type AwsIotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type AwsIotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type AwsIotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type AwsIotTopicRuleSpec struct {
	Sns              []AwsIotTopicRuleSpec `json:"sns"`
	Name             string                `json:"name"`
	Dynamodb         []AwsIotTopicRuleSpec `json:"dynamodb"`
	Kinesis          []AwsIotTopicRuleSpec `json:"kinesis"`
	S3               []AwsIotTopicRuleSpec `json:"s3"`
	Description      string                `json:"description"`
	Enabled          bool                  `json:"enabled"`
	Firehose         []AwsIotTopicRuleSpec `json:"firehose"`
	Lambda           []AwsIotTopicRuleSpec `json:"lambda"`
	Republish        []AwsIotTopicRuleSpec `json:"republish"`
	Sqs              []AwsIotTopicRuleSpec `json:"sqs"`
	Arn              string                `json:"arn"`
	Sql              string                `json:"sql"`
	SqlVersion       string                `json:"sql_version"`
	CloudwatchAlarm  []AwsIotTopicRuleSpec `json:"cloudwatch_alarm"`
	CloudwatchMetric []AwsIotTopicRuleSpec `json:"cloudwatch_metric"`
	Elasticsearch    []AwsIotTopicRuleSpec `json:"elasticsearch"`
}

type AwsIotTopicRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRuleList is a list of AwsIotTopicRules
type AwsIotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotTopicRule CRD objects
	Items []AwsIotTopicRule `json:"items,omitempty"`
}
