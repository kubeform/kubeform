package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KinesisFirehoseDeliveryStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisFirehoseDeliveryStreamSpec   `json:"spec,omitempty"`
	Status            KinesisFirehoseDeliveryStreamStatus `json:"status,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameterName" tf:"parameter_name"`
	ParameterValue string `json:"parameterValue" tf:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters []KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
	Type       string                                                                                                   `json:"type" tf:"type"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Processors []KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors `json:"processors,omitempty" tf:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration struct {
	// +optional
	BufferingInterval int `json:"bufferingInterval,omitempty" tf:"buffering_interval,omitempty"`
	// +optional
	BufferingSize int `json:"bufferingSize,omitempty" tf:"buffering_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	DomainArn                string                                                                                `json:"domainArn" tf:"domain_arn"`
	IndexName                string                                                                                `json:"indexName" tf:"index_name"`
	// +optional
	IndexRotationPeriod string `json:"indexRotationPeriod,omitempty" tf:"index_rotation_period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration []KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int    `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	// +optional
	S3BackupMode string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
	// +optional
	TypeName string `json:"typeName,omitempty" tf:"type_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe struct {
	// +optional
	TimestampFormats []string `json:"timestampFormats,omitempty" tf:"timestamp_formats,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe struct {
	// +optional
	CaseInsensitive bool `json:"caseInsensitive,omitempty" tf:"case_insensitive,omitempty"`
	// +optional
	ColumnToJSONKeyMappings map[string]string `json:"columnToJSONKeyMappings,omitempty" tf:"column_to_json_key_mappings,omitempty"`
	// +optional
	ConvertDotsInJSONKeysToUnderscores bool `json:"convertDotsInJSONKeysToUnderscores,omitempty" tf:"convert_dots_in_json_keys_to_underscores,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HiveJSONSerDe []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe `json:"hiveJSONSerDe,omitempty" tf:"hive_json_ser_de,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OpenXJSONSerDe []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe `json:"openXJSONSerDe,omitempty" tf:"open_x_json_ser_de,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Deserializer []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `json:"deserializer" tf:"deserializer"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe struct {
	// +optional
	BlockSizeBytes int `json:"blockSizeBytes,omitempty" tf:"block_size_bytes,omitempty"`
	// +optional
	BloomFilterColumns []string `json:"bloomFilterColumns,omitempty" tf:"bloom_filter_columns,omitempty"`
	// +optional
	BloomFilterFalsePositiveProbability json.Number `json:"bloomFilterFalsePositiveProbability,omitempty" tf:"bloom_filter_false_positive_probability,omitempty"`
	// +optional
	Compression string `json:"compression,omitempty" tf:"compression,omitempty"`
	// +optional
	DictionaryKeyThreshold json.Number `json:"dictionaryKeyThreshold,omitempty" tf:"dictionary_key_threshold,omitempty"`
	// +optional
	EnablePadding bool `json:"enablePadding,omitempty" tf:"enable_padding,omitempty"`
	// +optional
	FormatVersion string `json:"formatVersion,omitempty" tf:"format_version,omitempty"`
	// +optional
	PaddingTolerance json.Number `json:"paddingTolerance,omitempty" tf:"padding_tolerance,omitempty"`
	// +optional
	RowIndexStride int `json:"rowIndexStride,omitempty" tf:"row_index_stride,omitempty"`
	// +optional
	StripeSizeBytes int `json:"stripeSizeBytes,omitempty" tf:"stripe_size_bytes,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe struct {
	// +optional
	BlockSizeBytes int `json:"blockSizeBytes,omitempty" tf:"block_size_bytes,omitempty"`
	// +optional
	Compression string `json:"compression,omitempty" tf:"compression,omitempty"`
	// +optional
	EnableDictionaryCompression bool `json:"enableDictionaryCompression,omitempty" tf:"enable_dictionary_compression,omitempty"`
	// +optional
	MaxPaddingBytes int `json:"maxPaddingBytes,omitempty" tf:"max_padding_bytes,omitempty"`
	// +optional
	PageSizeBytes int `json:"pageSizeBytes,omitempty" tf:"page_size_bytes,omitempty"`
	// +optional
	WriterVersion string `json:"writerVersion,omitempty" tf:"writer_version,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OrcSerDe []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe `json:"orcSerDe,omitempty" tf:"orc_ser_de,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ParquetSerDe []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe `json:"parquetSerDe,omitempty" tf:"parquet_ser_de,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Serializer []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"serializer" tf:"serializer"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration struct {
	// +optional
	CatalogID    string `json:"catalogID,omitempty" tf:"catalog_id,omitempty"`
	DatabaseName string `json:"databaseName" tf:"database_name"`
	// +optional
	Region    string `json:"region,omitempty" tf:"region,omitempty"`
	RoleArn   string `json:"roleArn" tf:"role_arn"`
	TableName string `json:"tableName" tf:"table_name"`
	// +optional
	VersionID string `json:"versionID,omitempty" tf:"version_id,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	InputFormatConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration `json:"inputFormatConfiguration" tf:"input_format_configuration"`
	// +kubebuilder:validation:MaxItems=1
	OutputFormatConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration `json:"outputFormatConfiguration" tf:"output_format_configuration"`
	// +kubebuilder:validation:MaxItems=1
	SchemaConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration `json:"schemaConfiguration" tf:"schema_configuration"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameterName" tf:"parameter_name"`
	ParameterValue string `json:"parameterValue" tf:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
	Type       string                                                                                                `json:"type" tf:"type"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Processors []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors `json:"processors,omitempty" tf:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	// +optional
	BufferInterval int `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	CompressionFormat string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	// +optional
	BufferInterval int `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	CompressionFormat string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DataFormatConversionConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"dataFormatConversionConfiguration,omitempty" tf:"data_format_conversion_configuration,omitempty"`
	// +optional
	ErrorOutputPrefix string `json:"errorOutputPrefix,omitempty" tf:"error_output_prefix,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`
	RoleArn                 string                                                                            `json:"roleArn" tf:"role_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3BackupConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration `json:"s3BackupConfiguration,omitempty" tf:"s3_backup_configuration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration struct {
	KinesisStreamArn string `json:"kinesisStreamArn" tf:"kinesis_stream_arn"`
	RoleArn          string `json:"roleArn" tf:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameterName" tf:"parameter_name"`
	ParameterValue string `json:"parameterValue" tf:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
	Type       string                                                                                              `json:"type" tf:"type"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Processors []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors `json:"processors,omitempty" tf:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	// +optional
	BufferInterval int `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	CompressionFormat string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	ClusterJdbcurl           string                                                                           `json:"clusterJdbcurl" tf:"cluster_jdbcurl"`
	// +optional
	CopyOptions string `json:"copyOptions,omitempty" tf:"copy_options,omitempty"`
	// +optional
	DataTableColumns string `json:"dataTableColumns,omitempty" tf:"data_table_columns,omitempty"`
	DataTableName    string `json:"dataTableName" tf:"data_table_name"`
	Password         string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int    `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3BackupConfiguration []KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration `json:"s3BackupConfiguration,omitempty" tf:"s3_backup_configuration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
	Username     string `json:"username" tf:"username"`
}

type KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecS3Configuration struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	// +optional
	BufferInterval int `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	CompressionFormat string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	LogGroupName string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`
	// +optional
	LogStreamName string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameterName" tf:"parameter_name"`
	ParameterValue string `json:"parameterValue" tf:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters []KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
	Type       string                                                                                            `json:"type" tf:"type"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Processors []KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors `json:"processors,omitempty" tf:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchLoggingOptions []KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	HecAcknowledgmentTimeout int    `json:"hecAcknowledgmentTimeout,omitempty" tf:"hec_acknowledgment_timeout,omitempty"`
	HecEndpoint              string `json:"hecEndpoint" tf:"hec_endpoint"`
	// +optional
	HecEndpointType string `json:"hecEndpointType,omitempty" tf:"hec_endpoint_type,omitempty"`
	HecToken        string `json:"hecToken" tf:"hec_token"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration []KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Arn         string `json:"arn,omitempty" tf:"arn,omitempty"`
	Destination string `json:"destination" tf:"destination"`
	// +optional
	DestinationID string `json:"destinationID,omitempty" tf:"destination_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticsearchConfiguration []KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration `json:"elasticsearchConfiguration,omitempty" tf:"elasticsearch_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExtendedS3Configuration []KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"extendedS3Configuration,omitempty" tf:"extended_s3_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisSourceConfiguration []KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration `json:"kinesisSourceConfiguration,omitempty" tf:"kinesis_source_configuration,omitempty"`
	Name                       string                                                        `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedshiftConfiguration []KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"redshiftConfiguration,omitempty" tf:"redshift_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Configuration []KinesisFirehoseDeliveryStreamSpecS3Configuration `json:"s3Configuration,omitempty" tf:"s3_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SplunkConfiguration []KinesisFirehoseDeliveryStreamSpecSplunkConfiguration `json:"splunkConfiguration,omitempty" tf:"splunk_configuration,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VersionID string `json:"versionID,omitempty" tf:"version_id,omitempty"`
}

type KinesisFirehoseDeliveryStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KinesisFirehoseDeliveryStreamSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KinesisFirehoseDeliveryStreamList is a list of KinesisFirehoseDeliveryStreams
type KinesisFirehoseDeliveryStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KinesisFirehoseDeliveryStream CRD objects
	Items []KinesisFirehoseDeliveryStream `json:"items,omitempty"`
}
