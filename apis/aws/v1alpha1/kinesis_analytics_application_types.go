package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisAnalyticsApplicationSpec   `json:"spec,omitempty"`
	Status            KinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

type KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions struct {
	LogStreamArn string `json:"logStreamArn" tf:"log_stream_arn"`
	RoleArn      string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsKinesisFirehose struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsKinesisStream struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsParallelism struct {
	Count int `json:"count" tf:"count"`
}

type KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsProcessingConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Lambda []KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda `json:"lambda" tf:"lambda"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns struct {
	// +optional
	Mapping string `json:"mapping,omitempty" tf:"mapping,omitempty"`
	Name    string `json:"name" tf:"name"`
	SqlType string `json:"sqlType" tf:"sql_type"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"recordRowPath" tf:"record_row_path"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Csv []KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv `json:"csv,omitempty" tf:"csv,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Json []KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson `json:"json,omitempty" tf:"json,omitempty"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MappingParameters []KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type KinesisAnalyticsApplicationSpecInputsSchema struct {
	RecordColumns []KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns `json:"recordColumns" tf:"record_columns"`
	// +optional
	RecordEncoding string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RecordFormat []KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat `json:"recordFormat" tf:"record_format"`
}

type KinesisAnalyticsApplicationSpecInputs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisFirehose []KinesisAnalyticsApplicationSpecInputsKinesisFirehose `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisStream []KinesisAnalyticsApplicationSpecInputsKinesisStream `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`
	NamePrefix    string                                               `json:"namePrefix" tf:"name_prefix"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Parallelism []KinesisAnalyticsApplicationSpecInputsParallelism `json:"parallelism,omitempty" tf:"parallelism,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration []KinesisAnalyticsApplicationSpecInputsProcessingConfiguration `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Schema []KinesisAnalyticsApplicationSpecInputsSchema `json:"schema" tf:"schema"`
}

type KinesisAnalyticsApplicationSpecOutputsKinesisFirehose struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsKinesisStream struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsLambda struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
	RoleArn     string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsSchema struct {
	// +optional
	RecordFormatType string `json:"recordFormatType,omitempty" tf:"record_format_type,omitempty"`
}

type KinesisAnalyticsApplicationSpecOutputs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisFirehose []KinesisAnalyticsApplicationSpecOutputsKinesisFirehose `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisStream []KinesisAnalyticsApplicationSpecOutputsKinesisStream `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Lambda []KinesisAnalyticsApplicationSpecOutputsLambda `json:"lambda,omitempty" tf:"lambda,omitempty"`
	Name   string                                         `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	Schema []KinesisAnalyticsApplicationSpecOutputsSchema `json:"schema" tf:"schema"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesS3 struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	FileKey   string `json:"fileKey" tf:"file_key"`
	RoleArn   string `json:"roleArn" tf:"role_arn"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns struct {
	// +optional
	Mapping string `json:"mapping,omitempty" tf:"mapping,omitempty"`
	Name    string `json:"name" tf:"name"`
	SqlType string `json:"sqlType" tf:"sql_type"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"recordRowPath" tf:"record_row_path"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Csv []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv `json:"csv,omitempty" tf:"csv,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Json []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson `json:"json,omitempty" tf:"json,omitempty"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MappingParameters []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema struct {
	RecordColumns []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns `json:"recordColumns" tf:"record_columns"`
	// +optional
	RecordEncoding string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RecordFormat []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat `json:"recordFormat" tf:"record_format"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSources struct {
	// +kubebuilder:validation:MaxItems=1
	S3 []KinesisAnalyticsApplicationSpecReferenceDataSourcesS3 `json:"s3" tf:"s3"`
	// +kubebuilder:validation:MaxItems=1
	Schema    []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"schema" tf:"schema"`
	TableName string                                                      `json:"tableName" tf:"table_name"`
}

type KinesisAnalyticsApplicationSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CloudwatchLoggingOptions []KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`
	// +optional
	Code string `json:"code,omitempty" tf:"code,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Inputs []KinesisAnalyticsApplicationSpecInputs `json:"inputs,omitempty" tf:"inputs,omitempty"`
	Name   string                                  `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	Outputs []KinesisAnalyticsApplicationSpecOutputs `json:"outputs,omitempty" tf:"outputs,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReferenceDataSources []KinesisAnalyticsApplicationSpecReferenceDataSources `json:"referenceDataSources,omitempty" tf:"reference_data_sources,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type KinesisAnalyticsApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KinesisAnalyticsApplicationList is a list of KinesisAnalyticsApplications
type KinesisAnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KinesisAnalyticsApplication CRD objects
	Items []KinesisAnalyticsApplication `json:"items,omitempty"`
}
