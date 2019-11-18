---
title: MonitoringAlertPolicy
menu:
  docs_{{ .version }}:
    identifier: monitoringalertpolicy-google.kubeform.com
    name: MonitoringAlertPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitoringAlertPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitoringAlertPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitoringAlertPolicySpec](#monitoringalertpolicyspec)***||
| `status` | ***[MonitoringAlertPolicyStatus](#monitoringalertpolicystatus)***||
## MonitoringAlertPolicySpec

Appears on:[MonitoringAlertPolicy](#monitoringalertpolicy), [MonitoringAlertPolicyStatus](#monitoringalertpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `combiner` | ***string***||
| `conditions` | ***[[]MonitoringAlertPolicySpecConditions](#monitoringalertpolicyspecconditions)***||
| `creationRecord` | ***[[]MonitoringAlertPolicySpecCreationRecord](#monitoringalertpolicyspeccreationrecord)***| ***(Optional)*** |
| `displayName` | ***string***||
| `enabled` | ***bool***||
| `labels` | ***[]string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `notificationChannels` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditions

Appears on:[MonitoringAlertPolicySpec](#monitoringalertpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `conditionAbsent` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsent](#monitoringalertpolicyspecconditionsconditionabsent)***| ***(Optional)*** |
| `conditionThreshold` | ***[[]MonitoringAlertPolicySpecConditionsConditionThreshold](#monitoringalertpolicyspecconditionsconditionthreshold)***| ***(Optional)*** |
| `displayName` | ***string***||
| `name` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsent

Appears on:[MonitoringAlertPolicySpecConditions](#monitoringalertpolicyspecconditions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsentAggregations](#monitoringalertpolicyspecconditionsconditionabsentaggregations)***| ***(Optional)*** |
| `duration` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `trigger` | ***[[]MonitoringAlertPolicySpecConditionsConditionAbsentTrigger](#monitoringalertpolicyspecconditionsconditionabsenttrigger)***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsentAggregations

Appears on:[MonitoringAlertPolicySpecConditionsConditionAbsent](#monitoringalertpolicyspecconditionsconditionabsent)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionAbsentTrigger

Appears on:[MonitoringAlertPolicySpecConditionsConditionAbsent](#monitoringalertpolicyspecconditionsconditionabsent)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***| ***(Optional)*** |
| `percent` | ***float64***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThreshold

Appears on:[MonitoringAlertPolicySpecConditions](#monitoringalertpolicyspecconditions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdAggregations](#monitoringalertpolicyspecconditionsconditionthresholdaggregations)***| ***(Optional)*** |
| `comparison` | ***string***||
| `denominatorAggregations` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations](#monitoringalertpolicyspecconditionsconditionthresholddenominatoraggregations)***| ***(Optional)*** |
| `denominatorFilter` | ***string***| ***(Optional)*** |
| `duration` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `thresholdValue` | ***float64***| ***(Optional)*** |
| `trigger` | ***[[]MonitoringAlertPolicySpecConditionsConditionThresholdTrigger](#monitoringalertpolicyspecconditionsconditionthresholdtrigger)***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdAggregations

Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#monitoringalertpolicyspecconditionsconditionthreshold)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations

Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#monitoringalertpolicyspecconditionsconditionthreshold)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `alignmentPeriod` | ***string***| ***(Optional)*** |
| `crossSeriesReducer` | ***string***| ***(Optional)*** |
| `groupByFields` | ***[]string***| ***(Optional)*** |
| `perSeriesAligner` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicySpecConditionsConditionThresholdTrigger

Appears on:[MonitoringAlertPolicySpecConditionsConditionThreshold](#monitoringalertpolicyspecconditionsconditionthreshold)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***| ***(Optional)*** |
| `percent` | ***float64***| ***(Optional)*** |
## MonitoringAlertPolicySpecCreationRecord

Appears on:[MonitoringAlertPolicySpec](#monitoringalertpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `mutateTime` | ***string***| ***(Optional)*** |
| `mutatedBy` | ***string***| ***(Optional)*** |
## MonitoringAlertPolicyStatus

Appears on:[MonitoringAlertPolicy](#monitoringalertpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitoringAlertPolicySpec](#monitoringalertpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MonitoringAlertPolicyStatus](#monitoringalertpolicystatus)

---
