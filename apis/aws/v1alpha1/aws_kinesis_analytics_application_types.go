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

type AwsKinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKinesisAnalyticsApplicationSpec   `json:"spec,omitempty"`
	Status            AwsKinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

type AwsKinesisAnalyticsApplicationSpecCloudwatchLoggingOptions struct {
	Id           string `json:"id"`
	LogStreamArn string `json:"log_stream_arn"`
	RoleArn      string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecInputsKinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecInputsParallelism struct {
	Count int `json:"count"`
}

type AwsKinesisAnalyticsApplicationSpecInputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecInputsProcessingConfiguration struct {
	Lambda []AwsKinesisAnalyticsApplicationSpecInputsProcessingConfiguration `json:"lambda"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordColumns struct {
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"record_row_path"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters struct {
	Csv  []AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters `json:"csv"`
	Json []AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters `json:"json"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormat struct {
	MappingParameters []AwsKinesisAnalyticsApplicationSpecInputsSchemaRecordFormat `json:"mapping_parameters"`
	RecordFormatType  string                                                       `json:"record_format_type"`
}

type AwsKinesisAnalyticsApplicationSpecInputsSchema struct {
	RecordColumns  []AwsKinesisAnalyticsApplicationSpecInputsSchema `json:"record_columns"`
	RecordEncoding string                                           `json:"record_encoding"`
	RecordFormat   []AwsKinesisAnalyticsApplicationSpecInputsSchema `json:"record_format"`
}

type AwsKinesisAnalyticsApplicationSpecInputsStartingPositionConfiguration struct {
	StartingPosition string `json:"starting_position"`
}

type AwsKinesisAnalyticsApplicationSpecInputs struct {
	Id                            string                                     `json:"id"`
	KinesisFirehose               []AwsKinesisAnalyticsApplicationSpecInputs `json:"kinesis_firehose"`
	Parallelism                   []AwsKinesisAnalyticsApplicationSpecInputs `json:"parallelism"`
	StreamNames                   []string                                   `json:"stream_names"`
	KinesisStream                 []AwsKinesisAnalyticsApplicationSpecInputs `json:"kinesis_stream"`
	NamePrefix                    string                                     `json:"name_prefix"`
	ProcessingConfiguration       []AwsKinesisAnalyticsApplicationSpecInputs `json:"processing_configuration"`
	Schema                        []AwsKinesisAnalyticsApplicationSpecInputs `json:"schema"`
	StartingPositionConfiguration []AwsKinesisAnalyticsApplicationSpecInputs `json:"starting_position_configuration"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsLambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsSchema struct {
	RecordFormatType string `json:"record_format_type"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsKinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputs struct {
	Lambda          []AwsKinesisAnalyticsApplicationSpecOutputs `json:"lambda"`
	Name            string                                      `json:"name"`
	Schema          []AwsKinesisAnalyticsApplicationSpecOutputs `json:"schema"`
	Id              string                                      `json:"id"`
	KinesisFirehose []AwsKinesisAnalyticsApplicationSpecOutputs `json:"kinesis_firehose"`
	KinesisStream   []AwsKinesisAnalyticsApplicationSpecOutputs `json:"kinesis_stream"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"record_row_path"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	Json []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"json"`
	Csv  []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"csv"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat struct {
	MappingParameters []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat `json:"mapping_parameters"`
	RecordFormatType  string                                                                     `json:"record_format_type"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns struct {
	SqlType string `json:"sql_type"`
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema struct {
	RecordFormat   []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_format"`
	RecordColumns  []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_columns"`
	RecordEncoding string                                                         `json:"record_encoding"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesS3 struct {
	FileKey   string `json:"file_key"`
	RoleArn   string `json:"role_arn"`
	BucketArn string `json:"bucket_arn"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSources struct {
	Schema    []AwsKinesisAnalyticsApplicationSpecReferenceDataSources `json:"schema"`
	TableName string                                                   `json:"table_name"`
	Id        string                                                   `json:"id"`
	S3        []AwsKinesisAnalyticsApplicationSpecReferenceDataSources `json:"s3"`
}

type AwsKinesisAnalyticsApplicationSpec struct {
	CloudwatchLoggingOptions []AwsKinesisAnalyticsApplicationSpec `json:"cloudwatch_logging_options"`
	CreateTimestamp          string                               `json:"create_timestamp"`
	LastUpdateTimestamp      string                               `json:"last_update_timestamp"`
	Code                     string                               `json:"code"`
	Description              string                               `json:"description"`
	Status                   string                               `json:"status"`
	Version                  int                                  `json:"version"`
	Inputs                   []AwsKinesisAnalyticsApplicationSpec `json:"inputs"`
	Outputs                  []AwsKinesisAnalyticsApplicationSpec `json:"outputs"`
	Name                     string                               `json:"name"`
	Arn                      string                               `json:"arn"`
	ReferenceDataSources     []AwsKinesisAnalyticsApplicationSpec `json:"reference_data_sources"`
	Tags                     map[string]string                    `json:"tags"`
}

type AwsKinesisAnalyticsApplicationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKinesisAnalyticsApplicationList is a list of AwsKinesisAnalyticsApplications
type AwsKinesisAnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKinesisAnalyticsApplication CRD objects
	Items []AwsKinesisAnalyticsApplication `json:"items,omitempty"`
}
