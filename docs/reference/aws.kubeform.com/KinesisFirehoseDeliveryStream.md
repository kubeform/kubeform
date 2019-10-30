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
| `spec` | ***[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)***||
| `status` | ***[KinesisFirehoseDeliveryStreamStatus](#kinesisfirehosedeliverystreamstatus)***||
## KinesisFirehoseDeliveryStreamSpec

Appears on:[KinesisFirehoseDeliveryStream](#kinesisfirehosedeliverystream), [KinesisFirehoseDeliveryStreamStatus](#kinesisfirehosedeliverystreamstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `destination` | ***string***||
| `destinationID` | ***string***| ***(Optional)*** |
| `elasticsearchConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#kinesisfirehosedeliverystreamspecelasticsearchconfiguration)***| ***(Optional)*** |
| `extendedS3Configuration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#kinesisfirehosedeliverystreamspecextendeds3configuration)***| ***(Optional)*** |
| `kinesisSourceConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration](#kinesisfirehosedeliverystreamspeckinesissourceconfiguration)***| ***(Optional)*** |
| `name` | ***string***||
| `redshiftConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfiguration)***| ***(Optional)*** |
| `s3Configuration` | ***[[]KinesisFirehoseDeliveryStreamSpecS3Configuration](#kinesisfirehosedeliverystreamspecs3configuration)***| ***(Optional)*** |
| `splunkConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#kinesisfirehosedeliverystreamspecsplunkconfiguration)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `versionID` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bufferingInterval` | ***int***| ***(Optional)*** |
| `bufferingSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `domainArn` | ***string***||
| `indexName` | ***string***||
| `indexRotationPeriod` | ***string***| ***(Optional)*** |
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationprocessingconfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupMode` | ***string***| ***(Optional)*** |
| `typeName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#kinesisfirehosedeliverystreamspecelasticsearchconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration](#kinesisfirehosedeliverystreamspecelasticsearchconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationprocessingconfigurationprocessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationprocessingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationprocessingconfigurationprocessorsparameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecelasticsearchconfigurationprocessingconfigurationprocessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecextendeds3configurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `dataFormatConversionConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfiguration)***| ***(Optional)*** |
| `errorOutputPrefix` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationprocessingconfiguration)***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurations3backupconfiguration)***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#kinesisfirehosedeliverystreamspecextendeds3configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#kinesisfirehosedeliverystreamspecextendeds3configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `inputFormatConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfiguration)***||
| `outputFormatConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfiguration)***||
| `schemaConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationschemaconfiguration)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deserializer` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfigurationdeserializer)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hiveJSONSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfigurationdeserializerhivejsonserde)***| ***(Optional)*** |
| `openXJSONSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfigurationdeserializeropenxjsonserde)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJSONSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfigurationdeserializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `timestampFormats` | ***[]string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJSONSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationinputformatconfigurationdeserializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caseInsensitive` | ***bool***| ***(Optional)*** |
| `columnToJSONKeyMappings` | ***map[string]string***| ***(Optional)*** |
| `convertDotsInJSONKeysToUnderscores` | ***bool***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `serializer` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfigurationserializer)***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `orcSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfigurationserializerorcserde)***| ***(Optional)*** |
| `parquetSerDe` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfigurationserializerparquetserde)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfigurationserializer)

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

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfigurationoutputformatconfigurationserializer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `blockSizeBytes` | ***int***| ***(Optional)*** |
| `compression` | ***string***| ***(Optional)*** |
| `enableDictionaryCompression` | ***bool***| ***(Optional)*** |
| `maxPaddingBytes` | ***int***| ***(Optional)*** |
| `pageSizeBytes` | ***int***| ***(Optional)*** |
| `writerVersion` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationDataFormatConversionConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationdataformatconversionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `catalogID` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `region` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `tableName` | ***string***||
| `versionID` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#kinesisfirehosedeliverystreamspecextendeds3configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecextendeds3configurationprocessingconfigurationprocessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurationprocessingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters](#kinesisfirehosedeliverystreamspecextendeds3configurationprocessingconfigurationprocessorsparameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecextendeds3configurationprocessingconfigurationprocessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3Configuration](#kinesisfirehosedeliverystreamspecextendeds3configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecextendeds3configurations3backupconfigurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration](#kinesisfirehosedeliverystreamspecextendeds3configurations3backupconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `kinesisStreamArn` | ***string***||
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecredshiftconfigurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `clusterJdbcurl` | ***string***||
| `copyOptions` | ***string***| ***(Optional)*** |
| `dataTableColumns` | ***string***| ***(Optional)*** |
| `dataTableName` | ***string***||
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfigurationprocessingconfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***||
| `s3BackupConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfigurations3backupconfiguration)***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
| `username` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecredshiftconfigurationprocessingconfigurationprocessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfigurationprocessingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters](#kinesisfirehosedeliverystreamspecredshiftconfigurationprocessingconfigurationprocessorsparameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecredshiftconfigurationprocessingconfigurationprocessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecredshiftconfigurations3backupconfigurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration](#kinesisfirehosedeliverystreamspecredshiftconfigurations3backupconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecS3Configuration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `bufferInterval` | ***int***| ***(Optional)*** |
| `bufferSize` | ***int***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecs3configurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `compressionFormat` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
## KinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecS3Configuration](#kinesisfirehosedeliverystreamspecs3configuration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLoggingOptions` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions](#kinesisfirehosedeliverystreamspecsplunkconfigurationcloudwatchloggingoptions)***| ***(Optional)*** |
| `hecAcknowledgmentTimeout` | ***int***| ***(Optional)*** |
| `hecEndpoint` | ***string***||
| `hecEndpointType` | ***string***| ***(Optional)*** |
| `hecToken` | ***string***||
| `processingConfiguration` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecsplunkconfigurationprocessingconfiguration)***| ***(Optional)*** |
| `retryDuration` | ***int***| ***(Optional)*** |
| `s3BackupMode` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#kinesisfirehosedeliverystreamspecsplunkconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** |
| `logStreamName` | ***string***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfiguration](#kinesisfirehosedeliverystreamspecsplunkconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `processors` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecsplunkconfigurationprocessingconfigurationprocessors)***| ***(Optional)*** |
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration](#kinesisfirehosedeliverystreamspecsplunkconfigurationprocessingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameters` | ***[[]KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters](#kinesisfirehosedeliverystreamspecsplunkconfigurationprocessingconfigurationprocessorsparameters)***| ***(Optional)*** |
| `type` | ***string***||
## KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters

Appears on:[KinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors](#kinesisfirehosedeliverystreamspecsplunkconfigurationprocessingconfigurationprocessors)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameterName` | ***string***||
| `parameterValue` | ***string***||
## KinesisFirehoseDeliveryStreamStatus

Appears on:[KinesisFirehoseDeliveryStream](#kinesisfirehosedeliverystream)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KinesisFirehoseDeliveryStreamSpec](#kinesisfirehosedeliverystreamspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KinesisFirehoseDeliveryStreamStatus](#kinesisfirehosedeliverystreamstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `redshift_configuration.<index>.password` | ***string*** ||
