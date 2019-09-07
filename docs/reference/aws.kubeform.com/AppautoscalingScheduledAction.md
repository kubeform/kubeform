---
title: AppautoscalingScheduledAction
menu:
  docs_v0.0.1:
    identifier: appautoscalingscheduledaction-aws.kubeform.com
    name: AppautoscalingScheduledAction
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AppautoscalingScheduledAction
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppautoscalingScheduledAction` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppautoscalingScheduledActionSpec](#AppautoscalingScheduledActionSpec)***||
| `status` | ***[AppautoscalingScheduledActionStatus](#AppautoscalingScheduledActionStatus)***||
## AppautoscalingScheduledActionSpec
##### (Appears on:[AppautoscalingScheduledAction](#AppautoscalingScheduledAction), [AppautoscalingScheduledActionStatus](#AppautoscalingScheduledActionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `endTime` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceID` | ***string***||
| `scalableDimension` | ***string***| ***(Optional)*** |
| `scalableTargetAction` | ***[[]AppautoscalingScheduledActionSpecScalableTargetAction](#AppautoscalingScheduledActionSpecScalableTargetAction)***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
| `serviceNamespace` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
## AppautoscalingScheduledActionSpecScalableTargetAction
##### (Appears on:[AppautoscalingScheduledActionSpec](#AppautoscalingScheduledActionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxCapacity` | ***int***| ***(Optional)*** |
| `minCapacity` | ***int***| ***(Optional)*** |
## AppautoscalingScheduledActionStatus
##### (Appears on:[AppautoscalingScheduledAction](#AppautoscalingScheduledAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppautoscalingScheduledActionSpec](#AppautoscalingScheduledActionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
