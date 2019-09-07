---
title: IotTopicRule
menu:
  docs_v0.0.1:
    identifier: iottopicrule-aws.kubeform.com
    name: IotTopicRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IotTopicRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotTopicRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotTopicRuleSpec](#IotTopicRuleSpec)***||
| `status` | ***[IotTopicRuleStatus](#IotTopicRuleStatus)***||
## IotTopicRuleSpec
##### (Appears on:[IotTopicRule](#IotTopicRule), [IotTopicRuleStatus](#IotTopicRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudwatchAlarm` | ***[[]IotTopicRuleSpecCloudwatchAlarm](#IotTopicRuleSpecCloudwatchAlarm)***| ***(Optional)*** |
| `cloudwatchMetric` | ***[[]IotTopicRuleSpecCloudwatchMetric](#IotTopicRuleSpecCloudwatchMetric)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dynamodb` | ***[[]IotTopicRuleSpecDynamodb](#IotTopicRuleSpecDynamodb)***| ***(Optional)*** |
| `elasticsearch` | ***[[]IotTopicRuleSpecElasticsearch](#IotTopicRuleSpecElasticsearch)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `firehose` | ***[[]IotTopicRuleSpecFirehose](#IotTopicRuleSpecFirehose)***| ***(Optional)*** |
| `kinesis` | ***[[]IotTopicRuleSpecKinesis](#IotTopicRuleSpecKinesis)***| ***(Optional)*** |
| `lambda` | ***[[]IotTopicRuleSpecLambda](#IotTopicRuleSpecLambda)***| ***(Optional)*** |
| `name` | ***string***||
| `republish` | ***[[]IotTopicRuleSpecRepublish](#IotTopicRuleSpecRepublish)***| ***(Optional)*** |
| `s3` | ***[[]IotTopicRuleSpecS3](#IotTopicRuleSpecS3)***| ***(Optional)*** |
| `sns` | ***[[]IotTopicRuleSpecSns](#IotTopicRuleSpecSns)***| ***(Optional)*** |
| `sql` | ***string***||
| `sqlVersion` | ***string***||
| `sqs` | ***[[]IotTopicRuleSpecSqs](#IotTopicRuleSpecSqs)***| ***(Optional)*** |
## IotTopicRuleSpecCloudwatchAlarm
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alarmName` | ***string***||
| `roleArn` | ***string***||
| `stateReason` | ***string***||
| `stateValue` | ***string***||
## IotTopicRuleSpecCloudwatchMetric
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricName` | ***string***||
| `metricNamespace` | ***string***||
| `metricTimestamp` | ***string***| ***(Optional)*** |
| `metricUnit` | ***string***||
| `metricValue` | ***string***||
| `roleArn` | ***string***||
## IotTopicRuleSpecDynamodb
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hashKeyField` | ***string***||
| `hashKeyType` | ***string***| ***(Optional)*** |
| `hashKeyValue` | ***string***||
| `payloadField` | ***string***| ***(Optional)*** |
| `rangeKeyField` | ***string***| ***(Optional)*** |
| `rangeKeyType` | ***string***| ***(Optional)*** |
| `rangeKeyValue` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `tableName` | ***string***||
## IotTopicRuleSpecElasticsearch
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
| `ID` | ***string***||
| `index` | ***string***||
| `roleArn` | ***string***||
| `type` | ***string***||
## IotTopicRuleSpecFirehose
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deliveryStreamName` | ***string***||
| `roleArn` | ***string***||
| `separator` | ***string***| ***(Optional)*** |
## IotTopicRuleSpecKinesis
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `partitionKey` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `streamName` | ***string***||
## IotTopicRuleSpecLambda
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
## IotTopicRuleSpecRepublish
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `roleArn` | ***string***||
| `topic` | ***string***||
## IotTopicRuleSpecS3
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `key` | ***string***||
| `roleArn` | ***string***||
## IotTopicRuleSpecSns
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `messageFormat` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `targetArn` | ***string***||
## IotTopicRuleSpecSqs
##### (Appears on:[IotTopicRuleSpec](#IotTopicRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `queueURL` | ***string***||
| `roleArn` | ***string***||
| `useBase64` | ***bool***||
## IotTopicRuleStatus
##### (Appears on:[IotTopicRule](#IotTopicRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotTopicRuleSpec](#IotTopicRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
