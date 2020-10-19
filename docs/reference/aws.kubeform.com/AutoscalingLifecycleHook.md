---
title: AutoscalingLifecycleHook
menu:
  docs_{{ .version }}:
    identifier: autoscalinglifecyclehook-aws.kubeform.com
    name: AutoscalingLifecycleHook
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutoscalingLifecycleHook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingLifecycleHook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingLifecycleHookSpec](#autoscalinglifecyclehookspec)***||
| `status` | ***[AutoscalingLifecycleHookStatus](#autoscalinglifecyclehookstatus)***||
## AutoscalingLifecycleHookSpec

Appears on:[AutoscalingLifecycleHook](#autoscalinglifecyclehook), [AutoscalingLifecycleHookStatus](#autoscalinglifecyclehookstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscalingGroupName` | ***string***||
| `defaultResult` | ***string***| ***(Optional)*** |
| `heartbeatTimeout` | ***int64***| ***(Optional)*** |
| `lifecycleTransition` | ***string***||
| `name` | ***string***||
| `notificationMetadata` | ***string***| ***(Optional)*** |
| `notificationTargetArn` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***| ***(Optional)*** |
## AutoscalingLifecycleHookStatus

Appears on:[AutoscalingLifecycleHook](#autoscalinglifecyclehook)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingLifecycleHookSpec](#autoscalinglifecyclehookspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutoscalingLifecycleHookStatus](#autoscalinglifecyclehookstatus)

---
