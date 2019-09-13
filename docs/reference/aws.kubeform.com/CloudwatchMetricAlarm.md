---
title: CloudwatchMetricAlarm
menu:
  docs_{{ .version }}:
    identifier: cloudwatchmetricalarm-aws.kubeform.com
    name: CloudwatchMetricAlarm
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchMetricAlarm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchMetricAlarm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchMetricAlarmSpec](#CloudwatchMetricAlarmSpec)***||
| `status` | ***[CloudwatchMetricAlarmStatus](#CloudwatchMetricAlarmStatus)***||
## CloudwatchMetricAlarmSpec
##### (Appears on:[CloudwatchMetricAlarm](#CloudwatchMetricAlarm), [CloudwatchMetricAlarmStatus](#CloudwatchMetricAlarmStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `actionsEnabled` | ***bool***| ***(Optional)*** |
| `alarmActions` | ***[]string***| ***(Optional)*** |
| `alarmDescription` | ***string***| ***(Optional)*** |
| `alarmName` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `comparisonOperator` | ***string***||
| `datapointsToAlarm` | ***int***| ***(Optional)*** |
| `dimensions` | ***map[string]string***| ***(Optional)*** |
| `evaluateLowSampleCountPercentiles` | ***string***| ***(Optional)*** |
| `evaluationPeriods` | ***int***||
| `extendedStatistic` | ***string***| ***(Optional)*** |
| `insufficientDataActions` | ***[]string***| ***(Optional)*** |
| `metricName` | ***string***| ***(Optional)*** |
| `metricQuery` | ***[[]CloudwatchMetricAlarmSpecMetricQuery](#CloudwatchMetricAlarmSpecMetricQuery)***| ***(Optional)*** |
| `namespace` | ***string***| ***(Optional)*** |
| `okActions` | ***[]string***| ***(Optional)*** |
| `period` | ***int***| ***(Optional)*** |
| `statistic` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `threshold` | ***encoding/json.Number***||
| `treatMissingData` | ***string***| ***(Optional)*** |
| `unit` | ***string***| ***(Optional)*** |
## CloudwatchMetricAlarmSpecMetricQuery
##### (Appears on:[CloudwatchMetricAlarmSpec](#CloudwatchMetricAlarmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `expression` | ***string***| ***(Optional)*** |
| `ID` | ***string***||
| `label` | ***string***| ***(Optional)*** |
| `metric` | ***[[]CloudwatchMetricAlarmSpecMetricQueryMetric](#CloudwatchMetricAlarmSpecMetricQueryMetric)***| ***(Optional)*** |
| `returnData` | ***bool***| ***(Optional)*** |
## CloudwatchMetricAlarmSpecMetricQueryMetric
##### (Appears on:[CloudwatchMetricAlarmSpecMetricQuery](#CloudwatchMetricAlarmSpecMetricQuery))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dimensions` | ***map[string]string***| ***(Optional)*** |
| `metricName` | ***string***||
| `namespace` | ***string***| ***(Optional)*** |
| `period` | ***int***||
| `stat` | ***string***||
| `unit` | ***string***| ***(Optional)*** |
## CloudwatchMetricAlarmStatus
##### (Appears on:[CloudwatchMetricAlarm](#CloudwatchMetricAlarm))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchMetricAlarmSpec](#CloudwatchMetricAlarmSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
