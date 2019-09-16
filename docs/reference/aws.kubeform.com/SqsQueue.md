---
title: SqsQueue
menu:
  docs_{{ .version }}:
    identifier: sqsqueue-aws.kubeform.com
    name: SqsQueue
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqsQueue
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SqsQueue` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqsQueueSpec](#SqsQueueSpec)***||
| `status` | ***[SqsQueueStatus](#SqsQueueStatus)***||
## SqsQueueSpec

Appears on:[SqsQueue](#SqsQueue), [SqsQueueStatus](#SqsQueueStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `contentBasedDeduplication` | ***bool***| ***(Optional)*** |
| `delaySeconds` | ***int***| ***(Optional)*** |
| `fifoQueue` | ***bool***| ***(Optional)*** |
| `kmsDataKeyReusePeriodSeconds` | ***int***| ***(Optional)*** |
| `kmsMasterKeyID` | ***string***| ***(Optional)*** |
| `maxMessageSize` | ***int***| ***(Optional)*** |
| `messageRetentionSeconds` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `receiveWaitTimeSeconds` | ***int***| ***(Optional)*** |
| `redrivePolicy` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `visibilityTimeoutSeconds` | ***int***| ***(Optional)*** |
## SqsQueueStatus

Appears on:[SqsQueue](#SqsQueue)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqsQueueSpec](#SqsQueueSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
