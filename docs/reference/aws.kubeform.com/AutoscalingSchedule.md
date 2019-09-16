---
title: AutoscalingSchedule
menu:
  docs_{{ .version }}:
    identifier: autoscalingschedule-aws.kubeform.com
    name: AutoscalingSchedule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutoscalingSchedule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingSchedule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingScheduleSpec](#AutoscalingScheduleSpec)***||
| `status` | ***[AutoscalingScheduleStatus](#AutoscalingScheduleStatus)***||
## AutoscalingScheduleSpec

Appears on:[AutoscalingSchedule](#AutoscalingSchedule), [AutoscalingScheduleStatus](#AutoscalingScheduleStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `autoscalingGroupName` | ***string***||
| `desiredCapacity` | ***int***| ***(Optional)*** |
| `endTime` | ***string***| ***(Optional)*** |
| `maxSize` | ***int***| ***(Optional)*** |
| `minSize` | ***int***| ***(Optional)*** |
| `recurrence` | ***string***| ***(Optional)*** |
| `scheduledActionName` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
## AutoscalingScheduleStatus

Appears on:[AutoscalingSchedule](#AutoscalingSchedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingScheduleSpec](#AutoscalingScheduleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
