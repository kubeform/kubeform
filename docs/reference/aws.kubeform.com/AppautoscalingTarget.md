---
title: AppautoscalingTarget
menu:
  docs_{{ .version }}:
    identifier: appautoscalingtarget-aws.kubeform.com
    name: AppautoscalingTarget
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppautoscalingTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppautoscalingTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppautoscalingTargetSpec](#AppautoscalingTargetSpec)***||
| `status` | ***[AppautoscalingTargetStatus](#AppautoscalingTargetStatus)***||
## AppautoscalingTargetSpec
##### (Appears on:[AppautoscalingTarget](#AppautoscalingTarget), [AppautoscalingTargetStatus](#AppautoscalingTargetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `maxCapacity` | ***int***||
| `minCapacity` | ***int***||
| `resourceID` | ***string***||
| `roleArn` | ***string***| ***(Optional)*** |
| `scalableDimension` | ***string***||
| `serviceNamespace` | ***string***||
## AppautoscalingTargetStatus
##### (Appears on:[AppautoscalingTarget](#AppautoscalingTarget))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppautoscalingTargetSpec](#AppautoscalingTargetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
