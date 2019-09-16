---
title: BatchJobQueue
menu:
  docs_{{ .version }}:
    identifier: batchjobqueue-aws.kubeform.com
    name: BatchJobQueue
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BatchJobQueue
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchJobQueue` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchJobQueueSpec](#batchjobqueuespec)***||
| `status` | ***[BatchJobQueueStatus](#batchjobqueuestatus)***||
## BatchJobQueueSpec

Appears on:[BatchJobQueue](#batchjobqueue), [BatchJobQueueStatus](#batchjobqueuestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `computeEnvironments` | ***[]string***||
| `name` | ***string***||
| `priority` | ***int***||
| `state` | ***string***||
## BatchJobQueueStatus

Appears on:[BatchJobQueue](#batchjobqueue)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchJobQueueSpec](#batchjobqueuespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
