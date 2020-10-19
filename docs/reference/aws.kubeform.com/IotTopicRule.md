---
title: IotTopicRule
menu:
  docs_{{ .version }}:
    identifier: iottopicrule-aws.kubeform.com
    name: IotTopicRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IotTopicRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotTopicRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotTopicRuleSpec](#iottopicrulespec)***||
| `status` | ***[IotTopicRuleStatus](#iottopicrulestatus)***||
## IotTopicRuleSpec

Appears on:[IotTopicRule](#iottopicrule), [IotTopicRuleStatus](#iottopicrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudwatchAlarm` | ***[[]IotTopicRuleSpecCloudwatchAlarm](#iottopicrulespeccloudwatchalarm)***| ***(Optional)*** |
| `cloudwatchMetric` | ***[[]IotTopicRuleSpecCloudwatchMetric](#iottopicrulespeccloudwatchmetric)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dynamodb` | ***[[]IotTopicRuleSpecDynamodb](#iottopicrulespecdynamodb)***| ***(Optional)*** |
| `elasticsearch` | ***[[]IotTopicRuleSpecElasticsearch](#iottopicrulespecelasticsearch)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `firehose` | ***[[]IotTopicRuleSpecFirehose](#iottopicrulespecfirehose)***| ***(Optional)*** |
| `kinesis` | ***[[]IotTopicRuleSpecKinesis](#iottopicrulespeckinesis)***| ***(Optional)*** |
| `lambda` | ***[[]IotTopicRuleSpecLambda](#iottopicrulespeclambda)***| ***(Optional)*** |
| `name` | ***string***||
| `republish` | ***[[]IotTopicRuleSpecRepublish](#iottopicrulespecrepublish)***| ***(Optional)*** |
| `s3` | ***[[]IotTopicRuleSpecS3](#iottopicrulespecs3)***| ***(Optional)*** |
| `sns` | ***[[]IotTopicRuleSpecSns](#iottopicrulespecsns)***| ***(Optional)*** |
| `sql` | ***string***||
| `sqlVersion` | ***string***||
| `sqs` | ***[[]IotTopicRuleSpecSqs](#iottopicrulespecsqs)***| ***(Optional)*** |
## IotTopicRuleSpecCloudwatchAlarm

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `alarmName` | ***string***||
| `roleArn` | ***string***||
| `stateReason` | ***string***||
| `stateValue` | ***string***||
## IotTopicRuleSpecCloudwatchMetric

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricName` | ***string***||
| `metricNamespace` | ***string***||
| `metricTimestamp` | ***string***| ***(Optional)*** |
| `metricUnit` | ***string***||
| `metricValue` | ***string***||
| `roleArn` | ***string***||
## IotTopicRuleSpecDynamodb

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

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

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
| `ID` | ***string***||
| `index` | ***string***||
| `roleArn` | ***string***||
| `type` | ***string***||
## IotTopicRuleSpecFirehose

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deliveryStreamName` | ***string***||
| `roleArn` | ***string***||
| `separator` | ***string***| ***(Optional)*** |
## IotTopicRuleSpecKinesis

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `partitionKey` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `streamName` | ***string***||
## IotTopicRuleSpecLambda

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
## IotTopicRuleSpecRepublish

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `roleArn` | ***string***||
| `topic` | ***string***||
## IotTopicRuleSpecS3

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `key` | ***string***||
| `roleArn` | ***string***||
## IotTopicRuleSpecSns

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `messageFormat` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `targetArn` | ***string***||
## IotTopicRuleSpecSqs

Appears on:[IotTopicRuleSpec](#iottopicrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `queueURL` | ***string***||
| `roleArn` | ***string***||
| `useBase64` | ***bool***||
## IotTopicRuleStatus

Appears on:[IotTopicRule](#iottopicrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotTopicRuleSpec](#iottopicrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IotTopicRuleStatus](#iottopicrulestatus)

---
