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
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type AwsIotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type AwsIotTopicRuleSpecDynamodb struct {
	RangeKeyValue string `json:"range_key_value"`
	RoleArn       string `json:"role_arn"`
	HashKeyField  string `json:"hash_key_field"`
	RangeKeyField string `json:"range_key_field"`
	PayloadField  string `json:"payload_field"`
	RangeKeyType  string `json:"range_key_type"`
	TableName     string `json:"table_name"`
	HashKeyValue  string `json:"hash_key_value"`
	HashKeyType   string `json:"hash_key_type"`
}

type AwsIotTopicRuleSpecRepublish struct {
	Topic   string `json:"topic"`
	RoleArn string `json:"role_arn"`
}

type AwsIotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type AwsIotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type AwsIotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	Separator          string `json:"separator"`
}

type AwsIotTopicRuleSpecKinesis struct {
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
	PartitionKey string `json:"partition_key"`
}

type AwsIotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type AwsIotTopicRuleSpecSns struct {
	TargetArn     string `json:"target_arn"`
	RoleArn       string `json:"role_arn"`
	MessageFormat string `json:"message_format"`
}

type AwsIotTopicRuleSpecSqs struct {
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
}

type AwsIotTopicRuleSpec struct {
	Sql              string                `json:"sql"`
	S3               []AwsIotTopicRuleSpec `json:"s3"`
	Name             string                `json:"name"`
	CloudwatchMetric []AwsIotTopicRuleSpec `json:"cloudwatch_metric"`
	Dynamodb         []AwsIotTopicRuleSpec `json:"dynamodb"`
	Republish        []AwsIotTopicRuleSpec `json:"republish"`
	Description      string                `json:"description"`
	SqlVersion       string                `json:"sql_version"`
	Elasticsearch    []AwsIotTopicRuleSpec `json:"elasticsearch"`
	Arn              string                `json:"arn"`
	Enabled          bool                  `json:"enabled"`
	CloudwatchAlarm  []AwsIotTopicRuleSpec `json:"cloudwatch_alarm"`
	Firehose         []AwsIotTopicRuleSpec `json:"firehose"`
	Kinesis          []AwsIotTopicRuleSpec `json:"kinesis"`
	Lambda           []AwsIotTopicRuleSpec `json:"lambda"`
	Sns              []AwsIotTopicRuleSpec `json:"sns"`
	Sqs              []AwsIotTopicRuleSpec `json:"sqs"`
}

type AwsIotTopicRuleStatus struct {
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
