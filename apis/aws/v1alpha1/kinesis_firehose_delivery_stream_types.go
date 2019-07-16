package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KinesisFirehoseDeliveryStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisFirehoseDeliveryStreamSpec   `json:"spec,omitempty"`
	Status            KinesisFirehoseDeliveryStreamStatus `json:"status,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters *[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors `json:"parameters,omitempty"`
	Type       string                                                                                          `json:"type"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Processors *[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration `json:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration struct {
	// +optional
	BufferingInterval int `json:"buffering_interval,omitempty"`
	// +optional
	BufferingSize int    `json:"buffering_size,omitempty"`
	DomainArn     string `json:"domain_arn"`
	IndexName     string `json:"index_name"`
	// +optional
	IndexRotationPeriod string `json:"index_rotation_period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration *[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration `json:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int    `json:"retry_duration,omitempty"`
	RoleArn       string `json:"role_arn"`
	// +optional
	S3BackupMode string `json:"s3_backup_mode,omitempty"`
	// +optional
	TypeName string `json:"type_name,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe struct {
	// +optional
	TimestampFormats []string `json:"timestamp_formats,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe struct {
	// +optional
	CaseInsensitive bool `json:"case_insensitive,omitempty"`
	// +optional
	ColumnToJsonKeyMappings map[string]string `json:"column_to_json_key_mappings,omitempty"`
	// +optional
	ConvertDotsInJsonKeysToUnderscores bool `json:"convert_dots_in_json_keys_to_underscores,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HiveJsonSerDe *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `json:"hive_json_ser_de,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OpenXJsonSerDe *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `json:"open_x_json_ser_de,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Deserializer []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration `json:"deserializer"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe struct {
	// +optional
	BlockSizeBytes int `json:"block_size_bytes,omitempty"`
	// +optional
	BloomFilterColumns []string `json:"bloom_filter_columns,omitempty"`
	// +optional
	BloomFilterFalsePositiveProbability json.Number `json:"bloom_filter_false_positive_probability,omitempty"`
	// +optional
	Compression string `json:"compression,omitempty"`
	// +optional
	DictionaryKeyThreshold json.Number `json:"dictionary_key_threshold,omitempty"`
	// +optional
	EnablePadding bool `json:"enable_padding,omitempty"`
	// +optional
	FormatVersion string `json:"format_version,omitempty"`
	// +optional
	PaddingTolerance json.Number `json:"padding_tolerance,omitempty"`
	// +optional
	RowIndexStride int `json:"row_index_stride,omitempty"`
	// +optional
	StripeSizeBytes int `json:"stripe_size_bytes,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe struct {
	// +optional
	BlockSizeBytes int `json:"block_size_bytes,omitempty"`
	// +optional
	Compression string `json:"compression,omitempty"`
	// +optional
	EnableDictionaryCompression bool `json:"enable_dictionary_compression,omitempty"`
	// +optional
	MaxPaddingBytes int `json:"max_padding_bytes,omitempty"`
	// +optional
	PageSizeBytes int `json:"page_size_bytes,omitempty"`
	// +optional
	WriterVersion string `json:"writer_version,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OrcSerDe *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"orc_ser_de,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ParquetSerDe *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"parquet_ser_de,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Serializer []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration `json:"serializer"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration struct {
	DatabaseName string `json:"database_name"`
	RoleArn      string `json:"role_arn"`
	TableName    string `json:"table_name"`
	// +optional
	VersionId string `json:"version_id,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	InputFormatConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"input_format_configuration"`
	// +kubebuilder:validation:MaxItems=1
	OutputFormatConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"output_format_configuration"`
	// +kubebuilder:validation:MaxItems=1
	SchemaConfiguration []KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration `json:"schema_configuration"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors `json:"parameters,omitempty"`
	Type       string                                                                                       `json:"type"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Processors *[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration `json:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration struct {
	BucketArn string `json:"bucket_arn"`
	// +optional
	BufferInterval int `json:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"buffer_size,omitempty"`
	// +optional
	CompressionFormat string `json:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty"`
	RoleArn string `json:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration struct {
	BucketArn string `json:"bucket_arn"`
	// +optional
	BufferInterval int `json:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"buffer_size,omitempty"`
	// +optional
	CompressionFormat string `json:"compression_format,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DataFormatConversionConfiguration *[]KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"data_format_conversion_configuration,omitempty"`
	// +optional
	ErrorOutputPrefix string `json:"error_output_prefix,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration *[]KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"processing_configuration,omitempty"`
	RoleArn                 string                                                      `json:"role_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3BackupConfiguration *[]KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration `json:"s3_backup_configuration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3_backup_mode,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration struct {
	KinesisStreamArn string `json:"kinesis_stream_arn"`
	RoleArn          string `json:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters *[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors `json:"parameters,omitempty"`
	Type       string                                                                                     `json:"type"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Processors *[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration `json:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration struct {
	BucketArn string `json:"bucket_arn"`
	// +optional
	BufferInterval int `json:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"buffer_size,omitempty"`
	// +optional
	CompressionFormat string `json:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty"`
	RoleArn string `json:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration struct {
	ClusterJdbcurl string `json:"cluster_jdbcurl"`
	// +optional
	CopyOptions string `json:"copy_options,omitempty"`
	// +optional
	DataTableColumns string `json:"data_table_columns,omitempty"`
	DataTableName    string `json:"data_table_name"`
	Password         string `json:"password"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration *[]KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int    `json:"retry_duration,omitempty"`
	RoleArn       string `json:"role_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3BackupConfiguration *[]KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration `json:"s3_backup_configuration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3_backup_mode,omitempty"`
	Username     string `json:"username"`
}

type KinesisFirehoseDeliveryStreamSpecS3Configuration struct {
	BucketArn string `json:"bucket_arn"`
	// +optional
	BufferInterval int `json:"buffer_interval,omitempty"`
	// +optional
	BufferSize int `json:"buffer_size,omitempty"`
	// +optional
	CompressionFormat string `json:"compression_format,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	Prefix  string `json:"prefix,omitempty"`
	RoleArn string `json:"role_arn"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors struct {
	// +optional
	Parameters *[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors `json:"parameters,omitempty"`
	Type       string                                                                                   `json:"type"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Processors *[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration `json:"processors,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpecSplunkConfiguration struct {
	// +optional
	HecAcknowledgmentTimeout int    `json:"hec_acknowledgment_timeout,omitempty"`
	HecEndpoint              string `json:"hec_endpoint"`
	// +optional
	HecEndpointType string `json:"hec_endpoint_type,omitempty"`
	HecToken        string `json:"hec_token"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration *[]KinesisFirehoseDeliveryStreamSpecSplunkConfiguration `json:"processing_configuration,omitempty"`
	// +optional
	RetryDuration int `json:"retry_duration,omitempty"`
	// +optional
	S3BackupMode string `json:"s3_backup_mode,omitempty"`
}

type KinesisFirehoseDeliveryStreamSpec struct {
	Destination string `json:"destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticsearchConfiguration *[]KinesisFirehoseDeliveryStreamSpec `json:"elasticsearch_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExtendedS3Configuration *[]KinesisFirehoseDeliveryStreamSpec `json:"extended_s3_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisSourceConfiguration *[]KinesisFirehoseDeliveryStreamSpec `json:"kinesis_source_configuration,omitempty"`
	Name                       string                               `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedshiftConfiguration *[]KinesisFirehoseDeliveryStreamSpec `json:"redshift_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Configuration *[]KinesisFirehoseDeliveryStreamSpec `json:"s3_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SplunkConfiguration *[]KinesisFirehoseDeliveryStreamSpec `json:"splunk_configuration,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type KinesisFirehoseDeliveryStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
