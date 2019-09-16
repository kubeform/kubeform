---
title: KinesisFirehoseDeliveryStream
menu:
  docs_{{ .version }}:
    identifier: kinesisfirehosedeliverystream-aws.kubeform.com
    name: KinesisFirehoseDeliveryStream
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KinesisFirehoseDeliveryStream
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KinesisFirehoseDeliveryStream` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)***||
| `status` | ***[KinesisFirehoseDeliveryStreamStatus](#KinesisFirehoseDeliveryStreamStatus)***||
## KinesisFirehoseDeliveryStreamSpec

Appears on:[KinesisFirehoseDeliveryStream](#KinesisFirehoseDeliveryStream), [KinesisFirehoseDeliveryStreamStatus](#KinesisFirehoseDeliveryStreamStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `destination` | ***string***||
| `destinationID` | ***string***| ***(Optional)*** |
| `elasticsearchConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration)***| ***(Optional)*** |
| `extendedS3Configuration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration)***| ***(Optional)*** |
| `kinesisSourceConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration](#KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration)***| ***(Optional)*** |
| `name` | ***string***||
| `redshiftConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration)***| ***(Optional)*** |
| `s3Configuration` | ***[[]KinesisFirehoseDeliveryStreamSpecS3Configuration](#KinesisFirehoseDeliveryStreamSpecS3Configuration)***| ***(Optional)*** |
| `splunkConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#KinesisFirehoseDeliveryStreamSpecSplunkConfiguration)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `versionID` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bufferingInterval` | ***int***| ***(Optional)*** |
| `bufferingSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `domainArn` | ***string***||
| `indexName` | ***string***||
| `indexRotationPeriod` | ***string***| ***(Optional)*** |
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupMode` | ***string***| ***(Optional)*** |
| `typeName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `dataFormatConversionConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration)***| ***(Optional)*** |
| `errorOutputPrefix` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration)***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration)***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `inputFormatConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration)***||
| `outputFormatConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration)***||
| `schemaConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deserializer` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hiveJSONSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe)***| ***(Optional)*** |
| `openXJSONSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `timestampFormats` | ***[]string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caseInsensitive` | ***bool***| ***(Optional)*** |
| `columnToJSONKeyMappings` | ***map[string]string***| ***(Optional)*** |
| `convertDotsInJSONKeysToUnderscores` | ***bool***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `serializer` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `orcSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe)***| ***(Optional)*** |
| `parquetSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `blockSizeBytes` | ***int***| ***(Optional)*** |
| `bloomFilterColumns` | ***[]string***| ***(Optional)*** |
| `bloomFilterFalsePositiveProbability` | ***encoding/json.Number***| ***(Optional)*** |
| `compression` | ***string***| ***(Optional)*** |
| `dictionaryKeyThreshold` | ***encoding/json.Number***| ***(Optional)*** |
| `enablePadding` | ***bool***| ***(Optional)*** |
| `formatVersion` | ***string***| ***(Optional)*** |
| `paddingTolerance` | ***encoding/json.Number***| ***(Optional)*** |
| `rowIndexStride` | ***int***| ***(Optional)*** |
| `stripeSizeBytes` | ***int***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `blockSizeBytes` | ***int***| ***(Optional)*** |
| `compression` | ***string***| ***(Optional)*** |
| `enableDictionaryCompression` | ***bool***| ***(Optional)*** |
| `maxPaddingBytes` | ***int***| ***(Optional)*** |
| `pageSizeBytes` | ***int***| ***(Optional)*** |
| `writerVersion` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `catalogID` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `region` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `tableName` | ***string***||
| `versionID` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration](#KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `kinesisStreamArn` | ***string***||
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `clusterJdbcurl` | ***string***||
| `copyOptions` | ***string***| ***(Optional)*** |
| `dataTableColumns` | ***string***| ***(Optional)*** |
| `dataTableName` | ***string***||
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration)***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
| `username` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration](#KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecS3Configuration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecS3Configuration](#KinesisFirehoseDeliveryStreamSpecS3Configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions)***| ***(Optional)*** |
| `hecAcknowledgmentTimeout` | ***int***| ***(Optional)*** |
| `hecEndpoint` | ***string***||
| `hecEndpointType` | ***string***| ***(Optional)*** |
| `hecToken` | ***string***||
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#KinesisFirehoseDeliveryStreamSpecSplunkConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#KinesisFirehoseDeliveryStreamSpecSplunkConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors](#KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamStatus

Appears on:[KinesisFirehoseDeliveryStream](#KinesisFirehoseDeliveryStream)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KinesisFirehoseDeliveryStreamSpec](#KinesisFirehoseDeliveryStreamSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `redshift_configuration.<index>.password` | ***string*** ||
