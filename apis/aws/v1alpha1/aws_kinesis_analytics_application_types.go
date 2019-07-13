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

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesS3 struct {
	BucketArn string `json:"bucket_arn"`
	FileKey   string `json:"file_key"`
	RoleArn   string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns struct {
	SqlType string `json:"sql_type"`
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"record_row_path"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	Csv  []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"csv"`
	Json []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"json"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat struct {
	RecordFormatType  string                                                                     `json:"record_format_type"`
	MappingParameters []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat `json:"mapping_parameters"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema struct {
	RecordColumns  []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_columns"`
	RecordEncoding string                                                         `json:"record_encoding"`
	RecordFormat   []AwsKinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_format"`
}

type AwsKinesisAnalyticsApplicationSpecReferenceDataSources struct {
	TableName string                                                   `json:"table_name"`
	Id        string                                                   `json:"id"`
	S3        []AwsKinesisAnalyticsApplicationSpecReferenceDataSources `json:"s3"`
	Schema    []AwsKinesisAnalyticsApplicationSpecReferenceDataSources `json:"schema"`
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

type AwsKinesisAnalyticsApplicationSpecInputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecInputsParallelism struct {
	Count int `json:"count"`
}

type AwsKinesisAnalyticsApplicationSpecInputsStartingPositionConfiguration struct {
	StartingPosition string `json:"starting_position"`
}

type AwsKinesisAnalyticsApplicationSpecInputs struct {
	Id                            string                                     `json:"id"`
	KinesisFirehose               []AwsKinesisAnalyticsApplicationSpecInputs `json:"kinesis_firehose"`
	ProcessingConfiguration       []AwsKinesisAnalyticsApplicationSpecInputs `json:"processing_configuration"`
	Schema                        []AwsKinesisAnalyticsApplicationSpecInputs `json:"schema"`
	StreamNames                   []string                                   `json:"stream_names"`
	KinesisStream                 []AwsKinesisAnalyticsApplicationSpecInputs `json:"kinesis_stream"`
	NamePrefix                    string                                     `json:"name_prefix"`
	Parallelism                   []AwsKinesisAnalyticsApplicationSpecInputs `json:"parallelism"`
	StartingPositionConfiguration []AwsKinesisAnalyticsApplicationSpecInputs `json:"starting_position_configuration"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsKinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsLambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type AwsKinesisAnalyticsApplicationSpecOutputsSchema struct {
	RecordFormatType string `json:"record_format_type"`
}

type AwsKinesisAnalyticsApplicationSpecOutputs struct {
	Id              string                                      `json:"id"`
	KinesisFirehose []AwsKinesisAnalyticsApplicationSpecOutputs `json:"kinesis_firehose"`
	KinesisStream   []AwsKinesisAnalyticsApplicationSpecOutputs `json:"kinesis_stream"`
	Lambda          []AwsKinesisAnalyticsApplicationSpecOutputs `json:"lambda"`
	Name            string                                      `json:"name"`
	Schema          []AwsKinesisAnalyticsApplicationSpecOutputs `json:"schema"`
}

type AwsKinesisAnalyticsApplicationSpec struct {
	Tags                     map[string]string                    `json:"tags"`
	Name                     string                               `json:"name"`
	Arn                      string                               `json:"arn"`
	ReferenceDataSources     []AwsKinesisAnalyticsApplicationSpec `json:"reference_data_sources"`
	LastUpdateTimestamp      string                               `json:"last_update_timestamp"`
	Status                   string                               `json:"status"`
	Version                  int                                  `json:"version"`
	CloudwatchLoggingOptions []AwsKinesisAnalyticsApplicationSpec `json:"cloudwatch_logging_options"`
	Inputs                   []AwsKinesisAnalyticsApplicationSpec `json:"inputs"`
	Code                     string                               `json:"code"`
	CreateTimestamp          string                               `json:"create_timestamp"`
	Description              string                               `json:"description"`
	Outputs                  []AwsKinesisAnalyticsApplicationSpec `json:"outputs"`
}



type AwsKinesisAnalyticsApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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