---
title: SqsQueuePolicy
menu:
  docs_v0.0.1:
    identifier: sqsqueuepolicy-aws.kubeform.com
    name: SqsQueuePolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SqsQueuePolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SqsQueuePolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqsQueuePolicySpec](#SqsQueuePolicySpec)***||
| `status` | ***[SqsQueuePolicyStatus](#SqsQueuePolicyStatus)***||
## SqsQueuePolicySpec
##### (Appears on:[SqsQueuePolicy](#SqsQueuePolicy), [SqsQueuePolicyStatus](#SqsQueuePolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policy` | ***string***||
| `queueURL` | ***string***||
## SqsQueuePolicyStatus
##### (Appears on:[SqsQueuePolicy](#SqsQueuePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqsQueuePolicySpec](#SqsQueuePolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
