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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppautoscalingTargetSpec](#appautoscalingtargetspec)***||
| `status` | ***[AppautoscalingTargetStatus](#appautoscalingtargetstatus)***||
## AppautoscalingTargetSpec

Appears on:[AppautoscalingTarget](#appautoscalingtarget), [AppautoscalingTargetStatus](#appautoscalingtargetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `maxCapacity` | ***int64***||
| `minCapacity` | ***int64***||
| `resourceID` | ***string***||
| `roleArn` | ***string***| ***(Optional)*** |
| `scalableDimension` | ***string***||
| `serviceNamespace` | ***string***||
## AppautoscalingTargetStatus

Appears on:[AppautoscalingTarget](#appautoscalingtarget)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppautoscalingTargetSpec](#appautoscalingtargetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppautoscalingTargetStatus](#appautoscalingtargetstatus)

---
