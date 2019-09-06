## MonitoringAlertPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitoringAlertPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitoringAlertPolicySpec](#MonitoringAlertPolicySpec)***||
| `status` | ***[MonitoringAlertPolicyStatus](#MonitoringAlertPolicyStatus)***||
## MonitoringAlertPolicySpec
##### (Appears on:[MonitoringAlertPolicy](#MonitoringAlertPolicy), [MonitoringAlertPolicyStatus](#MonitoringAlertPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `combiner` | ***string***||
| `conditions` | ***[[]MonitoringAlertPolicySpecConditions](#MonitoringAlertPolicySpecConditions)***||
| `creationRecord` | ***[[]MonitoringAlertPolicySpecCreationRecord](#MonitoringAlertPolicySpecCreationRecord)***| ***(Optional)*** |
| `displayName` | ***string***||
| `enabled` | ***bool***||
| `labels` | ***[]string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `notificationChannels` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditions
##### (Appears on:[MonitoringAlertPolicySpec](#MonitoringAlertPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `conditionAbsent` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsent](#MonitoringAlertPolicySpecConditionsConditionAbsent)***| ***(Optional)*** |
| `conditionThreshold` | ***[[]MonitoringAlertPolicySpecConditionsConditionThreshold](#MonitoringAlertPolicySpecConditionsConditionThreshold)***| ***(Optional)*** |
| `displayName` | ***string***||
| `name` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsent
##### (Appears on:[MonitoringAlertPolicySpecConditions](#MonitoringAlertPolicySpecConditions))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `aggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsentAggregations](#MonitoringAlertPolicySpecConditionsConditionAbsentAggregations)***| ***(Optional)*** |
| `duration` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `trigger` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsentTrigger](#MonitoringAlertPolicySpecConditionsConditionAbsentTrigger)***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsentAggregations
##### (Appears on:[MonitoringAlertPolicySpecConditionsConditionAbsent](#MonitoringAlertPolicySpecConditionsConditionAbsent))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsentTrigger
##### (Appears on:[MonitoringAlertPolicySpecConditionsConditionAbsent](#MonitoringAlertPolicySpecConditionsConditionAbsent))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `percent` | ***encoding/json.Number***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThreshold
##### (Appears on:[MonitoringAlertPolicySpecConditions](#MonitoringAlertPolicySpecConditions))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `aggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdAggregations](#MonitoringAlertPolicySpecConditionsConditionThresholdAggregations)***| ***(Optional)*** |
| `comparison` | ***string***||
| `denominatorAggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations](#MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations)***| ***(Optional)*** |
| `denominatorFilter` | ***string***| ***(Optional)*** |
| `duration` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `thresholdValue` | ***encoding/json.Number***| ***(Optional)*** |
| `trigger` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdTrigger](#MonitoringAlertPolicySpecConditionsConditionThresholdTrigger)***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdAggregations
##### (Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#MonitoringAlertPolicySpecConditionsConditionThreshold))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations
##### (Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#MonitoringAlertPolicySpecConditionsConditionThreshold))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdTrigger
##### (Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#MonitoringAlertPolicySpecConditionsConditionThreshold))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `percent` | ***encoding/json.Number***| ***(Optional)*** |
## MonitoringAlertPolicySpecCreationRecord
##### (Appears on:[MonitoringAlertPolicySpec](#MonitoringAlertPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mutateTime` | ***string***| ***(Optional)*** |
| `mutatedBy` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicyStatus
##### (Appears on:[MonitoringAlertPolicy](#MonitoringAlertPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitoringAlertPolicySpec](#MonitoringAlertPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
