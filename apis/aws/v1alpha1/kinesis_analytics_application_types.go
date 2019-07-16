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

type KinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisAnalyticsApplicationSpec   `json:"spec,omitempty"`
	Status            KinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

type KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions struct {
	LogStreamArn string `json:"log_stream_arn"`
	RoleArn      string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsKinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsParallelism struct {
	Count int `json:"count"`
}

type KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecInputsProcessingConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Lambda []KinesisAnalyticsApplicationSpecInputsProcessingConfiguration `json:"lambda"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns struct {
	// +optional
	Mapping string `json:"mapping,omitempty"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"record_row_path"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Csv *[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters `json:"csv,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Json *[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters `json:"json,omitempty"`
}

type KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MappingParameters *[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat `json:"mapping_parameters,omitempty"`
}

type KinesisAnalyticsApplicationSpecInputsSchema struct {
	RecordColumns []KinesisAnalyticsApplicationSpecInputsSchema `json:"record_columns"`
	// +optional
	RecordEncoding string `json:"record_encoding,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RecordFormat []KinesisAnalyticsApplicationSpecInputsSchema `json:"record_format"`
}

type KinesisAnalyticsApplicationSpecInputs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisFirehose *[]KinesisAnalyticsApplicationSpecInputs `json:"kinesis_firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisStream *[]KinesisAnalyticsApplicationSpecInputs `json:"kinesis_stream,omitempty"`
	NamePrefix    string                                   `json:"name_prefix"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Parallelism *[]KinesisAnalyticsApplicationSpecInputs `json:"parallelism,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProcessingConfiguration *[]KinesisAnalyticsApplicationSpecInputs `json:"processing_configuration,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Schema []KinesisAnalyticsApplicationSpecInputs `json:"schema"`
}

type KinesisAnalyticsApplicationSpecOutputsKinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsKinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsLambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecOutputsSchema struct {
	// +optional
	RecordFormatType string `json:"record_format_type,omitempty"`
}

type KinesisAnalyticsApplicationSpecOutputs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisFirehose *[]KinesisAnalyticsApplicationSpecOutputs `json:"kinesis_firehose,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisStream *[]KinesisAnalyticsApplicationSpecOutputs `json:"kinesis_stream,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Lambda *[]KinesisAnalyticsApplicationSpecOutputs `json:"lambda,omitempty"`
	Name   string                                    `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	Schema []KinesisAnalyticsApplicationSpecOutputs `json:"schema"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesS3 struct {
	BucketArn string `json:"bucket_arn"`
	FileKey   string `json:"file_key"`
	RoleArn   string `json:"role_arn"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns struct {
	// +optional
	Mapping string `json:"mapping,omitempty"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson struct {
	RecordRowPath string `json:"record_row_path"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Csv *[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"csv,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Json *[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"json,omitempty"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MappingParameters *[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat `json:"mapping_parameters,omitempty"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema struct {
	RecordColumns []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_columns"`
	// +optional
	RecordEncoding string `json:"record_encoding,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	RecordFormat []KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema `json:"record_format"`
}

type KinesisAnalyticsApplicationSpecReferenceDataSources struct {
	// +kubebuilder:validation:MaxItems=1
	S3 []KinesisAnalyticsApplicationSpecReferenceDataSources `json:"s3"`
	// +kubebuilder:validation:MaxItems=1
	Schema    []KinesisAnalyticsApplicationSpecReferenceDataSources `json:"schema"`
	TableName string                                                `json:"table_name"`
}

type KinesisAnalyticsApplicationSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CloudwatchLoggingOptions *[]KinesisAnalyticsApplicationSpec `json:"cloudwatch_logging_options,omitempty"`
	// +optional
	Code string `json:"code,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Inputs *[]KinesisAnalyticsApplicationSpec `json:"inputs,omitempty"`
	Name   string                             `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	Outputs *[]KinesisAnalyticsApplicationSpec `json:"outputs,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReferenceDataSources *[]KinesisAnalyticsApplicationSpec `json:"reference_data_sources,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type KinesisAnalyticsApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
