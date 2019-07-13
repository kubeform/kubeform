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

type AwsKinesisFirehoseDeliveryStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKinesisFirehoseDeliveryStreamSpec   `json:"spec,omitempty"`
	Status            AwsKinesisFirehoseDeliveryStreamStatus `json:"status,omitempty"`
}

type AwsKinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration struct {
	KinesisStreamArn string `json:"kinesis_stream_arn"`
	RoleArn          string `json:"role_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecS3Configuration struct {
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecS3Configuration `json:"cloudwatch_logging_options"`
	BucketArn                string                                                `json:"bucket_arn"`
	BufferSize               int                                                   `json:"buffer_size"`
	BufferInterval           int                                                   `json:"buffer_interval"`
	CompressionFormat        string                                                `json:"compression_format"`
	KmsKeyArn                string                                                `json:"kms_key_arn"`
	RoleArn                  string                                                `json:"role_arn"`
	Prefix                   string                                                `json:"prefix"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors `json:"parameters"`
	Type       string                                                                                       `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                               `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions struct {
	LogStreamName string `json:"log_stream_name"`
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	LogStreamName string `json:"log_stream_name"`
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration struct {
	RoleArn                  string                                                                           `json:"role_arn"`
	Prefix                   string                                                                           `json:"prefix"`
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration `json:"cloudwatch_logging_options"`
	BucketArn                string                                                                           `json:"bucket_arn"`
	BufferSize               int                                                                              `json:"buffer_size"`
	BufferInterval           int                                                                              `json:"buffer_interval"`
	CompressionFormat        string                                                                           `json:"compression_format"`
	KmsKeyArn                string                                                                           `json:"kms_key_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration struct {
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"processing_configuration"`
	S3BackupMode             string                                                      `json:"s3_backup_mode"`
	RetryDuration            int                                                         `json:"retry_duration"`
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"cloudwatch_logging_options"`
	DataTableName            string                                                      `json:"data_table_name"`
	ClusterJdbcurl           string                                                      `json:"cluster_jdbcurl"`
	Username                 string                                                      `json:"username"`
	Password                 string                                                      `json:"password"`
	RoleArn                  string                                                      `json:"role_arn"`
	S3BackupConfiguration    []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"s3_backup_configuration"`
	CopyOptions              string                                                      `json:"copy_options"`
	DataTableColumns         string                                                      `json:"data_table_columns"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterValue string `json:"parameter_value"`
	ParameterName  string `json:"parameter_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors `json:"parameters"`
	Type       string                                                                                            `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                                    `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration struct {
	DomainArn                string                                                           `json:"domain_arn"`
	RoleArn                  string                                                           `json:"role_arn"`
	S3BackupMode             string                                                           `json:"s3_backup_mode"`
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration `json:"processing_configuration"`
	BufferingInterval        int                                                              `json:"buffering_interval"`
	IndexName                string                                                           `json:"index_name"`
	IndexRotationPeriod      string                                                           `json:"index_rotation_period"`
	RetryDuration            int                                                              `json:"retry_duration"`
	TypeName                 string                                                           `json:"type_name"`
	BufferingSize            int                                                              `json:"buffering_size"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterValue string `json:"parameter_value"`
	ParameterName  string `json:"parameter_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors `json:"parameters"`
	Type       string                                                                                     `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                             `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfiguration struct {
	HecAcknowledgmentTimeout int                                                       `json:"hec_acknowledgment_timeout"`
	HecEndpoint              string                                                    `json:"hec_endpoint"`
	HecEndpointType          string                                                    `json:"hec_endpoint_type"`
	HecToken                 string                                                    `json:"hec_token"`
	S3BackupMode             string                                                    `json:"s3_backup_mode"`
	RetryDuration            int                                                       `json:"retry_duration"`
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfiguration `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfiguration `json:"processing_configuration"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe struct {
	TimestampFormats []string `json:"timestamp_formats"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe struct {
	CaseInsensitive                    bool              `json:"case_insensitive"`
	ColumnToJsonKeyMappings            map[string]string `json:"column_to_json_key_mappings"`
	ConvertDotsInJsonKeysToUnderscores bool              `json:"convert_dots_in_json_keys_to_underscores"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	HiveJsonSerDe  []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `json:"hive_json_ser_de"`
	OpenXJsonSerDe []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `json:"open_x_json_ser_de"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration struct {
	Deserializer []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration `json:"deserializer"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe struct {
	BlockSizeBytes                      int      `json:"block_size_bytes"`
	BloomFilterColumns                  []string `json:"bloom_filter_columns"`
	BloomFilterFalsePositiveProbability float64  `json:"bloom_filter_false_positive_probability"`
	Compression                         string   `json:"compression"`
	DictionaryKeyThreshold              float64  `json:"dictionary_key_threshold"`
	EnablePadding                       bool     `json:"enable_padding"`
	FormatVersion                       string   `json:"format_version"`
	RowIndexStride                      int      `json:"row_index_stride"`
	PaddingTolerance                    float64  `json:"padding_tolerance"`
	StripeSizeBytes                     int      `json:"stripe_size_bytes"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe struct {
	BlockSizeBytes              int    `json:"block_size_bytes"`
	Compression                 string `json:"compression"`
	EnableDictionaryCompression bool   `json:"enable_dictionary_compression"`
	MaxPaddingBytes             int    `json:"max_padding_bytes"`
	PageSizeBytes               int    `json:"page_size_bytes"`
	WriterVersion               string `json:"writer_version"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer struct {
	OrcSerDe     []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"orc_ser_de"`
	ParquetSerDe []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"parquet_ser_de"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration struct {
	Serializer []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration `json:"serializer"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration struct {
	VersionId    string `json:"version_id"`
	CatalogId    string `json:"catalog_id"`
	DatabaseName string `json:"database_name"`
	Region       string `json:"region"`
	RoleArn      string `json:"role_arn"`
	TableName    string `json:"table_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration struct {
	InputFormatConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"input_format_configuration"`
	OutputFormatConfiguration []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"output_format_configuration"`
	SchemaConfiguration       []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"schema_configuration"`
	Enabled                   bool                                                                                           `json:"enabled"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration struct {
	CloudwatchLoggingOptions []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration `json:"cloudwatch_logging_options"`
	BucketArn                string                                                                             `json:"bucket_arn"`
	BufferSize               int                                                                                `json:"buffer_size"`
	BufferInterval           int                                                                                `json:"buffer_interval"`
	CompressionFormat        string                                                                             `json:"compression_format"`
	KmsKeyArn                string                                                                             `json:"kms_key_arn"`
	RoleArn                  string                                                                             `json:"role_arn"`
	Prefix                   string                                                                             `json:"prefix"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors `json:"parameters"`
	Type       string                                                                                         `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                                 `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration struct {
	DataFormatConversionConfiguration []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"data_format_conversion_configuration"`
	KmsKeyArn                         string                                                        `json:"kms_key_arn"`
	RoleArn                           string                                                        `json:"role_arn"`
	S3BackupConfiguration             []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"s3_backup_configuration"`
	ProcessingConfiguration           []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"processing_configuration"`
	CompressionFormat                 string                                                        `json:"compression_format"`
	BufferSize                        int                                                           `json:"buffer_size"`
	BufferInterval                    int                                                           `json:"buffer_interval"`
	ErrorOutputPrefix                 string                                                        `json:"error_output_prefix"`
	Prefix                            string                                                        `json:"prefix"`
	S3BackupMode                      string                                                        `json:"s3_backup_mode"`
	CloudwatchLoggingOptions          []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"cloudwatch_logging_options"`
	BucketArn                         string                                                        `json:"bucket_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpec struct {
	VersionId                  string                                 `json:"version_id"`
	DestinationId              string                                 `json:"destination_id"`
	KinesisSourceConfiguration []AwsKinesisFirehoseDeliveryStreamSpec `json:"kinesis_source_configuration"`
	Destination                string                                 `json:"destination"`
	S3Configuration            []AwsKinesisFirehoseDeliveryStreamSpec `json:"s3_configuration"`
	RedshiftConfiguration      []AwsKinesisFirehoseDeliveryStreamSpec `json:"redshift_configuration"`
	ElasticsearchConfiguration []AwsKinesisFirehoseDeliveryStreamSpec `json:"elasticsearch_configuration"`
	SplunkConfiguration        []AwsKinesisFirehoseDeliveryStreamSpec `json:"splunk_configuration"`
	Arn                        string                                 `json:"arn"`
	Name                       string                                 `json:"name"`
	Tags                       map[string]string                      `json:"tags"`
	ExtendedS3Configuration    []AwsKinesisFirehoseDeliveryStreamSpec `json:"extended_s3_configuration"`
}



type AwsKinesisFirehoseDeliveryStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStreams
type AwsKinesisFirehoseDeliveryStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKinesisFirehoseDeliveryStream CRD objects
	Items []AwsKinesisFirehoseDeliveryStream `json:"items,omitempty"`
}