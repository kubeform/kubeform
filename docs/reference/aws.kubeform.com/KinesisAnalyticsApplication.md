---
title: KinesisAnalyticsApplication
menu:
  docs_{{ .version }}:
    identifier: kinesisanalyticsapplication-aws.kubeform.com
    name: KinesisAnalyticsApplication
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KinesisAnalyticsApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KinesisAnalyticsApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec)***||
| `status` | ***[KinesisAnalyticsApplicationStatus](#KinesisAnalyticsApplicationStatus)***||
## KinesisAnalyticsApplicationSpec
##### (Appears on:[KinesisAnalyticsApplication](#KinesisAnalyticsApplication), [KinesisAnalyticsApplicationStatus](#KinesisAnalyticsApplicationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudwatchLoggingOptions` | ***[[]KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions](#KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions)***| ***(Optional)*** |
| `code` | ***string***| ***(Optional)*** |
| `createTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `inputs` | ***[[]KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs)***| ***(Optional)*** |
| `lastUpdateTimestamp` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `outputs` | ***[[]KinesisAnalyticsApplicationSpecOutputs](#KinesisAnalyticsApplicationSpecOutputs)***| ***(Optional)*** |
| `referenceDataSources` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSources](#KinesisAnalyticsApplicationSpecReferenceDataSources)***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecCloudwatchLoggingOptions
##### (Appears on:[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `logStreamArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecInputs
##### (Appears on:[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `kinesisFirehose` | ***[[]KinesisAnalyticsApplicationSpecInputsKinesisFirehose](#KinesisAnalyticsApplicationSpecInputsKinesisFirehose)***| ***(Optional)*** |
| `kinesisStream` | ***[[]KinesisAnalyticsApplicationSpecInputsKinesisStream](#KinesisAnalyticsApplicationSpecInputsKinesisStream)***| ***(Optional)*** |
| `namePrefix` | ***string***||
| `parallelism` | ***[[]KinesisAnalyticsApplicationSpecInputsParallelism](#KinesisAnalyticsApplicationSpecInputsParallelism)***| ***(Optional)*** |
| `processingConfiguration` | ***[[]KinesisAnalyticsApplicationSpecInputsProcessingConfiguration](#KinesisAnalyticsApplicationSpecInputsProcessingConfiguration)***| ***(Optional)*** |
| `schema` | ***[[]KinesisAnalyticsApplicationSpecInputsSchema](#KinesisAnalyticsApplicationSpecInputsSchema)***||
| `startingPositionConfiguration` | ***[[]KinesisAnalyticsApplicationSpecInputsStartingPositionConfiguration](#KinesisAnalyticsApplicationSpecInputsStartingPositionConfiguration)***| ***(Optional)*** |
| `streamNames` | ***[]string***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecInputsKinesisFirehose
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecInputsKinesisStream
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecInputsParallelism
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
## KinesisAnalyticsApplicationSpecInputsProcessingConfiguration
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `lambda` | ***[[]KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda](#KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda)***||
## KinesisAnalyticsApplicationSpecInputsProcessingConfigurationLambda
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsProcessingConfiguration](#KinesisAnalyticsApplicationSpecInputsProcessingConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecInputsSchema
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordColumns` | ***[[]KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns](#KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns)***||
| `recordEncoding` | ***string***| ***(Optional)*** |
| `recordFormat` | ***[[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat)***||
## KinesisAnalyticsApplicationSpecInputsSchemaRecordColumns
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsSchema](#KinesisAnalyticsApplicationSpecInputsSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mapping` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `sqlType` | ***string***||
## KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsSchema](#KinesisAnalyticsApplicationSpecInputsSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mappingParameters` | ***[[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters)***| ***(Optional)*** |
| `recordFormatType` | ***string***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormat))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `csv` | ***[[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv)***| ***(Optional)*** |
| `json` | ***[[]KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson)***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersCsv
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordColumnDelimiter` | ***string***||
| `recordRowDelimiter` | ***string***||
## KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParametersJson
##### (Appears on:[KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecInputsSchemaRecordFormatMappingParameters))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordRowPath` | ***string***||
## KinesisAnalyticsApplicationSpecInputsStartingPositionConfiguration
##### (Appears on:[KinesisAnalyticsApplicationSpecInputs](#KinesisAnalyticsApplicationSpecInputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `startingPosition` | ***string***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecOutputs
##### (Appears on:[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `kinesisFirehose` | ***[[]KinesisAnalyticsApplicationSpecOutputsKinesisFirehose](#KinesisAnalyticsApplicationSpecOutputsKinesisFirehose)***| ***(Optional)*** |
| `kinesisStream` | ***[[]KinesisAnalyticsApplicationSpecOutputsKinesisStream](#KinesisAnalyticsApplicationSpecOutputsKinesisStream)***| ***(Optional)*** |
| `lambda` | ***[[]KinesisAnalyticsApplicationSpecOutputsLambda](#KinesisAnalyticsApplicationSpecOutputsLambda)***| ***(Optional)*** |
| `name` | ***string***||
| `schema` | ***[[]KinesisAnalyticsApplicationSpecOutputsSchema](#KinesisAnalyticsApplicationSpecOutputsSchema)***||
## KinesisAnalyticsApplicationSpecOutputsKinesisFirehose
##### (Appears on:[KinesisAnalyticsApplicationSpecOutputs](#KinesisAnalyticsApplicationSpecOutputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecOutputsKinesisStream
##### (Appears on:[KinesisAnalyticsApplicationSpecOutputs](#KinesisAnalyticsApplicationSpecOutputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecOutputsLambda
##### (Appears on:[KinesisAnalyticsApplicationSpecOutputs](#KinesisAnalyticsApplicationSpecOutputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceArn` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecOutputsSchema
##### (Appears on:[KinesisAnalyticsApplicationSpecOutputs](#KinesisAnalyticsApplicationSpecOutputs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordFormatType` | ***string***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecReferenceDataSources
##### (Appears on:[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `s3` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesS3](#KinesisAnalyticsApplicationSpecReferenceDataSourcesS3)***||
| `schema` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema)***||
| `tableName` | ***string***||
## KinesisAnalyticsApplicationSpecReferenceDataSourcesS3
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSources](#KinesisAnalyticsApplicationSpecReferenceDataSources))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketArn` | ***string***||
| `fileKey` | ***string***||
| `roleArn` | ***string***||
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSources](#KinesisAnalyticsApplicationSpecReferenceDataSources))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordColumns` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns)***||
| `recordEncoding` | ***string***| ***(Optional)*** |
| `recordFormat` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat)***||
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordColumns
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mapping` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `sqlType` | ***string***||
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mappingParameters` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters)***| ***(Optional)*** |
| `recordFormatType` | ***string***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormat))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `csv` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv)***| ***(Optional)*** |
| `json` | ***[[]KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson)***| ***(Optional)*** |
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersCsv
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordColumnDelimiter` | ***string***||
| `recordRowDelimiter` | ***string***||
## KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParametersJson
##### (Appears on:[KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters](#KinesisAnalyticsApplicationSpecReferenceDataSourcesSchemaRecordFormatMappingParameters))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `recordRowPath` | ***string***||
## KinesisAnalyticsApplicationStatus
##### (Appears on:[KinesisAnalyticsApplication](#KinesisAnalyticsApplication))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KinesisAnalyticsApplicationSpec](#KinesisAnalyticsApplicationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
